package ajinomotopark

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/sites/common"
)

func TestNewParser(t *testing.T) {
	ctx := context.Background()
	tests := map[string]string{
		"https://park.ajinomoto.co.jp/recipe/card/706051/": `{"Title":"こだわり手作り！エビのチリソース炒め（干焼蝦仁）","Image":"https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/706051.jpeg","IngredientGroups":[{"Name":"","Children":[{"Name":"むきえび","Amount":"350g","Comment":""},{"Name":"片栗粉","Amount":"大さじ1・1/2","Comment":""},{"Name":"にんにくのみじん切り","Amount":"小さじ1","Comment":""},{"Name":"ねぎ","Amount":"1/2本","Comment":""},{"Name":"豆板醤","Amount":"小さじ1(5g)","Comment":""},{"Name":"酒","Amount":"大さじ1","Comment":""},{"Name":"サラダ油","Amount":"大さじ1","Comment":""},{"Name":"ごま油","Amount":"小さじ1","Comment":""},{"Name":"香菜","Amount":"少々","Comment":""}]},{"Name":"A","Children":[{"Name":"鶏がらスープ","Amount":"小さじ1","Comment":""},{"Name":"トマトケチャップ","Amount":"大さじ3","Comment":""},{"Name":"水","Amount":"大さじ5","Comment":""},{"Name":"片栗粉","Amount":"小さじ2","Comment":""},{"Name":"砂糖","Amount":"小さじ1","Comment":""},{"Name":"塩","Amount":"小さじ1/4","Comment":""}]}],"Steps":[{"Text":"1えびは背ワタを取り、水で洗い、水気を拭く。ねぎは粗みじん切りにする。","Images":["https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_0_0.jpeg","https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_0_1.jpeg"]},{"Text":"2ボウルにＡを入れて混ぜ合わせ、合わせ調味料を作る。","Images":null},{"Text":"3（１）のえびに片栗粉をまぶす。","Images":["https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_2_0.jpeg"]},{"Text":"4フライパンに油を熱し、にんにくを入れて香りが出るまで炒め、（３）のえびを加えてほぐすようにして炒める。えびの色が変わったら、「熟成豆板醤」を加えて炒める。","Images":["https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_3_0.jpeg"]},{"Text":"5（１）のねぎを加えてサッと炒め、酒をふり、（２）の合わせ調味料を加えてとろみがつくまで炒め合わせる。","Images":["https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706051_direction_4_0.jpeg"]},{"Text":"6器に盛り、ごま油をふり、香菜を飾る。","Images":null}]}`,
		"https://park.ajinomoto.co.jp/recipe/card/701300/": `{"Title":"豚肉・しめじ・小松菜のオイスターソース炒め","Image":"https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/701300.jpeg","IngredientGroups":[{"Name":"","Children":[{"Name":"豚こま切れ肉","Amount":"200g","Comment":""},{"Name":"しめじ","Amount":"1パック","Comment":""},{"Name":"小松菜","Amount":"150g","Comment":""},{"Name":"ごま油","Amount":"小さじ1","Comment":""}]},{"Name":"A","Children":[{"Name":"片栗粉","Amount":"大さじ1/2","Comment":""},{"Name":"塩","Amount":"少々","Comment":""},{"Name":"こしょう","Amount":"少々","Comment":""}]},{"Name":"B","Children":[{"Name":"オイスターソース","Amount":"大さじ1","Comment":""},{"Name":"酒","Amount":"大さじ1","Comment":""},{"Name":"鶏がらスープ","Amount":"小さじ1/3","Comment":""}]}],"Steps":[{"Text":"1豚肉はＡをもみ込む。しめじは小房に分け、小松菜は４ｃｍ長さに切る。","Images":null},{"Text":"2フライパンにごま油を熱し、（１）の豚肉をほぐしながら炒める。","Images":null},{"Text":"3肉の色が変わってきたら、（１）のしめじ・小松菜を加えてＢで調味し、小松菜がしんなりしたら火を止める。","Images":null}]}`,
		"https://park.ajinomoto.co.jp/recipe/card/700708/": `{"Title":"ミートボールの甘酢あんかけ","Image":"https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/700708.jpeg","IngredientGroups":[{"Name":"","Children":[{"Name":"豚ひき肉","Amount":"150g","Comment":""},{"Name":"パン粉","Amount":"大さじ2","Comment":""},{"Name":"酒","Amount":"小さじ1","Comment":""},{"Name":"サラダ油","Amount":"適量","Comment":""},{"Name":"レタス","Amount":"適量","Comment":""}]},{"Name":"A","Children":[{"Name":"玉ねぎのみじん切り","Amount":"1/4個分(50g)","Comment":""},{"Name":"にんにくのみじん切り","Amount":"1/2かけ分","Comment":""},{"Name":"卵","Amount":"1/2個","Comment":""},{"Name":"塩","Amount":"少々","Comment":""},{"Name":"味の素","Amount":"少々","Comment":""},{"Name":"こしょう","Amount":"少々","Comment":""}]},{"Name":"B","Children":[{"Name":"水","Amount":"1/3カップ","Comment":""},{"Name":"酢","Amount":"大さじ1・1/2","Comment":""},{"Name":"トマトケチャップ","Amount":"大さじ1","Comment":""},{"Name":"しょうゆ","Amount":"大さじ1/2","Comment":""},{"Name":"砂糖","Amount":"大さじ1/2","Comment":""},{"Name":"オイスターソース","Amount":"小さじ1/2","Comment":""},{"Name":"ごま油","Amount":"小さじ1/2","Comment":""},{"Name":"粉末中華スープ","Amount":"少々","Comment":""},{"Name":"片栗粉","Amount":"大さじ1/4","Comment":""}]}],"Steps":[{"Text":"1パン粉は酒をふって混ぜる。","Images":null},{"Text":"2ひき肉に（１）のパン粉、Ａを加えて粘りが出るまでよく練り混ぜ、ひと口大に丸める。１６０～１７０℃の揚げ油できつね色に揚げる。","Images":null},{"Text":"3小鍋にＢを入れてよく混ぜ、強火にかけて混ぜながら煮立て、とろみをつける。","Images":null},{"Text":"4器にレタスを敷き、（２）のミートボールを盛り、（３）のあんかけをかける。","Images":null}]}`,
		"https://park.ajinomoto.co.jp/recipe/card/702479/": `{"Title":"パラっと香ばしい！  \n                納豆チャーハン","Image":"https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/702479.jpeg","IngredientGroups":[{"Name":"","Children":[{"Name":"ご飯","Amount":"400g","Comment":""},{"Name":"納豆","Amount":"2パック","Comment":""},{"Name":"卵","Amount":"2個","Comment":""},{"Name":"ねぎ・粗みじん切り","Amount":"1/2本分","Comment":""},{"Name":"サラダ油","Amount":"大さじ3","Comment":""},{"Name":"ごま油","Amount":"小さじ1","Comment":""}]},{"Name":"A","Children":[{"Name":"しょうゆ","Amount":"大さじ1","Comment":""},{"Name":"鶏がらスープ","Amount":"大さじ1","Comment":""},{"Name":"こしょう","Amount":"少々","Comment":""}]}],"Steps":[{"Text":"1卵は溶きほぐしておく。","Images":null},{"Text":"2フライパンの中心に油を入れて熱し、（１）の溶き卵を油の中心に流し入れて包み込むように混ぜ、半熟状にする。","Images":["https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/702479_direction_1_0.jpeg"]},{"Text":"3ご飯を加えて卵をご飯の中に混ぜ込むようにして炒め合わせ、パラパラになってきたら、納豆、ねぎを加えてさらに炒める。","Images":["https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/702479_direction_2_0.jpeg","https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/702479_direction_2_1.jpeg"]},{"Text":"4納豆のネバネバが切れる程度に炒めたら、Ａを加え、仕上げにごま油を回し入れる。","Images":["https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/702479_direction_3_0.jpeg"]},{"Text":"＊納豆のネバネバがなくなるまでしっかり炒めましょう。","Images":null}]}`,
		"https://park.ajinomoto.co.jp/recipe/card/706078/": `{"Title":"ふんわり卵で絶品  \n                オムライス","Image":"https://park.ajinomoto.co.jp/wp-content/uploads/2018/03/706078.jpeg","IngredientGroups":[{"Name":"","Children":[{"Name":"鶏むね肉","Amount":"100g","Comment":""},{"Name":"玉ねぎ","Amount":"1/2個","Comment":""},{"Name":"温かいご飯","Amount":"300g","Comment":""},{"Name":"コンソメ","Amount":"小さじ1","Comment":"顆粒"},{"Name":"トマトケチャップ","Amount":"適量","Comment":""},{"Name":"卵","Amount":"4個","Comment":""},{"Name":"塩","Amount":"適量","Comment":""},{"Name":"こしょう","Amount":"適量","Comment":""},{"Name":"サラダ油","Amount":"適量","Comment":""},{"Name":"バター","Amount":"大さじ2","Comment":""},{"Name":"パセリ","Amount":"適量","Comment":""}]}],"Steps":[{"Text":"1鶏肉は１．５ｃｍ角に切り、塩・こしょう少々をふる。玉ねぎはみじん切りにする。","Images":["https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_0_0.jpeg"]},{"Text":"2フライパンに油小さじ１を熱し、（１）の鶏肉を炒める。焼き色がついたら、バター大さじ１、（１）の玉ねぎを加えてよく炒める。","Images":["https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_1_0.jpeg"]},{"Text":"3ご飯を加えて「コンソメ」をふり、混ぜながら炒める。トマトケチャップ大さじ２、塩・こしょう少々で味を調え、チキンライスを作る。","Images":["https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_2_0.jpeg"]},{"Text":"4小さめのボウルに卵２個を溶きほぐし、塩・こしょう少々を混ぜる。","Images":null},{"Text":"5フライパンに油、バター各大さじ１／２を熱し、（４）の溶き卵を一気に流し入れて全体をサッと混ぜる。半熟状になったら（３）のチキンライスの半量を中央にのせ、両端からヘラで折り曲げる。","Images":["https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_4_0.jpeg"]},{"Text":"6フライパンの片側に寄せ、皿に返して盛りつける。トマトケチャップ少々をかけ、パセリを飾る。もう一つも同様に作る。","Images":["https://park.ajinomoto.co.jp/wp-content/uploads/2021/08/706078_direction_5_0.jpeg"]},{"Text":"＊フライパンのフチを利用すると形よく整えられます。","Images":null}]}`,
	}

	for url, want := range tests {
		url := url
		want := want

		t.Run(url, func(t *testing.T) {
			t.Parallel()

			rcp, err := NewParser().Parse(ctx, url)
			if err != nil {
				t.Fatal(err)
			}

			if err := common.RecipeMustBe(*rcp, want); err != nil {
				t.Error(err)
			}
		})
	}
}
