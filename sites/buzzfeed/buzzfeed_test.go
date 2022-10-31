package buzzfeed

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://www.buzzfeed.com/jp/kazuyawakana/eggplant-ooba-mentsuyu-pickled-doubanjiang": {
		Title:    "ご飯がススム！なすと大葉のめんつゆ豆板醤漬け",
		Image:    "https://img.buzzfeed.com/buzzfeed-static/static/2022-05/27/6/asset/5c946aa4fc12/sub-buzz-3082-1653632598-1.jpg?downsize=700%3A%2A&output-quality=auto&output-format=auto",
		Servings: 1,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "なす", Amount: "2本", Comment: ""},
			{Group: "", Name: "大葉", Amount: "10枚", Comment: ""},
			{Group: "A", Name: "しょうがすりおろし", Amount: "小さじ1", Comment: ""},
			{Group: "A", Name: "ごま油", Amount: "大さじ1", Comment: ""},
			{Group: "A", Name: "めんつゆ（2倍濃縮）大さじ3", Amount: "", Comment: ""},
			{Group: "A", Name: "豆板醤", Amount: "小さじ2", Comment: ""},
			{Group: "A", Name: "白ごま", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "なすのへたをとり、縦に薄くスライスする。皿に並べラップをし、600Wの電子レンジで3分加熱する。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ボウルにAを入れよく混ぜる。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "保存容器に①のなすを並べ、その上に大葉を乗せる。何層にも重ねたら②をかけ冷蔵庫で1時間寝かせたら、完成！"}}},
		},
	},
	"https://www.buzzfeed.com/jp/kazuyawakana/burdock-pickled-in-ponzu-sauce": {
		Title:    "簡単さっぱり♪ごぼうのポン酢漬け",
		Image:    "https://img.buzzfeed.com/buzzfeed-static/static/2022-06/14/2/asset/0d424cf2ac90/sub-buzz-1340-1655172520-17.jpg?downsize=700%3A%2A&output-quality=auto&output-format=auto",
		Servings: 1,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "ごぼう", Amount: "1/4本", Comment: ""},
			{Group: "A", Name: "ポン酢", Amount: "大さじ2", Comment: ""},
			{Group: "A", Name: "砂糖", Amount: "大さじ1/2", Comment: ""},
			{Group: "A", Name: "たかのつめ輪切り", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ごぼうの皮を丸めたアルミホイルでそぎ落とす。よく洗い、食べやすい大きさに切る。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "水を入れた鍋に①を入れ、４分ゆでる。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ボウルに②とAを入れよく和える。冷蔵庫で1時間寝かせたら、完成！"}}},
		},
	},
	"https://www.buzzfeed.com/jp/maorikato/easy-pickled-eggplant": {
		Title:    "悪魔のなす漬け",
		Image:    "https://img.buzzfeed.com/buzzfeed-static/static/2022-06/13/6/asset/159c3b3a29f7/sub-buzz-7518-1655103572-17.jpg?downsize=700%3A%2A&output-quality=auto&output-format=auto",
		Servings: 1,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "なす1本", Amount: "", Comment: ""},
			{Group: "", Name: "ごま油", Amount: "小さじ1", Comment: ""},
			{Group: "A", Name: "長ねぎ", Amount: "1/4本", Comment: ""},
			{Group: "A", Name: "醤油", Amount: "大さじ1", Comment: ""},
			{Group: "A", Name: "水", Amount: "大さじ1", Comment: ""},
			{Group: "A", Name: "砂糖", Amount: "小さじ1\u3000", Comment: ""},
			{Group: "A", Name: "にんにく（すりおろし）", Amount: "小さじ1/2", Comment: ""},
			{Group: "A", Name: "輪切り唐辛子", Amount: "適量", Comment: ""},
			{Group: "A", Name: "白ごま", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "なすを縦半分に切り、皮面に斜めに浅く切り込みを入れ、3等分に切る。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンにごま油をひき、①をしんなりするまで焼く。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "保存袋にAと②を入れて1時間置いたら、完成！"}}},
		},
	},
	"https://www.buzzfeed.com/jp/maorikato/fried-shiso-leaf": {
		Title:    "大葉唐揚げ",
		Image:    "https://img.buzzfeed.com/buzzfeed-static/static/2022-06/10/3/asset/159c3b3a29f7/sub-buzz-2812-1654832843-1.jpg?downsize=700%3A%2A&output-quality=auto&output-format=auto",
		Servings: 1,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "鶏もも肉", Amount: "150g", Comment: ""},
			{Group: "A", Name: "しょうゆ", Amount: "小さじ2", Comment: ""},
			{Group: "A", Name: "みりん", Amount: "小さじ2", Comment: ""},
			{Group: "A", Name: "にんにく（すりおろし）", Amount: "小さじ1/2", Comment: ""},
			{Group: "", Name: "大葉", Amount: "5枚", Comment: ""},
			{Group: "", Name: "片栗粉", Amount: "適量", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "鶏もも肉を一口大に切る。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ボールに①、Aを入れて15分置く。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "水気を切った②に大葉を巻いて片栗粉をまぶし、フライパンにサラダ油を熱し、カリッとするまで3〜4分揚げたら、完成！"}}},
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
