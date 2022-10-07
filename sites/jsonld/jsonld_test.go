package jsonld

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/sites"
)

func TestNewParser(t *testing.T) {
	ctx := context.Background()
	tests := map[string]string{
		"https://s.recipe-blog.jp/profile/313934/recipe/1432314": `{"Title":"自家製ごまダレで、牛肉と水菜の簡単しゃぶしゃぶ","Image":"https://asset.recipe-blog.jp/cache/images/recipe/bc/ae/fe6575effb49833f63fea6b56510cf2f8e21bcae.640x640.cut.jpg","IngredientGroups":[{"Name":"","Children":[{"Name":"牛こま切れ","Amount":"３４０g","Comment":""},{"Name":"水菜","Amount":"１束","Comment":""},{"Name":"Ａごま","Amount":"大さじ１","Comment":""},{"Name":"Ａポン酢","Amount":"大さじ１","Comment":""},{"Name":"Ａマヨネーズ","Amount":"大さじ１","Comment":""},{"Name":"Ａ砂糖","Amount":"大さじ１","Comment":""},{"Name":"Ａ味噌","Amount":"小さじ２分の１","Comment":""},{"Name":"ポン酢","Amount":"適量","Comment":""}]}],"Steps":[{"Text":"水菜はよく洗い、3㎝幅に切り、熱湯で水菜を１分ほど茹でる。冷水に取り、水気をぎゅっと絞る。牛肉こま切れは、熱湯でさっと茹でて水気を切る。","Images":null},{"Text":"１をお皿に盛り付ける。","Images":null},{"Text":"Ａをよく混ぜ合わせ、つけダレにする。ポン酢もつけダレにする。","Images":null}]}`,
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
