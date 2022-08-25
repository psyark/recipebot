package recipebot

import (
	"context"
	"fmt"

	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites/united"
)

// BotのNotion側の処理を書く

const (
	RECIPE_DB_ID    = "ff24a40498c94ac3ac2fa8894ac0d489"
	RECIPE_ORIGINAL = "%5CiX%60"
	RECIPE_EVAL     = "Ha%3Ba"
	RECIPE_CATEGORY = "gmv%3A"
)

func (b *Bot) autoUpdateRecipePage(ctx context.Context, recipeURL string) (*notionapi.Page, error) {
	rcp, err := united.Parsers.Parse(ctx, recipeURL)
	if err != nil {
		return nil, err
	}

	opt := &notionapi.QueryDatabaseOptions{Filter: notionapi.PropertyFilter{
		Property: RECIPE_ORIGINAL,
		URL:      &notionapi.TextFilterCondition{Equals: recipeURL},
	}}

	pagi, err := b.notion.QueryDatabase(ctx, RECIPE_DB_ID, opt)
	if err != nil {
		return nil, err
	}

	var page notionapi.Page

	if len(pagi.Results) == 0 {
		opt := &notionapi.CreatePageOptions{
			Parent: &notionapi.Parent{Type: "database_id", DatabaseID: RECIPE_DB_ID},
			Properties: map[string]notionapi.PropertyValue{
				"title":         {Type: "title", Title: toRichTextArray(rcp.Title)},
				RECIPE_ORIGINAL: {Type: "url", URL: recipeURL},
				RECIPE_EVAL:     {Type: "select", Select: notionapi.SelectPropertyValueData{Name: "👀次作る"}},
			},
		}
		if rcp.GetEmoji() != "" {
			opt.Icon = &notionapi.FileOrEmoji{Type: "emoji", Emoji: rcp.GetEmoji()}
		}
		if rcp.Image != "" {
			opt.Cover = &notionapi.File{Type: "external", External: notionapi.ExternalFileData{URL: rcp.Image}}
		}

		if page2, err := b.notion.CreatePage(ctx, opt); err != nil {
			return nil, err
		} else {
			page = *page2
		}
	} else {
		page = pagi.Results[0]

		opt := &notionapi.UpdatePageOptions{}
		if page.Icon == nil && rcp.GetEmoji() != "" {
			opt.Icon = &notionapi.FileOrEmoji{Type: "emoji", Emoji: rcp.GetEmoji()}
		}
		if page.Cover == nil && rcp.Image != "" {
			opt.Cover = &notionapi.File{Type: "external", External: notionapi.ExternalFileData{URL: rcp.Image}}
		}

		if opt.Icon != nil || opt.Cover != nil {
			if _, err := b.notion.UpdatePage(ctx, page.ID, opt); err != nil {
				return nil, err
			}
		}
	}

	return &page, b.autoUpdateRecipePageContent(ctx, page.ID, rcp)
}

func (b *Bot) autoUpdateRecipePageContent(ctx context.Context, pageID string, rcp *recipe.Recipe) error {
	// 以前のブロックを削除
	pagi, err := b.notion.RetrieveBlockChildren(ctx, pageID)
	if err != nil {
		return err
	}

	if pagi.HasMore {
		return fmt.Errorf("autoUpdateRecipePageContent: Not implemented")
	}

	for _, block := range pagi.Results {
		_, err := b.notion.DeleteBlock(ctx, block.ID)
		if err != nil {
			return err
		}
	}

	opt := &notionapi.AppendBlockChildrenOptions{Children: ToNotionBlocks(rcp)}
	_, err = b.notion.AppendBlockChildren(ctx, pageID, opt)
	return err
}

func (b *Bot) updateCategory(ctx context.Context, pageID string, category string) (*notionapi.Page, error) {
	return b.notion.UpdatePage(ctx, pageID, &notionapi.UpdatePageOptions{
		Properties: map[string]notionapi.PropertyValue{
			"分類": {Type: "select", Select: notionapi.SelectPropertyValueData{Name: category}},
		},
	})
}
