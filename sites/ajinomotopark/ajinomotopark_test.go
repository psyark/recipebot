package ajinomotopark

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://park.ajinomoto.co.jp/recipe/card/706051/": {
		Title:    "こだわり手作り！エビのチリソース炒め（干焼蝦仁）",
		Image:    "https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/706051.jpeg",
		Servings: 4,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "むきえび", Amount: "350g", Comment: ""},
			{Group: "", Name: "片栗粉", Amount: "大さじ1・1/2", Comment: ""},
			{Group: "", Name: "にんにくのみじん切り", Amount: "小さじ1", Comment: ""},
			{Group: "", Name: "ねぎ", Amount: "1/2本", Comment: ""},
			{Group: "A", Name: "鶏がらスープ", Amount: "小さじ1", Comment: ""},
			{Group: "A", Name: "トマトケチャップ", Amount: "大さじ3", Comment: ""},
			{Group: "A", Name: "水", Amount: "大さじ5", Comment: ""},
			{Group: "A", Name: "片栗粉", Amount: "小さじ2", Comment: ""},
			{Group: "A", Name: "砂糖", Amount: "小さじ1", Comment: ""},
			{Group: "A", Name: "塩", Amount: "小さじ1/4", Comment: ""},
			{Group: "", Name: "豆板醤", Amount: "小さじ1(5g)", Comment: ""},
			{Group: "", Name: "酒", Amount: "大さじ1", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "大さじ1", Comment: ""},
			{Group: "", Name: "ごま油", Amount: "小さじ1", Comment: ""},
			{Group: "", Name: "香菜", Amount: "少々", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "えびは背ワタを取り、水で洗い、水気を拭く。ねぎは粗みじん切りにする。"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_0_0.jpeg"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_0_1.jpeg"},
			}},
			{Label: "", Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "ボウルにＡを入れて混ぜ合わせ、合わせ調味料を作る。"},
			}},
			{Label: "", Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "（１）のえびに片栗粉をまぶす。"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_2_0.jpeg"},
			}},
			{Label: "", Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "フライパンに油を熱し、にんにくを入れて香りが出るまで炒め、（３）のえびを加えてほぐすようにして炒める。えびの色が変わったら、「熟成豆板醤」を加えて炒める。"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_3_0.jpeg"},
			}},
			{Label: "", Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "（１）のねぎを加えてサッと炒め、酒をふり、（２）の合わせ調味料を加えてとろみがつくまで炒め合わせる。"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_4_0.jpeg"},
			}},
			{Label: "", Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "器に盛り、ごま油をふり、香菜を飾る。"},
			}},
		},
	},
	"https://park.ajinomoto.co.jp/recipe/card/701300/": {
		Title:    "豚肉・しめじ・小松菜のオイスターソース炒め",
		Image:    "https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/701300.jpeg",
		Servings: 2,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "豚こま切れ肉", Amount: "200g", Comment: ""},
			{Group: "A", Name: "片栗粉", Amount: "大さじ1/2", Comment: ""},
			{Group: "A", Name: "塩", Amount: "少々", Comment: ""},
			{Group: "A", Name: "こしょう", Amount: "少々", Comment: ""},
			{Group: "", Name: "しめじ", Amount: "1パック", Comment: ""},
			{Group: "", Name: "小松菜", Amount: "150g", Comment: ""},
			{Group: "B", Name: "オイスターソース", Amount: "大さじ1", Comment: ""},
			{Group: "B", Name: "酒", Amount: "大さじ1", Comment: ""},
			{Group: "B", Name: "鶏がらスープ", Amount: "小さじ1/3", Comment: ""},
			{Group: "", Name: "ごま油", Amount: "小さじ1", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "豚肉はＡをもみ込む。しめじは小房に分け、小松菜は４ｃｍ長さに切る。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンにごま油を熱し、（１）の豚肉をほぐしながら炒める。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "肉の色が変わってきたら、（１）のしめじ・小松菜を加えてＢで調味し、小松菜がしんなりしたら火を止める。"}}},
		},
	},
	"https://park.ajinomoto.co.jp/recipe/card/700708/": {
		Title:    "ミートボールの甘酢あんかけ",
		Image:    "https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/700708.jpeg",
		Servings: 2,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "豚ひき肉", Amount: "150g", Comment: ""},
			{Group: "", Name: "パン粉", Amount: "大さじ2", Comment: ""},
			{Group: "", Name: "酒", Amount: "小さじ1", Comment: ""},
			{Group: "A", Name: "玉ねぎのみじん切り", Amount: "1/4個分(50g)", Comment: ""},
			{Group: "A", Name: "にんにくのみじん切り", Amount: "1/2かけ分", Comment: ""},
			{Group: "A", Name: "卵", Amount: "1/2個", Comment: ""},
			{Group: "A", Name: "塩", Amount: "少々", Comment: ""},
			{Group: "A", Name: "味の素", Amount: "少々", Comment: ""},
			{Group: "A", Name: "こしょう", Amount: "少々", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "適量", Comment: ""},
			{Group: "B", Name: "水", Amount: "1/3カップ", Comment: ""},
			{Group: "B", Name: "酢", Amount: "大さじ1・1/2", Comment: ""},
			{Group: "B", Name: "トマトケチャップ", Amount: "大さじ1", Comment: ""},
			{Group: "B", Name: "しょうゆ", Amount: "大さじ1/2", Comment: ""},
			{Group: "B", Name: "砂糖", Amount: "大さじ1/2", Comment: ""},
			{Group: "B", Name: "オイスターソース", Amount: "小さじ1/2", Comment: ""},
			{Group: "B", Name: "ごま油", Amount: "小さじ1/2", Comment: ""},
			{Group: "B", Name: "粉末中華スープ", Amount: "少々", Comment: ""},
			{Group: "B", Name: "片栗粉", Amount: "大さじ1/4", Comment: ""},
			{Group: "", Name: "レタス", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "パン粉は酒をふって混ぜる。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ひき肉に（１）のパン粉、Ａを加えて粘りが出るまでよく練り混ぜ、ひと口大に丸める。１６０～１７０℃の揚げ油できつね色に揚げる。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "小鍋にＢを入れてよく混ぜ、強火にかけて混ぜながら煮立て、とろみをつける。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "器にレタスを敷き、（２）のミートボールを盛り、（３）のあんかけをかける。"}}},
		},
	},
	"https://park.ajinomoto.co.jp/recipe/card/702479/": {
		Title:    "納豆チャーハン",
		Image:    "https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/702479.jpeg",
		Servings: 2,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "ご飯", Amount: "400g", Comment: ""},
			{Group: "", Name: "納豆", Amount: "2パック", Comment: ""},
			{Group: "", Name: "卵", Amount: "2個", Comment: ""},
			{Group: "", Name: "ねぎ・粗みじん切り", Amount: "1/2本分", Comment: ""},
			{Group: "A", Name: "しょうゆ", Amount: "大さじ1", Comment: ""},
			{Group: "A", Name: "鶏がらスープ", Amount: "大さじ1", Comment: ""},
			{Group: "A", Name: "こしょう", Amount: "少々", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "大さじ3", Comment: ""},
			{Group: "", Name: "ごま油", Amount: "小さじ1", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "卵は溶きほぐしておく。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "フライパンの中心に油を入れて熱し、（１）の溶き卵を油の中心に流し入れて包み込むように混ぜ、半熟状にする。"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/702479_direction_1_0.jpeg"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "ご飯を加えて卵をご飯の中に混ぜ込むようにして炒め合わせ、パラパラになってきたら、納豆、ねぎを加えてさらに炒める。"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/702479_direction_2_0.jpeg"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/702479_direction_2_1.jpeg"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "納豆のネバネバが切れる程度に炒めたら、Ａを加え、仕上げにごま油を回し入れる。"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/702479_direction_3_0.jpeg"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "＊納豆のネバネバがなくなるまでしっかり炒めましょう。"},
			}},
		},
	},
	"https://park.ajinomoto.co.jp/recipe/card/706078/": {
		Title:    "オムライス",
		Image:    "https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/706078.jpeg",
		Servings: 2,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "鶏むね肉", Amount: "100g", Comment: ""},
			{Group: "", Name: "玉ねぎ", Amount: "1/2個", Comment: ""},
			{Group: "", Name: "温かいご飯", Amount: "300g", Comment: ""},
			{Group: "", Name: "コンソメ", Amount: "小さじ1", Comment: "顆粒"},
			{Group: "", Name: "トマトケチャップ", Amount: "適量", Comment: ""},
			{Group: "", Name: "卵", Amount: "4個", Comment: ""},
			{Group: "", Name: "塩", Amount: "適量", Comment: ""},
			{Group: "", Name: "こしょう", Amount: "適量", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "適量", Comment: ""},
			{Group: "", Name: "バター", Amount: "大さじ2", Comment: ""},
			{Group: "", Name: "パセリ", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "鶏肉は１．５ｃｍ角に切り、塩・こしょう少々をふる。玉ねぎはみじん切りにする。"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_0_0.jpeg"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "フライパンに油小さじ１を熱し、（１）の鶏肉を炒める。焼き色がついたら、バター大さじ１、（１）の玉ねぎを加えてよく炒める。"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_1_0.jpeg"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "ご飯を加えて「コンソメ」をふり、混ぜながら炒める。トマトケチャップ大さじ２、塩・こしょう少々で味を調え、チキンライスを作る。"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_2_0.jpeg"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "小さめのボウルに卵２個を溶きほぐし、塩・こしょう少々を混ぜる。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "フライパンに油、バター各大さじ１／２を熱し、（４）の溶き卵を一気に流し入れて全体をサッと混ぜる。半熟状になったら（３）のチキンライスの半量を中央にのせ、両端からヘラで折り曲げる。"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_4_0.jpeg"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "フライパンの片側に寄せ、皿に返して盛りつける。トマトケチャップ少々をかけ、パセリを飾る。もう一つも同様に作る。"},
				&rexch.ImageInstructionElement{URL: "https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_5_0.jpeg"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "＊フライパンのフチを利用すると形よく整えられます。"},
			}},
		},
	},
	"https://park.ajinomoto.co.jp/recipe/card/706119/": {
		Title:    "ホットケーキ",
		Image:    "https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/706119.jpeg",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "A", Name: "薄力粉", Amount: "150g", Comment: ""},
			{Group: "A", Name: "ベーキングパウダー", Amount: "小さじ2", Comment: ""},
			{Group: "A", Name: "砂糖", Amount: "40g", Comment: ""},
			{Group: "", Name: "卵", Amount: "1個", Comment: ""},
			{Group: "", Name: "牛乳", Amount: "130ml", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "少々", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ボウルにＡをふるい入れ、泡立て器で混ぜる。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "別のボウルに卵を割りほぐし、牛乳を加えて混ぜる。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "（１）のボウルに（２）を加え、泡立て器で粉っぽさがなくなるまで混ぜ、生地を作る。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンに油をしみ込ませたキッチンペーパーで油を薄く塗り、フライパンを熱する。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ぬれたふきんの上にフライパンを置いて熱を取り、再び弱火にかけ、（３）の生地をおたま１杯、上から中心に落とす（こうすると丸い形になる）。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "約３分焼き、表面にプツプツと穴がでてきたら裏返し、約２分弱火のまま焼き、取り出す。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "２枚目からは油をひかずにフライパンをぬれふきんの上に置いて熱を取り、弱火にかけて同様に焼く。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "器に盛り、好みでバター、メープルシロップをかける。"}}},
			{Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "＊ベーキングパウダー入りの生地はすぐに焼かないとふくらみが悪くなるので、\u3000混ぜたらすぐに焼きましょう。"}}},
		},
	},
	"https://park.ajinomoto.co.jp/recipe/card/710481/": {
		Title:    "だし香る鮭のちゃんちゃん焼き",
		Image:    "https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/710481.jpeg",
		Servings: 4,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "生ざけ", Amount: "4切れ", Comment: ""},
			{Group: "", Name: "和風だしの素", Amount: "小さじ1", Comment: ""},
			{Group: "", Name: "キャベツ", Amount: "200g", Comment: ""},
			{Group: "", Name: "もやし", Amount: "1袋", Comment: ""},
			{Group: "", Name: "玉ねぎ", Amount: "1個", Comment: ""},
			{Group: "", Name: "にんにく", Amount: "1かけ", Comment: ""},
			{Group: "A", Name: "みそ", Amount: "大さじ3", Comment: ""},
			{Group: "A", Name: "酒", Amount: "大さじ1", Comment: ""},
			{Group: "A", Name: "和風だしの素", Amount: "小さじ2", Comment: ""},
			{Group: "A", Name: "砂糖", Amount: "小さじ1", Comment: ""},
			{Group: "A", Name: "おろしにんにく", Amount: "適量", Comment: "チューブ"},
			{Group: "", Name: "バター", Amount: "10g", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "小さじ2", Comment: ""},
			{Group: "", Name: "一味唐がらし・好みで", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "さけの両面に「お塩控えめの・ほんだし」をふる。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "キャベツは３ｃｍ角くらいに切り、玉ねぎはタテ半分に切って芯を取り、５ｍｍ幅の薄切りにする。にんにくは薄切りにする。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "Ａを混ぜ合わせ、合わせみそを作る。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンに油小さじ１を熱し、（１）のさけを入れ、両面香ばしい焼き色がつくまで焼いていったん皿に取り出す。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンの汚れをキッチンペーパーで拭き取り、油小さじ１を入れて熱し、（２）のキャベツ・玉ねぎ・にんにく、もやしを加えて混ぜる。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "（３）の合わせみそを数か所にのせてフタをし、弱めの中火でムラなく焼けるように時々混ぜながら、７分ほど蒸し焼きにする。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンの中央をあけ、（４）のさけを戻し入れてフタをし、３分ほど蒸し焼きにし、さけの上にバターをちぎってのせる。好みで一味唐がらしをふる。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "＊食卓へフライパンのまま豪快に出しても、お皿に盛りつけても。"}}},
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
