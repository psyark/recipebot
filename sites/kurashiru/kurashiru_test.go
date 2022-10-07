package kurashiru

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/sites"
)

func TestNewParser(t *testing.T) {
	ctx := context.Background()
	tests := map[string]string{
		"https://www.kurashiru.com/recipes/9d2cec44-130f-49e4-880c-b1ed15eec39b": `{"Title":"オレンジチキン","Image":"https://video.kurashiru.com/production/videos/9d2cec44-130f-49e4-880c-b1ed15eec39b/compressed_thumbnail_square_large.jpg?1639367889","IngredientGroups":[{"Name":"","Children":[{"Name":"鶏もも肉","Amount":"300g","Comment":""},{"Name":"塩こしょう","Amount":"小さじ1/4","Comment":""},{"Name":"片栗粉","Amount":"大さじ2","Comment":""},{"Name":"揚げ油","Amount":"適量","Comment":""}]},{"Name":"ソース","Children":[{"Name":"１００％オレンジジュース","Amount":"30ml","Comment":""},{"Name":"しょうゆ","Amount":"大さじ2","Comment":""},{"Name":"酢","Amount":"大さじ2","Comment":""},{"Name":"はちみつ","Amount":"大さじ1","Comment":""},{"Name":"すりおろしニンニク","Amount":"小さじ1","Comment":""},{"Name":"すりおろし生姜","Amount":"小さじ1","Comment":""}]},{"Name":"付け合せ","Children":[{"Name":"ベビーリーフ","Amount":"30g","Comment":""}]}],"Steps":[{"Text":"鶏もも肉は一口大に切ります。塩こしょうをふり、片栗粉をまぶします。","Images":null},{"Text":"鍋底から3cm程の高さまで揚げ油を注ぎ、180℃に加熱します。1を入れ、鶏もも肉に火が通るまで5分ほど揚げたら油切りをします。","Images":null},{"Text":"ソースを作ります。フライパンにソースの材料を入れて中火で熱します。","Images":null},{"Text":"沸騰してから1分ほど中火で加熱し、2を入れます。よく絡め、全体に味がなじんだら火から下ろします。","Images":null},{"Text":"ベビーリーフと共にお皿に盛り付けてできあがりです。","Images":null}]}`,
		"https://www.kurashiru.com/recipes/e6c3ef62-8e77-4fed-9ab4-705a1ec78fd3": `{"Title":"とろーりおいしい！肉巻き半熟卵","Image":"https://video.kurashiru.com/production/videos/e6c3ef62-8e77-4fed-9ab4-705a1ec78fd3/compressed_thumbnail_square_large.jpg?1649747039","IngredientGroups":[{"Name":"","Children":[{"Name":"豚バラ肉","Amount":"200g","Comment":"スライス"},{"Name":"塩こしょう","Amount":"小さじ1/4","Comment":""},{"Name":"卵","Amount":"4個","Comment":"Ｍサイズ"},{"Name":"お湯","Amount":"1000ml","Comment":"卵をゆでる用"},{"Name":"薄力粉","Amount":"大さじ1","Comment":""},{"Name":"サラダ油","Amount":"小さじ2","Comment":""},{"Name":"レタス","Amount":"20g","Comment":""},{"Name":"ミニトマト","Amount":"2個","Comment":""}]},{"Name":"（Ａ）","Children":[{"Name":"酒","Amount":"大さじ1","Comment":""},{"Name":"砂糖","Amount":"小さじ2","Comment":""},{"Name":"みりん","Amount":"小さじ2","Comment":""},{"Name":"しょうゆ","Amount":"大さじ1.5","Comment":""}]}],"Steps":[{"Text":"卵は常温に戻しておきます。","Images":null},{"Text":"お湯を沸騰させ、卵を7分ゆでて流水で冷やし、殻を剥きます。","Images":null},{"Text":"豚バラ肉に塩こしょうをふり、1を巻いて、全体に薄力粉をまぶします。","Images":null},{"Text":"中火に熱したフライパンにサラダ油をひき、2を焼きます。","Images":null},{"Text":"全体に焼き色がつき豚バラ肉に火が通ったら、(A)を入れて、中火で全体を煮詰めて火から下ろします。","Images":null},{"Text":"お皿に盛り付け、レタス、ミニトマトを添えて完成です。","Images":null}]}`,
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
