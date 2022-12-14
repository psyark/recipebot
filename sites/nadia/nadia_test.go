package nadia

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://oceans-nadia.com/user/22780/recipe/294470": {
		Title:    "なすと豚バラの甘酢照り焼き＊おろし添え【#作り置き】",
		Image:    "https://asset.oceans-nadia.com/upload/save_image/69/69cc2952b5de.JPG",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "豚バラ薄切り肉", Amount: "200g", Comment: ""},
			{Group: "", Name: "なす", Amount: "3本(240g)", Comment: ""},
			{Group: "", Name: "大根", Amount: "4〜5cm", Comment: ""},
			{Group: "A", Name: "しょうゆ", Amount: "大さじ2", Comment: ""},
			{Group: "A", Name: "砂糖、みりん、酢", Amount: "各大さじ1.5", Comment: ""},
			{Group: "A", Name: "片栗粉", Amount: "小さじ1", Comment: ""},
			{Group: "", Name: "いり白ごま", Amount: "適量", Comment: ""},
			{Group: "", Name: "ごま油", Amount: "小さじ2", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "tips", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "★豚バラ薄切り肉は、しゃぶしゃぶ用を使用しました。長いものをご使用になる際は、食べやすく切ってください。\r\n\r\n★甘めの味付けです。お好みで砂糖やみりんの量を調整してください。\r\n\r\n★フライパンは、26cmのものを使用しました。\r\n\r\n★日持ちは、冷蔵庫で2〜3日です。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "なすはヘタを取り、一口サイズの乱切りにする。大根は皮をむいてすりおろす。"}, &rexch.ImageInstructionElement{URL: "https://asset.oceans-nadia.com/upload/save_image/0d/0d370c7887ae.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンにごま油を中火で熱し、豚バラ薄切り肉を炒める。豚肉の色が8割がた変わったらなすを加えてサッと炒める。全体に油が回ったら蓋をし、弱火で4〜5分蒸し焼きにする。"}, &rexch.ImageInstructionElement{URL: "https://asset.oceans-nadia.com/upload/save_image/2c/2c5fe2dec5e4.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ペーパータオルで余分な油を拭き取り、合わせたA しょうゆ大さじ2、砂糖、みりん、酢各大さじ1.5、片栗粉小さじ1を回し入れて炒める。とろみとツヤが出たら、仕上げにいり白ごまをふる。"}, &rexch.ImageInstructionElement{URL: "https://asset.oceans-nadia.com/upload/save_image/22/22474e5ec1f3.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "器に盛り、大根おろしを添えてお召し上がりください♪\n\nあれば青じそを添えると、さらにさっぱり食べられます( ´艸`)"}, &rexch.ImageInstructionElement{URL: "https://asset.oceans-nadia.com/upload/save_image/d7/d7abe2804c15.JPG"}}},
		},
	},
	"https://oceans-nadia.com/user/543935/recipe/416620": {
		Title:    "豚バラしそチーズ巻き",
		Image:    "https://asset.oceans-nadia.com/upload/save_image/2a/2a0ad183b6cbc6fdac19788d420d3f09.jpeg",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "豚バラ薄切り肉", Amount: "250g", Comment: ""},
			{Group: "", Name: "大葉", Amount: "6枚", Comment: ""},
			{Group: "", Name: "スライスチーズ", Amount: "3枚", Comment: ""},
			{Group: "", Name: "片栗粉", Amount: "大さじ1", Comment: ""},
			{Group: "", Name: "酒", Amount: "小さじ2", Comment: ""},
			{Group: "", Name: "塩、コショウ", Amount: "各少々", Comment: ""},
			{Group: "A", Name: "ポン酢醤油", Amount: "大さじ2", Comment: ""},
			{Group: "A", Name: "みりん", Amount: "大さじ1", Comment: ""},
			{Group: "A", Name: "にんにく（すりおろす）", Amount: "1片分", Comment: ""},
			{Group: "A", Name: "しょうが（すりおろす）", Amount: "1かけ分", Comment: ""},
			{Group: "", Name: "ごま油", Amount: "大さじ1", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "tips", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "☆片栗粉をまぶすことで具材と肉をくっつけて肉巻きの形が崩れにくくなります。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "豚バラ薄切り肉をまな板の上で広げて塩、コショウ、酒を振りかける。"}, &rexch.ImageInstructionElement{URL: "https://asset.oceans-nadia.com/upload/save_image/9c/9cdd0a381d641966e329ad2fcd20cc36.jpeg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "片栗粉を全体にまぶし、上に大葉、スライスチーズを重ねる。"}, &rexch.ImageInstructionElement{URL: "https://asset.oceans-nadia.com/upload/save_image/c6/c6224ba3add17f4331f2b8937e71de5e.jpeg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "手前から奥に少しきつめにしっかり巻く。"}, &rexch.ImageInstructionElement{URL: "https://asset.oceans-nadia.com/upload/save_image/4a/4a4b1068530a4b8cd5741ddf702bd741.jpeg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンを中火で加熱し、ごま油を入れる。②の巻き終わりを下にして約2分焼く。"}, &rexch.ImageInstructionElement{URL: "https://asset.oceans-nadia.com/upload/save_image/01/014384a22dbd8be00d592d54ffef0fb8.jpeg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "焦がさないように転がしながら全体に焼き色が付くまで焼く。キッチンペーパーで余分な油を拭き取りA ポン酢醤油大さじ2、みりん大さじ1、にんにく（すりおろす）1片分、しょうが（すりおろす）1かけ分を合わせたたれをかけて煮絡める。"}, &rexch.ImageInstructionElement{URL: "https://asset.oceans-nadia.com/upload/save_image/ec/ece364071408259a9cd8f6d047d68754.jpeg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "食べやすい大きさに切り器に盛り付ける。"}, &rexch.ImageInstructionElement{URL: "https://asset.oceans-nadia.com/upload/save_image/c5/c51fda39d4af0aad31246d0e01fa291b.jpeg"}}},
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
