package recipebot

import (
	"context"
	"fmt"
	"strings"

	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites/united"
)

type recipeService struct {
	client *notionapi.Client
}

func (s recipeService) GetRecipeByURL(ctx context.Context, url string) (*notionapi.Page, error) {
	opt := &notionapi.QueryDatabaseOptions{Filter: notionapi.PropertyFilter{
		Property: RECIPE_ORIGINAL,
		URL:      &notionapi.TextFilterCondition{Equals: url},
	}}

	pagi, err := s.client.QueryDatabase(ctx, RECIPE_DB_ID, opt)
	if err != nil {
		return nil, err
	}

	if len(pagi.Results) != 0 {
		return &pagi.Results[0], nil
	}
	return nil, nil
}

func (s recipeService) CreateRecipe(ctx context.Context, url string) (*notionapi.Page, error) {
	rcp, err := united.Parsers.Parse(ctx, url)
	if err != nil {
		return nil, err
	}

	opt := &notionapi.CreatePageOptions{
		Parent: &notionapi.Parent{Type: "database_id", DatabaseID: RECIPE_DB_ID},
		Properties: map[string]notionapi.PropertyValue{
			"title":         {Type: "title", Title: toRichTextArray(rcp.Title)},
			RECIPE_ORIGINAL: {Type: "url", URL: url},
			RECIPE_EVAL:     {Type: "select", Select: notionapi.SelectPropertyValueData{Name: "👀次作る"}},
		},
	}
	if rcp.GetEmoji() != "" {
		opt.Icon = &notionapi.FileOrEmoji{Type: "emoji", Emoji: rcp.GetEmoji()}
	}
	if rcp.Image != "" {
		opt.Cover = &notionapi.File{Type: "external", External: notionapi.ExternalFileData{URL: rcp.Image}}
	}

	page, err := s.client.CreatePage(ctx, opt)
	if err != nil {
		return nil, err
	}

	return page, s.updatePageContent(ctx, page.ID, rcp)
}

func (s recipeService) UpdateRecipe(ctx context.Context, pageID string) error {
	page, err := s.client.RetrievePage(ctx, pageID)
	if err != nil {
		return fmt.Errorf("recipeService.client.RetrievePage: %w", err)
	}

	piop, err := s.client.RetrievePagePropertyItem(ctx, page.ID, RECIPE_ORIGINAL)
	if err != nil {
		return fmt.Errorf("recipeService.client.RetrievePagePropertyItem: %w", err)
	}

	rcp, err := united.Parsers.Parse(ctx, piop.PropertyItem.URL)
	if err != nil {
		return fmt.Errorf("united.Parsers.Parse: %w", err)
	}

	opt := &notionapi.UpdatePageOptions{}
	if page.Icon == nil && rcp.GetEmoji() != "" {
		opt.Icon = &notionapi.FileOrEmoji{Type: "emoji", Emoji: rcp.GetEmoji()}
	}
	if page.Cover == nil && rcp.Image != "" {
		opt.Cover = &notionapi.File{Type: "external", External: notionapi.ExternalFileData{URL: rcp.Image}}
	}

	if opt.Icon != nil || opt.Cover != nil {
		if _, err := s.client.UpdatePage(ctx, page.ID, opt); err != nil {
			return fmt.Errorf("recipeService.client.UpdatePage: %w", err)
		}
	}

	return s.updatePageContent(ctx, page.ID, rcp)
}

func (s recipeService) updatePageContent(ctx context.Context, pageID string, rcp *recipe.Recipe) error {
	// 以前のブロックを削除
	pagi, err := s.client.RetrieveBlockChildren(ctx, pageID)
	if err != nil {
		return fmt.Errorf("recipeService.client.RetrieveBlockChildren: %w", err)
	}

	if pagi.HasMore {
		return fmt.Errorf("updatePageContent: Not implemented")
	}

	for _, block := range pagi.Results {
		_, err := s.client.DeleteBlock(ctx, block.ID)
		if err != nil {
			return fmt.Errorf("recipeService.client.DeleteBlock: %w", err)
		}
	}

	// 新しいブロックを作成
	opt := &notionapi.AppendBlockChildrenOptions{Children: Recipe(*rcp).NotionBlocks()}
	if _, err = s.client.AppendBlockChildren(ctx, pageID, opt); err != nil {
		return fmt.Errorf("recipeService.client.AppendBlockChildren: %w", err)
	}
	return nil
}

