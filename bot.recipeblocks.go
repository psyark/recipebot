package recipebot

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/slack-go/slack"
)

func (b *Bot) getRecipeBlocks(ctx context.Context, pageID string) ([]slack.Block, error) {
	var pageTitle string
	var pageURL string
	var thumbnail *slack.Accessory
	var category string
	var categories []string

	// カテゴリーの選択肢の取得
	if db, err := b.notion.RetrieveDatabase(ctx, RECIPE_DB_ID); err != nil {
		return nil, err
	} else {
		for _, prop := range db.Properties {
			if prop.ID == RECIPE_CATEGORY {
				for _, opt := range prop.Select.Options {
					categories = append(categories, opt.Name)
				}
			}
		}
	}

	// 現在のカテゴリーの取得
	if c, err := b.notion.RetrievePagePropertyItem(ctx, pageID, RECIPE_CATEGORY); err != nil {
		return nil, err
	} else {
		category = c.PropertyItem.Select.Name
	}

	// タイトルの取得
	if t, err := b.notion.RetrievePagePropertyItem(ctx, pageID, "title"); err != nil {
		return nil, err
	} else {
		for _, item := range t.PropertyItemPagination.Results {
			pageTitle += item.Title.Text.Content
		}
		if pageTitle == "" {
			pageTitle = "無題"
		}
	}

	// ページの取得
	if page, err := b.notion.RetrievePage(ctx, pageID); err != nil {
		return nil, err
	} else {
		pageURL = page.URL
		if page.Icon != nil {
			pageTitle = page.Icon.Emoji + pageTitle
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

func (b *Bot) getRecipeBlocks_CategoryBlock(pageID string, categories []string, category string) slack.Block {
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

func (b *Bot) getRecipeBlocks_MenuBlock(pageID string) slack.Block {
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

func (b *Bot) getRecipeBlocks_RebuildBlock(pageID string) slack.Block {
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
