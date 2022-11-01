package cookpad

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
	"golang.org/x/text/width"

	"github.com/PuerkitoBio/goquery"
)

type parser struct{}

var (
	igrRegex = regexp.MustCompile(`^(.*)（(.+)）$`)
	ngWords  = []string{"クックパッドニュース", "感謝", "発売", "掲載", "検索", "話題", "ありがとう", "年"}
)

func (p *parser) Parse(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://cookpad.com/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rex := &rexch.Recipe{
		Title: strings.TrimSpace(doc.Find(`h1.recipe-title`).Text()),
		Image: getSrc(doc.Find(`#main-photo img`)),
	}

	fmt.Println(width.Fold.String(doc.Find(`.servings_for`).Text()))

	doc.Find(`.ingredient_row`).Each(func(i int, s *goquery.Selection) {
		igd := rexch.NewIngredient(s.Find(`.ingredient_name`).Text(), s.Find(`.ingredient_quantity`).Text())
		if igd.Name != "" {
			if match := igrRegex.FindStringSubmatch(igd.Name); len(match) != 0 {
				igd.Name = match[1]
				igd.Comment = match[2]
			}
			rex.Ingredients = append(rex.Ingredients, *igd)
		}
	})

	doc.Find(`li.step, li.step_last`).Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Find(`.step_text`).Text())

		// "Invalid image url." エラーとなるため画像は未サポート（CookpadのCDNがContent-Typeを送っていないの関係ある？）

		// s.Find(`.image img`).Each(func(i int, s *goquery.Selection) {
		// 	src := getSrc(s)
		// 	stp.Images = append(stp.Images, src)
		// })

		for _, ngword := range ngWords {
			if strings.Contains(text, ngword) {
				return
			}
		}
		if len([]rune(text)) < 3 {
			return
		}

		ist := rexch.Instruction{}
		ist.AddText(text)
		rex.Instructions = append(rex.Instructions, ist)
	})

	return rex, nil
}

func NewParser() sites.Parser {
	return &parser{}
}

func getSrc(s *goquery.Selection) string {
	return s.AttrOr("data-large-photo", s.AttrOr("src", ""))
}
