package kurashiru

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
	if !strings.HasPrefix(url, "https://www.kurashiru.com/recipes/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rex := &rexch.Recipe{
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
			idg := rexch.NewIngredient(s.Find(`.ingredient-name`).Text(), s.Find(`.ingredient-quantity-amount`).Text())
			groupName2 := groupName

			for _, prefix := range []string{"（Ａ）", "（Ｂ）", "（Ｃ）"} {
				if strings.HasPrefix(idg.Name, prefix) {
					idg.Name = strings.TrimPrefix(idg.Name, prefix)
					groupName2 = prefix
					break
				}
			}

			idg.Group = groupName2

			rex.Ingredients = append(rex.Ingredients, *idg)
		}
	})

	doc.Find(`li.instruction-list-item`).Each(func(i int, s *goquery.Selection) {
		ist := rexch.Instruction{}
		ist.AddText(strings.TrimSpace(s.Find(`.content`).Text()))
		rex.Instructions = append(rex.Instructions, ist)
	})

	return rex, nil
}

func NewParser() sites.Parser2 {
	return &parser{}
}
