package foodie

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://mi-journey.jp/foodie/62677/": {
		Title:    "まっすぐなエビフライの作り方。プリッとジューシーに仕上げるプロのコツ",
		Image:    "https://mi-journey.jp/foodie/wp-content/uploads/2020/06/200404ebifurai1.jpg",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "えび", Amount: "8尾※ブラックタイガーなど", Comment: "無頭"},
			{Group: "", Name: "片栗粉、塩、こしょう", Amount: "各少々", Comment: ""},
			{Group: "", Name: "小麦粉、溶き卵、生パン粉、揚げ油", Amount: "各適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "えびの殻を剥く"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "背わたを除く"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "塩、片栗粉で軽く揉む"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "水洗いし、水気を拭く"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "尾先と剣先を切って水を出す"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "腹に切り目を入れる"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "塩、こしょうをふり、20分ほどおく"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "卵液を作る"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "衣をつける"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "パン粉をつける"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "揚げる"}}},
		},
	},
	"https://mi-journey.jp/foodie/59727/": {
		Title:    "鮭のムニエルのレシピ～洋食店のように仕上げる焼き方のコツ 【シェフ直伝】",
		Image:    "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere12.jpg",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "生鮭", Amount: "2切れ", Comment: ""},
			{Group: "", Name: "塩", Amount: "少々", Comment: ""},
			{Group: "", Name: "小麦粉", Amount: "適量", Comment: ""},
			{Group: "", Name: "オリーブオイル", Amount: "大さじ1", Comment: ""},
			{Group: "", Name: "バター", Amount: "15g", Comment: ""},
			{Group: "【ソース】", Name: "バター", Amount: "60g", Comment: ""},
			{Group: "【ソース】", Name: "しょうゆ", Amount: "小さじ2", Comment: ""},
			{Group: "【ソース】", Name: "にんにく", Amount: "1/2かけ分", Comment: "みじん切り"},
			{Group: "【ソース】", Name: "トマト", Amount: "1/2個分(約30g)", Comment: "湯むきして種を除き、角切りにしたもの"},
			{Group: "【ソース】", Name: "レモン", Amount: "1/4個分(約7g)", Comment: "種と薄皮を除き角切りにしたもの"},
			{Group: "【ソース】", Name: "ケイパー", Amount: "大さじ2", Comment: "酢漬け"},
			{Group: "【ソース】", Name: "パセリ", Amount: "各適量", Comment: "みじん切り）、レモンの皮（好みで"},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "①鮭の両面に塩をふり、10分ほどおく"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "②ペーパータオルで水けをふき取る"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "③鮭に小麦粉をまんべんなくまぶす"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "④冷たいフライパンにオリーブオイルを入れ、鮭を皮目から入れる"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "⑤弱火にかけ、鮭をじっくり焼く"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "⑥皮目に焼き色がついたら身を焼く"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "⑦バターを加える"}}},
		},
	},
	"https://mi-journey.jp/foodie/57058/": {
		Title:    "【シェフ直伝】オムレツのレシピ ホテルのように美しく作るコツ",
		Image:    "https://mi-journey.jp/foodie/wp-content/uploads/2019/04/190413omelette1_.jpg",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "卵", Amount: "2〜3個", Comment: ""},
			{Group: "", Name: "バター", Amount: "10g", Comment: ""},
			{Group: "", Name: "塩、こしょう", Amount: "各少々", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "卵液を作る"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "卵液を漉す"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンを火にかけ、バターを溶かす"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "卵液を流し入れる"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ゴムべらで混ぜながら、半熟状になるまで火を通す"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "濡れ布巾の上でフライパンを叩く"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンにこびりついた卵の端を処理する"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "手前から卵を包む"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "反対側の卵を包む"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンを奥に寄せて成形する"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ゴムべらでひっくり返し、強火にかけ、表面を固める"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "赤ワインを煮詰めてソースを作る"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ケチャップを加える"}}},
		},
	},
	"https://mi-journey.jp/foodie/52916/": {
		Title:    "【シェフ直伝】本格リゾットのレシピ。生米をアルデンテに仕上げるテクニック",
		Image:    "https://mi-journey.jp/foodie/wp-content/uploads/2018/10/1810_28_top.jpg",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "米", Amount: "1合(150g)", Comment: ""},
			{Group: "", Name: "オリーブオイル", Amount: "大さじ2", Comment: ""},
			{Group: "", Name: "ブイヨン", Amount: "3カップ※顆粒ブイヨンを湯に溶かしたもの", Comment: ""},
			{Group: "", Name: "ローリエ", Amount: "2枚", Comment: ""},
			{Group: "", Name: "白ワイン", Amount: "大さじ1", Comment: ""},
			{Group: "", Name: "バター", Amount: "20g", Comment: ""},
			{Group: "", Name: "パルミジャーノ・レッジャーノチーズ", Amount: "大さじ2", Comment: ""},
			{Group: "", Name: "塩", Amount: "少々", Comment: ""},
			{Group: "", Name: "粗びき黒こしょう", Amount: "少々", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "オリーブオイルで米を炒めながら、隣でブイヨンに白ワインとあればローリエを加えて温める"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "お玉2杯分のブイヨンと白ワインを入れて、米を炊き始める"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ブイヨンをつぎ足しながら、20分炊く"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "米がふくらみ、粒が立って、艶々してきたら、味見をする"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "火を止めてから、バターとパルミジャーノチーズを加えて余熱で混ぜ、塩で味を調えてから、器に盛り、粗びき黒こしょうをふる。"}}},
		},
	},
	"https://mi-journey.jp/foodie/20667/": {
		Title:    "お家で簡単！フライパンで本格パエリアレシピ ～プロと家庭はここが違った！作り方の3つのポイント",
		Image:    "https://mi-journey.jp/foodie/wp-content/uploads/2016/02/0303_paella_top.jpg",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "いか", Amount: "小1パイ分(100g)", Comment: ""},
			{Group: "", Name: "白身魚", Amount: "小1枚(100g)", Comment: "切り身"},
			{Group: "", Name: "玉ねぎ", Amount: "40g", Comment: ""},
			{Group: "", Name: "にんじん", Amount: "20g", Comment: ""},
			{Group: "", Name: "セロリ", Amount: "20g", Comment: ""},
			{Group: "", Name: "にんにく", Amount: "2片", Comment: ""},
			{Group: "", Name: "トマト缶", Amount: "1/2カップ(100g)", Comment: "ホール"},
			{Group: "", Name: "有頭えび", Amount: "5尾(200g)", Comment: ""},
			{Group: "", Name: "あさり", Amount: "15~16個(100g)", Comment: "砂抜き済"},
			{Group: "", Name: "水", Amount: "600ml(※)", Comment: ""},
			{Group: "", Name: "塩", Amount: "小さじ1/2", Comment: ""},
			{Group: "", Name: "サフラン", Amount: "10本", Comment: ""},
			{Group: "", Name: "米", Amount: "1合(180ml)", Comment: ""},
			{Group: "", Name: "レモン、イタリアンパセリ", Amount: "お好みで", Comment: ""},
			{Group: "", Name: "オリーブ油", Amount: "1/4カップ(50ml)", Comment: ""},
		},
		Instructions: []rexch.Instruction{},
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
