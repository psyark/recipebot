package united

import (
	"context"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites"
	"github.com/psyark/recipebot/sites/ajinomotopark"
	"github.com/psyark/recipebot/sites/buzzfeed"
	"github.com/psyark/recipebot/sites/cookpad"
	"github.com/psyark/recipebot/sites/dancyu"
	"github.com/psyark/recipebot/sites/delishkitchen"
	"github.com/psyark/recipebot/sites/foodie"
	"github.com/psyark/recipebot/sites/jsonld"
	"github.com/psyark/recipebot/sites/kikkoman"
	"github.com/psyark/recipebot/sites/kurashiru"
	"github.com/psyark/recipebot/sites/lettuceclub"
	"github.com/psyark/recipebot/sites/macaroni"
	"github.com/psyark/recipebot/sites/nadia"
	"github.com/psyark/recipebot/sites/orangepage"
	"github.com/psyark/recipebot/sites/sbfoods"
	"github.com/psyark/recipebot/sites/sirogohan"
)

func NewParser() sites.Parser {
	return &parsers{
		ajinomotopark.NewParser(),
		buzzfeed.NewParser(),
		cookpad.NewParser(),
		dancyu.NewParser(),
		delishkitchen.NewParser(),
		foodie.NewParser(),
		kikkoman.NewParser(),
		kurashiru.NewParser(),
		lettuceclub.NewParser(),
		macaroni.NewParser(),
		nadia.NewParser(),
		orangepage.NewParser(),
		sbfoods.NewParser(),
		sirogohan.NewParser(),
		jsonld.NewParser(), // 最後
	}
}

type parsers []sites.Parser

func (p parsers) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	for _, c := range p {
		rcp, err := c.Parse(ctx, url)
		switch err {
		case nil:
			return rcp, nil
		case sites.ErrUnsupportedURL:
			continue
		default:
			return nil, err
		}
	}
	return nil, sites.ErrUnsupportedURL
}
