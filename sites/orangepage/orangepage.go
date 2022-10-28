package orangepage

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

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	rex, err := p.Parse2(ctx, url)
	if err != nil {
		return nil, err
	}
	return rex.BackCompat(), nil
}

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://www.orangepage.net/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rex := &rexch.Recipe{
		Title: strings.TrimSpace(doc.Find(`h1.articleTitle`).Text()),
		Image: doc.Find(`.articleDetailImg img`).AttrOr("src", ""),
	}

	body := doc.Find("#opDailyBody")

	// brを改行に
	body.Find("br").Each(func(i int, s *goquery.Selection) {
		s.ReplaceWithHtml("\n")
	})

	mode := "intro"
	for _, line := range strings.Split(body.Text(), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			switch mode {
			case "intro":
				if strings.Contains(line, "材料") {
					// fmt.Println(mode, line)
					mode = "ingredients"
				}
			case "ingredients":
				if strings.Contains(line, "作り方") {
					mode = "instructions"
				} else {
					parts := strings.SplitN(line, "……", 2)
					if len(parts) == 1 {
						parts = append(parts, "")
					}
					rex.Ingredients = append(rex.Ingredients, *rexch.NewIngredient(parts[0], parts[1]))
				}
			case "instructions":
				if strings.Contains(line, "関連記事") {
					mode = "outro"
				} else {
					inst := rexch.Instruction{Elements: []rexch.InstructionElement{
						&rexch.TextInstructionElement{Text: line},
					}}
					rex.Instructions = append(rex.Instructions, inst)
				}
			}
		}
	}

	if match := servingsRegex.FindStringSubmatch(doc.Find(`.bigTitle_quantity`).Text()); len(match) != 0 {
		i, _ := strconv.Atoi(match[1])
		rex.Servings = i
	}

	return rex, nil
}

func NewParser() sites.Parser2 {
	return &parser{}
}
