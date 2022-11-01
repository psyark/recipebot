package delishkitchen

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://delishkitchen.tv/recipes/236903854006862303": {
		Title:    "牛すじ肉の下処理",
		Image:    "https://image.delishkitchen.tv/recipe/236903854006862303/1.jpg?version=1624426802&w=460",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "牛すじ肉", Amount: "300g", Comment: ""},
			{Group: "", Name: "長ねぎ[青い部分]", Amount: "1本分", Comment: ""},
			{Group: "", Name: "しょうが(薄切り)", Amount: "3枚", Comment: ""},
			{Group: "", Name: "酒", Amount: "50cc", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "鍋に牛すじ肉、牛すじ肉が浸かるくらいの水(分量外:適量)を入れてわかし、弱火で10分程煮る。"}, &rexch.ImageInstructionElement{URL: "https://media.delishkitchen.tv/recipe/236903854006862303/steps/1.jpg?version=1647222138"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "湯を切り、アクと一緒に余分な脂とアクを流水で洗い流し、食べやすい大きさに切る。"}, &rexch.ImageInstructionElement{URL: "https://media.delishkitchen.tv/recipe/236903854006862303/steps/2.jpg?version=1647222138"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "鍋も一度洗い、牛すじ肉を戻し入れ、牛すじ肉が浸かるくらいの水(分量外:適量)、ねぎ、しょうが、酒を入れてわかし、アクを取り除きながら1時間程弱火で煮る。"}, &rexch.ImageInstructionElement{URL: "https://media.delishkitchen.tv/recipe/236903854006862303/steps/3.jpg?version=1647222138"}}},
		},
	},
	"https://delishkitchen.tv/recipes/148173434692567529": {
		Title:    "鶏むね肉と夏野菜の酢豚風",
		Image:    "https://image.delishkitchen.tv/recipe/148173434692567529/1.jpg?version=1624338242&w=460",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "鶏むね肉", Amount: "1枚(300g)", Comment: ""},
			{Group: "", Name: "ピーマン", Amount: "2個", Comment: ""},
			{Group: "", Name: "なす", Amount: "1本", Comment: ""},
			{Group: "", Name: "玉ねぎ", Amount: "1/4個", Comment: ""},
			{Group: "", Name: "酒", Amount: "大さじ1", Comment: ""},
			{Group: "", Name: "片栗粉", Amount: "大さじ1", Comment: ""},
			{Group: "", Name: "ごま油", Amount: "大さじ3", Comment: ""},
			{Group: "☆たれ", Name: "酒", Amount: "大さじ1", Comment: ""},
			{Group: "☆たれ", Name: "砂糖", Amount: "小さじ1", Comment: ""},
			{Group: "☆たれ", Name: "酢", Amount: "大さじ2", Comment: ""},
			{Group: "☆たれ", Name: "しょうゆ", Amount: "大さじ1", Comment: ""},
			{Group: "☆たれ", Name: "ケチャップ", Amount: "大さじ2", Comment: ""},
			{Group: "☆たれ", Name: "鶏ガラスープの素", Amount: "小さじ1", Comment: ""},
			{Group: "☆たれ", Name: "片栗粉", Amount: "小さじ1", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "鶏むね肉は食べやすい大きさに切る。ボウルに入れて酒、片栗粉を加えてもみこむ。"}, &rexch.ImageInstructionElement{URL: "https://media.delishkitchen.tv/recipe/148173434692567529/steps/1.jpg?version=1624338235"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ピーマンは半分に切って種とわたを取り除き、食べやすい大きさに切る。なすはヘタをとり、食べやすい大きさに切る。玉ねぎは放射状に3等分に切る。"}, &rexch.ImageInstructionElement{URL: "https://media.delishkitchen.tv/recipe/148173434692567529/steps/2.jpg?version=1624338235"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "別のボウルに☆を入れて混ぜる(たれ)。"}, &rexch.ImageInstructionElement{URL: "https://media.delishkitchen.tv/recipe/148173434692567529/steps/3.jpg?version=1624338235"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンにごま油を入れて熱し、なす、玉ねぎを入れてしんなりするまで中火で炒め、ピーマンを加えて油がなじむ程度に炒めて取り出す。"}, &rexch.ImageInstructionElement{URL: "https://media.delishkitchen.tv/recipe/148173434692567529/steps/4.jpg?version=1624338235"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "同じフライパンに鶏むね肉を入れ、焼き色がつくまで中火で焼き、上下を返し、弱火で肉に火が通るまで焼く。"}, &rexch.ImageInstructionElement{URL: "https://media.delishkitchen.tv/recipe/148173434692567529/steps/5.jpg?version=1624338235"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "野菜を戻し入れ、たれを加えてとろみがつくまで炒める。"}, &rexch.ImageInstructionElement{URL: "https://media.delishkitchen.tv/recipe/148173434692567529/steps/6.jpg?version=1624338235"}}},
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
