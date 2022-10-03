package kikkoman

import (
	"context"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

type parser struct{}

var debrandMap = map[string]string{
	"キッコーマン特選丸大豆しょうゆ":        "醤油",
	"キッコーマンいつでも新鮮しぼりたて生しょうゆ": "醤油",
	"マンジョウ濃厚熟成本みりん":          "みりん",
	"マンジョウ国産米こだわり仕込み料理の清酒":   "料理酒",
}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	if !strings.HasPrefix(url, "https://www.kikkoman.co.jp/homecook/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rcp := &recipe.Recipe{
		Title: strings.TrimSpace(doc.Find(`.elem-heading-lv1`).Text()),
		Image: sites.ResolvePath(url, doc.Find(`.main-movBox img`).AttrOr("src", "")),
	}

	groupName := ""
	doc.Find(`.ingredients > *`).Each(func(i int, s *goquery.Selection) {
		switch s.AttrOr("class", "") {
		case "elem-heading-lv2": // 料理名
		case "ingredients-form--list": // 材料
			rcp.AddIngredient(groupName, recipe.Ingredient{
				Name:   debrand(strings.TrimSpace(s.Find("dt").Text())),
				Amount: strings.TrimSpace(s.Find("dd").Text()),
			})
		case "elem-heading-lv4": // 材料グループ名
			groupName = strings.TrimSpace(s.Text())
		}
	})

	doc.Find(`.instruction`).Each(func(i int, s *goquery.Selection) {
		rcp.Steps = append(rcp.Steps, recipe.Step{Text: strings.TrimSpace(s.Text())})
	})

	return rcp, nil
}

func debrand(name string) string {
	if debranded, ok := debrandMap[name]; ok {
		return debranded
	}
	return name
}

func NewParser() sites.Parser {
	return &parser{}
}
