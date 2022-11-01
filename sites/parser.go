package sites

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kylelemons/godebug/pretty"
	"github.com/psyark/recipebot/rexch"
	"golang.org/x/text/width"
)

var (
	errUnmatch        = errors.New("unmatch")
	ErrUnsupportedURL = errors.New("unsupported url")
	servingsRegex     = regexp.MustCompile(`(\d+)([~〜]\d+)?\s*(?:人分|servings)`)
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

func ParseServings(src string) (int, bool) {
	src = width.Fold.String(src)
	if match := servingsRegex.FindStringSubmatch(src); len(match) != 0 {
		i, _ := strconv.Atoi(match[1])
		return i, true
	}
	return 0, false
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
