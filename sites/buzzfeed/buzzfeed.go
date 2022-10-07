package buzzfeed

import (
	"context"
	"regexp"
	"strconv"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

var (
	servingsRegex = regexp.MustCompile(`(\d+)人分`)
	stepRegex     = regexp.MustCompile(`^[①-⑩]\s*`)
	groupRegex    = regexp.MustCompile(`^([ABC])(.+)$`)
)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	rex, err := p.Parse2(ctx, url)
	if err != nil {
		return nil, err
	}
	return rex.BackCompat(), nil
}

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://www.buzzfeed.com/jp/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rex := &rexch.Recipe{
		Title: strings.TrimSpace(doc.Find(`h2.subbuzz__title`).Eq(0).Text()),
		Image: sites.ResolvePath(url, doc.Find(`img.subbuzz-picture`).Eq(0).AttrOr("src", "")),
	}

	doc.Find(".subbuzz__description").Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		if strings.Contains(t, "材料") && strings.Contains(t, "作り方") {
			mode := ""
			s.Children().Each(func(i int, s *goquery.Selection) {
				t := s.Text()
				if mode == "" {
					if match := servingsRegex.FindStringSubmatch(t); len(match) != 0 {
						i, _ := strconv.Atoi(match[1])
						rex.Servings = i
					}
				}
				if mode == "" && t == "材料：" {
					mode = "inde"
					return
				} else if mode == "inde" && t == "作り方：" {
					mode = "step"
					return
				}
				if mode == "inde" && t != "" {
					igd := rexch.Ingredient{}
					pair := strings.SplitN(t, "　", 2)
					if len(pair) == 2 {
						igd.Name = pair[0]
						igd.Amount = pair[1]
					} else {
						igd.Name = t
					}

					if match := groupRegex.FindStringSubmatch(igd.Name); len(match) == 3 {
						igd.Group = match[1]
						igd.Name = match[2]
					}

					rex.Ingredients = append(rex.Ingredients, igd)
				}
				if mode == "step" && t != "" {
					rex.Instructions = append(rex.Instructions, rexch.Instruction{
						Elements: []rexch.InstructionElement{
							&rexch.TextInstructionElement{Text: stepRegex.ReplaceAllString(t, "")},
						},
					})
				}
			})
		}
	})

	return rex, nil
}

func NewParser() sites.Parser2 {
	return &parser{}
}
