package sites

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/psyark/recipebot/recipe"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var (
	errUnmatch        = errors.New("unmatch")
	ErrUnsupportedURL = errors.New("unsupported url")
)

type Parser interface {
	Parse(ctx context.Context, url string) (*recipe.Recipe, error)
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
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(indent(want), indent(string(got)), false)
		diffs = dmp.DiffCleanupSemantic(diffs)

		return fmt.Errorf("%w: %v", errUnmatch, dmp.DiffPrettyText(diffs))
	}
	return nil
}

func indent(src string) string {
	var x recipe.Recipe
	json.Unmarshal([]byte(src), &x)
	dst, _ := json.MarshalIndent(x, "", "  ")
	return string(dst)
}
