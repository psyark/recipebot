package delishkitchen

import (
	"context"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	if !strings.HasPrefix(url, "https://delishkitchen.tv/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rcp := &recipe.Recipe{
		Title: strings.TrimSpace(doc.Find(`h1.title`).Text()),
		Image: doc.Find(`div.delish-main-player video`).AttrOr("poster", ""),
	}

	doc.Find(`ul.ingredient-list`).Each(func(i int, s *goquery.Selection) {
		group := ""
		s.Find(`li`).Each(func(i int, s *goquery.Selection) {
			switch s.AttrOr("class", "") {
			case "ingredient-group__header":
				group = strings.TrimSpace(s.Text())
			case "ingredient":
				rcp.AddIngredient(group, recipe.Ingredient{
					Name:   strings.TrimSpace(s.Find(`.ingredient-name`).Text()),
					Amount: strings.TrimSpace(s.Find(`.ingredient-serving`).Text()),
				})
			}
		})
	})

	doc.Find(`li.step`).Each(func(i int, s *goquery.Selection) {
		rcp.Steps = append(rcp.Steps, recipe.Step{
			Text: strings.TrimSpace(s.Find(`p.step-desc`).Text()),
			Images: s.Find(`video`).Map(func(i int, s *goquery.Selection) string {
				return s.AttrOr("poster", "")
			}),
		})
	})

	return rcp, nil
}

func NewParser() sites.Parser {
	return &parser{}
}
