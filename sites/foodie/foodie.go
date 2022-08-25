package foodie

import (
	"context"
	"regexp"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

type parser struct{}

var idgRegex = regexp.MustCompile(`(.+?)(?:（(.+?)）)?…+(.+)`)
var stepRegex = regexp.MustCompile(`^(\d+)\.　(.+)`)

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	if !strings.HasPrefix(url, "https://mi-journey.jp/foodie/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rcp := &recipe.Recipe{
		Title: strings.TrimSpace(doc.Find(`h1.main_title`).Text()),
		Image: doc.Find(`#contents .main_pic img`).AttrOr("src", ""),
	}

	doc.Find(`#contents > *`).Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if strings.HasPrefix(text, "＜材料＞") || strings.HasPrefix(text, "［材料］") {
			groupName := ""
			s.NextAll().EachWithBreak(func(i int, s *goquery.Selection) bool {
				switch goquery.NodeName(s) {
				case "ul":
					s.Children().Each(func(i int, s *goquery.Selection) {
						parts := idgRegex.FindStringSubmatch(strings.TrimSpace(s.Text()))
						igr := recipe.Ingredient{
							Name:    strings.TrimSpace(parts[1]),
							Amount:  strings.TrimSpace(parts[3]),
							Comment: strings.TrimSpace(parts[2]),
						}
						if igr.Name != "" {
							rcp.AddIngredient(groupName, igr)
						}
					})
				case "p":
					groupName = strings.TrimSpace(s.Text())
				case "h3":
					return false
				}
				return true
			})
		}
	})

	doc.Find(`#contents > *`).Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text == "＜作り方＞" || text == "［作り方］" {
			var curStep *recipe.Step
			s.NextAll().EachWithBreak(func(i int, s *goquery.Selection) bool {
				switch goquery.NodeName(s) {
				case "h3":
					text := strings.TrimSpace(s.Text())
					rcp.Steps = append(rcp.Steps, recipe.Step{Text: text})
					curStep = &rcp.Steps[len(rcp.Steps)-1]
					return true
				case "p":
					text := strings.TrimSpace(s.Text())
					match := stepRegex.FindStringSubmatch(text)
					if match != nil {
						rcp.Steps = append(rcp.Steps, recipe.Step{Text: match[2]})
						curStep = &rcp.Steps[len(rcp.Steps)-1]
					} else {
						curStep.Text += "\n" + text
						imgs := s.Find("img").Map(func(i int, s *goquery.Selection) string { return s.AttrOr("src", "") })
						curStep.Images = append(curStep.Images, imgs...)
					}
					return true
				}
				return false
			})
		}
	})

	doc.Find(`li.step`).Each(func(i int, s *goquery.Selection) {
		stp := recipe.Step{
			Text: strings.TrimSpace(s.Find(`.step_text`).Text()),
		}

		// "Invalid image url." エラーとなるため画像は未サポート（CookpadのCDNがContent-Typeを送っていないの関係ある？）

		// s.Find(`.image img`).Each(func(i int, s *goquery.Selection) {
		// 	src := getSrc(s)
		// 	stp.Images = append(stp.Images, src)
		// })

		if strings.HasSuffix(stp.Text, "クックパッドニュース") {
			return
		}
		if strings.Contains(stp.Text, "感謝") {
			return
		}

		rcp.Steps = append(rcp.Steps, stp)
	})

	return rcp, nil
}

func NewParser() sites.Parser {
	return &parser{}
}
