package sbfoods

import (
	"context"
	"regexp"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

var stepRegex = regexp.MustCompile(`^【[１２３４５６７８９０]+】`)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	if !strings.HasPrefix(url, "https://www.sbfoods.co.jp/recipe/detail/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rcp := &recipe.Recipe{
		Title: strings.TrimSpace(doc.Find(`h1`).Text()),
		Image: doc.Find(`.detail-img img`).AttrOr("src", ""),
	}

	groupName := ""
	doc.Find(`ul.list-ingredient li`).Each(func(i int, s *goquery.Selection) {
		if s.Find(".data").Length() == 0 {
			groupName = strings.TrimSpace(s.Find(".title").Text())
		} else {
			rcp.AddIngredient(groupName, recipe.Ingredient{
				Name:   debrand(strings.TrimSpace(s.Find(".title").Text())),
				Amount: strings.TrimSpace(s.Find(".data").Text()),
			})
		}
	})

	doc.Find(`#box-howto li`).Each(func(i int, s *goquery.Selection) {
		rcp.Steps = append(rcp.Steps, recipe.Step{Text: stepRegex.ReplaceAllString(strings.TrimSpace(s.Text()), "")})
	})

	return rcp, nil
}

func debrand(name string) string {
	return strings.TrimPrefix(name, "S&B")
}

func NewParser() sites.Parser {
	return &parser{}
}
