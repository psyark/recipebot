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
	"github.com/dave/jennifer/jen"
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

func MigrateTest(packageName string, tests map[string]string) {
	f := jen.NewFile(packageName)

	{
		dict := jen.Dict{}
		for url := range tests {
			dict[jen.Lit(url)] = jen.Nil()
		}
		f.Var().Id("tests").Op("=").Map(jen.String()).Op("*").Qual("github.com/psyark/recipebot/rexch", "Recipe").Block(dict)
	}

	f.Save(packageName + "_2_test.go")

	// func TestNewParser(t *testing.T) {
	// 	ctx := context.Background()

	// 	for url, want := range tests {
	// 		url := url
	// 		want := want

	// 		t.Run(url, func(t *testing.T) {
	// 			t.Parallel()

	// 			rex, err := NewParser().Parse2(ctx, url)
	// 			if err != nil {
	// 				t.Fatal(err)
	// 			}

	// 			if err := sites.RecipeMustBe2(want, rex); err != nil {
	// 				t.Error(err)
	// 			}
	// 		})
	// 	}
	// }

}
