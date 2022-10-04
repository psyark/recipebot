package slack

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/mvdan/xurls"
	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/service/notion"
	"github.com/psyark/slackbot"
	"github.com/slack-go/slack"
)

type ofType string

const (
	// botChannelID     = "D03SNU2C80H"
	botMemberID      = "U03SCN7MYEQ"
	cookingChannelID = "C03SNSP9HNV" // #料理 チャンネル

	ofTypeRebuildRecipe     = ofType("rebuildRecipe")
	ofTypeUpdateIngredients = ofType("updateIngredients")
)

type Service struct {
	notion         *notion.Service
	client         *slack.Client
	actionOverflow string
}

func New(slackClient *slack.Client, notionClient *notionapi.Client, registry *slackbot.Registry) *Service {
	var svc *Service
	svc = &Service{
		notion:         notion.New(notionClient),
		client:         slackClient,
		actionOverflow: registry.GetActionID("overflow", func(args *slackbot.BlockActionHandlerArgs) error { return svc.onOverflow(args) }),
	}
	return svc
}

func (b *Service) OnError(args *slackbot.ErrorHandlerArgs) {
	b.client.PostMessage(cookingChannelID, slack.MsgOptionText(fmt.Sprintf("⚠️ %v", args.Err.Error()), true))
}

func (b *Service) OnCallbackMessage(args *slackbot.MessageHandlerArgs) error {
	ctx := context.Background()
	event := args.MessageEvent

	if args.Request.Header.Get("X-Slack-Retry-Num") != "" {
		return nil // リトライは無視
	} else if event.User == botMemberID {
		return nil // 自分のメッセージは無視
	} else if event.Text == "" {
		return nil // テキストが空のメッセージ（URLプレビュー削除とかで送られてくるっぽい？）は無視
	}

	if url := xurls.Strict.FindString(event.Text); url != "" {
		if strings.Contains(url, "|") {
			url = strings.Split(url, "|")[0]
		}

		_, timestamp, err := b.client.PostMessage(event.Channel, slack.MsgOptionText(":hourglass_flowing_sand:", false))
		if err != nil {
			return err
		}

		page, err := b.autoUpdateRecipePage(ctx, url)
		if err != nil {
			return fmt.Errorf("autoUpdateRecipePage: %w", err)
		}

		if err := b.UpdateRecipeBlocks(ctx, event.Channel, timestamp, page.ID); err != nil {
			return fmt.Errorf("postRecipeBlocks: %w", err)
		}
	}

	return nil
}

// Deprecated:
func (b *Service) autoUpdateRecipePage(ctx context.Context, recipeURL string) (*notionapi.Page, error) {
	// レシピページを取得
	page, err := b.notion.GetRecipeByURL(ctx, recipeURL)
	if err != nil {
		return nil, fmt.Errorf("Bot.GetRecipeByURL: %w", err)
	}

	if page != nil {
		return page, nil
	}

	// レシピページがなければ作成
	return b.notion.CreateRecipe(ctx, recipeURL)
}

func (s *Service) UpdateRecipeBlocks(ctx context.Context, channelID string, timestamp string, pageID string) error {
	blocks, err := s.getRecipeBlocks(ctx, pageID)
	if err != nil {
		return fmt.Errorf("getRecipeBlocks: %w", err)
	}

	_, _, _, err = s.client.UpdateMessage(channelID, timestamp, slack.MsgOptionBlocks(blocks...))
	if err != nil {
		return fmt.Errorf("updateMessage: %w", err)
	}

	return nil
}

func (s *Service) onOverflow(args *slackbot.BlockActionHandlerArgs) error {
	ctx := context.Background()

	ofArgs := OverflowArgs{}
	if err := json.Unmarshal([]byte(args.BlockAction.SelectedOption.Value), &ofArgs); err != nil {
		return err
	}

	switch ofArgs.Type {
	case ofTypeRebuildRecipe:
		return s.notion.UpdateRecipe(ctx, ofArgs.PageID)

	case ofTypeUpdateIngredients:
		stockMap, err := s.notion.GetStockMap(ctx)
		if err != nil {
			return err
		}

		result, err := s.notion.UpdateRecipeIngredients(ctx, ofArgs.PageID, stockMap)
		if err != nil {
			return err
		}

		foundItems := []string{}
		notFoundItems := []string{}
		for name, found := range result {
			if found {
				foundItems = append(foundItems, name)
			} else {
				notFoundItems = append(notFoundItems, name)
			}
		}

		if len(foundItems) != 0 {
			sort.Strings(foundItems)
			_, _, err := s.client.PostMessage(args.InteractionCallback.Channel.ID, slack.MsgOptionText(fmt.Sprintf("材料を設定しました: %v", foundItems), true))
			if err != nil {
				return err
			}
		}
		if len(notFoundItems) != 0 {
			sort.Strings(notFoundItems)
			_, _, err := s.client.PostMessage(args.InteractionCallback.Channel.ID, slack.MsgOptionText(fmt.Sprintf("材料が見つかりませんでした: %v", notFoundItems), true))
			if err != nil {
				return err
			}
		}
		return nil
	default:
		return fmt.Errorf("unknown ofType: %v", ofArgs.Type)
	}
}

func (b *Service) getRecipeBlocks(ctx context.Context, pageID string) ([]slack.Block, error) {
	var pageURL string
	var thumbnail *slack.Accessory

	// タイトルの取得
	pageTitle, err := b.notion.GetRecipeTitle(ctx, pageID)
	if err != nil {
		return nil, err
	}
	if pageTitle == "" {
		pageTitle = "無題"
	}

	// ページの取得
	if page, err := b.notion.RetrievePage(ctx, pageID); err != nil {
		return nil, err
	} else {
		pageURL = page.URL
		if page.Icon != nil {
			if emoji, ok := page.Icon.(*notionapi.Emoji); ok {
				pageTitle = emoji.Emoji + pageTitle
			}
		}
		if page.Cover != nil {
			if page.Cover.External != nil {
				thumbnail = slack.NewAccessory(slack.NewImageBlockElement(page.Cover.External.URL, "レシピの写真"))
			} else if page.Cover.File != nil {
				thumbnail = slack.NewAccessory(slack.NewImageBlockElement(page.Cover.File.URL, "レシピの写真"))
			}
		}
	}

	return []slack.Block{
		slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("*<%v|%v>*", pageURL, pageTitle), false, false),
			nil,
			thumbnail,
		),
		slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, "*このレシピの操作*", false, false),
			nil,
			slack.NewAccessory(slack.NewOverflowBlockElement(
				b.actionOverflow,
				slack.NewOptionBlockObject(OverflowArgs{ofTypeRebuildRecipe, pageID}.String(), slack.NewTextBlockObject(slack.PlainTextType, "レシピを再取得", false, false), nil),
				slack.NewOptionBlockObject(OverflowArgs{ofTypeUpdateIngredients, pageID}.String(), slack.NewTextBlockObject(slack.PlainTextType, "主な材料を更新", false, false), nil),
			)),
		),
	}, nil
}

type OverflowArgs struct {
	Type   ofType `json:"type"`
	PageID string `json:"page_id"`
}

func (a OverflowArgs) String() string {
	data, _ := json.Marshal(a)
	return string(data)
}