func (s recipeService) SetRecipeCategory(ctx context.Context, pageID string, category string) error {
	_, err := s.client.UpdatePage(ctx, pageID, &notionapi.UpdatePageOptions{
		Properties: map[string]notionapi.PropertyValue{
			RECIPE_CATEGORY: {Type: "select", Select: notionapi.SelectPropertyValueData{Name: category}},
		},
	})
	if err != nil {
		return fmt.Errorf("recipeService.client.UpdatePage: %w", err)
	}
	return nil
}

type Recipe recipe.Recipe

func (rcp Recipe) NotionBlocks() []notionapi.Block {
	indices := []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "🔟", "🔢"}

	blocks := []notionapi.Block{
		{
			Object: "block",
			Type:   "synced_block",
			SyncedBlock: &notionapi.SyncedBlockBlocks{
				SyncedFrom: &notionapi.SyncedFrom{Type: "block_id", BlockID: RECIPE_HEADER_ID},
			},
		},
		rcp.toHeading1("材料"),
	}
	for _, group := range rcp.IngredientGroups {
		if group.Name != "" {
			blocks = append(blocks, rcp.toHeading3(group.Name))
		}

		width := group.LongestNameWidth() + 1
		for _, idg := range group.Children {
			todo := rcp.toToDo(idg.Name + strings.Repeat("　", width-idg.NameWidth()) + idg.Amount)
			if idg.Comment != "" {
				comment := toRichTextArray(" （" + idg.Comment + "）")
				comment[0].Annotations = &notionapi.Annotations{Color: "green"}
				todo.ToDo.RichText = append(todo.ToDo.RichText, comment...)
			}
			blocks = append(blocks, todo)
		}
	}
	blocks = append(blocks, rcp.toHeading1("手順"))
	for idx, stp := range rcp.Steps {
		emoji := "🔢"
		if idx < len(indices) {
			emoji = indices[idx]
		}
		blocks = append(blocks, rcp.toCallout(stp.Text, emoji, stp.Images))
	}
	return blocks
}

func (rcp Recipe) toHeading1(str string) notionapi.Block {
	return notionapi.Block{
		Object:   "block",
		Type:     "heading_1",
		Heading1: notionapi.HeadingBlockData{RichText: toRichTextArray(str), Color: "default"},
	}
}

func (rcp Recipe) toHeading3(str string) notionapi.Block {
	return notionapi.Block{
		Object:   "block",
		Type:     "heading_3",
		Heading3: notionapi.HeadingBlockData{RichText: toRichTextArray(str), Color: "default"},
	}
}

func (rcp Recipe) toToDo(str string) notionapi.Block {
	return notionapi.Block{
		Object: "block",
		Type:   "to_do",
		ToDo:   notionapi.ToDoBlockData{RichText: toRichTextArray(str), Color: "default"},
	}
}

func (rcp Recipe) toCallout(str string, emoji string, images []string) notionapi.Block {
	block := notionapi.Block{
		Object: "block",
		Type:   "callout",
		Callout: notionapi.CalloutBlockData{
			RichText: toRichTextArray(str),
			Icon:     &notionapi.FileOrEmoji{Type: "emoji", Emoji: emoji},
			Color:    "gray_background",
		},
	}
	for _, url := range images {
		block.Callout.Children = append(block.Callout.Children, rcp.toImage(url))
	}
	return block
}

func (rcp Recipe) toImage(url string) notionapi.Block {
	return notionapi.Block{
		Object: "block",
		Type:   "image",
		Image: notionapi.File{
			Type:     "external",
			External: notionapi.ExternalFileData{URL: url},
		},
	}
}

func toRichTextArray(text string) []notionapi.RichText {
	return []notionapi.RichText{{Type: "text", Text: notionapi.Text{Content: text}}}
}
