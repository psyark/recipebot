package notion

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/sites"
	"github.com/psyark/recipebot/sites/united"
	"golang.org/x/sync/errgroup"
)

var client *notionapi.Client

func init() {
	if err := godotenv.Load("../../test.env"); err != nil {
		panic(err)
	}

	client = notionapi.NewClient(os.Getenv("NOTION_API_KEY"))
}

func TestService(t *testing.T) {
	ctx := context.Background()
	service := New(client)

	stockMap, err := service.GetStockMap(ctx)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v件のストック\n", len(stockMap))

	opt := &notionapi.QueryDatabaseOptions{
		Filter: notionapi.CompoundFilter{
			And: []notionapi.Filter{
				notionapi.PropertyFilter{
					Property: recipe_ingredients,
					Relation: &notionapi.RelationFilterCondition{IsEmpty: true},
				},
				notionapi.PropertyFilter{
					Property: recipe_original,
					URL:      &notionapi.TextFilterCondition{IsNotEmpty: true},
				},
			},
		},
		PageSize: 100,
	}

	recipes, err := client.QueryDatabase(ctx, recipe_db_id, opt)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%v件のレシピ\n", len(recipes.Results))

	eg := errgroup.Group{}
	for i, recipePage := range recipes.Results {
		recipePage := recipePage
		i := i
		eg.Go(func() error {
			piop, err := client.RetrievePagePropertyItem(ctx, recipePage.ID, recipe_original)
			if err != nil {
				return err
			}

			rcp, err := united.Parsers.Parse(ctx, piop.PropertyItem.URL)
			if errors.Is(err, sites.ErrUnsupportedURL) {
				fmt.Printf("%2d, unsupported: %v\n", i, piop.PropertyItem.URL)
				return nil
			} else if err != nil {
				return fmt.Errorf("%v: %w", piop.PropertyItem.URL, err)
			}

			title, err := service.GetRecipeTitle(ctx, recipePage.ID)
			if err != nil {
				return err
			}

			stockFound := []string{}
			stockNotFound := []string{}
			stockRelation := []notionapi.PageReference{}

			for _, g := range rcp.IngredientGroups {
				for _, idg := range g.Children {
					if id, ok := stockMap[idg.Name]; ok {
						stockFound = append(stockFound, idg.Name)
						stockRelation = append(stockRelation, notionapi.PageReference{ID: id})
					} else {
						stockNotFound = append(stockNotFound, idg.Name)
					}
				}
			}

			if len(stockRelation) != 0 {
				opt := &notionapi.UpdatePageOptions{
					Properties: map[string]notionapi.PropertyValue{
						recipe_ingredients: {Type: "relation", Relation: stockRelation},
					},
				}
				if _, err := client.UpdatePage(ctx, recipePage.ID, opt); err != nil {
					return err
				}

				fmt.Printf("%2d %v %v の材料 %v を設定しました (%v は見つかりませんでした）\n", i, recipePage.URL, title, stockFound, stockNotFound)
			} else if len(stockNotFound) != 0 {
				fmt.Printf("%2d %v %v の材料は一つも見つかりませんでした (%v）\n", i, recipePage.URL, title, stockNotFound)
			} else {
				fmt.Printf("%2d %v %v の材料は設定されていません\n", i, recipePage.URL, title)
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
