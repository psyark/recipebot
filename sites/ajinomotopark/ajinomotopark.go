package ajinomotopark

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

var servingsRegex = regexp.MustCompile(`（(\d+)人分）`)

type parser struct{}

var debrandMap = map[string]string{
	"「AJINOMOTO オリーブオイル」":     "オリーブオイル",
	"「AJINOMOTO ごま油好きの純正ごま油」": "ごま油",
	"「AJINOMOTO さらさらキャノーラ油」":  "サラダ油",
	"「AJINOMOTO サラダ油」":        "サラダ油",
	"「Cook Do」オイスターソース":       "オイスターソース",
	"「Cook Do」熟成豆板醤":          "豆板醤",
	"「丸鶏がらスープ」":               "鶏がらスープ",
	"「瀬戸のほんじお」":               "塩",
	"「味の素KKコンソメ」固形タイプ":        "コンソメ（固形）",
	"「味の素KKコンソメ」顆粒タイプ":        "コンソメ（顆粒）",
	"「味の素KK中華あじ」":             "粉末中華スープ",
	"うま味調味料「味の素®」":            "味の素",
}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	rex, err := p.Parse2(ctx, url)
	if err != nil {
		return nil, err
	}
	return rex.BackCompat(), nil
}

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://park.ajinomoto.co.jp/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rex := &rexch.Recipe{
		Title: strings.TrimSpace(doc.Find(`h1.recipeTitle`).Text()),
		Image: doc.Find(`.recipeImageArea img`).AttrOr("src", ""),
	}

	// "ふっくら焼ける！  \n                ホットケーキ" -> "ホットケーキ"
	if parts := strings.SplitN(rex.Title, "\n", 2); len(parts) == 2 {
		rex.Title = strings.TrimSpace(parts[1])
	}

	if match := servingsRegex.FindStringSubmatch(doc.Find(`.bigTitle_quantity`).Text()); len(match) != 0 {
		i, _ := strconv.Atoi(match[1])
		rex.Servings = i
	}

	doc.Find(`.recipeMaterialList dl dt`).Each(func(i int, s *goquery.Selection) {
		idg := rexch.NewIngredient(debrand(strings.TrimSpace(s.Text())), s.Next().Text())

		if className := s.AttrOr("class", ""); strings.HasPrefix(className, "ico") {
			idg.Group = strings.TrimPrefix(className, "ico")
		}

		rex.Ingredients = append(rex.Ingredients, *idg)
	})

	doc.Find(`#makeList ol li`).Each(func(i int, s *goquery.Selection) {
		s.Find(`.num`).Remove() // 番号を削除
		ist := rexch.Instruction{
			Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: strings.TrimSpace(s.Text())},
			},
		}
		s.Find("img").Each(func(i int, s *goquery.Selection) {
			ist.Elements = append(ist.Elements, &rexch.ImageInstructionElement{URL: s.AttrOr("src", "")})
		})
		rex.Instructions = append(rex.Instructions, ist)
	})

	return rex, nil
}

func debrand(name string) string {
	if debranded, ok := debrandMap[name]; ok {
		return debranded
	}
	return name
}

func NewParser() sites.Parser2 {
	return &parser{}
}
