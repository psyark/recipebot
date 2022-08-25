package buzzfeed

import (
	"context"
	"regexp"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

var (
	stepRegex  = regexp.MustCompile(`^[①-⑩]\s*`)
	groupRegex = regexp.MustCompile(`^([ABC])(.+)$`)
)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	if !strings.HasPrefix(url, "https://www.buzzfeed.com/jp/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rcp := &recipe.Recipe{
		Title: strings.TrimSpace(doc.Find(`h2.subbuzz__title`).Eq(0).Text()),
		Image: sites.ResolvePath(url, doc.Find(`img.subbuzz-picture`).Eq(0).AttrOr("src", "")),
	}

	doc.Find(".subbuzz__description").Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		if strings.Contains(t, "材料") && strings.Contains(t, "作り方") {
			mode := ""
			s.Children().Each(func(i int, s *goquery.Selection) {
				t := s.Text()
				if mode == "" && t == "材料：" {
					mode = "inde"
					return
				} else if mode == "inde" && t == "作り方：" {
					mode = "step"
					return
				}
				if mode == "inde" && t != "" {
					igd := recipe.Ingredient{}
					pair := strings.SplitN(t, "　", 2)
					if len(pair) == 2 {
						igd.Name = pair[0]
						igd.Amount = pair[1]
					} else {
						igd.Name = t
					}

					if match := groupRegex.FindStringSubmatch(igd.Name); len(match) == 3 {
						igd.Name = match[2]
						rcp.AddIngredient(match[1], igd)
					} else {
						rcp.AddIngredient("", igd)
					}
				}
				if mode == "step" && t != "" {
					rcp.Steps = append(rcp.Steps, recipe.Step{
						Text: stepRegex.ReplaceAllString(t, ""),
					})
				}
			})
		}
	})

	return rcp, nil
}

func NewParser() sites.Parser {
	return &parser{}
}
