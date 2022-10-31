package delishkitchen

import (
	"context"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	rex, err := p.Parse2(ctx, url)
	if err != nil {
		return nil, err
	}
	return rex.BackCompat(), nil
}

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://delishkitchen.tv/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rex := &rexch.Recipe{
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
				igd := rexch.Ingredient{
					Group:  group,
					Name:   strings.TrimSpace(s.Find(`.ingredient-name`).Text()),
					Amount: strings.TrimSpace(s.Find(`.ingredient-serving`).Text()),
				}
				rex.Ingredients = append(rex.Ingredients, igd)
			}
		})
	})

	doc.Find(`li.step`).Each(func(i int, s *goquery.Selection) {
		ist := rexch.Instruction{}
		ist.AddText(strings.TrimSpace(s.Find(`p.step-desc`).Text()))
		s.Find(`video`).Each(func(i int, s *goquery.Selection) {
			ist.AddImage(s.AttrOr("poster", ""))
		})
		rex.Instructions = append(rex.Instructions, ist)
	})

	return rex, nil
}

func NewParser() sites.Parser2 {
	return &parser{}
}
