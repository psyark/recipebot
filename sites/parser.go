package sites

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kylelemons/godebug/pretty"
	"github.com/psyark/recipebot/rexch"
)

var (
	errUnmatch        = errors.New("unmatch")
	ErrUnsupportedURL = errors.New("unsupported url")
)

type Parser interface {
	Parse(ctx context.Context, url string) (*rexch.Recipe, error)
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

	defer res.Body.Close()

	return goquery.NewDocumentFromReader(res.Body)
}

func ResolvePath(baseURL, path string) string {
	if strings.HasPrefix(path, "/") {
		u, _ := url.Parse(baseURL)
		return u.Scheme + "://" + u.Host + path
	}
	return path
}

func RecipeMustBe2(want, got *rexch.Recipe) error {
	if want == nil {
		// bytes, _ := json.MarshalIndent(got, "", "  ")
		return fmt.Errorf("%w: %#v", errUnmatch, got)
	}
	if !reflect.DeepEqual(want, got) {
		return fmt.Errorf("%w: %v", errUnmatch, pretty.Compare(want, got))
	}
	return nil
}
