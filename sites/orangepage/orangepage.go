package orangepage

import (
	"context"
	"errors"
	"regexp"

	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var (
	servingsRegex    = regexp.MustCompile(`(\d+)(?:～\d+)?人分`)
	instructionRegex = regexp.MustCompile(`^\(\d+\)`)
)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	rex, err := p.Parse2(ctx, url)
	if err != nil {
		return nil, err
	}
	return rex.BackCompat(), nil
}

func (p *parser) Parse2(ctx context.Context, url string) (*rexch.Recipe, error) {
	if rex, err := p.ParseYMSR(ctx, url); !errors.Is(err, sites.ErrUnsupportedURL) {
		return rex, err
	}
	if rex, err := p.ParseRecipes(ctx, url); !errors.Is(err, sites.ErrUnsupportedURL) {
		return rex, err
	}

	return nil, sites.ErrUnsupportedURL
}

func NewParser() sites.Parser2 {
	return &parser{}
}
