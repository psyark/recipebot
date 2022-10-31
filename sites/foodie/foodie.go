// Package foodie はfoodieに特化したパーサです
// foodieはJSONLDを提供していますが、内容があまりにも酷いので特別な対応を要します
package foodie

import (
	"context"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/psyark/jsonld"
	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var (
	nameAmountRegex = regexp.MustCompile(`([^…]+)…+([^…]+)`)
	instRegex       = regexp.MustCompile(`^\d+\.?`)
)

type parser struct{}

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://mi-journey.jp/foodie/") {
		return nil, sites.ErrUnsupportedURL
	}

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
			for _, ingr := range ldRcp.RecipeIngredient {
				switch ingr := ingr.(type) {
				case string:
					group := ""
					ingr = strings.ReplaceAll(ingr, "\r", "\n") // 改行コード統一
					for _, line := range strings.Split(ingr, "\n") {
						line = strings.TrimSpace(line)
						if line != "" {
							line = strings.TrimPrefix(line, "・")
							line = strings.TrimPrefix(line, "〇")

							if match := nameAmountRegex.FindStringSubmatch(line); len(match) != 0 {
								igd := rexch.NewIngredient(match[1], match[2])
								igd.Group = group
								rex.Ingredients = append(rex.Ingredients, *igd)
							} else if strings.HasPrefix(line, "【") {
								group = line
							}
						}
					}
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
					ist.AddText(instRegex.ReplaceAllString(strings.TrimSpace(inst), ""))
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
