package macaroni

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/psyark/jsonld"
	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"

	"github.com/PuerkitoBio/goquery"
)

type parser struct{}

var (
	groupRegex    = regexp.MustCompile(`^([ＡＢ])　(.+)$`)
	servingsRegex = regexp.MustCompile(`(\d+)(〜\d+)?人分`)
	newInstRegex  = regexp.MustCompile(`^(\d\. |[①-⑳])`)
)

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://macaro-ni.jp/") {
		return nil, sites.ErrUnsupportedURL
	}

	rex := &rexch.Recipe{}
	if err := p.parseURL(ctx, url, rex); err != nil {
		return nil, err
	}
	return rex, nil
}

func (p *parser) parseURL(ctx context.Context, url string, rex *rexch.Recipe) error {
	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return err
	}

	doc.Find(`script[type="application/ld+json"]`).Each(func(i int, s *goquery.Selection) {
		obj, err := jsonld.DecodeObject([]byte(s.Text()))
		if err != nil {
			return
		}

		if obj, ok := obj.(*jsonld.Recipe); ok {
			for _, text := range obj.Name {
				rex.Title = text.(string)
				break
			}
			for _, image := range obj.Image {
				rex.Image = image.(string)
				break
			}
			for _, igd := range obj.RecipeIngredient {
				parts := strings.Split(igd.(string), " ")

				group := ""
				if len(parts) == 3 {
					group = parts[0]
					parts = parts[1:]
				}

				igd := rexch.NewIngredient(parts[0], parts[1])
				igd.Group = group
				rex.Ingredients = append(rex.Ingredients, *igd)
			}
			for _, ins := range obj.RecipeInstructions {
				if ins, ok := ins.(*jsonld.HowToStep); ok {
					inst := rexch.Instruction{}
					for _, text := range ins.Text {
						inst.AddText(text.(string))
					}
					for _, image := range ins.Image {
						inst.AddImage(image.(string))
					}
					rex.Instructions = append(rex.Instructions, inst)
				}
			}
		}
	})

	if len(rex.Ingredients) != 0 && len(rex.Instructions) != 0 {
		return nil
	}

	mode := "intro"
	group := ""
	ins := rexch.Instruction{}

	for {
		body := doc.Find(".articleShow__body")
		body.Find(".articleShow__nutrition").Remove()       // 栄養情報
		body.Find(".articleShow__contentsLink").Remove()    // 他レシピへのリンク
		body.Find(".articleShow__contentsPhotoBy").Remove() // Photo by
		body.Find("script").Remove()
		body.Find("img").Each(func(i int, s *goquery.Selection) {
			s.ReplaceWithHtml(fmt.Sprintf("\nIMAGE:%s\n", s.AttrOr("data-original", s.AttrOr("src", ""))))
		})
		body.Find("br").ReplaceWithHtml("\n")

		for _, line := range strings.Split(body.Text(), "\n") {
			line = strings.TrimSpace(line)
			if line != "" {
				if mode == "intro" {
					if strings.Contains(line, "材料") {
						mode = "ingredients"
						if match := servingsRegex.FindStringSubmatch(line); len(match) != 0 {
							i, _ := strconv.Atoi(match[1])
							rex.Servings = i
						}
					}
					continue
				}
				if mode == "ingredients" {
					if strings.HasPrefix(line, "IMAGE:") {
					} else if strings.Contains(line, "下準備") {
						mode = "instructions"
					} else if strings.Contains(line, "作り方") {
						mode = "instructions"
						continue
					} else {
						var fields []string
						if strings.Contains(line, "……") {
							fields = strings.SplitN(line, "……", 2)
						} else if strings.Contains(line, "　") {
							fields = strings.SplitN(line, "　", 2)
						} else if strings.HasPrefix(line, "〈") || strings.HasPrefix(line, "＜") {
							group = line
							continue
						} else {
							fields = []string{line}
						}

						if len(fields) < 2 {
							fields = append(fields, "")
						}
						igd := rexch.NewIngredient(strings.TrimPrefix(fields[0], "・"), fields[1])
						if match := groupRegex.FindStringSubmatch(igd.Name); len(match) != 0 {
							igd.Group = match[1]
							igd.Name = match[2]
						} else {
							igd.Group = group
						}
						rex.Ingredients = append(rex.Ingredients, *igd)
					}
				}
				if mode == "instructions" {
					if strings.Contains(line, "作り方") {
						continue
					}
					if strings.HasPrefix(line, "IMAGE:") {
						ins.AddImage(strings.TrimPrefix(line, "IMAGE:"))
					} else {
						if len(ins.Elements) != 0 && newInstRegex.MatchString(line) {
							rex.Instructions = append(rex.Instructions, ins)
							ins = rexch.Instruction{}
						}
						ins.AddText(line)
					}
				}
			} else {
				group = ""
			}
		}

		if doc.Find(".articleShow__nextPage").Length() != 0 {
			doc, err = sites.NewDocumentFromURL(ctx, doc.Find(".articleShow__nextPage a").AttrOr("href", ""))
			if err != nil {
				return err
			}
		} else {
			break
		}
	}

	rex.Instructions = append(rex.Instructions, ins)

	return nil
}

func NewParser() sites.Parser2 {
	return &parser{}
}
