package orangepage

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://www.orangepage.net/ymsr/news/daily/posts/5716": {
		Title:    "玉ねぎ好きの飛田和緒さん直伝『丸ごと玉ねぎのとろとろカレー』がおいしすぎ！",
		Image:    "https://images.orangepage.net/media/article/5716/images/main_626f06d7f88bb6c7d7a2b0815792876c.jpg?d=960x540",
		Servings: 2,
		Ingredients: []rexch.Ingredient{
			{Name: "玉ねぎ", Amount: "5個(約500g)", Comment: "小"},
			{Name: "ブロックベーコン", Amount: "40g"},
			{Name: "にんにくのみじん切り", Amount: "1かけ分"},
			{Name: "しょうがのみじん切り", Amount: "1かけ分"},
			{Name: "カレー粉", Amount: "小さじ2"},
			{Name: "あればガラムマサラ", Amount: "小さじ1/2"},
			{Name: "塩", Amount: "小さじ1/3~1/2"},
			{Name: "しょうゆ", Amount: "小さじ2"},
			{Name: "温かいご飯", Amount: "適宜"},
			{Name: "バター", Amount: "20g"},
		},
		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "(1)\n玉ねぎ1個はみじん切りにする。残りは、根元を薄く切り落とす。ベーコンは5mm角の棒状に切る。"},
				&rexch.ImageInstructionElement{URL: "https://images.orangepage.net/media/article/5716/images/e760acc36976d8c8f5c2a530f9b75dda42ce1dc0.jpg?w=1200"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "(2)\n鍋ににんにく、しょうが、バターを入れ、弱火にかける。香りが立ったらみじん切りの玉ねぎを加えて炒める。途中ふたをして蒸し焼きにしながら、全体に色づくまで15~20分炒める。\nPOINT みじん切りの玉ねぎは、あめ色になるまでじっくり炒めて。"},
				&rexch.ImageInstructionElement{URL: "https://images.orangepage.net/media/article/5716/images/0503a1ed88865d0fdd207a2712bc08ee356893a5.jpg?w=1200"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "(3)\nベーコンとカレー粉、あればガラムマサラを加えて炒め、カレーの香りが立ったら水300mlと残りの玉ねぎを加える。弱めの中火にして煮立ったらふたをし、玉ねぎが柔らかくなるまで30分ほど煮る。途中、煮汁が煮つまりすぎていたら、適宜水をたす。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "(4)\n玉ねぎにすっと竹串が通るくらいになったら、味をみてから塩としょうゆを加える。ご飯を器に盛り、カレーをかける。\nメインの具は玉ねぎとベーコンのみと、本当にシンプル。\nけれどちょっとしたひと工夫で、いつものカレーがこんなそそる一皿になるんですねー。\nこれは絶対に試してみるべき!\n『丸ごと玉ねぎのとろとろカレー』ぜひ作ってみて下さいねー♪\n(『オレンジページCooking2022夕飯 夕飯、即決、迷わない。』より)"},
			}},
		},
	},
	"https://www.orangepage.net/ymsr/news/daily/posts/5763": {
		Title:    "【白メシが美味い！ 秋おかず】絶品『れんこんと豚肉の甘辛揚げ』のレシピ",
		Image:    "https://images.orangepage.net/media/article/5763/images/main_cf8cf4f59904c49016298e414deccdef.jpg?d=960x540",
		Servings: 2,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "れんこん", Amount: "1節(約150g)"},
			{Group: "", Name: "豚ロース肉", Amount: "2枚(約200g)", Comment: "とんカツ用"},
			{Group: "", Name: "〈甘辛だれ〉", Amount: ""},
			{Group: "", Name: "砂糖", Amount: "小さじ2"},
			{Group: "", Name: "酢", Amount: "小さじ2"},
			{Group: "", Name: "しょうゆ", Amount: "大さじ1"},
			{Group: "", Name: "酒", Amount: ""},
			{Group: "", Name: "しょうゆ", Amount: ""},
			{Group: "", Name: "小麦粉", Amount: ""},
			{Group: "", Name: "サラダ油", Amount: ""},
		},

		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "(1)材料の下ごしらえをする\nれんこんはよく洗い、皮つきのまま幅1cmの輪切りにする。豚肉は大きめの一口大に切る。ボールに豚肉を入れ、酒、しょうゆ各大さじ1/2をもみ込み、15分おく。汁けをかるくきって小麦粉を薄くまぶす。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "(2)フライパンで揚げる\nフライパンにサラダ油を高さ2cmほど入れて低温(約160℃。乾いた菜箸の先を底に当てると、細かい泡がゆっくりと揺れながら出る程度)に熱し、れんこんを入れてきつね色になるまで5~6分揚げ、油をきる。油を中温(約170℃。乾いた菜箸の先を底に当てると、細かい泡がシュワシュワッとまっすぐ出る程度)に熱し、豚肉を2~3分揚げ、油をきる。"},
				&rexch.ImageInstructionElement{URL: "https://images.orangepage.net/media/article/5763/images/19cfa836913fbdcd96a345a31461def0b4f25862.jpg?w=1200"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "(3)たれをからめる\nボールにたれの材料を混ぜ合わせ、れんこんと豚肉を加えてからめる。たれごと器に盛る。\nとんカツ用肉を使うので食べごたえも満点!\n「食欲の秋」という通り、本当に箸が止まらなくなりそうなおいしさです。\nたれをバウンドさせたご飯といっしょにかき込む至福のひとときを、ぜひご体感ください!\n(『オレンジページ』2022年10月17日号より)"},
			}},
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
