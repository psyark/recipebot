package orangepage

import (
	"context"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

func (p *parser) ParseRecipes(ctx context.Context, url string) (*rexch.Recipe, error) {
	if !strings.HasPrefix(url, "https://www.orangepage.net/recipes/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	rex := &rexch.Recipe{
		Title: strings.TrimSpace(doc.Find(`h1.recipesTitle`).Text()),
		Image: doc.Find(`.recipesDetailImg img`).AttrOr("src", ""),
	}

	if servings, ok := sites.ParseServings(doc.Find(`h2.IngredientsTit`).Text()); ok {
		rex.Servings = servings
	}

	doc.Find("[itemprop=recipeIngredient]").Each(func(i int, s *goquery.Selection) {
		parts := strings.SplitN(s.Text(), "　", 2)
		if len(parts) == 1 {
			parts = append(parts, "")
		}
		rex.Ingredients = append(rex.Ingredients, *rexch.NewIngredient(parts[0], parts[1]))
	})

	doc.Find(".instructionList li").Each(func(i int, s *goquery.Selection) {
		inst := rexch.Instruction{}
		inst.AddText(strings.TrimSpace(s.Text()))
		s.Find("img").Each(func(i int, s *goquery.Selection) {
			src := s.AttrOr("src", "")
			if strings.HasSuffix(src, "_w150hf.jpg") {
				src = strings.TrimSuffix(src, "_w150hf.jpg") + ".jpg"
			}
			inst.AddImage(src)
		})
		rex.Instructions = append(rex.Instructions, inst)
	})

	return rex, nil
}
