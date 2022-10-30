package macaroni

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type parser struct{}

var idgRegex = regexp.MustCompile(`・(.+?)(?:（(.+?)）)?……(.+)`)

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	rex, err := p.Parse2(ctx, url)
	if err != nil {
		return nil, err
	}
	return rex.BackCompat(), nil
}

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://macaro-ni.jp/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rex := &rexch.Recipe{}

	doc.Find(`script[type="application/ld+json"]`).Each(func(i int, s *goquery.Selection) {
		tmp := Recipe{}
		if err := json.Unmarshal([]byte(s.Text()), &tmp); err != nil {
			return
		}

		if tmp.Type == "Recipe" && len(tmp.RecipeIngredient) != 0 && len(tmp.RecipeInstructions) != 0 {
			rex.Title = tmp.Name
			rex.Image = tmp.Image[0]
			for _, idg := range tmp.RecipeIngredient {
				parts := strings.Split(idg, " ")

				group := ""
				if len(parts) == 3 {
					group = parts[0]
					parts = parts[1:]
				}

				igd := rexch.NewIngredient(parts[0], parts[1])
				igd.Group = group
				rex.Ingredients = append(rex.Ingredients, *igd)
			}
			for _, ins := range tmp.RecipeInstructions {
				inst := rexch.Instruction{}
				inst.AddText(ins.Text)
				if ins.Image != "" {
					inst.AddImage(ins.Image)
				}
				rex.Instructions = append(rex.Instructions, inst)
			}
		}
	})

	if len(rex.Instructions) != 0 {
		return rex, nil
	}

	rex.Title = strings.TrimSpace(doc.Find(`h1.articleInfo__title`).Text())
	rex.Image = doc.Find(`img.articleInfo__thumbnail`).AttrOr("src", "")

	var curStep *rexch.Instruction
	parseSteps := func(i int, s *goquery.Selection) {
		switch s.AttrOr("class", "") {
		case "articleShow__contentsHeading":
			rex.Instructions = append(rex.Instructions, rexch.Instruction{})
			curStep = &rex.Instructions[len(rex.Instructions)-1]
			curStep.AddText(strings.TrimSpace(s.Text()))
		case "articleShow__contentsText":
			if curStep == nil {
				rex.Instructions = append(rex.Instructions, rexch.Instruction{})
				curStep = &rex.Instructions[len(rex.Instructions)-1]
			}
			for _, line := range parseCTB(s.Find(`.articleShow__contentsTextBody`).Get(0)) {
				curStep.AddText(strings.TrimSpace(line))
			}
		case "articleShow__contentsImage":
			s.Find("img").Each(func(i int, s *goquery.Selection) {
				curStep.AddImage(s.AttrOr("data-original", s.AttrOr("src", "")))
			})
		}
	}

	doc.Find(`.articleShow__contents`).Each(func(i int, s *goquery.Selection) {
		if strings.HasPrefix(strings.TrimSpace(s.Find(".articleShow__contentsHeading").Text()), "材料") {
			groupName := ""
			for _, line := range parseCTB(s.NextAll().Find(".articleShow__contentsTextBody").Get(0)) {
				if line == "" {
					groupName = "" // 改行が連続したらグループ解除
				} else if strings.HasPrefix(line, "・") {
					match := idgRegex.FindStringSubmatch(line)
					igd := rexch.Ingredient{Group: groupName, Name: match[1], Amount: match[3], Comment: match[2]}
					rex.Ingredients = append(rex.Ingredients, igd)
				} else {
					groupName = line
				}
			}
		} else if strings.TrimSpace(s.Find(".articleShow__contentsHeading").Text()) == "作り方" {
			s.NextAll().Find("div").Each(parseSteps)
		}
	})

	for doc.Find(".articleShow__nextPage").Length() != 0 {
		doc, err = sites.NewDocumentFromURL(ctx, doc.Find(".articleShow__nextPage a").AttrOr("href", ""))
		if err != nil {
			return nil, err
		}
		doc.Find(`.articleShow__contents div`).Each(parseSteps)
	}

	return rex, nil
}

func parseCTB(ctb *html.Node) []string {
	lines := ""
	for c := ctb.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			lines += c.Data
		} else if c.Type == html.ElementNode {
			if c.Data == "br" {
				lines += "\n"
			} else {
				lines += goquery.NewDocumentFromNode(c).Text()
			}
		}
	}
	return strings.Split(lines, "\n")
}

func NewParser() sites.Parser2 {
	return &parser{}
}
