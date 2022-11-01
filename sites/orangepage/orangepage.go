package orangepage

import (
	"context"
	"errors"
	"regexp"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var (
	instructionRegex = regexp.MustCompile(`^\(\d+\)`)
)

type parser struct{}

func (p *parser) Parse(ctx context.Context, url string) (*rexch.Recipe, error) {
	if rex, err := p.ParseYMSR(ctx, url); !errors.Is(err, sites.ErrUnsupportedURL) {
		return rex, err
	}
	if rex, err := p.ParseRecipes(ctx, url); !errors.Is(err, sites.ErrUnsupportedURL) {
		return rex, err
	}

	return nil, sites.ErrUnsupportedURL
}

func NewParser() sites.Parser {
	return &parser{}
}
