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
	"github.com/slack-go/slack/slackevents"
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

		return b.ReactMessageWithURL(event, url)
	}

	return nil
}

// ReactMessageWithURL はURL付きのメッセージに反応します
func (s *Service) ReactMessageWithURL(event *slackevents.MessageEvent, url string) error {
	ctx := context.Background()

	// 砂時計のプレースホルダを出しておく
	_, msgTs, err := s.client.PostMessage(event.Channel, slack.MsgOptionText(":hourglass_flowing_sand:", false))
	if err != nil {
		return err
	}

	// URLに対応するレシピページを探す
	page, err := s.notion.GetRecipeByURL(ctx, url)
	if err != nil {
		return err
	}

	// レシピページがなければ作成
	if page == nil {
		page, err = s.notion.CreateRecipe(ctx, url)
		if err != nil {
			return err
		}
	}

	// プレースホルダを更新
	blocks, err := s.getRecipeBlocks(ctx, page)
	if err != nil {
		return err
	}

	_, _, _, err = s.client.UpdateMessage(event.Channel, msgTs, slack.MsgOptionBlocks(blocks...))
	return err
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

func (b *Service) getRecipeBlocks(ctx context.Context, page *notionapi.Page) ([]slack.Block, error) {
	var pageURL string
	var thumbnail *slack.Accessory

	// タイトルの取得
	pageTitle, err := b.notion.GetRecipeTitle(ctx, page.ID)
	if err != nil {
		return nil, err
	}
	if pageTitle == "" {
		pageTitle = "無題"
	}

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

	return []slack.Block{
		slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("*<%v|%v>*", strings.ReplaceAll(pageURL, "\n", ""), pageTitle), false, false),
			nil,
			thumbnail,
		),
		slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, "*このレシピの操作*", false, false),
			nil,
			slack.NewAccessory(slack.NewOverflowBlockElement(
				b.actionOverflow,
				slack.NewOptionBlockObject(OverflowArgs{ofTypeRebuildRecipe, page.ID}.String(), slack.NewTextBlockObject(slack.PlainTextType, "レシピを再取得", false, false), nil),
				slack.NewOptionBlockObject(OverflowArgs{ofTypeUpdateIngredients, page.ID}.String(), slack.NewTextBlockObject(slack.PlainTextType, "主な材料を更新", false, false), nil),
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
