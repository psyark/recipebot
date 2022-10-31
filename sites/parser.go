package sites

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
	"reflect"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/kylelemons/godebug/pretty"
	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/rexch"
	"golang.org/x/sync/errgroup"
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

func RecipeMustBe(got recipe.Recipe, wantStr string) error {
	want := recipe.Recipe{}
	json.Unmarshal([]byte(wantStr), &want)

	if wantStr == "" {
		gotBytes, _ := json.Marshal(got)
		fmt.Println(string(gotBytes))
		return nil
	}

	if !reflect.DeepEqual(want, got) {
		return fmt.Errorf("%w: %v", errUnmatch, pretty.Compare(want, got))
	}
	return nil
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

func MigrateTest(parser Parser2, tests map[string]string) {
	pc, file, _, ok := runtime.Caller(1)
	if !ok {
		panic("caller")
	}

	fn := runtime.FuncForPC(pc)
	packageName := strings.TrimSuffix(path.Base(fn.Name()), path.Ext(path.Base(fn.Name())))
	// file = strings.Replace(file, "_test.go", "_2_test.go", 1)

	f, err := os.OpenFile(file, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	lines := []string{}
	eg := errgroup.Group{}
	ctx := context.Background()

	for url := range tests {
		url := url
		eg.Go(func() error {
			rex, err := parser.Parse2(ctx, url)
			if err != nil {
				return err
			}

			lines = append(lines, fmt.Sprintf("%q: %#v,\n", url, rex))
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		panic(err)
	}

	fmt.Fprintf(f, `package %s

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
%s}

func TestNewParser(t *testing.T) {
	ctx := context.Background()

	for url, want := range tests {
		url := url
		want := want

		t.Run(url, func(t *testing.T) {
			t.Parallel()

			rex, err := NewParser().Parse2(ctx, url)
			if err != nil {
				t.Fatal(err)
			}

			if err := sites.RecipeMustBe2(want, rex); err != nil {
				t.Error(err)
			}
		})
	}
}
`, packageName, strings.Join(lines, ""))

}
