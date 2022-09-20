package notion

import (
	"context"
	"fmt"
	"strings"

	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites/united"
)

const (
	recipe_db_id            = "ff24a40498c94ac3ac2fa8894ac0d489"
	recipe_original         = "%5CiX%60"
	recipe_eval             = "Ha%3Ba"
	recipe_category         = "gmv%3A"
	recipe_shared_header_id = "60a4999c-b1fa-4e3d-9d6b-48034ad7b675"
)

type Service struct {
	client *notionapi.Client
}

func New(client *notionapi.Client) *Service {
	return &Service{client: client}
}

// GetRecipeCategories は レシピ データベースからカテゴリーの選択肢を取得して返します
func (s Service) GetRecipeCategories(ctx context.Context) ([]string, error) {
	categories := []string{}
	if db, err := s.client.RetrieveDatabase(ctx, recipe_db_id); err != nil {
		return nil, err
	} else {
		for _, prop := range db.Properties {
			if prop.ID == recipe_category {
				for _, opt := range prop.Select.Options {
					categories = append(categories, opt.Name)
				}
			}
		}
	}
	return categories, nil
}

// GetRecipeCategory は レシピ データベースのページからカテゴリーを取得して返します
func (s Service) GetRecipeCategory(ctx context.Context, pageID string) (string, error) {
	// 現在のカテゴリーの取得
	if piop, err := s.client.RetrievePagePropertyItem(ctx, pageID, recipe_category); err != nil {
		return "", err
	} else {
		return piop.PropertyItem.Select.Name, nil
	}
}

// GetRecipeTitle は レシピ データベースのページからタイトルを取得して返します
func (s Service) GetRecipeTitle(ctx context.Context, pageID string) (string, error) {
	// タイトルの取得
	if piop, err := s.client.RetrievePagePropertyItem(ctx, pageID, "title"); err != nil {
		return "", err
	} else {
		title := ""
		for _, item := range piop.PropertyItemPagination.Results {
			title += item.Title.Text.Content
		}
		return title, nil
	}
}

// RetrievePage は単純な RetrievePage APIの呼び出しです
func (s Service) RetrievePage(ctx context.Context, pageID string) (*notionapi.Page, error) {
	return s.client.RetrievePage(ctx, pageID)
}

func (s Service) GetRecipeByURL(ctx context.Context, url string) (*notionapi.Page, error) {
	opt := &notionapi.QueryDatabaseOptions{Filter: notionapi.PropertyFilter{
		Property: recipe_original,
		URL:      &notionapi.TextFilterCondition{Equals: url},
	}}

	pagi, err := s.client.QueryDatabase(ctx, recipe_db_id, opt)
	if err != nil {
		return nil, err
	}

	if len(pagi.Results) != 0 {
		return &pagi.Results[0], nil
	}
	return nil, nil
}

func (s Service) CreateRecipe(ctx context.Context, url string) (*notionapi.Page, error) {
	rcp, err := united.Parsers.Parse(ctx, url)
	if err != nil {
		return nil, err
	}

	opt := &notionapi.CreatePageOptions{
		Parent: &notionapi.Parent{Type: "database_id", DatabaseID: recipe_db_id},
		Properties: map[string]notionapi.PropertyValue{
			"title":         {Type: "title", Title: toRichTextArray(rcp.Title)},
			recipe_original: {Type: "url", URL: url},
			recipe_eval:     {Type: "select", Select: notionapi.SelectPropertyValueData{Name: "👀次作る"}},
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

func (s Service) UpdateRecipe(ctx context.Context, pageID string) error {
	page, err := s.client.RetrievePage(ctx, pageID)
	if err != nil {
		return fmt.Errorf("recipeService.client.RetrievePage: %w", err)
	}

	piop, err := s.client.RetrievePagePropertyItem(ctx, page.ID, recipe_original)
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

func (s Service) updatePageContent(ctx context.Context, pageID string, rcp *recipe.Recipe) error {
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

func (s Service) SetRecipeCategory(ctx context.Context, pageID string, category string) error {
	_, err := s.client.UpdatePage(ctx, pageID, &notionapi.UpdatePageOptions{
		Properties: map[string]notionapi.PropertyValue{
			recipe_category: {Type: "select", Select: notionapi.SelectPropertyValueData{Name: category}},
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
				SyncedFrom: &notionapi.SyncedFrom{Type: "block_id", BlockID: recipe_shared_header_id},
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
