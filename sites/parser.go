package sites

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/rexch"
	"github.com/sergi/go-diff/diffmatchpatch"
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
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(indent(want), indent(string(got)), false)
		diffs = dmp.DiffCleanupSemantic(diffs)

		return fmt.Errorf("%w: %v", errUnmatch, DiffPrettyText(diffs))
	}
	return nil
}

func RecipeMustBe2(rex *rexch.Recipe, want string) error {
	got, _ := json.Marshal(rex)
	if want == "" {
		fmt.Println(string(got))
		return nil
	}

	if want != string(got) {
		dmp := diffmatchpatch.New()
		diffs := dmp.DiffMain(indent2(want), indent2(string(got)), false)
		diffs = dmp.DiffCleanupSemantic(diffs)

		return fmt.Errorf("%w: %v", errUnmatch, DiffPrettyText(diffs))
	}
	return nil
}

func indent(src string) string {
	var x recipe.Recipe
	json.Unmarshal([]byte(src), &x)
	dst, _ := json.MarshalIndent(x, "", "  ")
	return string(dst)
}

func indent2(src string) string {
	var x rexch.Recipe
	json.Unmarshal([]byte(src), &x)
	dst, _ := json.MarshalIndent(x, "", "  ")
	return string(dst)
}

func DiffPrettyText(diffs []diffmatchpatch.Diff) string {
	var buff bytes.Buffer
	for _, diff := range diffs {
		text := diff.Text

		switch diff.Type {
		case diffmatchpatch.DiffInsert:
			_, _ = buff.WriteString("🍣")
			_, _ = buff.WriteString(text)
			_, _ = buff.WriteString("🍣")
		case diffmatchpatch.DiffDelete:
			_, _ = buff.WriteString("🍰")
			_, _ = buff.WriteString(text)
			_, _ = buff.WriteString("🍰")
		case diffmatchpatch.DiffEqual:
			_, _ = buff.WriteString(text)
		}
	}

	return buff.String()
}
