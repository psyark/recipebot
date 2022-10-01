package slack

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mvdan/xurls"
	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/service/notion"
	"github.com/psyark/slackbot"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

const (
	// botChannelID     = "D03SNU2C80H"
	botMemberID      = "U03SCN7MYEQ"
	cookingChannelID = "C03SNSP9HNV" // #料理 チャンネル
)

type Service struct {
	notionService *notion.Service
	client        *slack.Client

	actionSetCategory string
	actionCreateMenu  string
	actionRebuild     string
}

func New(slackClient *slack.Client, notionClient *notionapi.Client, registry *slackbot.HandlerRegistry) *Service {
	var svc *Service
	svc = &Service{
		notionService: notion.New(notionClient),
		client:        slackClient,

		actionCreateMenu: registry.GetActionID("create_menu", func(callback *slack.InteractionCallback, action *slack.BlockAction) error {
			return svc.onCreateMenu(callback, action)
		}),
		actionSetCategory: registry.GetActionID("set_category", func(callback *slack.InteractionCallback, action *slack.BlockAction) error {
			return svc.onSetCategory(callback, action)
		}),
		actionRebuild: registry.GetActionID("rebuild", func(callback *slack.InteractionCallback, action *slack.BlockAction) error {
			return svc.onRebuild(callback, action)
		}),
	}

	return svc
}

func (b *Service) OnError(w http.ResponseWriter, r *http.Request, err error) {
	b.client.PostMessage(cookingChannelID, slack.MsgOptionText(fmt.Sprintf("⚠️ %v", err.Error()), true))
}

func (b *Service) OnCallbackMessage(req *http.Request, event *slackevents.MessageEvent) error {
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

		if err := b.client.AddReaction("thumbsup", ref); err != nil {
			return fmt.Errorf("addReaction: %w", err)
		}

		page, err := b.autoUpdateRecipePage(ctx, url)
		if err != nil {
			return fmt.Errorf("autoUpdateRecipePage: %w", err)
		}

		if err := b.PostRecipeBlocks(ctx, event.Channel, page.ID); err != nil {
			return fmt.Errorf("postRecipeBlocks: %w", err)
		}
		return nil
	} else {
		if err := b.client.AddReaction("thinking_face", ref); err != nil {
			return fmt.Errorf("addReaction(channel=%v, ts=%v, text=%v) = %w", event.Channel, event.TimeStamp, event.Text, err)
		}
		return nil
	}
}

// Deprecated:
func (b *Service) autoUpdateRecipePage(ctx context.Context, recipeURL string) (*notionapi.Page, error) {
	// レシピページを取得
	page, err := b.notionService.GetRecipeByURL(ctx, recipeURL)
	if err != nil {
		return nil, fmt.Errorf("Bot.GetRecipeByURL: %w", err)
	}

	if page != nil {
		return page, nil
	}

	// レシピページがなければ作成
	return b.notionService.CreateRecipe(ctx, recipeURL)
}

func (s *Service) PostRecipeBlocks(ctx context.Context, channelID string, pageID string) error {
	blocks, err := s.getRecipeBlocks(ctx, pageID)
	if err != nil {
		return fmt.Errorf("getRecipeBlocks: %w", err)
	}

	_, _, err = s.client.PostMessage(channelID, slack.MsgOptionBlocks(blocks...))
	if err != nil {
		return fmt.Errorf("postMessage: %w", err)
	}

	return nil
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

func (b *Service) onCreateMenu(callback *slack.InteractionCallback, action *slack.BlockAction) error {
	_, _, err := b.client.PostMessage(callback.Channel.ID, slack.MsgOptionText(fmt.Sprintf("未実装: %v", action.Value), true))
	return err
}

func (b *Service) onSetCategory(callback *slack.InteractionCallback, action *slack.BlockAction) error {
	ctx := context.Background()

	pair := [2]string{}
	if err := json.Unmarshal([]byte(action.SelectedOption.Value), &pair); err != nil {
		return fmt.Errorf("json.Unmarshal: %w", err)
	}

	if err := b.notionService.SetRecipeCategory(ctx, pair[0], pair[1]); err != nil {
		return fmt.Errorf("Bot.SetRecipeCategory: %w", err)
	}

	return b.UpdateRecipeBlocks(ctx, callback.Channel.ID, callback.Message.Timestamp, pair[0])
}

func (b *Service) onRebuild(callback *slack.InteractionCallback, action *slack.BlockAction) error {
	ctx := context.Background()
	return b.notionService.UpdateRecipe(ctx, action.Value)
}

func (b *Service) getRecipeBlocks(ctx context.Context, pageID string) ([]slack.Block, error) {
	var pageURL string
	var thumbnail *slack.Accessory

	// カテゴリーの選択肢の取得
	categories, err := b.notionService.GetRecipeCategories(ctx)
	if err != nil {
		return nil, err
	}

	// 現在のカテゴリーの取得
	category, err := b.notionService.GetRecipeCategory(ctx, pageID)
	if err != nil {
		return nil, err
	}

	// タイトルの取得
	pageTitle, err := b.notionService.GetRecipeTitle(ctx, pageID)
	if err != nil {
		return nil, err
	}
	if pageTitle == "" {
		pageTitle = "無題"
	}

	// ページの取得
	if page, err := b.notionService.RetrievePage(ctx, pageID); err != nil {
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
		slack.NewDividerBlock(),
		slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, "*このレシピの操作*", false, false), nil, nil),
		b.getRecipeBlocks_CategoryBlock(pageID, categories, category),
		b.getRecipeBlocks_MenuBlock(pageID),
		b.getRecipeBlocks_RebuildBlock(pageID),
	}, nil
}

func (b *Service) getRecipeBlocks_CategoryBlock(pageID string, categories []string, category string) slack.Block {
	var catOptions []*slack.OptionBlockObject
	var initialOption *slack.OptionBlockObject

	for _, c := range categories {
		val, _ := json.Marshal([]string{pageID, c})
		opt := slack.NewOptionBlockObject(string(val), slack.NewTextBlockObject(slack.PlainTextType, c, true, false), nil)
		catOptions = append(catOptions, opt)
		if c == category {
			initialOption = opt
		}
	}

	selectBlock := slack.NewOptionsSelectBlockElement(
		slack.OptTypeStatic,
		slack.NewTextBlockObject(slack.PlainTextType, "分類", true, false),
		b.actionSetCategory,
		catOptions...,
	)
	selectBlock.InitialOption = initialOption

	return slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "分類を設定する", false, false),
		nil,
		slack.NewAccessory(selectBlock),
	)
}

func (b *Service) getRecipeBlocks_MenuBlock(pageID string) slack.Block {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "<https://www.notion.so/80cf0a5ec25c4b7489f00594362f6e3b|🍽️献立表>に追加する", false, false),
		nil,
		slack.NewAccessory(slack.NewButtonBlockElement(
			b.actionCreateMenu,
			pageID,
			slack.NewTextBlockObject(slack.PlainTextType, "献立表に追加", true, false),
		)),
	)
}

func (b *Service) getRecipeBlocks_RebuildBlock(pageID string) slack.Block {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "再取得して作り直す", false, false),
		nil,
		slack.NewAccessory(slack.NewButtonBlockElement(
			b.actionRebuild,
			pageID,
			slack.NewTextBlockObject(slack.PlainTextType, "作り直す", true, false),
		)),
	)
}
