package recipebot

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/slack-go/slack"
)

type RecipeBlocksInfo struct {
	PageID         string
	PageURL        string
	Title          string
	ImageURL       string
	ShowCategoryUI bool
}

func (b *Bot) GetRecipeBlocksInfo(ctx context.Context, pageID string) (*RecipeBlocksInfo, error) {
	info := &RecipeBlocksInfo{PageID: pageID}

	if category, err := b.notion.RetrievePagePropertyItem(ctx, pageID, RECIPE_CATEGORY); err != nil {
		return nil, err
	} else {
		info.ShowCategoryUI = category.PropertyItem.Select.Name == ""
	}

	if title, err := b.notion.RetrievePagePropertyItem(ctx, pageID, "title"); err != nil {
		return nil, err
	} else {
		for _, item := range title.PropertyItemPagination.Results {
			info.Title += item.Title.Text.Content
		}
		if info.Title == "" {
			info.Title = "無題"
		}
	}

	if page, err := b.notion.RetrievePage(ctx, pageID); err != nil {
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

// CreateRecipeBlocks はレシピページのSlack Blocksを作成します
func CreateRecipeBlocks(info *RecipeBlocksInfo) []slack.Block {
	var thumbnail *slack.Accessory
	if info.ImageURL != "" {
		thumbnail = slack.NewAccessory(slack.NewImageBlockElement(info.ImageURL, "thumbnail"))
	}

	blocks := []slack.Block{
		slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, "レシピを作ったよ", false, false), nil, nil),
		slack.NewDividerBlock(),
		slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("*<%v|%v>*", info.PageURL, info.Title), false, false),
			nil,
			thumbnail,
		),
		slack.NewDividerBlock(),
		slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, "*操作*", false, false), nil, nil),
	}

	if info.ShowCategoryUI {
		catOptions := []*slack.OptionBlockObject{}
		for _, c := range categories {
			val, _ := json.Marshal([]string{info.PageID, c})
			opt := slack.NewOptionBlockObject(string(val), slack.NewTextBlockObject(slack.PlainTextType, c, true, false), nil)
			catOptions = append(catOptions, opt)
		}
		blocks = append(blocks, slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, "このレシピの分類を設定", false, false),
			nil,
			slack.NewAccessory(slack.NewOptionsSelectBlockElement(
				slack.OptTypeStatic,
				slack.NewTextBlockObject(slack.PlainTextType, "分類", true, false),
				actionSetCategory,
				catOptions...,
			)),
		))
	}

	blocks = append(blocks, slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "このレシピを<https://www.notion.so/80cf0a5ec25c4b7489f00594362f6e3b|🍽️献立表>に追加する", false, false),
		nil,
		slack.NewAccessory(slack.NewButtonBlockElement(
			actionCreateMenu,
			info.PageID,
			slack.NewTextBlockObject(slack.PlainTextType, "献立表に追加", true, false),
		)),
	))
	return blocks
}
