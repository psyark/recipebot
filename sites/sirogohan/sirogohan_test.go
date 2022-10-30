package sirogohan

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://www.sirogohan.com/recipe/nikudouhu/": {
		Title:    "肉豆腐",
		Image:    "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig6717.JPG",
		Servings: 3,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "牛こま切れ肉", Amount: "200g", Comment: ""},
			{Group: "", Name: "木綿豆腐", Amount: "1と1/2丁(計450gほど)", Comment: ""},
			{Group: "", Name: "玉ねぎ", Amount: "1個", Comment: ""},
			{Group: "", Name: "椎茸", Amount: "4枚", Comment: ""},
			{Group: "", Name: "えのき茸", Amount: "100g", Comment: ""},
			{Group: "A", Name: "醤油", Amount: "大さじ5", Comment: ""},
			{Group: "A", Name: "酒", Amount: "大さじ4", Comment: ""},
			{Group: "A", Name: "みりん", Amount: "大さじ4", Comment: ""},
			{Group: "A", Name: "砂糖", Amount: "大さじ2", Comment: ""},
			{Group: "A", Name: "水", Amount: "100ml", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "肉豆腐の下ごしらえ\n\nこの肉豆腐は26㎝ほどのフライパンで作る工程で紹介しています。フライパンの体積を考えると木綿豆腐１と1/2丁くらいがちょうどよいので、半端ですが一度上記分量で作ってみてください。\n\nはじめに木綿豆腐は1丁を8等分ほどに切って、キッチンペーパーの上にしばらく置いて水を吸わせます。\n\n水気を切ることで煮汁が水っぽくなったりせず、味もしみ込みやすくなります。置く時間は10分ほどでOKです。\n\n\t\n\t\t\n\t\t\n\t\n\n※木綿豆腐の水切りをする時間がないときは、より水気が少ない「焼き豆腐」を買ってきて作っても。\n\n玉ねぎは10～12等分ほどのくし切りにします。椎茸は軸を切り落として半分に、えのき茸は石づきを切り落として食べやすい大きさに手で割いておきます。"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig1.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig4.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig2.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig3.JPG"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "肉豆腐の味付け/レシピ\n\n肉豆腐の味付けは牛肉や野菜からの出汁が出るので、特にだし汁は必要ありません。Aの調味料（醬油大さじ5、みりんと酒が各大さじ4、砂糖大さじ2、水100ml）を合わせて箸で混ぜて軽く溶かし、フライパンを火にかける前に、切った玉ねぎを入れておきます。\n\n\t\n\t\t\n\t\t\n\t\n\nフライパンを中火にかけて沸くのを待ち、煮汁が沸いたら玉ねぎを端に寄せて牛肉を入れるスペースを作ります。\n\n\t\n\t\t\n\t\t\n\t\n\n煮汁が出た空いたスペースに、牛肉を入れて箸でほぐしながら全体的に色が変わるまで火を通します。火が通った牛肉は箸ですくい上げて取り出します。\n\n\t\n\t\t\n\t\t\n\t\n\n※肉豆腐は豆腐や玉ねぎなどやわらかい食感の具材が多いので、牛肉が固くならないよう、はじめにさっと火を通して取り出し、最後に戻し入れて温めるという火の入れ方にしています。\n\n牛肉を取り出したら、玉ねぎを端に寄せたまま豆腐をそっと入れます（玉ねぎは重なってもよいので、豆腐が重ならないように）。豆腐がすべて入れば、上にきのこを広げるようにのせます。\n\n\t\n\t\t\n\t\t\n\t\n\nきのこを上にのせて少し煮るとしんなりしてくるので、きのこをフライパンの端や、豆腐と豆腐の隙間など、入れ込むことができる部分に入れます。\n\n\t\n\t\t\n\t\t\n\t\n\nきのこがある程度煮汁に浸かれば落し蓋をします（落し蓋をすれば自然とカサは減るのですべてが浸かる必要はありません）。蓋の下の煮汁がぐつぐつと沸く火加減で、10～12分煮ます。\n\n\t\n\t\t\n\t\t\n\t\n\n10分ほど煮れば、玉ねぎやきのこもしっかりとやわらかく煮上がり、煮汁も適度に煮詰まってきます。\nきのこや玉ねぎを豆腐の上に移動させ、先に取り出しておいた牛肉を煮汁に戻し入れ、1～2分煮てしっかり温めます。\n\nこれで出来上がり、煮汁をたっぷりかけて盛り付けてください。"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig5.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig6.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig7.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig8.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig9.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig10.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig0.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig12.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig13.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig14.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig15.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig16.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig17.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig18.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig6717small.JPG"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "肉豆腐を丼ぶりにしても！\n\nご飯の上に豆腐、牛肉、野菜をのせて、煮汁をたっぷりかけると牛丼のような美味しい丼ものにもなります。\n好みで一味唐辛子をふりかけても美味しいです。2日目などの食べ方としてぜひ！"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig19.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig20.JPG"}}},
		},
	},
	"https://www.sirogohan.com/sp/recipe/butabaradaikon/amp/": {
		Title:    "豚バラ大根",
		Image:    "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/abutabaradaikon76112.JPG",
		Servings: 2,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "大根", Amount: "600g(約1/2本)", Comment: ""},
			{Group: "", Name: "豚バラ肉薄切り", Amount: "150g", Comment: ""},
			{Group: "", Name: "生姜", Amount: "20gほど", Comment: ""},
			{Group: "", Name: "粗びき黒胡椒", Amount: "好みで少々", Comment: "仕上げ用"},
			{Group: "", Name: "ごま油", Amount: "小さじ1", Comment: ""},
			{Group: "", Name: "だし汁", Amount: "400ml", Comment: ""},
			{Group: "", Name: "砂糖", Amount: "大さじ1と1/2", Comment: ""},
			{Group: "", Name: "醤油", Amount: "大さじ2", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "豚バラ大根の下ごしらえ\n\n用意するものは大根、豚バラ肉（薄切り）、生姜で、大根600ｇに対して豚バラ肉は150ｇほどでOKです。\n煮物にする大根は皮近くにある筋をむき取った方が口当たりがよいので、大根は包丁やピーラーで皮を厚めにむき取ります。\n\n\t\n\t\t\n\t\t\n\t\n\n皮をむいた大根は2㎝幅ほどの半月切りにします（先端側は火が通りやすいので3㎝幅ほどにしています）。\n豚バラ肉は3～4㎝幅に、生姜は皮をむいてせん切りにします。"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon1.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon2.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon3.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon4.JPG"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "豚バラ大根のレシピ/作り方\n\n大きめのフライパンか鍋を用意し、ごま油小さじ1と生姜を入れて中火にかけ、香りが立ってくれば豚バラ肉を入れてほぐしながら炒めます。\n\n\t\n\t\t\n\t\t\n\t\n\n豚バラ肉の色が変わってほぼ火が通れば塩ひとつまみとこしょう少々（各分量外）を加えて下味をつけます。\n続けて大根を加え、中火のまま1分ほど大根と豚バラ肉を炒め合わせます。\n\n\t\n\t\t\n\t\t\n\t\n\nだし汁400mlをそそぎ入れ、沸いてきたら軽くアクをすくい取ります。\n\n\t\n\t\t\n\t\t\n\t\n\n※豚肉などの具材からだしが出るのでだし汁は薄めでもOKです。レシピ下の補足に各種だし汁に関してのリンクを貼っています。\n\nここからは落し蓋をして中で煮汁を対流させながら15分煮ます。\n※落とし蓋の下では下の写真のように煮汁がグツグツと煮立つ状態にすることが大切です。弱火ではなく弱めの中火くらいの火加減で煮ます。\n\n\t\n\t\t\n\t\t\n\t\n\n15分経てば砂糖大さじ1と1/2を加え、煮汁をかけるなどして溶かし混ぜます。\n砂糖を入れたら5分煮るのですが、竹串を大きめの大根に刺してみてすっと刺されば落とし蓋を外して5分煮ていきます。\n\n\t\n\t\t\n\t\t\n\t\n\n※大根にしっかり火が通っていないようなら5分の間は落とし蓋をしたままで煮て、次の醬油を加える時に蓋を外してください。\n\n続けて醤油大さじ2を加え、同じく煮汁をまわしかけるなどして溶かし混ぜます。\nこの段階で煮汁が半分以下になっているのですが、ここから煮汁がさらに少なくなるまで5～7分ほど煮詰めます。\n\n\t\n\t\t\n\t\t\n\t\n\n煮汁が少ないので途中3～4度くらいフライパンをふるか箸で大根の上下を返すなどして、煮汁のしみ込みを均一にするとよいです。\n煮汁がフライパンの底に少し残って、大根や豚肉にしっかり煮汁がからむくらいになれば完成です。器に盛って好みで粗びき黒胡椒などを散らしていただきましょう！"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon5.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon6.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon7.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon8.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon9.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon10.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon11.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon12.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon13.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon14.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon15.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon16.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon17.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/butabaradaikon/butabaradaikon18.JPG"}}},
		},
	},
	"https://www.sirogohan.com/recipe/tukemono-hakusai/": {
		Title:    "簡単！白菜の浅漬け",
		Image:    "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusaiyoko.JPG",
		Servings: 0, // TODO
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "白菜", Amount: "300g(約1/8玉)", Comment: ""},
			{Group: "", Name: "塩", Amount: "野菜の重さの2%", Comment: ""},
			{Group: "", Name: "昆布", Amount: "野菜の重さの0.5~1%ほど", Comment: ""},
			{Group: "", Name: "食べるときに醤油やごまを適量", Amount: "", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "浅漬けの塩加減の基本\n\n※手順は『浅漬けのきほん』がベース。\n3種の漬け方から、“バットを使った漬け方”で紹介しています（ボウルや浅漬け器でも可）。\n\nまず、塩の量は野菜の重さに対して2％の塩を用意してください。\n今回なら白菜300ｇに、塩は2％の6ｇ（塩小さじ1ほど）となります。\n\n\t\n\t\t\n\t\t\n\t\n\n※慣れないうちは、きちんと野菜と塩の重さを量って作るのがおすすめです。また、塩は細かさによって同じ小さじ1でも重さが異なります！さらさらした精製塩などは小さじ1が6ｇ、粒の大きな塩なら小さじ1が4〜5ｇほどに。"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai1.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai2.JPG"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "白菜の浅漬けの下ごしらえ\n\n白菜は根元を切り落としてから水洗いします。特に根元側に土が残っていることがあるので注意して洗うとよいです。洗ったものは3〜4㎝幅に切ります。\n\n\n\t\n\t\t\n\t\t\n\t\t\n\t\n\n一緒に漬け込む昆布は、うま味をプラスできるので、ぜひ少量用意して合わせてみてください。野菜の重さに対して0.5～1％ほどが目安なので、今回は昆布2～3ｇが目安となります。\n※時間があれば昆布を切る前にさっと昆布を直火であぶると柔らかくなって、細く切りやすいです（バリバリとかけらが飛び散りにくくなります）。\n\n\n\t\n\t\t\n\t\t\n\t\n\n\nバットで漬けこむときに袋を使うのですが、袋の中だけでは塩や昆布がよく混ざらないので、袋に入れる前にボウルで混ぜるとよいです。\n（下の写真では手で混ぜていますが、箸を使ってもOK）分量の塩を加えて混ぜたら、昆布を足してジップロックなどの袋に移します。"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/hakusaisai1.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/hakusaisai2.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai6.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai4.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai5.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai7.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai8.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai9.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai10.JPG"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "白菜の漬物の漬け方/レシピ\n\n白菜を袋に入れたらバットに平らに広げ、3〜4倍の重しをかけます。冷蔵庫に移して、3〜4時間以上塩をなじませましょう。\n\n\t\n\t\t\n\t\t\n\t\t\n\t\t\n\t\n\n※丸い石を重石に使ったので、均一に加重されるよう上にもバットをかましています。平たい1㎏の砂糖やお皿などならバットなしで直に重石をしても。\n\n3～4時間ほどで白菜から水気が出てしんなりしてくるので、そのくらいからが食べごろです。\n半日～1日以上置いて、芯の部分にも塩をしっかり馴染ませても美味しいので、好みの加減まで塩をなじませてみてください。\n\n\t\n\t\t\n\t\t\n\t\n\n好みの漬け加減になれば重石を取り除き、軽く空気を抜いて冷蔵庫でそのまま保存します。\n食べるときは食べる分だけを取り出し、汁気をギュッとしぼり出してから器に盛り付けます。好みでほんの少しの醤油やごまをプラスすると美味しさが引き立ちます！"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai11.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai12.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai13.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai14.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai15.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai16.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusai17.JPG"}, &rexch.ImageInstructionElement{URL: "https://www.sirogohan.com/_files/recipe/images/tukemono-hakusai/tukemono-hakusaiyokosmall2.JPG"}}},
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
