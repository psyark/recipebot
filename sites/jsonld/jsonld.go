package jsonld

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/psyark/jsonld"
	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"
)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	var rcp *recipe.Recipe

	doc.Find(`script[type="application/ld+json"]`).EachWithBreak(func(i int, s *goquery.Selection) bool {
		jsonStr := s.Text()

		obj, err := jsonld.DecodeObject([]byte(jsonStr))
		if err != nil {
			// パースエラーに続いて有効なレシピが得られる場合があるので無視
			return true
		}

		if ldRcp, ok := obj.(*jsonld.Recipe); ok {
			rcp = &recipe.Recipe{}

			for _, text := range ldRcp.Name {
				if text, ok := text.(string); ok {
					rcp.Title = text
				}
			}
			for _, text := range ldRcp.Image {
				if text, ok := text.(string); ok {
					rcp.Image = text
				}
			}
			for _, text := range ldRcp.RecipeIngredient {
				if text, ok := text.(string); ok {
					fields := strings.SplitN(text, " ", 2)
					ingr := recipe.Ingredient{Name: fields[0]}
					if len(fields) == 2 {
						ingr.Amount = fields[1]
					}
					rcp.AddIngredient("", ingr)
				}
			}
			for _, inst := range ldRcp.RecipeInstructions {
				step := recipe.Step{}
				switch inst := inst.(type) {
				case *jsonld.HowToStep:
					for _, text := range inst.Text {
						if text, ok := text.(string); ok {
							step.Text += text
						}
					}
				case string:
					step.Text = strings.TrimSpace(inst)
				}
				rcp.Steps = append(rcp.Steps, step)
			}

			return false // 1ページに複数レシピがある場合があるので必ず1個目で中止
		}

		return true
	})

	if rcp == nil {
		return nil, sites.ErrUnsupportedURL
	}

	return rcp, nil
}

func NewParser() sites.Parser {
	return &parser{}
}
