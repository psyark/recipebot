package cookpad

import (
	"context"
	"regexp"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/width"
)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	if !strings.HasPrefix(url, "https://cookpad.com/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rcp := &recipe.Recipe{
		Title: strings.TrimSpace(doc.Find(`h1.recipe-title`).Text()),
		Image: getSrc(doc.Find(`#main-photo img`)),
	}

	igrRegex := regexp.MustCompile(`^(.*)（(.+)）$`)

	doc.Find(`.ingredient_row`).Each(func(i int, s *goquery.Selection) {
		igr := recipe.Ingredient{
			Name:   strings.TrimSpace(width.Widen.String(s.Find(`.ingredient_name`).Text())),
			Amount: strings.TrimSpace(width.Fold.String(s.Find(`.ingredient_quantity`).Text())),
		}
		if igr.Name != "" {
			if match := igrRegex.FindStringSubmatch(igr.Name); len(match) != 0 {
				igr.Name = match[1]
				igr.Comment = match[2]
			}

			rcp.AddIngredient("", igr)
		}
	})

	doc.Find(`li.step, li.step_last`).Each(func(i int, s *goquery.Selection) {
		stp := recipe.Step{
			Text: strings.TrimSpace(s.Find(`.step_text`).Text()),
		}

		// "Invalid image url." エラーとなるため画像は未サポート（CookpadのCDNがContent-Typeを送っていないの関係ある？）

		// s.Find(`.image img`).Each(func(i int, s *goquery.Selection) {
		// 	src := getSrc(s)
		// 	stp.Images = append(stp.Images, src)
		// })

		if strings.Contains(stp.Text, "クックパッドニュース") {
			return
		}
		if strings.Contains(stp.Text, "感謝") {
			return
		}
		if strings.Contains(stp.Text, "発売") {
			return
		}
		if strings.Contains(stp.Text, "掲載") {
			return
		}
		if strings.Contains(stp.Text, "検索") {
			return
		}
		if strings.Contains(stp.Text, "話題") {
			return
		}
		if strings.Contains(stp.Text, "ありがとう") {
			return
		}
		if strings.Contains(stp.Text, "年") {
			return
		}
		if len([]rune(stp.Text)) < 3 {
			return
		}

		rcp.Steps = append(rcp.Steps, stp)
	})

	return rcp, nil
}

func NewParser() sites.Parser {
	return &parser{}
}

func getSrc(s *goquery.Selection) string {
	return s.AttrOr("data-large-photo", s.AttrOr("src", ""))
}
