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
	botMemberID      = "U03SCN7MYEQ"
	botChannelID     = "D03SNU2C80H"
	RECIPE_DB_ID     = "ff24a40498c94ac3ac2fa8894ac0d489"
	RECIPE_ORIGINAL  = "%5CiX%60"
	RECIPE_EVAL      = "Ha%3Ba"
	RECIPE_CATEGORY  = "gmv%3A"
	RECIPE_HEADER_ID = "60a4999c-b1fa-4e3d-9d6b-48034ad7b675"
)

// Bot はGoogle Cloud Functionsへの応答を行うクラスです
type Bot struct {
	recipeService
	slack  *slack.Client
	notion *notionapi.Client

	actionSetCategory string
	actionCreateMenu  string
	actionRebuild     string
}

func NewBot(slackClient *slack.Client, notionClient *notionapi.Client, hr slackbot.HandlerRegistry) *Bot {
	bot := &Bot{
		recipeService: recipeService{
			client: notionClient,
		},
		slack:  slackClient,
		notion: notionClient,
	}

	bot.actionCreateMenu = hr.GetActionID("create_menu", bot.onCreateMenu)
	bot.actionSetCategory = hr.GetActionID("set_category", bot.onSetCategory)
	bot.actionRebuild = hr.GetActionID("rebuild", bot.onRebuild)

	return bot
}

