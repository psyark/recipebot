package lettuceclub

import (
	"context"
	"strings"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

type parser struct{}

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://www.lettuceclub.net/recipe/dish/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rex := &rexch.Recipe{
		Title: strings.TrimSpace(doc.Find(`h1.c-heading2__title`).Text()),
		Image: sites.ResolvePath(url, doc.Find(`.p-mainvisual__item img`).AttrOr("src", "")),
	}

	doc.Find(`ul.c-textbox__ingredients li`).Each(func(i int, s *goquery.Selection) {
		parts := strings.Split(strings.TrimSpace(s.Text()), "…")
		igd := rexch.Ingredient{Name: parts[0], Amount: parts[1]}
		rex.Ingredients = append(rex.Ingredients, igd)
	})

	doc.Find(`#step ol li`).Each(func(i int, s *goquery.Selection) {
		ist := rexch.Instruction{}
		ist.AddText(strings.TrimSpace(s.Text()))
		rex.Instructions = append(rex.Instructions, ist)
	})

	return rex, nil
}

func NewParser() sites.Parser2 {
	return &parser{}
}
