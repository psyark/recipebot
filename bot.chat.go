package recipebot

import (
	"encoding/json"
	"fmt"

	"github.com/slack-go/slack"
)

type RecipeBlocksInfo struct {
	PageID     string
	PageURL    string
	Title      string
	ImageURL   string
	Category   string
	Categories []string
}

// ToSlackBlocks はレシピページのSlack Blocksを作成します
func (info *RecipeBlocksInfo) ToSlackBlocks(actionSetCategory, actionCreateMenu, actionRebuild string) []slack.Block {
	return []slack.Block{
		slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, "レシピを作ったよ", false, false), nil, nil),
		slack.NewDividerBlock(),
		slack.NewSectionBlock(
			slack.NewTextBlockObject(slack.MarkdownType, fmt.Sprintf("*<%v|%v>*", info.PageURL, info.Title), false, false),
			nil,
			info.getThumbnail(),
		),
		slack.NewDividerBlock(),
		slack.NewSectionBlock(slack.NewTextBlockObject(slack.MarkdownType, "*このレシピの操作*", false, false), nil, nil),
		info.getCategoryBlock(actionSetCategory),
		info.getMenuBlock(actionCreateMenu),
		info.getRebuildBlock(actionRebuild),
	}
}

func (info *RecipeBlocksInfo) getThumbnail() *slack.Accessory {
	if info.ImageURL != "" {
		return slack.NewAccessory(slack.NewImageBlockElement(info.ImageURL, "thumbnail"))
	}
	return nil
}

func (info *RecipeBlocksInfo) getCategoryBlock(actionSetCategory string) slack.Block {
	catOptions := []*slack.OptionBlockObject{}
	var initialOption *slack.OptionBlockObject
	for _, c := range info.Categories {
		val, _ := json.Marshal([]string{info.PageID, c})
		opt := slack.NewOptionBlockObject(string(val), slack.NewTextBlockObject(slack.PlainTextType, c, true, false), nil)
		catOptions = append(catOptions, opt)
		if c == info.Category {
			initialOption = opt
		}
	}
	selectBlock := slack.NewOptionsSelectBlockElement(
		slack.OptTypeStatic,
		slack.NewTextBlockObject(slack.PlainTextType, "分類", true, false),
		actionSetCategory,
		catOptions...,
	)
	if initialOption != nil {
		selectBlock.InitialOption = initialOption
	}
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "分類を設定する", false, false),
		nil,
		slack.NewAccessory(selectBlock),
	)
}

func (info *RecipeBlocksInfo) getMenuBlock(actionCreateMenu string) slack.Block {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "<https://www.notion.so/80cf0a5ec25c4b7489f00594362f6e3b|🍽️献立表>に追加する", false, false),
		nil,
		slack.NewAccessory(slack.NewButtonBlockElement(
			actionCreateMenu,
			info.PageID,
			slack.NewTextBlockObject(slack.PlainTextType, "献立表に追加", true, false),
		)),
	)
}

func (info *RecipeBlocksInfo) getRebuildBlock(actionRebuild string) slack.Block {
	return slack.NewSectionBlock(
		slack.NewTextBlockObject(slack.MarkdownType, "再取得して作り直す", false, false),
		nil,
		slack.NewAccessory(slack.NewButtonBlockElement(
			actionRebuild,
			info.PageID,
			slack.NewTextBlockObject(slack.PlainTextType, "作り直す", true, false),
		)),
	)
}
