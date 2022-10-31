package lettuceclub

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://www.lettuceclub.net/recipe/dish/24626/": {
		Title:    "アマトリチャーナ",
		Image:    "https://www.lettuceclub.net/i/R1/img/dish/1/S20170925009002A2_000.jpg?w=450",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "スパゲッティ(1.6mm)", Amount: "160〜200g", Comment: ""},
			{Group: "", Name: "玉ねぎ", Amount: "1/2個", Comment: ""},
			{Group: "", Name: "ベーコン", Amount: "4枚", Comment: ""},
			{Group: "", Name: "オリーブ油", Amount: "大さじ3", Comment: ""},
			{Group: "", Name: "にんにくのみじん切り", Amount: "小さじ2", Comment: ""},
			{Group: "", Name: "ホールトマト缶", Amount: "1缶(約400g)", Comment: ""},
			{Group: "", Name: "塩", Amount: "小さじ1/4", Comment: ""},
			{Group: "", Name: "こしょう", Amount: "少々", Comment: ""},
			{Group: "", Name: "粉チーズ", Amount: "適量", Comment: ""},
			{Group: "", Name: "粗塩", Amount: "大さじ1", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "鍋に湯1.6Lを沸かし、粗塩大さじ1を加える。スパゲッティを加えてさっと混ぜ、袋の表示より1〜2分短くゆでる。玉ねぎは縦薄切りにする。ベーコンは1cm幅に切る。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンにオリーブ油大さじ3、にんにくのみじん切り小さじ2、玉ねぎを入れて中火にかけ、しんなりするまで約2分炒める。ベーコンを加えてさっと炒め、ホールトマト缶を加えて潰す。時々混ぜながら3〜4分煮て、塩小さじ1/4、こしょう少々をふる。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "スパゲッティの湯をきって2に加え、さっとあえる。器に盛り、粉チーズ適量をふる。"}}},
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
