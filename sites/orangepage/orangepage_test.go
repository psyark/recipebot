package orangepage

import (
	"context"
	"fmt"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://www.orangepage.net/ymsr/news/daily/posts/5552": {
		Title:       "【レシピあり】豆腐の『めんつゆ漬け』が簡単＆おいしすぎる！",
		Image:       "https://images.orangepage.net/media/article/5552/images/main_e58d54aa086fef082e40c486dda80244.jpg?d=960x540",
		Servings:    2,
		Ingredients: nil,
		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "(1)めんつゆを作る。鍋にだし汁1カップを入れて中火で煮立てる。みりん1/3カップ、しょうゆ1/2カップを加えてひと煮立ちさせ、火を止めてさます。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "(2)絹ごし豆腐1丁(300〜350g)は横半分に切り、保存容器に入れる。殻をむいた半熟ゆで卵(沸騰してから7分ゆでたもの)2個を加え、(1)を注ぐ。ふたをして冷蔵庫に入れ、半日ほど漬ける。\n※豆腐と卵がめんつゆにしっかり浸かる大きさの保存容器に入れてください。残った漬け汁は煮ものなどに使えます。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "(3)器に温かいご飯どんぶり2杯分(360〜400g)を等分に盛り、豆腐とゆで卵をのせ、漬け汁適宜をかける。好みで練りわさび適宜をのせる。\nう~ん、めんつゆのうまみがしみた豆腐が絶品……!\n甘じょっぱい味つけに、わさびがピリッときいて、もうたまりません。\nひんやりした豆腐と、あったかいご飯の組み合わせが、また最高!\nとろりとした半熟味たまといっしょに食べれば、至福のおいしさですよ♪\nさくっと済ませたい平日ランチや、飲んだあとのシメにもおすすめ。\n『めんつゆ漬け豆腐めし』、ぜひ作ってみてくださいね♪\n(『オレンジページ』2022年9月2日号より)"},
			}},
		},
	},
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
	"https://www.orangepage.net/ymsr/news/daily/posts/5925": {
		Title:    "焦がしたうまみが最高！「じゃがいもとベーコンの塩だけ煮込み」が秋冬じゅう作りたいおいしさ。",
		Image:    "https://images.orangepage.net/media/article/5925/images/3bcb73701558ddc363d8d896452fa378d3ad8547.jpg?w=1135",
		Servings: 2,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "じゃがいも", Amount: "3個(約450g)", Comment: ""},
			{Group: "", Name: "ベーコン", Amount: "120g", Comment: "ブロック"},
			{Group: "", Name: "ミニトマト", Amount: "12個", Comment: ""},
			{Group: "", Name: "玉ねぎ", Amount: "1/4個(約50g)", Comment: ""},
			{Group: "", Name: "ローリエ", Amount: "1枚", Comment: ""},
			{Group: "", Name: "オリーブオイル", Amount: "", Comment: ""},
			{Group: "", Name: "酒", Amount: "", Comment: ""},
			{Group: "", Name: "塩", Amount: "", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "(1)材料の下ごしらえをする\nじゃがいもは皮をむき、幅1.5cmに切る。水に5分ほどさらし、水けをきる。玉ねぎは縦に薄切りにする。ミニトマトはへたを取る。ベーコンは幅1cmに切る。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "(2)こんがりと焼く\nフライパンにオリーブオイル大さじ1を強めの中火で熱し、じゃがいもとベーコンを並べ入れて、こんがりとするまで両面を2分くらいずつ焼く。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "(3)煮込む\n酒大さじ2をふり、玉ねぎを加えてさっと炒める。水1と1/2カップとローリエ、ミニトマトを加える。煮立ったらふたをし、弱火で10~15分煮込む。塩小さじ1/3を加え、中火でひと煮する。\n以上! こちら、煮込みだけどフライパンで作れちゃうのもうれしいですよね。\nちなみに煮込む時間は、じゃがいもの食感の好みで調節しちゃってOK。フレンチマスタードをつけて食べてもおいしいですよ。\nこれからの季節にリピートしたくなる「じゃがいもとかたまりベーコンの塩味煮込み」、早速本日いかがですか?\n(『オレンジページ』2022年11月2日号より)"}}},
		},
	},
	"https://www.orangepage.net/recipes/detail_302394": {
		Title:    "じゃがいものガレット",
		Image:    "https://production-orp.s3.amazonaws.com/uploads/recipes/image/0000302394/20200907150806_w300hf.jpg",
		Servings: 2,
		Ingredients: []rexch.Ingredient{
			{Name: "じゃがいも", Amount: "4個(約500g)"},
			{Name: "仕上げ用の塩", Amount: "適宜", Comment: "あれば粒が大きめのもの"},
			{Name: "塩", Amount: "サラダ油 粗びき黒こしょう"},
		},
		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "じゃがいもは皮をむき、スライサーで細切りにする（なければ包丁でせん切りにする）。塩小さじ1/3をふり、混ぜる。フライパンにサラダ油大さじ3をひき、じゃがいもを全体に広げ入れる。"},
				&rexch.ImageInstructionElement{URL: "https://production-orp.s3.amazonaws.com/uploads/recipe_mades/image/0000051362/20200907150916.jpg"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "強火にかけ、フライ返しで全体をときどき押さえながら2分ほど焼く。パチパチと音がしてきたら中火にし、こんがりと焼き色がつくまで7～8分焼く。"},
				&rexch.ImageInstructionElement{URL: "https://production-orp.s3.amazonaws.com/uploads/recipe_mades/image/0000051363/20200907150916.jpg"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "火を止め、フライパンを少し傾けて油をため、その状態のまま、フライ返しをすきまに差し入れてひっくり返す。こうすると、油がはねにくくなる。"},
				&rexch.ImageInstructionElement{URL: "https://production-orp.s3.amazonaws.com/uploads/recipe_mades/image/0000051364/20200907150916.jpg"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "中火にかけ、サラダ油大さじ2をフライパンの縁から回し入れる。こんがりと焼き色がつくまで7～8分焼く。切り分けて器に盛り、仕上げ用の塩と粗びき黒こしょう適宜をふる。"},
				&rexch.ImageInstructionElement{URL: "https://production-orp.s3.amazonaws.com/uploads/recipe_mades/image/0000051365/20200907150917.jpg"},
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

			rex, err := NewParser().Parse(ctx, url)
			if err != nil {
				t.Fatal(err)
			}

			if err := sites.RecipeMustBe2(want, rex); err != nil {
				t.Fatal(fmt.Errorf("%v: %w", url, err))
			}
		})
	}
}
