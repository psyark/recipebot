package foodie

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/sites"
)

func TestNewParser(t *testing.T) {
	ctx := context.Background()
	tests := map[string]string{
		"https://mi-journey.jp/foodie/52916/": `{"Title":"【シェフ直伝】本格リゾットのレシピ。生米をアルデンテに仕上げるテクニック","Image":"https://mi-journey.jp/foodie/wp-content/uploads/2018/10/1810_28_top.jpg","IngredientGroups":[{"Name":"","Children":[{"Name":"米","Amount":"1合(150g)","Comment":""},{"Name":"オリーブオイル","Amount":"大さじ2","Comment":""},{"Name":"ブイヨン","Amount":"3カップ※顆粒ブイヨンを湯に溶かしたもの","Comment":""},{"Name":"ローリエ","Amount":"2枚","Comment":""},{"Name":"白ワイン","Amount":"大さじ1","Comment":""},{"Name":"バター","Amount":"20g","Comment":""},{"Name":"パルミジャーノ・レッジャーノチーズ","Amount":"大さじ2","Comment":""},{"Name":"塩","Amount":"少々","Comment":""},{"Name":"粗びき黒こしょう","Amount":"少々","Comment":""}]}],"Steps":[{"Text":"オリーブオイルで米を炒めながら、隣でブイヨンに白ワインとあればローリエを加えて温める","Images":null},{"Text":"お玉2杯分のブイヨンと白ワインを入れて、米を炊き始める","Images":null},{"Text":"ブイヨンをつぎ足しながら、20分炊く","Images":null},{"Text":"米がふくらみ、粒が立って、艶々してきたら、味見をする","Images":null},{"Text":"火を止めてから、バターとパルミジャーノチーズを加えて余熱で混ぜ、塩で味を調えてから、器に盛り、粗びき黒こしょうをふる。","Images":null}]}`,
		"https://mi-journey.jp/foodie/57058/": `{"Title":"【シェフ直伝】オムレツのレシピ ホテルのように美しく作るコツ","Image":"https://mi-journey.jp/foodie/wp-content/uploads/2019/04/190413omelette1_.jpg","IngredientGroups":[{"Name":"","Children":[{"Name":"卵","Amount":"2〜3個","Comment":""},{"Name":"バター","Amount":"10g","Comment":""},{"Name":"塩、こしょう","Amount":"各少々","Comment":""}]}],"Steps":[{"Text":"卵液を作る","Images":null},{"Text":"卵液を漉す","Images":null},{"Text":"フライパンを火にかけ、バターを溶かす","Images":null},{"Text":"卵液を流し入れる","Images":null},{"Text":"ゴムべらで混ぜながら、半熟状になるまで火を通す","Images":null},{"Text":"濡れ布巾の上でフライパンを叩く","Images":null},{"Text":"フライパンにこびりついた卵の端を処理する","Images":null},{"Text":"手前から卵を包む","Images":null},{"Text":"反対側の卵を包む","Images":null},{"Text":"フライパンを奥に寄せて成形する","Images":null},{"Text":"ゴムべらでひっくり返し、強火にかけ、表面を固める","Images":null},{"Text":"赤ワインを煮詰めてソースを作る","Images":null},{"Text":"ケチャップを加える","Images":null}]}`,
		"https://mi-journey.jp/foodie/20667/": `{"Title":"お家で簡単！フライパンで本格パエリアレシピ ～プロと家庭はここが違った！作り方の3つのポイント","Image":"https://mi-journey.jp/foodie/wp-content/uploads/2016/02/0303_paella_top.jpg","IngredientGroups":[{"Name":"","Children":[{"Name":"いか","Amount":"小1パイ分(100g)","Comment":""},{"Name":"白身魚","Amount":"小1枚(100g)","Comment":"切り身"},{"Name":"玉ねぎ","Amount":"40g","Comment":""},{"Name":"にんじん","Amount":"20g","Comment":""},{"Name":"セロリ","Amount":"20g","Comment":""},{"Name":"にんにく","Amount":"2片","Comment":""},{"Name":"トマト缶","Amount":"1/2カップ(100g)","Comment":"ホール"},{"Name":"有頭えび","Amount":"5尾(200g)","Comment":""},{"Name":"あさり","Amount":"15~16個(100g)","Comment":"砂抜き済"},{"Name":"水","Amount":"600ml(※)","Comment":""},{"Name":"塩","Amount":"小さじ1/2","Comment":""},{"Name":"サフラン","Amount":"10本","Comment":""},{"Name":"米","Amount":"1合(180ml)","Comment":""},{"Name":"レモン、イタリアンパセリ","Amount":"お好みで","Comment":""},{"Name":"オリーブ油","Amount":"1/4カップ(50ml)","Comment":""}]}],"Steps":null}`,
		"https://mi-journey.jp/foodie/62677/": `{"Title":"まっすぐなエビフライの作り方。プリッとジューシーに仕上げるプロのコツ","Image":"https://mi-journey.jp/foodie/wp-content/uploads/2020/06/200404ebifurai1.jpg","IngredientGroups":[{"Name":"","Children":[{"Name":"えび","Amount":"8尾※ブラックタイガーなど","Comment":"無頭"},{"Name":"片栗粉、塩、こしょう","Amount":"各少々","Comment":""},{"Name":"小麦粉、溶き卵、生パン粉、揚げ油","Amount":"各適量","Comment":""}]}],"Steps":[{"Text":"えびの殻を剥く","Images":null},{"Text":"背わたを除く","Images":null},{"Text":"塩、片栗粉で軽く揉む","Images":null},{"Text":"水洗いし、水気を拭く","Images":null},{"Text":"尾先と剣先を切って水を出す","Images":null},{"Text":"腹に切り目を入れる","Images":null},{"Text":"塩、こしょうをふり、20分ほどおく","Images":null},{"Text":"卵液を作る","Images":null},{"Text":"衣をつける","Images":null},{"Text":"パン粉をつける","Images":null},{"Text":"揚げる","Images":null}]}`,
		"https://mi-journey.jp/foodie/59727/": `{"Title":"鮭のムニエルのレシピ～洋食店のように仕上げる焼き方のコツ 【シェフ直伝】","Image":"https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere12.jpg","IngredientGroups":[{"Name":"","Children":[{"Name":"生鮭","Amount":"2切れ","Comment":""},{"Name":"塩","Amount":"少々","Comment":""},{"Name":"小麦粉","Amount":"適量","Comment":""},{"Name":"オリーブオイル","Amount":"大さじ1","Comment":""},{"Name":"バター","Amount":"15g","Comment":""}]},{"Name":"【ソース】","Children":[{"Name":"バター","Amount":"60g","Comment":""},{"Name":"しょうゆ","Amount":"小さじ2","Comment":""},{"Name":"にんにく","Amount":"1/2かけ分","Comment":"みじん切り"},{"Name":"トマト","Amount":"1/2個分(約30g)","Comment":"湯むきして種を除き、角切りにしたもの"},{"Name":"レモン","Amount":"1/4個分(約7g)","Comment":"種と薄皮を除き角切りにしたもの"},{"Name":"ケイパー","Amount":"大さじ2","Comment":"酢漬け"},{"Name":"パセリ","Amount":"各適量","Comment":"みじん切り）、レモンの皮（好みで"}]}],"Steps":[{"Text":"①鮭の両面に塩をふり、10分ほどおく","Images":null},{"Text":"②ペーパータオルで水けをふき取る","Images":null},{"Text":"③鮭に小麦粉をまんべんなくまぶす","Images":null},{"Text":"④冷たいフライパンにオリーブオイルを入れ、鮭を皮目から入れる","Images":null},{"Text":"⑤弱火にかけ、鮭をじっくり焼く","Images":null},{"Text":"⑥皮目に焼き色がついたら身を焼く","Images":null},{"Text":"⑦バターを加える","Images":null}]}`,
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

			if err := sites.RecipeMustBe(*rcp, want); err != nil {
				t.Error(err)
			}
		})
	}
}
