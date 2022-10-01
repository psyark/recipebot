package united

import (
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
	"github.com/psyark/recipebot/sites/sbfoods"
	"github.com/psyark/recipebot/sites/sirogohan"
)

var Parsers = sites.Parsers{
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
	sbfoods.NewParser(),
	sirogohan.NewParser(),
	jsonld.NewParser(),
}
