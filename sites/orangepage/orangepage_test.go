package orangepage

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://www.orangepage.net/ymsr/news/daily/posts/5716": {
		Title: "玉ねぎ好きの飛田和緒さん直伝『丸ごと玉ねぎのとろとろカレー』がおいしすぎ！",
		Image: "https://images.orangepage.net/media/article/5716/images/main_626f06d7f88bb6c7d7a2b0815792876c.jpg?d=960x540",
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "玉ねぎ", Amount: "", Comment: "小）……５個（約５００ｇ"},
			{Group: "", Name: "ブロックベーコン……４０ｇ", Amount: "", Comment: ""},
			{Group: "", Name: "にんにくのみじん切り……１かけ分", Amount: "", Comment: ""},
			{Group: "", Name: "しょうがのみじん切り……１かけ分", Amount: "", Comment: ""},
			{Group: "", Name: "カレー粉……小さじ２", Amount: "", Comment: ""},
			{Group: "", Name: "あればガラムマサラ……小さじ１／２", Amount: "", Comment: ""},
			{Group: "", Name: "塩……小さじ１／３～１／２", Amount: "", Comment: ""},
			{Group: "", Name: "しょうゆ……小さじ２", Amount: "", Comment: ""},
			{Group: "", Name: "温かいご飯……適宜", Amount: "", Comment: ""},
			{Group: "", Name: "バター……２０ｇ", Amount: "", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "（1）"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "玉ねぎ1個はみじん切りにする。残りは、根元を薄く切り落とす。ベーコンは5mm角の棒状に切る。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "（2）"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "鍋ににんにく、しょうが、バターを入れ、弱火にかける。香りが立ったらみじん切りの玉ねぎを加えて炒める。途中ふたをして蒸し焼きにしながら、全体に色づくまで15～20分炒める。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "POINT\u3000みじん切りの玉ねぎは、あめ色になるまでじっくり炒めて。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "（3）"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "ベーコンとカレー粉、あればガラムマサラを加えて炒め、カレーの香りが立ったら水300mlと残りの玉ねぎを加える。弱めの中火にして煮立ったらふたをし、玉ねぎが柔らかくなるまで30分ほど煮る。途中、煮汁が煮つまりすぎていたら、適宜水をたす。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "（4）"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "玉ねぎにすっと竹串が通るくらいになったら、味をみてから塩としょうゆを加える。ご飯を器に盛り、カレーをかける。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "メインの具は玉ねぎとベーコンのみと、本当にシンプル。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "けれどちょっとしたひと工夫で、いつものカレーがこんなそそる一皿になるんですねー。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "これは絶対に試してみるべき！"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "『丸ごと玉ねぎのとろとろカレー』ぜひ作ってみて下さいねー♪"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "（『オレンジページCooking2022夕飯\u3000夕飯、即決、迷わない。』より）"},
			}},
		},
	},
	"https://www.orangepage.net/ymsr/news/daily/posts/5763": {
		Title: "【白メシが美味い！ 秋おかず】絶品『れんこんと豚肉の甘辛揚げ』のレシピ",
		Image: "https://images.orangepage.net/media/article/5763/images/main_cf8cf4f59904c49016298e414deccdef.jpg?d=960x540",
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "れんこん……１節", Amount: "", Comment: "約１５０ｇ"},
			{Group: "", Name: "豚ロース肉", Amount: "", Comment: "とんカツ用）……２枚（約２００ｇ"},
			{Group: "", Name: "〈甘辛だれ〉", Amount: "", Comment: ""},
			{Group: "", Name: "砂糖……小さじ２", Amount: "", Comment: ""},
			{Group: "", Name: "酢……小さじ２", Amount: "", Comment: ""},
			{Group: "", Name: "しょうゆ……大さじ１", Amount: "", Comment: ""},
			{Group: "", Name: "酒", Amount: "", Comment: ""},
			{Group: "", Name: "しょうゆ", Amount: "", Comment: ""},
			{Group: "", Name: "小麦粉", Amount: "", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "", Comment: ""},
		},

		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "（1）材料の下ごしらえをする"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "れんこんはよく洗い、皮つきのまま幅1cmの輪切りにする。豚肉は大きめの一口大に切る。ボールに豚肉を入れ、酒、しょうゆ各大さじ1/2をもみ込み、15分おく。汁けをかるくきって小麦粉を薄くまぶす。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "（2）フライパンで揚げる"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "フライパンにサラダ油を高さ2cmほど入れて低温（約160℃。乾いた菜箸の先を底に当てると、細かい泡がゆっくりと揺れながら出る程度）に熱し、れんこんを入れてきつね色になるまで5～6分揚げ、油をきる。油を中温（約170℃。乾いた菜箸の先を底に当てると、細かい泡がシュワシュワッとまっすぐ出る程度）に熱し、豚肉を2～3分揚げ、油をきる。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "（3）たれをからめる"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "ボールにたれの材料を混ぜ合わせ、れんこんと豚肉を加えてからめる。たれごと器に盛る。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "とんカツ用肉を使うので食べごたえも満点！"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "「食欲の秋」という通り、本当に箸が止まらなくなりそうなおいしさです。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "たれをバウンドさせたご飯といっしょにかき込む至福のひとときを、ぜひご体感ください！"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "（『オレンジページ』2022年10月17日号より）"},
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
