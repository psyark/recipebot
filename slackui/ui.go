package slackui

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/mvdan/xurls"
	"github.com/psyark/notionapi"
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
	rebuildRecipeButton     *rebuildRecipeButton
	updateIngredientsButton *updateIngredientsButton
	buttons                 BlockActionReacters
	modals                  ViewSubmissionReacters
}

func New(coreService *core.Service, slackClient *slack.Client) *UI {
	ui := &UI{
		coreService: coreService,
		slackClient: slackClient,
	}
	ui.rebuildRecipeButton = &rebuildRecipeButton{ui}
	ui.updateIngredientsButton = &updateIngredientsButton{ui}
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

func (ui *UI) handleViewSubmission(req *http.Request, callback *slack.InteractionCallback) (*slack.ViewSubmissionResponse, error) {
	ok, resp, err := ui.modals.React(callback)
	if err != nil || ok {
		return resp, err
	}

	// if callback.View.CallbackID != ignore {
	// 	return nil, fmt.Errorf("view submission unhandled: %#v", callback.View.CallbackID)
	// }
	return nil, nil
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
	blocks, err := s.getRecipeBlocks(ctx, page)
	if err != nil {
		return err
	}

	_, _, _, err = s.slackClient.UpdateMessage(event.Channel, msgTs, slack.MsgOptionBlocks(blocks...))
	return err
}

func (b *UI) getRecipeBlocks(ctx context.Context, page *notionapi.Page) ([]slack.Block, error) {
	var pageURL string
	var thumbnail *slack.Accessory

	// タイトルの取得
	pageTitle, err := b.coreService.GetRecipeTitle(ctx, page.ID)
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
			slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("*<%v|%v>*", pageURL, strings.ReplaceAll(pageTitle, "\n", " ")), false, false),
			nil,
			thumbnail,
		),
		slack.NewActionBlock(
			"",
			b.rebuildRecipeButton.Render(page.ID),
			b.updateIngredientsButton.Render(page.ID),
		),
	}, nil
}
