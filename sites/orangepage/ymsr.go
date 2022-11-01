package orangepage

import (
	"context"
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
	"golang.org/x/text/width"
)

func (p *parser) ParseYMSR(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://www.orangepage.net/ymsr/news/daily/posts/") {
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

	// script除去
	body.Find("script").Remove()
	body.Find("style").Remove()

	// brを改行に
	body.Find("br").ReplaceWithHtml("\n")

	// 画像をURLに
	body.Find("img").Each(func(i int, s *goquery.Selection) {
		if src, ok := s.Attr("src"); ok {
			s.ReplaceWithHtml(fmt.Sprintf("\nIMAGE:%s\n", src))
		}
	})

	mode := "intro"
	for _, line := range strings.Split(body.Text(), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			switch mode {
			case "intro":
				if strings.Contains(line, "材料") {
					if servings, ok := sites.ParseServings(line); ok {
						rex.Servings = servings
					}

					if strings.Contains(line, "作り方") { // 材料を飛ばす場合がある
						mode = "instructions"
					} else {
						mode = "ingredients"
					}
				}
			case "ingredients":
				if strings.HasPrefix(line, "IMAGE:") {
					// skip
				} else if strings.Contains(line, "作り方") {
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
					line = width.Fold.String(line)
					if instructionRegex.MatchString(line) || len(rex.Instructions) == 0 {
						rex.Instructions = append(rex.Instructions, rexch.Instruction{})
					}

					lastInst := &rex.Instructions[len(rex.Instructions)-1]
					if strings.HasPrefix(line, "IMAGE:") {
						lastInst.AddImage(strings.TrimPrefix(line, "IMAGE:"))
					} else {
						lastInst.AddText(line)
					}
				}
			}
		}
	}

	return rex, nil
}
