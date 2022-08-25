package lettuceclub

import (
	"context"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	if !strings.HasPrefix(url, "https://www.lettuceclub.net/recipe/dish/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rcp := &recipe.Recipe{
		Title: strings.TrimSpace(doc.Find(`h1.c-heading2__title`).Text()),
		Image: sites.ResolvePath(url, doc.Find(`.p-mainvisual__item img`).AttrOr("src", "")),
	}

	doc.Find(`ul.c-textbox__ingredients li`).Each(func(i int, s *goquery.Selection) {
		parts := strings.Split(strings.TrimSpace(s.Text()), "…")
		rcp.AddIngredient("", recipe.Ingredient{Name: parts[0], Amount: parts[1]})
	})

	doc.Find(`#step ol li`).Each(func(i int, s *goquery.Selection) {
		rcp.Steps = append(rcp.Steps, recipe.Step{Text: strings.TrimSpace(s.Text())})
	})

	return rcp, nil
}

func NewParser() sites.Parser {
	return &parser{}
}
