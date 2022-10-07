package lettuceclub

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/sites"
)

func TestNewParser(t *testing.T) {
	ctx := context.Background()
	tests := map[string]string{
		"https://www.lettuceclub.net/recipe/dish/24626/": `{"Title":"アマトリチャーナ","Image":"https://www.lettuceclub.net/i/R1/img/dish/1/S20170925009002A2_000.jpg?w=450","IngredientGroups":[{"Name":"","Children":[{"Name":"スパゲッティ(1.6mm)","Amount":"160〜200g","Comment":""},{"Name":"玉ねぎ","Amount":"1/2個","Comment":""},{"Name":"ベーコン","Amount":"4枚","Comment":""},{"Name":"オリーブ油","Amount":"大さじ3","Comment":""},{"Name":"にんにくのみじん切り","Amount":"小さじ2","Comment":""},{"Name":"ホールトマト缶","Amount":"1缶(約400g)","Comment":""},{"Name":"塩","Amount":"小さじ1/4","Comment":""},{"Name":"こしょう","Amount":"少々","Comment":""},{"Name":"粉チーズ","Amount":"適量","Comment":""},{"Name":"粗塩","Amount":"大さじ1","Comment":""}]}],"Steps":[{"Text":"鍋に湯1.6Lを沸かし、粗塩大さじ1を加える。スパゲッティを加えてさっと混ぜ、袋の表示より1〜2分短くゆでる。玉ねぎは縦薄切りにする。ベーコンは1cm幅に切る。","Images":null},{"Text":"フライパンにオリーブ油大さじ3、にんにくのみじん切り小さじ2、玉ねぎを入れて中火にかけ、しんなりするまで約2分炒める。ベーコンを加えてさっと炒め、ホールトマト缶を加えて潰す。時々混ぜながら3〜4分煮て、塩小さじ1/4、こしょう少々をふる。","Images":null},{"Text":"スパゲッティの湯をきって2に加え、さっとあえる。器に盛り、粉チーズ適量をふる。","Images":null}]}`,
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
