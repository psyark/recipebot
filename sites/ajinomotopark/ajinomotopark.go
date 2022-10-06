package ajinomotopark

import (
	"context"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

type parser struct{}

var debrandMap = map[string]string{
	"「AJINOMOTO オリーブオイル」":     "オリーブオイル",
	"「AJINOMOTO ごま油好きの純正ごま油」": "ごま油",
	"「AJINOMOTO さらさらキャノーラ油」":  "サラダ油",
	"「AJINOMOTO サラダ油」":        "サラダ油",
	"「Cook Do」オイスターソース":       "オイスターソース",
	"「Cook Do」熟成豆板醤":          "豆板醤",
	"「丸鶏がらスープ」":               "鶏がらスープ",
	"「瀬戸のほんじお」":               "塩",
	"「味の素KKコンソメ」固形タイプ":        "コンソメ（固形）",
	"「味の素KKコンソメ」顆粒タイプ":        "コンソメ（顆粒）",
	"「味の素KK中華あじ」":             "粉末中華スープ",
	"うま味調味料「味の素®」":            "味の素",
}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	if !strings.HasPrefix(url, "https://park.ajinomoto.co.jp/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rcp := &recipe.Recipe{
		Title: strings.TrimSpace(doc.Find(`h1.recipeTitle`).Text()),
		Image: doc.Find(`.recipeImageArea img`).AttrOr("src", ""),
	}

	doc.Find(`.recipeMaterialList dl dt`).Each(func(i int, s *goquery.Selection) {
		groupName := ""
		className := s.AttrOr("class", "")
		if strings.HasPrefix(className, "ico") {
			groupName = strings.TrimPrefix(className, "ico")
		}

		idg := recipe.GetIngredient(debrand(strings.TrimSpace(s.Text())), s.Next().Text())
		rcp.AddIngredient(groupName, idg)
	})

	doc.Find(`#makeList ol li`).Each(func(i int, s *goquery.Selection) {
		stp := recipe.Step{Text: strings.TrimSpace(s.Text())}
		s.Find("img").Each(func(i int, s *goquery.Selection) {
			stp.Images = append(stp.Images, s.AttrOr("src", ""))
		})
		rcp.Steps = append(rcp.Steps, stp)
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
