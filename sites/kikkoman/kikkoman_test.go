package kikkoman

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://www.kikkoman.co.jp/homecook/search/recipe/00004691/index.html": {
		Title:    "基本の肉じゃが【味しみ！定番人気和食】",
		Image:    "https://www.kikkoman.co.jp/homecook/search/recipe/img/00004691.jpg",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "じゃがいも", Amount: "３個", Comment: ""},
			{Group: "", Name: "玉ねぎ", Amount: "1/2個", Comment: ""},
			{Group: "", Name: "にんじん", Amount: "1/2本", Comment: ""},
			{Group: "", Name: "牛肩肉(薄切り・切り落とし)", Amount: "１００ｇ", Comment: ""},
			{Group: "", Name: "しらたき", Amount: "１００ｇ", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "小さじ２", Comment: ""},
			{Group: "", Name: "かつおだし", Amount: "１と1/2カップ", Comment: ""},
			{Group: "", Name: "醤油", Amount: "大さじ２", Comment: ""},
			{Group: "", Name: "みりん", Amount: "大さじ３", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "じゃがいもはひと口大に切って水にさらし、水気をきる。玉ねぎはくし形切り、にんじんは乱切りにする。牛肉はひと口大に切る。しらたきはゆでて食べやすく切る。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "鍋にサラダ油を熱して（１）の玉ねぎを炒め、牛肉を加えてさらに炒める。にんじん、じゃがいも、しらたきも入れて炒め合わせる。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "かつおだしを注ぎ、沸騰したらアクを取り、しょうゆ、みりんを加えて落しぶたをする。沸騰したら弱火で１５分くらい煮る。"}}},
		},
	},
	"https://www.kikkoman.co.jp/homecook/search/recipe/00052848/index.html": {
		Title:    "クセになるおいしさ！えのきのカリカリ焼き",
		Image:    "https://www.kikkoman.co.jp/homecook/search/recipe/img/00052848.jpg",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "えのきたけ", Amount: "２００ｇ", Comment: ""},
			{Group: "", Name: "片栗粉", Amount: "大さじ４", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "大さじ２", Comment: ""},
			{Group: "(A)", Name: "醤油", Amount: "大さじ１", Comment: ""},
			{Group: "(A)", Name: "料理酒", Amount: "大さじ１", Comment: ""},
			{Group: "(A)", Name: "おろししょうが", Amount: "小さじ１", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "えのきたけは石づきを落とし、根元がつながったままの状態で縦に１０個に裂く。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "バットに（Ａ）を入れて混ぜ合わせ、（１）を並べて全体にからめ１０分程つけておく。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "汁気をかるくきって片栗粉をまぶす。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンに油を中火で熱し、（３）を並べる。時々フライ返しなどで押さえながら、カリッとするまで両面焼く。"}}},
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

			rex, err := NewParser().Parse2(ctx, url)
			if err != nil {
				t.Fatal(err)
			}

			if err := sites.RecipeMustBe2(want, rex); err != nil {
				t.Error(err)
			}
		})
	}
}
