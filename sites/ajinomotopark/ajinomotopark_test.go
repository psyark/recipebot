package ajinomotopark

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/sites"
)

func TestNewParser(t *testing.T) {
	ctx := context.Background()
	tests := map[string]string{
		"https://park.ajinomoto.co.jp/recipe/card/706051/": `{"title":"こだわり手作り！エビのチリソース炒め（干焼蝦仁）","image":"https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/706051.jpeg","servings":4,"ingredients":[{"name":"むきえび","amount":"350g"},{"name":"片栗粉","amount":"大さじ1・1/2"},{"name":"にんにくのみじん切り","amount":"小さじ1"},{"name":"ねぎ","amount":"1/2本"},{"group":"A","name":"鶏がらスープ","amount":"小さじ1"},{"group":"A","name":"トマトケチャップ","amount":"大さじ3"},{"group":"A","name":"水","amount":"大さじ5"},{"group":"A","name":"片栗粉","amount":"小さじ2"},{"group":"A","name":"砂糖","amount":"小さじ1"},{"group":"A","name":"塩","amount":"小さじ1/4"},{"name":"豆板醤","amount":"小さじ1(5g)"},{"name":"酒","amount":"大さじ1"},{"name":"サラダ油","amount":"大さじ1"},{"name":"ごま油","amount":"小さじ1"},{"name":"香菜","amount":"少々"}],"instructions":[{"elements":[{"text":"1えびは背ワタを取り、水で洗い、水気を拭く。ねぎは粗みじん切りにする。"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_0_0.jpeg"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_0_1.jpeg"}]},{"elements":[{"text":"2ボウルにＡを入れて混ぜ合わせ、合わせ調味料を作る。"}]},{"elements":[{"text":"3（１）のえびに片栗粉をまぶす。"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_2_0.jpeg"}]},{"elements":[{"text":"4フライパンに油を熱し、にんにくを入れて香りが出るまで炒め、（３）のえびを加えてほぐすようにして炒める。えびの色が変わったら、「熟成豆板醤」を加えて炒める。"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_3_0.jpeg"}]},{"elements":[{"text":"5（１）のねぎを加えてサッと炒め、酒をふり、（２）の合わせ調味料を加えてとろみがつくまで炒め合わせる。"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_4_0.jpeg"}]},{"elements":[{"text":"6器に盛り、ごま油をふり、香菜を飾る。"}]}]}`,
		"https://park.ajinomoto.co.jp/recipe/card/701300/": `{"title":"豚肉・しめじ・小松菜のオイスターソース炒め","image":"https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/701300.jpeg","servings":2,"ingredients":[{"name":"豚こま切れ肉","amount":"200g"},{"group":"A","name":"片栗粉","amount":"大さじ1/2"},{"group":"A","name":"塩","amount":"少々"},{"group":"A","name":"こしょう","amount":"少々"},{"name":"しめじ","amount":"1パック"},{"name":"小松菜","amount":"150g"},{"group":"B","name":"オイスターソース","amount":"大さじ1"},{"group":"B","name":"酒","amount":"大さじ1"},{"group":"B","name":"鶏がらスープ","amount":"小さじ1/3"},{"name":"ごま油","amount":"小さじ1"}],"instructions":[{"elements":[{"text":"1豚肉はＡをもみ込む。しめじは小房に分け、小松菜は４ｃｍ長さに切る。"}]},{"elements":[{"text":"2フライパンにごま油を熱し、（１）の豚肉をほぐしながら炒める。"}]},{"elements":[{"text":"3肉の色が変わってきたら、（１）のしめじ・小松菜を加えてＢで調味し、小松菜がしんなりしたら火を止める。"}]}]}`,
		"https://park.ajinomoto.co.jp/recipe/card/700708/": `{"title":"ミートボールの甘酢あんかけ","image":"https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/700708.jpeg","servings":2,"ingredients":[{"name":"豚ひき肉","amount":"150g"},{"name":"パン粉","amount":"大さじ2"},{"name":"酒","amount":"小さじ1"},{"group":"A","name":"玉ねぎのみじん切り","amount":"1/4個分(50g)"},{"group":"A","name":"にんにくのみじん切り","amount":"1/2かけ分"},{"group":"A","name":"卵","amount":"1/2個"},{"group":"A","name":"塩","amount":"少々"},{"group":"A","name":"味の素","amount":"少々"},{"group":"A","name":"こしょう","amount":"少々"},{"name":"サラダ油","amount":"適量"},{"group":"B","name":"水","amount":"1/3カップ"},{"group":"B","name":"酢","amount":"大さじ1・1/2"},{"group":"B","name":"トマトケチャップ","amount":"大さじ1"},{"group":"B","name":"しょうゆ","amount":"大さじ1/2"},{"group":"B","name":"砂糖","amount":"大さじ1/2"},{"group":"B","name":"オイスターソース","amount":"小さじ1/2"},{"group":"B","name":"ごま油","amount":"小さじ1/2"},{"group":"B","name":"粉末中華スープ","amount":"少々"},{"group":"B","name":"片栗粉","amount":"大さじ1/4"},{"name":"レタス","amount":"適量"}],"instructions":[{"elements":[{"text":"1パン粉は酒をふって混ぜる。"}]},{"elements":[{"text":"2ひき肉に（１）のパン粉、Ａを加えて粘りが出るまでよく練り混ぜ、ひと口大に丸める。１６０～１７０℃の揚げ油できつね色に揚げる。"}]},{"elements":[{"text":"3小鍋にＢを入れてよく混ぜ、強火にかけて混ぜながら煮立て、とろみをつける。"}]},{"elements":[{"text":"4器にレタスを敷き、（２）のミートボールを盛り、（３）のあんかけをかける。"}]}]}`,
		"https://park.ajinomoto.co.jp/recipe/card/702479/": `{"title":"パラっと香ばしい！  \n                納豆チャーハン","image":"https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/702479.jpeg","servings":2,"ingredients":[{"name":"ご飯","amount":"400g"},{"name":"納豆","amount":"2パック"},{"name":"卵","amount":"2個"},{"name":"ねぎ・粗みじん切り","amount":"1/2本分"},{"group":"A","name":"しょうゆ","amount":"大さじ1"},{"group":"A","name":"鶏がらスープ","amount":"大さじ1"},{"group":"A","name":"こしょう","amount":"少々"},{"name":"サラダ油","amount":"大さじ3"},{"name":"ごま油","amount":"小さじ1"}],"instructions":[{"elements":[{"text":"1卵は溶きほぐしておく。"}]},{"elements":[{"text":"2フライパンの中心に油を入れて熱し、（１）の溶き卵を油の中心に流し入れて包み込むように混ぜ、半熟状にする。"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/702479_direction_1_0.jpeg"}]},{"elements":[{"text":"3ご飯を加えて卵をご飯の中に混ぜ込むようにして炒め合わせ、パラパラになってきたら、納豆、ねぎを加えてさらに炒める。"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/702479_direction_2_0.jpeg"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/702479_direction_2_1.jpeg"}]},{"elements":[{"text":"4納豆のネバネバが切れる程度に炒めたら、Ａを加え、仕上げにごま油を回し入れる。"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/702479_direction_3_0.jpeg"}]},{"elements":[{"text":"＊納豆のネバネバがなくなるまでしっかり炒めましょう。"}]}]}`,
		"https://park.ajinomoto.co.jp/recipe/card/706078/": `{"title":"ふんわり卵で絶品  \n                オムライス","image":"https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/706078.jpeg","servings":2,"ingredients":[{"name":"鶏むね肉","amount":"100g"},{"name":"玉ねぎ","amount":"1/2個"},{"name":"温かいご飯","amount":"300g"},{"name":"コンソメ","amount":"小さじ1","comment":"顆粒"},{"name":"トマトケチャップ","amount":"適量"},{"name":"卵","amount":"4個"},{"name":"塩","amount":"適量"},{"name":"こしょう","amount":"適量"},{"name":"サラダ油","amount":"適量"},{"name":"バター","amount":"大さじ2"},{"name":"パセリ","amount":"適量"}],"instructions":[{"elements":[{"text":"1鶏肉は１．５ｃｍ角に切り、塩・こしょう少々をふる。玉ねぎはみじん切りにする。"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_0_0.jpeg"}]},{"elements":[{"text":"2フライパンに油小さじ１を熱し、（１）の鶏肉を炒める。焼き色がついたら、バター大さじ１、（１）の玉ねぎを加えてよく炒める。"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_1_0.jpeg"}]},{"elements":[{"text":"3ご飯を加えて「コンソメ」をふり、混ぜながら炒める。トマトケチャップ大さじ２、塩・こしょう少々で味を調え、チキンライスを作る。"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_2_0.jpeg"}]},{"elements":[{"text":"4小さめのボウルに卵２個を溶きほぐし、塩・こしょう少々を混ぜる。"}]},{"elements":[{"text":"5フライパンに油、バター各大さじ１／２を熱し、（４）の溶き卵を一気に流し入れて全体をサッと混ぜる。半熟状になったら（３）のチキンライスの半量を中央にのせ、両端からヘラで折り曲げる。"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_4_0.jpeg"}]},{"elements":[{"text":"6フライパンの片側に寄せ、皿に返して盛りつける。トマトケチャップ少々をかけ、パセリを飾る。もう一つも同様に作る。"},{"url":"https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_5_0.jpeg"}]},{"elements":[{"text":"＊フライパンのフチを利用すると形よく整えられます。"}]}]}`,
	}

	for url, want := range tests {
		url := url
		want := want

		t.Run(url, func(t *testing.T) {
			t.Parallel()

			rex, err := NewParser().Parse2(ctx, url)
			if err != nil {
				t.Fatal(err)
			}

			if err := sites.RecipeMustBe2(rex, want); err != nil {
				t.Error(err)
			}
		})
	}
}
