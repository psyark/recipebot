package dancyu

import (
	"context"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/psyark/jsonld"
	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var (
	titleRegex   = regexp.MustCompile(`"(.+?)"`)
	commentRegex = regexp.MustCompile(`^(.+)（(.+)）$`)
	stepRegex    = regexp.MustCompile(`^\d+\s*`)
)

type parser struct{}

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://dancyu.jp/recipe/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rex := &rexch.Recipe{}
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
					rex.Title = name
				}
			}

			for _, image := range rcp2.Image {
				switch image := image.(type) {
				case string:
					rex.Image = sites.ResolvePath(url, strings.TrimSpace(image))
				case *jsonld.ImageObject:
					for _, imgURL := range image.Url {
						if imgURL, ok := imgURL.(string); ok {
							rex.Image = sites.ResolvePath(url, strings.TrimSpace(imgURL))
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
				p.parseIngrediens3(ingredients, rex)
			} else if hasPrefixA > 1 {
				p.parseIngrediens2(ingredients, rex)
			} else {
				p.parseIngrediens1(ingredients, rex)
			}
			return false // 1ページに複数レシピがある場合があるので必ず1個目で中止
		}

		return true
	})

	if parseError != nil {
		return nil, parseError
	}

	// タイトルの引用符内側
	if match := titleRegex.FindStringSubmatch(rex.Title); len(match) == 2 {
		rex.Title = match[1]
	}

	// 手順の画像がJSON-LDに含まれないのでHTMLをパース
	const stepSelector = `.snippet__freearea__subtitle-number,.block__snippet__freeText`
	doc.Find(`div.block__snippet`).ChildrenFiltered(stepSelector).Each(func(i int, s *goquery.Selection) {
		ist := rexch.Instruction{}
		ist.AddText(stepRegex.ReplaceAllString(strings.TrimSpace(s.Text()), ""))
		s.NextUntil(stepSelector).Each(func(i int, s *goquery.Selection) {
			cls := s.AttrOr("class", "")
			if cls == "snippet__freearea__simple-text" {
				ist.AddText(s.Text())
			} else if strings.Contains(cls, "block__snippet__column") {
				s.Find("img").Each(func(i int, s *goquery.Selection) {
					ist.AddImage(sites.ResolvePath(url, s.AttrOr("data-src", "")))
				})
			}
		})
		rex.Instructions = append(rex.Instructions, ist)
	})

	return rex, nil
}

// パターン1
// https://dancyu.jp/recipe/2021_00005129.html
// 中黒は第2レベル、"A" は列挙の開始時に置かれる
func (p *parser) parseIngrediens1(ingredients []string, rex *rexch.Recipe) {
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

		igd := rexch.Ingredient{Name: th, Amount: td}
		if strings.HasPrefix(igd.Name, "・") {
			igd.Name = strings.TrimSpace(strings.TrimPrefix(igd.Name, "・"))
		} else {
			group = ""
		}
		if match := commentRegex.FindStringSubmatch(igd.Amount); len(match) == 3 {
			igd.Amount = match[1]
			igd.Comment = match[2]
		}
		igd.Group = group
		rex.Ingredients = append(rex.Ingredients, igd)
	}
}

// パターン2
// https://dancyu.jp/recipe/2019_00001391.html
// 中黒はトップレベル、"A" は各アイテムに付く
func (p *parser) parseIngrediens2(ingredients []string, rex *rexch.Recipe) {
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

		igd := rexch.Ingredient{Name: th, Amount: td}

		if match := commentRegex.FindStringSubmatch(igd.Amount); len(match) == 3 {
			igd.Amount = match[1]
			igd.Comment = match[2]
		}

		igd.Group = group

		rex.Ingredients = append(rex.Ingredients, igd)
	}
}

// パターン3
// https://dancyu.jp/recipe/2020_00003873.html
// "A" はグループ名に付くが中黒は使わない
func (p *parser) parseIngrediens3(ingredients []string, rex *rexch.Recipe) {
	group := ""
	for _, item := range ingredients {
		pair := strings.SplitN(item, "：", 2)
		th := strings.TrimSpace(pair[0])
		td := strings.TrimSpace(pair[1])

		if strings.HasPrefix(th, "A ") || strings.HasPrefix(th, "B ") || strings.HasPrefix(th, "C ") {
			group = th
		} else {
			igd := rexch.Ingredient{Name: th, Amount: td}

			if match := commentRegex.FindStringSubmatch(igd.Amount); len(match) == 3 {
				igd.Amount = match[1]
				igd.Comment = match[2]
			}

			igd.Group = group
			rex.Ingredients = append(rex.Ingredients, igd)
		}
	}
}

func NewParser() sites.Parser2 {
	return &parser{}
}
