package sites

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/psyark/recipebot/recipe"

	"github.com/PuerkitoBio/goquery"
)

var ErrUnsupportedURL = fmt.Errorf("unsupported url")

type Parser interface {
	Parse(ctx context.Context, url string) (*recipe.Recipe, error)
}

type Parsers []Parser

func (p Parsers) Parse(ctx context.Context, url string) (*recipe.Recipe, error) {
	for _, c := range p {
		rcp, err := c.Parse(ctx, url)
		switch err {
		case nil:
			return rcp, nil
		case ErrUnsupportedURL:
			continue
		default:
			return nil, err
		}
	}
	return nil, ErrUnsupportedURL
}

func NewDocumentFromURL(ctx context.Context, url string) (*goquery.Document, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return goquery.NewDocumentFromResponse(res)
}

func ResolvePath(baseURL, path string) string {
	if strings.HasPrefix(path, "/") {
		u, _ := url.Parse(baseURL)
		return u.Scheme + "://" + u.Host + path
	}
	return path
}
