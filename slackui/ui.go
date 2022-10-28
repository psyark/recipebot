package slackui

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/mvdan/xurls"
	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/async"
	"github.com/psyark/recipebot/core"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

const (
	// botChannelID     = "D03SNU2C80H"
	botMemberID      = "U03SCN7MYEQ"
	cookingChannelID = "C03SNSP9HNV" // #料理 チャンネル
)

type UI struct {
	coreService             *core.Service
	slackClient             *slack.Client
	rebuildRecipeButton     *asyncButton
	updateIngredientsButton *asyncButton
	buttons                 BlockActionReacters
}

func New(coreService *core.Service, slackClient *slack.Client) *UI {
	ui := &UI{
		coreService: coreService,
		slackClient: slackClient,
	}
	ui.rebuildRecipeButton = &asyncButton{ui: ui, actionID: "rebuildRecipe", label: "レシピを再構築", asyncType: async.TypeRebuildRecipe}
	ui.updateIngredientsButton = &asyncButton{ui: ui, actionID: "updateIngredients", label: "主な材料を更新", asyncType: async.TypeUpdateIngredients}
	ui.buttons = []BlockActionReacter{
		ui.rebuildRecipeButton,
		ui.updateIngredientsButton,
	}
	return ui
}

func (b *UI) handleCallbackMessage(req *http.Request, event *slackevents.MessageEvent) error {
	if req.Header.Get("X-Slack-Retry-Num") != "" {
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

func (ui *UI) handleBlockAction(req *http.Request, callback *slack.InteractionCallback, action *slack.BlockAction) error {
	ok, err := ui.buttons.React(callback, action)
	if err != nil || ok {
		return err
	}

	// if action.ActionID != ignore {
	// 	return fmt.Errorf("block action unhandled: %v", action.ActionID)
	// }
	return nil
}

// ReactMessageWithURL はURL付きのメッセージに反応します
func (s *UI) ReactMessageWithURL(event *slackevents.MessageEvent, url string) error {
	ctx := context.Background()

	// 砂時計のプレースホルダを出しておく
	_, msgTs, err := s.slackClient.PostMessage(event.Channel, slack.MsgOptionText(":hourglass_flowing_sand:", false))
	if err != nil {
		return err
	}

	// URLに対応するレシピページを探す
	page, err := s.coreService.GetRecipeByURL(ctx, url)
	if err != nil {
		return err
	}

	// レシピページがなければ作成
	if page == nil {
		page, err = s.coreService.CreateRecipe(ctx, url)
		if err != nil {
			return err
		}
	}

	// プレースホルダを更新
	return s.UpdateRecipeMessage(ctx, event.Channel, msgTs, page, nil)
}

// UpdateRecipeMessageは指定したメッセージを新たなレシピメッセージで更新します
func (ui *UI) UpdateRecipeMessage(ctx context.Context, channelID, timestamp string, page *notionapi.Page, option *RecipeMessageOption) error {
	var thumbnail slack.BlockElement

	if option == nil {
		option = &RecipeMessageOption{}
	}

	pageTitle := page.Properties.Get("title").Title.PlainText()
	if pageTitle == "" {
		pageTitle = "無題"
	}

	if page.Icon != nil {
		if emoji, ok := page.Icon.(*notionapi.Emoji); ok {
			pageTitle = emoji.Emoji + pageTitle
		}
	}
	if page.Cover != nil {
		if page.Cover.External != nil {
			thumbnail = slack.NewImageBlockElement(page.Cover.External.URL, "レシピの写真")
		} else if page.Cover.File != nil {
			thumbnail = slack.NewImageBlockElement(page.Cover.File.URL, "レシピの写真")
		}
	}

	blocks := []slack.Block{
		section(
			mrkdwn(fmt.Sprintf("*<%v|%v>*", page.URL, strings.ReplaceAll(pageTitle, "\n", " "))),
			thumbnail,
		),
		slack.NewActionBlock(
			"",
			ui.rebuildRecipeButton.Render(page.ID),
			ui.updateIngredientsButton.Render(page.ID),
		),
	}

	if strings.TrimSpace(option.AdditionalText) != "" {
		blocks = append(blocks, section(plain(strings.TrimSpace(option.AdditionalText)), nil))
	}

	_, _, _, err := ui.slackClient.UpdateMessage(channelID, timestamp, slack.MsgOptionBlocks(blocks...))
	return err
}

type RecipeMessageOption struct {
	AdditionalText string
}
