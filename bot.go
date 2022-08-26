package recipebot

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mvdan/xurls"
	"github.com/psyark/notionapi"
	"github.com/psyark/slackbot"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
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
	slackbot.RegisterHandler(bot)
}

// SlackのCallbackMessageへの応答
func (b *Bot) OnCallbackMessage(req *http.Request, event *slackevents.MessageEvent) {
	if err := b.onCallbackMessage(req, event); err != nil {
		b.slack.PostMessage(botChannelID, slack.MsgOptionText(err.Error(), true))
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
			return err
		}

		page, err := b.autoUpdateRecipePage(ctx, url)
		if err != nil {
			return err
		}

		return b.PostRecipeBlocks(ctx, event.Channel, page.ID)
	} else {
		return b.slack.AddReaction("thinking_face", ref)
	}
}

// Deprecated:
func (b *Bot) autoUpdateRecipePage(ctx context.Context, recipeURL string) (*notionapi.Page, error) {
	// レシピページを取得
	page, err := b.GetRecipeByURL(ctx, recipeURL)
	if err != nil {
		return nil, err
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
		b.slack.PostMessage(botChannelID, slack.MsgOptionText(err.Error(), true))
	}
}

func (b *Bot) onBlockActions(req *http.Request, event *slack.InteractionCallback) error {
	for _, ba := range event.ActionCallback.BlockActions {
		switch ba.ActionID {
		case actionCreateMenu:
			ctx := context.Background()
			page, err := b.notion.RetrievePage(ctx, ba.Value)
			if err != nil {
				return err
			}
			if _, _, err := b.slack.PostMessage(event.Channel.ID, slack.MsgOptionText(page.URL, true)); err != nil {
				return err
			}

		case actionSetCategory:
			ctx := context.Background()

			pair := [2]string{}
			if err := json.Unmarshal([]byte(ba.SelectedOption.Value), &pair); err != nil {
				return err
			}

			if err := b.SetRecipeCategory(ctx, pair[0], pair[1]); err != nil {
				return err
			}

			return b.UpdateRecipeBlocks(ctx, event.Channel.ID, event.Message.Timestamp, pair[0])

		case actionRebuild:
			ctx := context.Background()
			return b.UpdateRecipe(ctx, ba.Value)
		default:
			return fmt.Errorf("unsupported action: %v", ba.ActionID)
		}
	}
	return nil
}
