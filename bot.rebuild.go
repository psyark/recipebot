package recipebot

import (
	"context"

	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/sites/united"
)

func (b *MyBot) RespondRebuild(pageID string) error {
	ctx := context.Background()

	page, err := b.notion.RetrievePage(ctx, pageID)
	if err != nil {
		return err
	}

	piop, err := b.notion.RetrievePagePropertyItem(ctx, page.ID, RECIPE_ORIGINAL)
	if err != nil {
		return err
	}

	rcp, err := united.Parsers.Parse(ctx, piop.PropertyItem.URL)
	if err != nil {
		return err
	}

	opt := &notionapi.UpdatePageOptions{}
	if page.Icon == nil && rcp.GetEmoji() != "" {
		opt.Icon = &notionapi.FileOrEmoji{Type: "emoji", Emoji: rcp.GetEmoji()}
	}
	if page.Cover == nil && rcp.Image != "" {
		opt.Cover = &notionapi.File{Type: "external", External: notionapi.ExternalFileData{URL: rcp.Image}}
	}

	if opt.Icon != nil || opt.Cover != nil {
		if _, err := b.notion.UpdatePage(ctx, page.ID, opt); err != nil {
			return err
		}
	}

	return b.autoUpdateRecipePageContent(ctx, page.ID, rcp)
}
