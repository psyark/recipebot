package kikkoman

import (
	"context"
	"strings"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

type parser struct{}

var debrandMap = map[string]string{
	"キッコーマンいつでも新鮮しぼりたて生しょうゆ": "醤油",
	"キッコーマン特選丸大豆しょうゆ":        "醤油",
	"マンジョウ国産米こだわり仕込み料理の清酒":   "料理酒",
	"マンジョウ濃厚熟成本みりん":          "みりん",
	"マンジョウ米麹こだわり仕込み本みりん":     "みりん",
}

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://www.kikkoman.co.jp/homecook/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rex := &rexch.Recipe{
		Title: strings.TrimSpace(doc.Find(`.elem-heading-lv1`).Text()),
		Image: sites.ResolvePath(url, doc.Find(`img.photo`).AttrOr("src", "")),
	}

	groupName := ""
	doc.Find(`.ingredients > *`).Each(func(i int, s *goquery.Selection) {
		switch s.AttrOr("class", "") {
		case "elem-heading-lv2": // 料理名
		case "ingredients-form--list": // 材料
			igd := rexch.Ingredient{
				Group:  groupName,
				Name:   debrand(strings.TrimSpace(s.Find("dt").Text())),
				Amount: strings.TrimSpace(s.Find("dd").Text()),
			}
			rex.Ingredients = append(rex.Ingredients, igd)
		case "elem-heading-lv4": // 材料グループ名
			groupName = strings.TrimSpace(s.Text())
		}
	})

	doc.Find(`.instruction`).Each(func(i int, s *goquery.Selection) {
		ist := rexch.Instruction{}
		ist.AddText(strings.TrimSpace(s.Text()))
		rex.Instructions = append(rex.Instructions, ist)
	})

	return rex, nil
}

func debrand(name string) string {
	if debranded, ok := debrandMap[name]; ok {
		return debranded
	}
	return name
}

func NewParser() sites.Parser2 {
	return &parser{}
}
