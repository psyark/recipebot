package sirogohan

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

var servingsRegex = regexp.MustCompile(`(\d+)(?:～.+)?人分`)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	rex, err := p.Parse2(ctx, url)
	if err != nil {
		return nil, err
	}
	return rex.BackCompat(), nil
}

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	url = p.normalizeURL(url)

	if !strings.HasPrefix(url, "https://www.sirogohan.com/recipe/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rex := &rexch.Recipe{
		// h1#recipe-name には余分な文言が入っている場合があるので使わない
		Title: strings.TrimSuffix(strings.TrimSpace(doc.Find(`.howto .recipe-sttl`).Text()), "の作り方"),
		Image: doc.Find(`p#recipe-main img`).AttrOr("src", ""),
	}

	if match := (servingsRegex.FindStringSubmatch(doc.Find(`.material-ttl`).Text())); match != nil {
		if servings, err := strconv.Atoi(match[1]); err == nil {
			rex.Servings = servings
		}
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
			igd := rexch.NewIngredient(strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]))
			igd.Group = groupName
			rex.Ingredients = append(rex.Ingredients, *igd)
		})
	})

	doc.Find(`.howto-block`).Not(`#recipe_movie`).Each(func(i int, s *goquery.Selection) {
		ist := rexch.Instruction{}
		ist.Elements = append(ist.Elements, &rexch.TextInstructionElement{
			Text: strings.TrimSpace(s.Text()),
		})
		s.Find("img").Each(func(i int, s *goquery.Selection) {
			ist.Elements = append(ist.Elements, &rexch.ImageInstructionElement{
				URL: sites.ResolvePath(url, s.AttrOr("src", "")),
			})
		})
		rex.Instructions = append(rex.Instructions, ist)
	})

	return rex, nil
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

func NewParser() sites.Parser2 {
	return &parser{}
}
