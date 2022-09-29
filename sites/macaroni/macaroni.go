package macaroni

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type parser struct{}

var idgRegex = regexp.MustCompile(`・(.+?)(?:（(.+?)）)?……(.+)`)

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	if !strings.HasPrefix(url, "https://macaro-ni.jp/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rcp := &recipe.Recipe{}

	doc.Find(`script[type="application/ld+json"]`).Each(func(i int, s *goquery.Selection) {
		tmp := Recipe{}
		if err := json.Unmarshal([]byte(s.Text()), &tmp); err != nil {
			return
		}

		if tmp.Type == "Recipe" && len(tmp.RecipeIngredient) != 0 && len(tmp.RecipeInstructions) != 0 {
			rcp.Title = tmp.Name
			rcp.Image = tmp.Image[0]
			for _, idg := range tmp.RecipeIngredient {
				parts := strings.Split(idg, " ")
				if len(parts) == 3 {
					rcp.AddIngredient(parts[0], recipe.Ingredient{Name: parts[1], Amount: parts[2]})
				} else {
					rcp.AddIngredient("", recipe.Ingredient{Name: parts[0], Amount: parts[1]})
				}
			}
			for _, ins := range tmp.RecipeInstructions {
				step := recipe.Step{Text: ins.Text}
				if ins.Image != "" {
					step.Images = append(step.Images, ins.Image)
				}
				rcp.Steps = append(rcp.Steps, step)
			}
		}
	})

	if len(rcp.Steps) != 0 {
		return rcp, nil
	}

	rcp.Title = strings.TrimSpace(doc.Find(`h1.articleInfo__title`).Text())
	rcp.Image = doc.Find(`img.articleInfo__thumbnail`).AttrOr("src", "")

	var curStep *recipe.Step
	parseSteps := func(i int, s *goquery.Selection) {
		switch s.AttrOr("class", "") {
		case "articleShow__contentsHeading":
			rcp.Steps = append(rcp.Steps, recipe.Step{Text: strings.TrimSpace(s.Text())})
			curStep = &rcp.Steps[len(rcp.Steps)-1]
		case "articleShow__contentsText":
			if curStep == nil {
				rcp.Steps = append(rcp.Steps, recipe.Step{})
				curStep = &rcp.Steps[len(rcp.Steps)-1]
			}
			for _, line := range parseCTB(s.Find(`.articleShow__contentsTextBody`).Get(0)) {
				curStep.Text += "\n" + line
			}
		case "articleShow__contentsImage":
			s.Find("img").Each(func(i int, s *goquery.Selection) {
				curStep.Images = append(curStep.Images, s.AttrOr("data-original", s.AttrOr("src", "")))
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
					rcp.AddIngredient(groupName, recipe.Ingredient{Name: match[1], Amount: match[3], Comment: match[2]})
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

	return rcp, nil
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

func NewParser() sites.Parser {
	return &parser{}
}
