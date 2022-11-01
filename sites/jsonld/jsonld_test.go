package jsonld

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://s.recipe-blog.jp/profile/313934/recipe/1432314": {
		Title:    "自家製ごまダレで、牛肉と水菜の簡単しゃぶしゃぶ",
		Image:    "https://asset.recipe-blog.jp/cache/images/recipe/bc/ae/fe6575effb49833f63fea6b56510cf2f8e21bcae.640x640.cut.jpg",
		Servings: 3,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "牛こま切れ", Amount: "340g", Comment: ""},
			{Group: "", Name: "水菜", Amount: "1束", Comment: ""},
			{Group: "Ａ", Name: "ごま", Amount: "大さじ1", Comment: ""},
			{Group: "Ａ", Name: "ポン酢", Amount: "大さじ1", Comment: ""},
			{Group: "Ａ", Name: "マヨネーズ", Amount: "大さじ1", Comment: ""},
			{Group: "Ａ", Name: "砂糖", Amount: "大さじ1", Comment: ""},
			{Group: "Ａ", Name: "味噌", Amount: "小さじ2分の1", Comment: ""},
			{Group: "", Name: "ポン酢", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "水菜はよく洗い、3㎝幅に切り、熱湯で水菜を１分ほど茹でる。冷水に取り、水気をぎゅっと絞る。牛肉こま切れは、熱湯でさっと茹でて水気を切る。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "１をお皿に盛り付ける。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "Ａをよく混ぜ合わせ、つけダレにする。ポン酢もつけダレにする。"}}},
		},
	},
}

func TestNewParser(t *testing.T) {
	ctx := context.Background()

	for url, want := range tests {
		url := url
		want := want

		t.Run(url, func(t *testing.T) {
			t.Parallel()

			rex, err := NewParser().Parse(ctx, url)
			if err != nil {
				t.Fatal(err)
			}

			if err := sites.RecipeMustBe2(want, rex); err != nil {
				t.Error(err)
			}
		})
	}
}
