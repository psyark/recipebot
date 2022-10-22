package sirogohan

import (
	"context"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	url = p.normalizeURL(url)

	if !strings.HasPrefix(url, "https://www.sirogohan.com/recipe/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rcp := &recipe.Recipe{
		// h1#recipe-name には余分な文言が入っている場合があるので使わない
		Title: strings.TrimSuffix(strings.TrimSpace(doc.Find(`.howto .recipe-sttl`).Text()), "の作り方"),
		Image: doc.Find(`p#recipe-main img`).AttrOr("src", ""),
	}

	doc.Find(`.material ul`).Each(func(i int, s *goquery.Selection) {
		groupName := ""
		classes := " " + s.AttrOr("class", "") + " "
		if strings.Contains(classes, " a-list ") {
			groupName = "A"
		} else if strings.Contains(classes, " b-list ") {
			groupName = "B"
		} else if strings.Contains(classes, " c-list ") {
			groupName = "C"
		}
		s.Find(`li`).Each(func(i int, s *goquery.Selection) {
			parts := strings.Split(s.Text(), "　…　")
			if len(parts) == 1 {
				parts = append(parts, "")
			}
			rcp.AddIngredient(groupName, recipe.Ingredient{
				Name:   strings.TrimSpace(parts[0]),
				Amount: strings.TrimSpace(parts[1]),
			})
		})
	})

	doc.Find(`.howto-block`).Not(`#recipe_movie`).Each(func(i int, s *goquery.Selection) {
		rcp.Steps = append(rcp.Steps, recipe.Step{
			Text: strings.TrimSpace(s.Text()),
			Images: s.Find("img").Map(func(i int, s *goquery.Selection) string {
				return sites.ResolvePath(url, s.AttrOr("src", ""))
			}),
		})
	})

	return rcp, nil
}

func (p *parser) normalizeURL(url string) string {
	if strings.HasPrefix(url, "https://www.sirogohan.com/sp/") {
		url = "https://www.sirogohan.com/" + strings.TrimPrefix(url, "https://www.sirogohan.com/sp/")
	}
	if strings.HasSuffix(url, "/amp/") {
		url = strings.TrimSuffix(url, "/amp/") + "/"
	}
	return url
}

func NewParser() sites.Parser {
	return &parser{}
}
