package nadia

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/microcosm-cc/bluemonday"
	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"
)

const assetURL = "https://asset.oceans-nadia.com/"

var bmp = bluemonday.StrictPolicy()

type parser struct{}

type nd struct {
	Props struct {
		PageProps struct {
			Data struct {
				PublishedRecipe Recipe `json:"publishedRecipe"`
			} `json:"data"`
		} `json:"pageProps"`
	} `json:"props"`
}
type Recipe struct {
	Title        string        `json:"title"`
	ImageSet     []Image       `json:"imageSet"`
	Ingredients  []Ingredient  `json:"ingredients"`
	Instructions []Instruction `json:"instructions"`
	Tips         string        `json:"tips"`
	// BunryoPeople       interface{} `json:"bunryoPeople"`
	// Comment            interface{} `json:"comment"`
	// CookTime           interface{} `json:"cookTime"`
	// CookTimeMemo       interface{} `json:"cookTimeMemo"`
	// FavoriteCount      interface{} `json:"favoriteCount"`
	// Id                 interface{} `json:"id"`
	// IsSponsorRecipe    interface{} `json:"isSponsorRecipe"`
	// Modified           interface{} `json:"modified"`
	// PrepTime           interface{} `json:"prepTime"`
	// PublishedDate      interface{} `json:"publishedDate"`
	// RecipeType         interface{} `json:"recipeType"`
	// SpecialSite        interface{} `json:"specialSite"`
	// SpecialSiteId      interface{} `json:"specialSiteId"`
	// SponsorRecipe      interface{} `json:"sponsorRecipe"`
	// Tags               interface{} `json:"tags"`
	// User               interface{} `json:"user"`
	// VideoPublishedDate interface{} `json:"videoPublishedDate"`
	// VideoUrl           interface{} `json:"videoUrl"`
	// Yield              interface{} `json:"yield"`
}
type Image struct {
	Filename string `json:"filename"`
	Path     string `json:"path"`
}
type Ingredient struct {
	Kubun  *string `json:"kubun"`
	Name   string  `json:"name"`
	Amount string  `json:"amount"`
	Memo   string  `json:"memo"`
}
type Instruction struct {
	Comment  string `json:"comment"`
	ImageSet Image  `json:"imageSet"`
}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	if !strings.HasPrefix(url, "https://oceans-nadia.com/") {
		return nil, sites.ErrUnsupportedURL
	}

	doc, err := sites.NewDocumentFromURL(ctx, url)
	if err != nil {
		return nil, err
	}

	n := nd{}
	json.Unmarshal([]byte(doc.Find(`#__NEXT_DATA__`).Text()), &n)

	nr := n.Props.PageProps.Data.PublishedRecipe

	rcp := &recipe.Recipe{Title: nr.Title}
	for _, is := range nr.ImageSet {
		rcp.Image = sites.ResolvePath(assetURL, is.Path)
	}
	for _, id := range nr.Ingredients {
		idg := recipe.Ingredient{Name: id.Name, Amount: id.Amount, Comment: id.Memo}
		if id.Kubun != nil {
			rcp.AddIngredient(*id.Kubun, idg)
		} else {
			rcp.AddIngredient("", idg)
		}
	}
	if nr.Tips != "" {
		rcp.Steps = append(rcp.Steps, recipe.Step{Text: nr.Tips})
	}
	for _, in := range nr.Instructions {
		if in.Comment != "" {
			step := recipe.Step{Text: bmp.Sanitize(in.Comment)}
			if in.ImageSet.Path != "" {
				step.Images = append(step.Images, sites.ResolvePath(assetURL, in.ImageSet.Path))
			}
			rcp.Steps = append(rcp.Steps, step)
		}
	}

	return rcp, nil
}

func NewParser() sites.Parser {
	return &parser{}
}
