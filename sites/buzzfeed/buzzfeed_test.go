package buzzfeed

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/sites"
)

func TestNewParser(t *testing.T) {
	ctx := context.Background()
	tests := map[string]string{
		"https://www.buzzfeed.com/jp/kazuyawakana/eggplant-ooba-mentsuyu-pickled-doubanjiang": `{"title":"【めんつゆ×豆板醤で激ウマ！】ご飯がススム！なすと大葉のめんつゆ豆板醤漬け","image":"https://img.buzzfeed.com/buzzfeed-static/static/2022-05/27/6/asset/5c946aa4fc12/sub-buzz-3082-1653632598-1.jpg?downsize=700%3A%2A\u0026output-quality=auto\u0026output-format=auto","servings":1,"ingredients":[{"name":"なす","amount":"2本"},{"name":"大葉","amount":"10枚"},{"group":"A","name":"しょうがすりおろし","amount":"小さじ1"},{"group":"A","name":"ごま油","amount":"大さじ1"},{"group":"A","name":"めんつゆ（2倍濃縮）大さじ3"},{"group":"A","name":"豆板醤","amount":"小さじ2"},{"group":"A","name":"白ごま","amount":"適量"}],"instructions":[{"elements":[{"text":"なすのへたをとり、縦に薄くスライスする。皿に並べラップをし、600Wの電子レンジで3分加熱する。"}]},{"elements":[{"text":"ボウルにAを入れよく混ぜる。"}]},{"elements":[{"text":"保存容器に①のなすを並べ、その上に大葉を乗せる。何層にも重ねたら②をかけ冷蔵庫で1時間寝かせたら、完成！"}]}]}`,
		"https://www.buzzfeed.com/jp/kazuyawakana/burdock-pickled-in-ponzu-sauce":             `{"title":"【美味しすぎて作り置きにならん…！】簡単さっぱり♪ごぼうのポン酢漬け","image":"https://img.buzzfeed.com/buzzfeed-static/static/2022-06/14/2/asset/0d424cf2ac90/sub-buzz-1340-1655172520-17.jpg?downsize=700%3A%2A\u0026output-quality=auto\u0026output-format=auto","servings":1,"ingredients":[{"name":"ごぼう","amount":"1/4本"},{"group":"A","name":"ポン酢","amount":"大さじ2"},{"group":"A","name":"砂糖","amount":"大さじ1/2"},{"group":"A","name":"たかのつめ輪切り","amount":"適量"}],"instructions":[{"elements":[{"text":"ごぼうの皮を丸めたアルミホイルでそぎ落とす。よく洗い、食べやすい大きさに切る。"}]},{"elements":[{"text":"水を入れた鍋に①を入れ、４分ゆでる。"}]},{"elements":[{"text":"ボウルに②とAを入れよく和える。冷蔵庫で1時間寝かせたら、完成！"}]}]}`,
		"https://www.buzzfeed.com/jp/maorikato/easy-pickled-eggplant":                         `{"title":"悪魔のなす漬け","image":"https://img.buzzfeed.com/buzzfeed-static/static/2022-06/13/6/asset/159c3b3a29f7/sub-buzz-7518-1655103572-17.jpg?downsize=700%3A%2A\u0026output-quality=auto\u0026output-format=auto","servings":1,"ingredients":[{"name":"なす1本"},{"name":"ごま油","amount":"小さじ1"},{"group":"A","name":"長ねぎ","amount":"1/4本"},{"group":"A","name":"醤油","amount":"大さじ1"},{"group":"A","name":"水","amount":"大さじ1"},{"group":"A","name":"砂糖","amount":"小さじ1　"},{"group":"A","name":"にんにく（すりおろし）","amount":"小さじ1/2"},{"group":"A","name":"輪切り唐辛子","amount":"適量"},{"group":"A","name":"白ごま","amount":"適量"}],"instructions":[{"elements":[{"text":"なすを縦半分に切り、皮面に斜めに浅く切り込みを入れ、3等分に切る。"}]},{"elements":[{"text":"フライパンにごま油をひき、①をしんなりするまで焼く。"}]},{"elements":[{"text":"保存袋にAと②を入れて1時間置いたら、完成！"}]}]}`,
		"https://www.buzzfeed.com/jp/maorikato/fried-shiso-leaf":                              `{"title":"大葉唐揚げ","image":"https://img.buzzfeed.com/buzzfeed-static/static/2022-06/10/3/asset/159c3b3a29f7/sub-buzz-2812-1654832843-1.jpg?downsize=700%3A%2A\u0026output-quality=auto\u0026output-format=auto","servings":1,"ingredients":[{"name":"鶏もも肉","amount":"150g"},{"group":"A","name":"しょうゆ","amount":"小さじ2"},{"group":"A","name":"みりん","amount":"小さじ2"},{"group":"A","name":"にんにく（すりおろし）","amount":"小さじ1/2"},{"name":"大葉","amount":"5枚"},{"name":"片栗粉","amount":"適量"},{"name":"サラダ油","amount":"適量"}],"instructions":[{"elements":[{"text":"鶏もも肉を一口大に切る。"}]},{"elements":[{"text":"ボールに①、Aを入れて15分置く。"}]},{"elements":[{"text":"水気を切った②に大葉を巻いて片栗粉をまぶし、フライパンにサラダ油を熱し、カリッとするまで3〜4分揚げたら、完成！"}]}]}`,
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
