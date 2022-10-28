package sites

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kylelemons/godebug/pretty"
	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/rexch"
)

var (
	errUnmatch        = errors.New("unmatch")
	ErrUnsupportedURL = errors.New("unsupported url")
)

type Parser interface {
	Parse(ctx context.Context, url string) (*recipe.Recipe, error)
}

type Parser2 interface {
	Parser
	Parse2(ctx context.Context, url string) (*rexch.Recipe, error)
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

func RecipeMustBe(rcp recipe.Recipe, want string) error {
	got, _ := json.Marshal(rcp)
	if want == "" {
		fmt.Println(string(got))
		return nil
	}

	if want != string(got) {
		return errUnmatch
		// // dmp := diffmatchpatch.New()
		// // diffs := dmp.DiffMain(indent(want), indent(string(got)), false)
		// // diffs = dmp.DiffCleanupSemantic(diffs)

		// return fmt.Errorf("%w: %v", errUnmatch, DiffPrettyText(diffs))
	}
	return nil
}

func RecipeMustBe2(want, got *rexch.Recipe) error {
	if !reflect.DeepEqual(want, got) {
		return fmt.Errorf("%w: %v", errUnmatch, pretty.Compare(want, got))
	}
	return nil
}
