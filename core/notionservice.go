package core

import (
	"context"
	"fmt"
	"regexp"

	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites/united"
)

const (
	recipe_db_id            = "ff24a40498c94ac3ac2fa8894ac0d489"
	recipe_original         = "%5CiX%60"
	recipe_eval             = "Ha%3Ba"
	recipe_category         = "gmv%3A"
	recipe_ingredients      = "%5C~%7C%40"
	recipe_shared_header_id = "60a4999c-b1fa-4e3d-9d6b-48034ad7b675"
	stock_db_id             = "923bfcb7c9014273b417ddc966fd17b8"
	stock_regex             = "jpb%3D"
	stock_nolink            = "xy_~"
)

var unitedParser = united.NewParser()

type Service struct {
	client *notionapi.Client
}

func New(client *notionapi.Client) *Service {
	return &Service{client: client}
}

// GetRecipeCategories は レシピ データベースからカテゴリーの選択肢を取得して返します
func (s *Service) GetRecipeCategories(ctx context.Context) ([]string, error) {
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
func (s *Service) GetRecipeCategory(ctx context.Context, pageID string) (string, error) {
	// 現在のカテゴリーの取得
	if piop, err := s.client.RetrievePagePropertyItem(ctx, pageID, recipe_category); err != nil {
		return "", err
	} else {
		if option := piop.(*notionapi.PropertyItem).Select; option != nil {
			return option.Name, nil
		} else {
			return "", nil
		}
	}
}

// GetRecipeTitle は レシピ データベースのページからタイトルを取得して返します
// func (s *Service) GetRecipeTitle(ctx context.Context, pageID string) (string, error) {
// 	// タイトルの取得
// 	if piop, err := s.client.RetrievePagePropertyItem(ctx, pageID, "title"); err != nil {
// 		return "", err
// 	} else {
// 		title := ""
// 		for _, item := range piop.(*notionapi.PropertyItemPagination).Results {
// 			title += item.Title.Text.Content
// 		}
// 		return title, nil
// 	}
// }

// RetrievePage は単純な RetrievePage APIの呼び出しです
func (s *Service) RetrievePage(ctx context.Context, pageID string) (*notionapi.Page, error) {
	return s.client.RetrievePage(ctx, pageID)
}

func (s *Service) GetRecipeByURL(ctx context.Context, url string) (*notionapi.Page, error) {
	opt := &notionapi.QueryDatabaseOptions{Filter: &notionapi.PropertyFilter{
		Property: recipe_original,
		URL:      &notionapi.TextFilterCondition{Equals: &url},
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

func (s *Service) CreateRecipe(ctx context.Context, url string) (*notionapi.Page, error) {
	rcp, err := unitedParser.Parse(ctx, url)
	if err != nil {
		return nil, err
	}

	opt := &notionapi.CreatePageOptions{
		Parent: &notionapi.Parent{Type: "database_id", DatabaseID: recipe_db_id},
		Properties: map[string]notionapi.PropertyValue{
			"title":         {Type: "title", Title: toRichTextArray(rcp.Title)},
			recipe_original: {Type: "url", URL: &url},
			recipe_eval:     {Type: "select", Select: &notionapi.SelectOption{Name: "👀次作る"}},
		},
	}
	if rcp.GetEmoji() != "" {
		opt.Icon = &notionapi.Emoji{Type: "emoji", Emoji: rcp.GetEmoji()}
	}
	if rcp.Image != "" {
		opt.Cover = &notionapi.File{Type: "external", External: &notionapi.ExternalFileData{URL: rcp.Image}}
	}

	page, err := s.client.CreatePage(ctx, opt)
	if err != nil {
		return nil, err
	}

	return page, s.updatePageContent(ctx, page.ID, rcp)
}

func (s *Service) UpdateRecipe(ctx context.Context, pageID string) error {
	page, err := s.client.RetrievePage(ctx, pageID)
	if err != nil {
		return fmt.Errorf("recipeService.client.RetrievePage: %w", err)
	}

	title := ""
	url := ""
	for _, pv := range page.Properties {
		switch pv.ID {
		case "title":
			for _, t := range pv.Title {
				title += t.PlainText
			}
		case recipe_original:
			if pv.URL == nil {
				return fmt.Errorf("url unset")
			}
			url = *pv.URL
		}
	}

	rcp, err := unitedParser.Parse(ctx, url)
	if err != nil {
		return err
	}

	opt := &notionapi.UpdatePageOptions{}
	if title == "" && rcp.Title != "" {
		if opt.Properties == nil {
			opt.Properties = map[string]notionapi.PropertyValue{}
		}
		opt.Properties["title"] = notionapi.PropertyValue{Type: "title", Title: toRichTextArray(rcp.Title)}
	}
	if page.Icon == nil && rcp.GetEmoji() != "" {
		opt.Icon = &notionapi.Emoji{Type: "emoji", Emoji: rcp.GetEmoji()}
	}
	if page.Cover == nil && rcp.Image != "" {
		opt.Cover = &notionapi.File{Type: "external", External: &notionapi.ExternalFileData{URL: rcp.Image}}
	}

	if opt.Icon != nil || opt.Cover != nil || len(opt.Properties) != 0 {
		if _, err := s.client.UpdatePage(ctx, page.ID, opt); err != nil {
			return err
		}
	}

	return s.updatePageContent(ctx, page.ID, rcp)
}

func (s *Service) UpdateRecipeIngredients(ctx context.Context, pageID string, stockMap map[*regexp.Regexp]string) (map[string]bool, error) {
	piop, err := s.client.RetrievePagePropertyItem(ctx, pageID, recipe_original)
	if err != nil {
		return nil, err
	}

	rcp, err := unitedParser.Parse(ctx, piop.(*notionapi.PropertyItem).URL)
	if err != nil {
		return nil, err
	}

	stockRelation := []notionapi.PageReference{}
	foundMap := map[string]bool{}
	for _, g := range rcp.IngredientGroups {
		for _, igd := range g.Children {
			found := false
			for regex, pageID := range stockMap {
				if regex.MatchString(igd.Name) {
					found = true
					if pageID != "" { // リンクしない
						foundMap[igd.Name] = true
						stockRelation = append(stockRelation, notionapi.PageReference{ID: pageID})
					}
					break
				}
			}
			if !found {
				foundMap[igd.Name] = false
			}
		}
	}

	if len(stockRelation) != 0 {
		opt := &notionapi.UpdatePageOptions{
			Properties: map[string]notionapi.PropertyValue{recipe_ingredients: {Type: "relation", Relation: stockRelation}},
		}
		if _, err := s.client.UpdatePage(ctx, pageID, opt); err != nil {
			return nil, err
		}
	}

	return foundMap, nil
}

func (s *Service) updatePageContent(ctx context.Context, pageID string, rcp *recipe.Recipe) error {
	// 以前のブロックを削除
	pagi, err := s.client.RetrieveBlockChildren(ctx, pageID)
	if err != nil {
		return fmt.Errorf("retrieveBlockChildren: %w", err)
	}

	if pagi.HasMore {
		return fmt.Errorf("updatePageContent: Not implemented")
	}

	for _, block := range pagi.Results {
		_, err := s.client.DeleteBlock(ctx, block.ID)
		if err != nil {
			return fmt.Errorf("deleteBlock: %w", err)
		}
	}

	// 新しいブロックを作成
	opt := &notionapi.AppendBlockChildrenOptions{Children: toBlocks(*rcp)}
	if _, err = s.client.AppendBlockChildren(ctx, pageID, opt); err != nil {
		return fmt.Errorf("appendBlockChildren: %w", err)
	}
	return nil
}

func (s *Service) SetRecipeCategory(ctx context.Context, pageID string, category string) error {
	_, err := s.client.UpdatePage(ctx, pageID, &notionapi.UpdatePageOptions{
		Properties: map[string]notionapi.PropertyValue{
			recipe_category: {Type: "select", Select: &notionapi.SelectOption{Name: category}},
		},
	})
	if err != nil {
		return fmt.Errorf("recipeService.client.UpdatePage: %w", err)
	}
	return nil
}

// GetStockMap は 食材ストック の材料名からIDを引くマップを返します
func (s *Service) GetStockMap(ctx context.Context) (map[*regexp.Regexp]string, error) {
	opt := &notionapi.QueryDatabaseOptions{
		PageSize: 200,
	}

	pagi, err := s.client.QueryDatabase(ctx, stock_db_id, opt)
	if err != nil {
		return nil, err
	}

	stockMap := map[*regexp.Regexp]string{}

	for _, page := range pagi.Results {
		title := page.Properties.Get("title").Title.PlainText()

		regexStr := page.Properties.Get(stock_regex).RichText.PlainText()
		if regexStr == "" {
			regexStr = fmt.Sprintf("^%v$", title)
		}
		regex, err := regexp.Compile(regexStr)
		if err != nil {
			return nil, err
		}

		// リンクしない
		if page.Properties.Get(stock_nolink).Checkbox {
			stockMap[regex] = ""
		} else {
			stockMap[regex] = page.ID
		}
	}

	return stockMap, nil
}
