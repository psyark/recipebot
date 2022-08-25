package kurashiru

import (
	"context"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	if !strings.HasPrefix(url, "https://www.kurashiru.com/recipes/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rcp := &recipe.Recipe{
		Title: strings.TrimSuffix(doc.Find(`h1.title`).Text(), "　レシピ・作り方"),
		Image: doc.Find(`.main-video video`).AttrOr("poster", ""),
	}

	groupName := ""
	doc.Find(`li.ingredient-list-item`).Each(func(i int, s *goquery.Selection) {
		switch s.AttrOr("class", "") {
		case "ingredient-list-item group-title":
			groupName = strings.TrimSpace(s.Find(".ingredient-title").Text())
		case "ingredient-list-item":
			groupName = ""
			fallthrough
		case "ingredient-list-item group-item":
			rcp.AddIngredient(groupName, recipe.Ingredient{
				Name:   strings.TrimSpace(s.Find(`.ingredient-name`).Text()),
				Amount: strings.TrimSpace(s.Find(`.ingredient-quantity-amount`).Text()),
			})
		}
	})

	doc.Find(`li.instruction-list-item`).Each(func(i int, s *goquery.Selection) {
		rcp.Steps = append(rcp.Steps, recipe.Step{Text: strings.TrimSpace(s.Find(`.content`).Text())})
	})

	return rcp, nil
}

func NewParser() sites.Parser {
	return &parser{}
}
