package jsonld

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/psyark/jsonld"
	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

type parser struct{}

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	var rex *rexch.Recipe

	doc.Find(`script[type="application/ld+json"]`).EachWithBreak(func(i int, s *goquery.Selection) bool {
		jsonStr := s.Text()

		obj, err := jsonld.DecodeObject([]byte(jsonStr))
		if err != nil {
			// パースエラーに続いて有効なレシピが得られる場合があるので無視
			return true
		}

		if ldRcp, ok := obj.(*jsonld.Recipe); ok {
			rex = &rexch.Recipe{}

			for _, text := range ldRcp.Name {
				if text, ok := text.(string); ok {
					rex.Title = text
				}
			}
			for _, text := range ldRcp.Image {
				if text, ok := text.(string); ok {
					rex.Image = text
				}
			}
			for _, text := range ldRcp.RecipeIngredient {
				if text, ok := text.(string); ok {
					fields := strings.SplitN(text, " ", 2)
					igd := rexch.Ingredient{Name: fields[0]}
					if len(fields) == 2 {
						igd.Amount = fields[1]
					}
					rex.Ingredients = append(rex.Ingredients, igd)
				}
			}
			for _, inst := range ldRcp.RecipeInstructions {
				ist := rexch.Instruction{}
				switch inst := inst.(type) {
				case *jsonld.HowToStep:
					for _, text := range inst.Text {
						if text, ok := text.(string); ok {
							ist.AddText(text)
						}
					}
				case string:
					ist.AddText(strings.TrimSpace(inst))
				}
				rex.Instructions = append(rex.Instructions, ist)
			}

			return false // 1ページに複数レシピがある場合があるので必ず1個目で中止
		}

		return true
	})

	if rex == nil {
		return nil, sites.ErrUnsupportedURL
	}

	return rex, nil
}

func NewParser() sites.Parser2 {
	return &parser{}
}
