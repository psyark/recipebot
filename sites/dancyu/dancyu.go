package dancyu

import (
	"context"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/psyark/jsonld"
	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"
)

var (
	titleRegex   = regexp.MustCompile(`"(.+?)"`)
	commentRegex = regexp.MustCompile(`^(.+)（(.+)）$`)
	stepRegex    = regexp.MustCompile(`^\d+\s*`)
)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	if !strings.HasPrefix(url, "https://dancyu.jp/recipe/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rcp := &recipe.Recipe{}
	var parseError error

	doc.Find(`script[type="application/ld+json"]`).EachWithBreak(func(i int, s *goquery.Selection) bool {
		jsonStr := s.Text()
		jsonStr = strings.ReplaceAll(jsonStr, "\n", " ") // 改行文字を含んではならない https://dancyu.jp/recipe/2022_00006083.html
		jsonStr = strings.ReplaceAll(jsonStr, "\t", " ") // タブ文字を含んではならない https://dancyu.jp/recipe/2020_00003873.html

		obj2, err := jsonld.DecodeObject([]byte(jsonStr))
		if err != nil {
			parseError = err
			return false
		}

		if rcp2, ok := obj2.(*jsonld.Recipe); ok {
			for _, name := range rcp2.Name {
				if name, ok := name.(string); ok {
					rcp.Title = name
				}
			}

			for _, image := range rcp2.Image {
				switch image := image.(type) {
				case string:
					rcp.Image = strings.TrimSpace(image)
				case *jsonld.ImageObject:
					for _, url := range image.Url {
						if url, ok := url.(string); ok {
							rcp.Image = strings.TrimSpace(url)
						}
					}
				}
			}

			ingredients := []string{}
			hasPrefixA := 0
			hasPrefixDot := 0
			for _, ing := range rcp2.RecipeIngredient {
				if ing, ok := ing.(string); ok {
					ing := strings.ReplaceAll(ing, "&emsp;", " ")
					if strings.HasPrefix(ing, "A ") {
						hasPrefixA++
					}
					if strings.HasPrefix(ing, "・") {
						hasPrefixDot++
					}
					ingredients = append(ingredients, ing)
				}
			}

			if hasPrefixA == 1 && hasPrefixDot == 0 {
				p.parseIngrediens3(ingredients, rcp)
			} else if hasPrefixA > 1 {
				p.parseIngrediens2(ingredients, rcp)
			} else {
				p.parseIngrediens1(ingredients, rcp)
			}
			return false // 1ページに複数レシピがある場合があるので必ず1個目で中止
		}

		return true
	})

	if parseError != nil {
		return nil, parseError
	}

	// タイトルの引用符内側
	if match := titleRegex.FindStringSubmatch(rcp.Title); len(match) == 2 {
		rcp.Title = match[1]
	}

	// 手順の画像がJSON-LDに含まれないのでHTMLをパース
	const stepSelector = `.snippet__freearea__subtitle-number,.block__snippet__freeText`
	doc.Find(`div.block__snippet`).ChildrenFiltered(stepSelector).Each(func(i int, s *goquery.Selection) {
		step := recipe.Step{Text: stepRegex.ReplaceAllString(strings.TrimSpace(s.Text()), "")}
		s.NextUntil(stepSelector).Each(func(i int, s *goquery.Selection) {
			cls := s.AttrOr("class", "")
			if cls == "snippet__freearea__simple-text" {
				step.Text += "\n" + s.Text()
			} else if strings.Contains(cls, "block__snippet__column") {
				s.Find("img").Each(func(i int, s *goquery.Selection) {
					step.Images = append(step.Images, sites.ResolvePath(url, s.AttrOr("data-src", "")))
				})
			}
		})
		rcp.Steps = append(rcp.Steps, step)
	})

	return rcp, nil
}

// パターン1
// https://dancyu.jp/recipe/2021_00005129.html
// 中黒は第2レベル、"A" は列挙の開始時に置かれる
func (p *parser) parseIngrediens1(ingredients []string, rcp *recipe.Recipe) {
	group := ""
	for i, item := range ingredients {
		pair := strings.SplitN(item, "：", 2)
		th := strings.TrimSpace(pair[0])
		td := strings.TrimSpace(pair[1])

		if !strings.HasPrefix(th, "・") && i < len(ingredients)-1 {
			if strings.HasPrefix(ingredients[i+1], "・") {
				group = th + td
				continue
			}
		}

		idg := recipe.Ingredient{Name: th, Amount: td}
		if strings.HasPrefix(idg.Name, "・") {
			idg.Name = strings.TrimSpace(strings.TrimPrefix(idg.Name, "・"))
		} else {
			group = ""
		}
		if match := commentRegex.FindStringSubmatch(idg.Amount); len(match) == 3 {
			idg.Amount = match[1]
			idg.Comment = match[2]
		}
		rcp.AddIngredient(group, idg)
	}
}

// パターン2
// https://dancyu.jp/recipe/2019_00001391.html
// 中黒はトップレベル、"A" は各アイテムに付く
func (p *parser) parseIngrediens2(ingredients []string, rcp *recipe.Recipe) {
	for _, item := range ingredients {
		group := ""
		pair := strings.SplitN(item, "：", 2)
		th := strings.TrimSpace(pair[0])
		td := strings.TrimSpace(pair[1])

		if strings.HasPrefix(th, "・") {
			th = strings.TrimSpace(strings.TrimPrefix(th, "・"))
		} else if strings.Contains(th, " ") {
			pair := strings.SplitN(th, " ", 2)
			group = pair[0]
			th = pair[1]
		}

		idg := recipe.Ingredient{Name: th, Amount: td}

		if match := commentRegex.FindStringSubmatch(idg.Amount); len(match) == 3 {
			idg.Amount = match[1]
			idg.Comment = match[2]
		}

		rcp.AddIngredient(group, idg)
	}
}

// パターン3
// https://dancyu.jp/recipe/2020_00003873.html
// "A" はグループ名に付くが中黒は使わない
func (p *parser) parseIngrediens3(ingredients []string, rcp *recipe.Recipe) {
	group := ""
	for _, item := range ingredients {
		pair := strings.SplitN(item, "：", 2)
		th := strings.TrimSpace(pair[0])
		td := strings.TrimSpace(pair[1])

		if strings.HasPrefix(th, "A ") || strings.HasPrefix(th, "B ") || strings.HasPrefix(th, "C ") {
			group = th
		} else {
			idg := recipe.Ingredient{Name: th, Amount: td}

			if match := commentRegex.FindStringSubmatch(idg.Amount); len(match) == 3 {
				idg.Amount = match[1]
				idg.Comment = match[2]
			}

			rcp.AddIngredient(group, idg)
		}
	}
}

func NewParser() sites.Parser {
	return &parser{}
}
