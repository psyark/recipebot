package cookpad

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://cookpad.com/recipe/2032296": {
		Title:    "トロトロ！牛すじカレー",
		Image:    "https://img.cpcdn.com/recipes/2032296/m/27f0f55f9acb7eb68be9ecdadc8efd2a?u=1252112&p=1353214123",
		Servings: 0, // TODO
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "牛すじ肉", Amount: "500g", Comment: ""},
			{Group: "", Name: "玉ねぎ", Amount: "1個", Comment: "大"},
			{Group: "", Name: "ソフリット", Amount: "適量", Comment: "ビーフシチュー・イタリアンのレシピを参照して下さい"},
			{Group: "", Name: "赤ワイン", Amount: "200cc", Comment: ""},
			{Group: "", Name: "水", Amount: "1500cc", Comment: ""},
			{Group: "", Name: "バター", Amount: "大1", Comment: ""},
			{Group: "", Name: "昆布だし", Amount: "1スティック", Comment: "顆粒"},
			{Group: "", Name: "トマト缶", Amount: "1缶", Comment: ""},
			{Group: "", Name: "ローリエの葉", Amount: "3枚", Comment: ""},
			{Group: "", Name: "カレールー", Amount: "6キューブ", Comment: ""},
			{Group: "", Name: "カレー粉", Amount: "大2", Comment: ""},
			{Group: "", Name: "ウスターソース", Amount: "大1", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "写真の材料を用意して（ソフリットが無ければ玉ねぎ大2個をスライスにし、人参１本をさいの目切りにして下さい）"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "オリーブオイル（サラダ油でも可）大２をひいて牛すじを入れ軽く塩コショウして焦げ目が付くまでしっかり焼いて下さい。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "赤ワインを入れてアルコール分を飛ばします。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "鍋に移してソフリット（無ければ人参）を入れ赤ワインを煮詰めます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "煮詰めてる間に、ワインを拭いたフライパンにバター大１を入れ玉ねぎを炒めます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "玉ねぎも鍋に入れて良く混ぜて更に煮詰めて乳化させます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "水を入れて沸騰したら灰汁を綺麗に取って下さい。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "昆布ダシとトマト缶とローリエの葉を入れて蓋をせずに、2時間弱火でコトコト煮込みます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "一旦火を止めカレールーとカレー粉、ウスターソースを入れて更に１０分煮込みます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "味見をして好みの辛さに調整して完成！"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "白ご飯に乗せて牛スジカレーの完成"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "白ご飯の上に半熟卵を乗せて（簡単トロトロ！オムライス参照して下さい）オムカレーの完成！"}}},
		},
	},
	"https://cookpad.com/recipe/1948575": {
		Title:    "ビーフシチュー・イタリアン",
		Image:    "https://img.cpcdn.com/recipes/1948575/m/a76fc44f26fe07abfdf08902f966df60?u=1252112&p=1347192368",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "牛肉（バラ）有ればスネ肉", Amount: "1キロ", Comment: ""},
			{Group: "", Name: "玉ねぎ", Amount: "1個", Comment: "みじん切り"},
			{Group: "", Name: "人参", Amount: "2本", Comment: "みじん切り"},
			{Group: "", Name: "セロリ", Amount: "1本", Comment: "みじん切り"},
			{Group: "", Name: "赤ワイン", Amount: "500cc", Comment: ""},
			{Group: "", Name: "トマト缶", Amount: "1個", Comment: ""},
			{Group: "", Name: "塩コショウ", Amount: "適量", Comment: ""},
			{Group: "", Name: "小麦粉", Amount: "適量", Comment: "薄力粉"},
			{Group: "", Name: "粉チーズ", Amount: "お好みで", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "まずはソフリットを作ります。フライパンにみじん切りにした玉ねぎ、人参、セロリをオリーブオイル（大３）で炒めます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "水分が無くなるまで炒めます。本格的なソフリットは形が無くなるまで炒めますが、煮込んだら一緒なのでこの程度でもＯＫです。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "多めに作ってタッパーで冷凍すれば他の料理でも使えます。今回は半分を使いました。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "牛肉を大きめにカットして塩コショウし小麦粉を振ります。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "オリーブオイルで全面に香ばしく焼き色を付けます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "別の鍋に肉を移し肉汁が残ったフライパンに赤ワインを入れアルコールを飛ばします。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "肉にソフリットと赤ワインを入れ強火で鍋を揺すりながら煮詰めていきます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "これくらいまで煮詰めます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "水をヒタヒタになるまで入れ（約１０００cc）トマト缶を入れて強火で熱し沸騰したら弱火で蓋をせずにコトコト煮ます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "2時間以上煮たらこれくらいまで水分が無くなります。塩コショウで味を整えて好きな柔らかさになったら出来上がり。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "器に移してお好みで粉チーズを振りかけて下さい。"}}},
		},
	},
	"https://cookpad.com/recipe/1885344": {
		Title:    "マーマレードですっきり甘☆豚のスペアリブ",
		Image:    "https://img.cpcdn.com/recipes/1885344/m/d7fdaff65e9e0694d16801432cc6ea89?u=529143&p=1545602123",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "豚のスペアリブ", Amount: "400~500g", Comment: "骨付き肉"},
			{Group: "", Name: "★オレンジマーマレード", Amount: "80g", Comment: ""},
			{Group: "", Name: "★しょうゆ", Amount: "40~50cc", Comment: "こいくち"},
			{Group: "", Name: "★砂糖", Amount: "大さじ1~", Comment: ""},
			{Group: "", Name: "※水", Amount: "150~200cc", Comment: ""},
			{Group: "", Name: "☆サラダ油", Amount: "適宜", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "厚手の鍋にサラダ油を熱し、豚の骨付き肉の表面に焼き目をつけます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "※の水を加えて、煮立ったらアクをとります。★のマーマレードとしょうゆを加えて中火で煮ます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "煮汁が半分くらいになったら、味をみて、砂糖を加えてください。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "煮汁がほとんどなくなるまで、煮からめながら、照りがでるように仕上げます。"}}},
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
