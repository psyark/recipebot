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
	notion *notion.Service
	client *slack.Client

	// actionSetCategory       string
	// actionCreateMenu        string
	// actionRebuild           string
	// actionUpdateIngredients string
	actionOverflow string
}

func New(slackClient *slack.Client, notionClient *notionapi.Client, registry *slackbot.Registry) *Service {
	var svc *Service
	svc = &Service{
		notion: notion.New(notionClient),
		client: slackClient,
		// actionSetCategory:       registry.GetActionID("set_category", func(args *slackbot.BlockActionHandlerArgs) error { return svc.onSetCategory(args) }),
		// actionCreateMenu:        registry.GetActionID("create_menu", func(args *slackbot.BlockActionHandlerArgs) error { return svc.onCreateMenu(args) }),
		// actionRebuild:           registry.GetActionID("rebuild", func(args *slackbot.BlockActionHandlerArgs) error { return svc.onRebuild(args) }),
		// actionUpdateIngredients: registry.GetActionID("update_ingredients", func(args *slackbot.BlockActionHandlerArgs) error { return svc.onUpdateIngredients(args) }),
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

// func (s *Service) PostRecipeBlocks(ctx context.Context, channelID string, pageID string) error {
// 	blocks, err := s.getRecipeBlocks(ctx, pageID)
// 	if err != nil {
// 		return fmt.Errorf("getRecipeBlocks: %w", err)
// 	}

// 	_, _, err = s.client.PostMessage(channelID, slack.MsgOptionBlocks(blocks...))
// 	if err != nil {
// 		return fmt.Errorf("postMessage: %w", err)
// 	}

// 	return nil
// }

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

// func (b *Service) onCreateMenu(args *slackbot.BlockActionHandlerArgs) error {
// 	opt := slack.MsgOptionText(fmt.Sprintf("未実装: %v", args.BlockAction.Value), true)
// 	_, _, err := b.client.PostMessage(args.InteractionCallback.Channel.ID, opt)
// 	return err
// }

// func (b *Service) onSetCategory(args *slackbot.BlockActionHandlerArgs) error {
// 	ctx := context.Background()
// 	callback := args.InteractionCallback
// 	action := args.BlockAction

// 	pair := [2]string{}
// 	if err := json.Unmarshal([]byte(action.SelectedOption.Value), &pair); err != nil {
// 		return fmt.Errorf("json.Unmarshal: %w", err)
// 	}

// 	if err := b.notionService.SetRecipeCategory(ctx, pair[0], pair[1]); err != nil {
// 		return fmt.Errorf("Bot.SetRecipeCategory: %w", err)
// 	}

// 	return b.UpdateRecipeBlocks(ctx, callback.Channel.ID, callback.Message.Timestamp, pair[0])
// }

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

	// カテゴリーの選択肢の取得
	// categories, err := b.notionService.GetRecipeCategories(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	// 現在のカテゴリーの取得
	// category, err := b.notionService.GetRecipeCategory(ctx, pageID)
	// if err != nil {
	// 	return nil, err
	// }

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
			if page.Cover.External.URL != "" {
				thumbnail = slack.NewAccessory(slack.NewImageBlockElement(page.Cover.External.URL, "レシピの写真"))
			} else if page.Cover.File.URL != "" {
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
		// slack.NewActionBlock(
		// 	"", // ブロックID未設定
		// 	slack.NewButtonBlockElement(b.actionRebuild, pageID, slack.NewTextBlockObject(slack.PlainTextType, "レシピを再取得", false, false)),
		// 	slack.NewButtonBlockElement(b.actionUpdateIngredients, pageID, slack.NewTextBlockObject(slack.PlainTextType, "主な材料を更新", false, false)),
		// ),
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