func init() {
	router := slackbot.New()

	bot := NewBot(
		slack.New(os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN")),
		notionapi.NewClient(os.Getenv("NOTION_API_KEY")),
		router,
	)

	router.Error = func(w http.ResponseWriter, r *http.Request, err error) {
		bot.slack.PostMessage(botChannelID, slack.MsgOptionText(fmt.Sprintf("⚠️ %v", err.Error()), true))
	}
	router.Message = bot.onCallbackMessage

	functions.HTTP("main", router.Route)
}

func (b *Bot) onCallbackMessage(req *http.Request, event *slackevents.MessageEvent) error {
	if req.Header.Get("X-Slack-Retry-Num") != "" {
		return nil // リトライは無視
	} else if event.User == botMemberID {
		return nil // 自分のメッセージは無視
	} else if event.Text == "" {
		return nil // テキストが空のメッセージ（URLプレビュー削除とかで送られてくるっぽい？）は無視
	}

	ctx := context.Background()
	ref := slack.NewRefToMessage(event.Channel, event.TimeStamp)
	if url := xurls.Strict.FindString(event.Text); url != "" {
		if strings.Contains(url, "|") {
			url = strings.Split(url, "|")[0]
		}

		if err := b.slack.AddReaction("thumbsup", ref); err != nil {
			return fmt.Errorf("Bot.slack.AddReaction: %w", err)
		}

		page, err := b.autoUpdateRecipePage(ctx, url)
		if err != nil {
			return fmt.Errorf("Bot.autoUpdateRecipePage: %w", err)
		}

		if err := b.PostRecipeBlocks(ctx, event.Channel, page.ID); err != nil {
			return fmt.Errorf("Bot.PostRecipeBlocks: %w", err)
		}
		return nil
	} else {
		if err := b.slack.AddReaction("thinking_face", ref); err != nil {
			return fmt.Errorf("Bot.slack.AddReaction(channel=%v, ts=%v, text=%v) = %w", event.Channel, event.TimeStamp, event.Text, err)
		}
		return nil
	}
}

// Deprecated:
func (b *Bot) autoUpdateRecipePage(ctx context.Context, recipeURL string) (*notionapi.Page, error) {
	// レシピページを取得
	page, err := b.GetRecipeByURL(ctx, recipeURL)
	if err != nil {
		return nil, fmt.Errorf("Bot.GetRecipeByURL: %w", err)
	}

	if page != nil {
		return page, nil
	}

	// レシピページがなければ作成
	return b.CreateRecipe(ctx, recipeURL)
}

func (s *Bot) getRecipeBlocksInfo(ctx context.Context, pageID string) (*RecipeBlocksInfo, error) {
	info := &RecipeBlocksInfo{PageID: pageID}

	if db, err := s.notion.RetrieveDatabase(ctx, RECIPE_DB_ID); err != nil {
		return nil, err
	} else {
		for _, prop := range db.Properties {
			if prop.ID == RECIPE_CATEGORY {
				for _, opt := range prop.Select.Options {
					info.Categories = append(info.Categories, opt.Name)
				}
			}
		}
	}

	if category, err := s.notion.RetrievePagePropertyItem(ctx, pageID, RECIPE_CATEGORY); err != nil {
		return nil, err
	} else {
		info.Category = category.PropertyItem.Select.Name
	}

	if title, err := s.notion.RetrievePagePropertyItem(ctx, pageID, "title"); err != nil {
		return nil, err
	} else {
		for _, item := range title.PropertyItemPagination.Results {
			info.Title += item.Title.Text.Content
		}
		if info.Title == "" {
			info.Title = "無題"
		}
	}

	if page, err := s.notion.RetrievePage(ctx, pageID); err != nil {
		return nil, err
	} else {
		info.PageURL = page.URL
		if page.Icon != nil {
			info.Title = page.Icon.Emoji + info.Title
		}
		if page.Cover != nil {
			if page.Cover.External.URL != "" {
				info.ImageURL = page.Cover.External.URL
			} else if page.Cover.File.URL != "" {
				info.ImageURL = page.Cover.File.URL
			}
		}
	}

	return info, nil
}

func (s *Bot) PostRecipeBlocks(ctx context.Context, channelID string, pageID string) error {
	rbi, err := s.getRecipeBlocksInfo(ctx, pageID)
	if err != nil {
		return fmt.Errorf("chatService.getRecipeBlocksInfo: %w", err)
	}

	_, _, err = s.slack.PostMessage(channelID, slack.MsgOptionBlocks(rbi.ToSlackBlocks(s.actionSetCategory, s.actionCreateMenu, s.actionRebuild)...))
	if err != nil {
		return fmt.Errorf("chatService.slack.PostMessage: %w", err)
	}

	return nil
}

func (s *Bot) UpdateRecipeBlocks(ctx context.Context, channelID string, timestamp string, pageID string) error {
	rbi, err := s.getRecipeBlocksInfo(ctx, pageID)
	if err != nil {
		return fmt.Errorf("chatService.getRecipeBlocksInfo: %w", err)
	}

	_, _, _, err = s.slack.UpdateMessage(channelID, timestamp, slack.MsgOptionBlocks(rbi.ToSlackBlocks(s.actionSetCategory, s.actionCreateMenu, s.actionRebuild)...))
	if err != nil {
		return fmt.Errorf("chatService.slack.UpdateMessage: %w", err)
	}

	return nil
}

func (b *Bot) onCreateMenu(callback *slack.InteractionCallback, action *slack.BlockAction) error {
	ctx := context.Background()
	page, err := b.notion.RetrievePage(ctx, action.Value)
	if err != nil {
		return fmt.Errorf("Bot.notion.RetrievePage: %w", err)
	}
	_, _, err = b.slack.PostMessage(callback.Channel.ID, slack.MsgOptionText(page.URL, true))
	return err
}

func (b *Bot) onSetCategory(callback *slack.InteractionCallback, action *slack.BlockAction) error {
	ctx := context.Background()

	pair := [2]string{}
	if err := json.Unmarshal([]byte(action.SelectedOption.Value), &pair); err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	if err := b.SetRecipeCategory(ctx, pair[0], pair[1]); err != nil {
		return fmt.Errorf("Bot.SetRecipeCategory: %w", err)
	}

	return b.UpdateRecipeBlocks(ctx, callback.Channel.ID, callback.Message.Timestamp, pair[0])
}

func (b *Bot) onRebuild(callback *slack.InteractionCallback, action *slack.BlockAction) error {
	ctx := context.Background()
	return b.UpdateRecipe(ctx, action.Value)
}
