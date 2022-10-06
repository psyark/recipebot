package common

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/psyark/recipebot/recipe"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var errUnmatch = errors.New("unmatch")

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
