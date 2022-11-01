package jsonld

import (
	"context"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/psyark/jsonld"
	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

type parser struct{}

var (
	commonGroupRegex = regexp.MustCompile(`^([ABCＡＢＣ])(?:\s*)(.+)$`)
)

func (p *parser) Parse(ctx context.Context, url string) (*rexch.Recipe, error) {
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
					break
				}
			}
			for _, text := range ldRcp.Image {
				if text, ok := text.(string); ok {
					rex.Image = text
					break
				}
			}

			for _, yield := range ldRcp.RecipeYield {
				if yield, ok := yield.(string); ok {
					if servings, ok := sites.ParseServings(yield); ok {
						rex.Servings = servings
					}
				}
			}

			for _, text := range ldRcp.RecipeIngredient {
				if text, ok := text.(string); ok {
					fields := strings.SplitN(text, " ", 2)
					if len(fields) < 2 {
						fields = append(fields, "")
					}
					igd := rexch.NewIngredient(fields[0], fields[1])
					rex.Ingredients = append(rex.Ingredients, *igd)
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
					for _, url := range inst.Image {
						if url, ok := url.(string); ok {
							ist.AddImage(url)
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

	{ // 既知のプリフィックスに基づいてグループ化
		groupCount := map[string]int{}
		for _, igd := range rex.Ingredients {
			if match := commonGroupRegex.FindStringSubmatch(igd.Name); len(match) != 0 {
				groupCount[match[1]]++
			}
		}
		for i := range rex.Ingredients {
			igd := &rex.Ingredients[i]
			if match := commonGroupRegex.FindStringSubmatch(igd.Name); len(match) != 0 && groupCount[match[1]] >= 2 {
				igd.Group = match[1]
				igd.Name = match[2]
			}
		}
	}

	if rex == nil {
		return nil, sites.ErrUnsupportedURL
	}

	return rex, nil
}

func NewParser() sites.Parser {
	return &parser{}
}
