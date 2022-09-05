package recipebot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/mvdan/xurls"
	"github.com/psyark/notionapi"
	"github.com/psyark/slackbot"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
	"golang.org/x/xerrors"
)

/*
Google Cloud Console
https://console.cloud.google.com/functions/list?project=notion-recipe-importer

Slack #recipe
https://app.slack.com/client/T03S6UY399B/C03RVMW9S3Z

recipebot
https://api.slack.com/apps/A03SNSS0S81
*/

const (
	actionSetCategory = "set_category"
	actionCreateMenu  = "create_menu"
	actionRebuild     = "rebuild"
	botMemberID       = "U03SCN7MYEQ"
	botChannelID      = "D03SNU2C80H"
	RECIPE_DB_ID      = "ff24a40498c94ac3ac2fa8894ac0d489"
	RECIPE_ORIGINAL   = "%5CiX%60"
	RECIPE_EVAL       = "Ha%3Ba"
	RECIPE_CATEGORY   = "gmv%3A"
	RECIPE_HEADER_ID  = "60a4999c-b1fa-4e3d-9d6b-48034ad7b675"
)

// Bot はGoogle Cloud Functionsへの応答を行うクラスです
type Bot struct {
	recipeService
	chatService
	slack  *slack.Client
	notion *notionapi.Client
}

func NewBot(slackClient *slack.Client, notionClient *notionapi.Client) *Bot {
	return &Bot{
		recipeService: recipeService{
			client: notionClient,
		},
		chatService: chatService{
			notion: notionClient,
			slack:  slackClient,
		},
		slack:  slackClient,
		notion: notionClient,
	}
}

func init() {
	bot := NewBot(
		slack.New(os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN")),
		notionapi.NewClient(os.Getenv("NOTION_API_KEY")),
	)

	x := slackbot.New(bot)
	functions.HTTP("main", x.Handler)
	// slackbot.RegisterHTTP("main", bot)
}

// SlackのCallbackMessageへの応答
func (b *Bot) OnCallbackMessage(req *http.Request, event *slackevents.MessageEvent) {
	if err := b.onCallbackMessage(req, event); err != nil {
		b.slack.PostMessage(botChannelID, slack.MsgOptionText(fmt.Sprintf("OnCallbackMessage: %+v", err), true))
	}
}

func (b *Bot) onCallbackMessage(req *http.Request, event *slackevents.MessageEvent) error {
	if req.Header.Get("X-Slack-Retry-Num") != "" {
		return nil // リトライは無視
	} else if event.User == botMemberID {
		return nil // 自分のメッセージは無視
	}

	ctx := context.Background()
	ref := slack.NewRefToMessage(event.Channel, event.TimeStamp)
	if url := xurls.Strict.FindString(event.Text); url != "" {
		if strings.Contains(url, "|") {
			url = strings.Split(url, "|")[0]
		}

		if err := b.slack.AddReaction("thumbsup", ref); err != nil {
			return xerrors.Errorf("Bot.slack.AddReaction: %w", err)
		}

		page, err := b.autoUpdateRecipePage(ctx, url)
		if err != nil {
			return xerrors.Errorf("Bot.autoUpdateRecipePage: %w", err)
		}

		if err := b.PostRecipeBlocks(ctx, event.Channel, page.ID); err != nil {
			return xerrors.Errorf("Bot.PostRecipeBlocks: %w", err)
		}
		return nil
	} else {
		if err := b.slack.AddReaction("thinking_face", ref); err != nil {
			return xerrors.Errorf("Bot.slack.AddReaction(channel=%v, ts=%v, text=%v) = %w", event.Channel, event.TimeStamp, event.Text, err)
		}
		return nil
	}
}

// Deprecated:
func (b *Bot) autoUpdateRecipePage(ctx context.Context, recipeURL string) (*notionapi.Page, error) {
	// レシピページを取得
	page, err := b.GetRecipeByURL(ctx, recipeURL)
	if err != nil {
		return nil, xerrors.Errorf("Bot.GetRecipeByURL: %w", err)
	}

	if page != nil {
		return page, nil
	}

	// レシピページがなければ作成
	return b.CreateRecipe(ctx, recipeURL)
}

// SlackのBlockActionsへの応答
func (b *Bot) OnBlockActions(req *http.Request, event *slack.InteractionCallback) {
	if err := b.onBlockActions(req, event); err != nil {
		b.slack.PostMessage(botChannelID, slack.MsgOptionText(fmt.Sprintf("OnBlockActions: %+v", err), true))
	}
}

func (b *Bot) onBlockActions(req *http.Request, event *slack.InteractionCallback) error {
	for _, ba := range event.ActionCallback.BlockActions {
		switch ba.ActionID {
		case actionCreateMenu:
			ctx := context.Background()
			page, err := b.notion.RetrievePage(ctx, ba.Value)
			if err != nil {
				return xerrors.Errorf("Bot.notion.RetrievePage: %w", err)
			}
			if _, _, err := b.slack.PostMessage(event.Channel.ID, slack.MsgOptionText(page.URL, true)); err != nil {
				return xerrors.Errorf("Bot.slack.PostMessage: %w", err)
			}

		case actionSetCategory:
			ctx := context.Background()

			pair := [2]string{}
			if err := json.Unmarshal([]byte(ba.SelectedOption.Value), &pair); err != nil {
				return xerrors.Errorf("json.Unmarshal: %w", err)
			}

			if err := b.SetRecipeCategory(ctx, pair[0], pair[1]); err != nil {
				return xerrors.Errorf("Bot.SetRecipeCategory: %w", err)
			}

			return b.UpdateRecipeBlocks(ctx, event.Channel.ID, event.Message.Timestamp, pair[0])

		case actionRebuild:
			ctx := context.Background()
			return b.UpdateRecipe(ctx, ba.Value)
		default:
			return xerrors.Errorf("unsupported action: %v", ba.ActionID)
		}
	}
	return nil
}
