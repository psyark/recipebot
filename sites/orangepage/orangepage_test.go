package orangepage

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/sites"
)

func TestNewParser(t *testing.T) {
	ctx := context.Background()
	tests := map[string]string{
		"https://www.orangepage.net/ymsr/news/daily/posts/5716": `{"title":"玉ねぎ好きの飛田和緒さん直伝『丸ごと玉ねぎのとろとろカレー』がおいしすぎ！","image":"https://images.orangepage.net/media/article/5716/images/main_626f06d7f88bb6c7d7a2b0815792876c.jpg?d=960x540","ingredients":[{"name":"玉ねぎ","comment":"小）……５個（約５００ｇ"},{"name":"ブロックベーコン……４０ｇ"},{"name":"にんにくのみじん切り……１かけ分"},{"name":"しょうがのみじん切り……１かけ分"},{"name":"カレー粉……小さじ２"},{"name":"あればガラムマサラ……小さじ１／２"},{"name":"塩……小さじ１／３～１／２"},{"name":"しょうゆ……小さじ２"},{"name":"温かいご飯……適宜"},{"name":"バター……２０ｇ"}],"instructions":[{"elements":[{"text":"（1）"}]},{"elements":[{"text":"玉ねぎ1個はみじん切りにする。残りは、根元を薄く切り落とす。ベーコンは5mm角の棒状に切る。"}]},{"elements":[{"text":"（2）"}]},{"elements":[{"text":"鍋ににんにく、しょうが、バターを入れ、弱火にかける。香りが立ったらみじん切りの玉ねぎを加えて炒める。途中ふたをして蒸し焼きにしながら、全体に色づくまで15～20分炒める。"}]},{"elements":[{"text":"POINT　みじん切りの玉ねぎは、あめ色になるまでじっくり炒めて。"}]},{"elements":[{"text":"（3）"}]},{"elements":[{"text":"ベーコンとカレー粉、あればガラムマサラを加えて炒め、カレーの香りが立ったら水300mlと残りの玉ねぎを加える。弱めの中火にして煮立ったらふたをし、玉ねぎが柔らかくなるまで30分ほど煮る。途中、煮汁が煮つまりすぎていたら、適宜水をたす。"}]},{"elements":[{"text":"（4）"}]},{"elements":[{"text":"玉ねぎにすっと竹串が通るくらいになったら、味をみてから塩としょうゆを加える。ご飯を器に盛り、カレーをかける。"}]},{"elements":[{"text":"メインの具は玉ねぎとベーコンのみと、本当にシンプル。"}]},{"elements":[{"text":"けれどちょっとしたひと工夫で、いつものカレーがこんなそそる一皿になるんですねー。"}]},{"elements":[{"text":"これは絶対に試してみるべき！"}]},{"elements":[{"text":"『丸ごと玉ねぎのとろとろカレー』ぜひ作ってみて下さいねー♪"}]},{"elements":[{"text":"（『オレンジページCooking2022夕飯　夕飯、即決、迷わない。』より）"}]}]}`,
		"https://www.orangepage.net/ymsr/news/daily/posts/5763": `{"title":"【白メシが美味い！ 秋おかず】絶品『れんこんと豚肉の甘辛揚げ』のレシピ","image":"https://images.orangepage.net/media/article/5763/images/main_cf8cf4f59904c49016298e414deccdef.jpg?d=960x540","ingredients":[{"name":"れんこん……１節","comment":"約１５０ｇ"},{"name":"豚ロース肉","comment":"とんカツ用）……２枚（約２００ｇ"},{"name":"〈甘辛だれ〉"},{"name":"砂糖……小さじ２"},{"name":"酢……小さじ２"},{"name":"しょうゆ……大さじ１"},{"name":"酒"},{"name":"しょうゆ"},{"name":"小麦粉"},{"name":"サラダ油"}],"instructions":[{"elements":[{"text":"（1）材料の下ごしらえをする"}]},{"elements":[{"text":"れんこんはよく洗い、皮つきのまま幅1cmの輪切りにする。豚肉は大きめの一口大に切る。ボールに豚肉を入れ、酒、しょうゆ各大さじ1/2をもみ込み、15分おく。汁けをかるくきって小麦粉を薄くまぶす。"}]},{"elements":[{"text":"（2）フライパンで揚げる"}]},{"elements":[{"text":"フライパンにサラダ油を高さ2cmほど入れて低温（約160℃。乾いた菜箸の先を底に当てると、細かい泡がゆっくりと揺れながら出る程度）に熱し、れんこんを入れてきつね色になるまで5～6分揚げ、油をきる。油を中温（約170℃。乾いた菜箸の先を底に当てると、細かい泡がシュワシュワッとまっすぐ出る程度）に熱し、豚肉を2～3分揚げ、油をきる。"}]},{"elements":[{"text":"（3）たれをからめる"}]},{"elements":[{"text":"ボールにたれの材料を混ぜ合わせ、れんこんと豚肉を加えてからめる。たれごと器に盛る。"}]},{"elements":[{"text":"とんカツ用肉を使うので食べごたえも満点！"}]},{"elements":[{"text":"「食欲の秋」という通り、本当に箸が止まらなくなりそうなおいしさです。"}]},{"elements":[{"text":"たれをバウンドさせたご飯といっしょにかき込む至福のひとときを、ぜひご体感ください！"}]},{"elements":[{"text":"（『オレンジページ』2022年10月17日号より）"}]}]}`,
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
