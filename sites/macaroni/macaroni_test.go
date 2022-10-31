package macaroni

import (
	"context"
	"testing"

	"github.com/psyark/recipebot/rexch"
	"github.com/psyark/recipebot/sites"
)

var tests = map[string]*rexch.Recipe{
	"https://macaro-ni.jp/109611": {
		Title:    "とろ〜り半熟卵で作る「ウフ・アン・ムーレット」【秋元さくらシェフ直伝】",
		Image:    "https://cdn.macaro-ni.jp/image/summary/109/109611/BN1C2JEIn3tfIUOSJZUGolyDUYOXsDN6BY3HjDBp.jpg?p=1x1",
		Servings: 2,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "玉ねぎ", Amount: "1/6個", Comment: "みじん切り"},
			{Group: "", Name: "マッシュルーム", Amount: "6個", Comment: "薄切り"},
			{Group: "", Name: "ベーコン", Amount: "40g", Comment: "５ｍｍ幅・拍子切り"},
			{Group: "", Name: "にんにく", Amount: "少々", Comment: ""},
			{Group: "", Name: "赤ワイン", Amount: "400cc", Comment: ""},
			{Group: "", Name: "水溶きコーンスターチ", Amount: "小さじ1杯(コーンスターチ:小さじ1/2杯、水:小さじ1/2杯)", Comment: ""},
			{Group: "", Name: "オリーブオイル", Amount: "大さじ1杯", Comment: ""},
			{Group: "", Name: "バター", Amount: "20g", Comment: "有塩"},
			{Group: "", Name: "塩", Amount: "適量", Comment: ""},
			{Group: "〈ポーチドエッグ 〉", Name: "卵", Amount: "2個", Comment: ""},
			{Group: "〈ポーチドエッグ 〉", Name: "水", Amount: "2000cc", Comment: ""},
			{Group: "〈ポーチドエッグ 〉", Name: "酢", Amount: "大さじ1杯", Comment: ""},
			{Group: "", Name: "クルトン", Amount: "適宜", Comment: ""},
			{Group: "", Name: "パセリ", Amount: "適宜", Comment: "みじん切り"},
		},
		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "1. 具材を炒める"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/2tsacMfe0uGbQGt9gxXOW3T3IvU9NEmH1JkEYSgA.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "鍋にオリーブオイルを入れ、玉ねぎ、にんにくを加えて炒めます。にんにくの香りを引き出しながら、玉ねぎがくったりするまで炒めます。\n「鍋肌の温度が上がるまでは強火にしましょう。ジリジリっとした音が聞こえてきたら中火にします。今回は家庭料理なので玉ねぎを使用しますが、本来はエシャロットを使います。\n炒めるときの注意点は、あまりぐるぐるかきまぜないこと。触れば触るほど鍋中の温度は下がるので、かき混ぜるのは時々にしてくださいね」"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/fAtKszCFK5BW2is92bvWKEdQWUtUcHojthu1RdYm.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "バター、マッシュルーム、ベーコンを加えて、さらに炒めます。\n「素材の水分を出すために通常は塩をふりますが、今回はなし。ソース作りの過程でワインをかなり煮詰めていくので、ここで塩分を入れると塩味が強くなりがちなんです」"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/TEPWPU31mmym2HBQSIZm6lmrGrdl3eJQy0Dl5wVa.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "「マッシュルームから水分が出てきます。この水分を生かしながら、蒸すようにじっくり炒めて、それぞれの素材の甘みを引き出しましょう」"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/dw3TGr0oexZd1sSUDzPOgu5cWQC0QNgAx2WvgSNy.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "「甘味は旨味でもあります。きのこのグルタミン酸、ベーコンのアミノ酸の旨味をじっくり濃縮するのがポイントです。1/3量になるまで炒めてくださいね」"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "2. 赤ワインを加えて煮詰める"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/sYfFTga1KGecbXr8uIGb3ka79k95SO3MgkmSpDUK.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "赤ワインを加えて、中火でコトコト煮詰めていきます。赤ワインが1/6量になるぐらいが目安です。\n「このお料理はブルゴーニュ地方のスペシャリテなので、同じ土地で生まれたブルゴーニュのワインだといいですね！高いワインを使うともちろんおいしいソースができますが、赤ワインならどんなものでもOKです」"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "3. ポーチドエッグを作る"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/L8YTIVGx5JJyRLInfLsJd70agBynPRiMUYHk258J.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "鍋に2リットルの熱湯を沸かし、大さじ1杯の酢を加え、おたまを入れてぐるぐる混ぜて渦を作ります。ポーチドエッグは1個ずつ作るので、小ぶりな容器に卵を1個割り入れます。\n「塩も加えるレシピもありますが、シンプルに酢だけを使いましょう。いろいろな組み合わせで試しましたが、大きな違いはありませんでした。それなら簡単なほうがいいですよね」"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/rzRaq7X2XjKCD2cSBTkenp347VSjqMzGYYQfEzze.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "鍋の中央にできた渦の中に卵を落とし、1分半ほどゆでます。\n「成功の秘訣は、よく冷えた新しい卵を使うことです。お鍋の大きさによりますが、卵の高さの3倍量のお湯は必須ですね」"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/EVuQOquPuyDyI5r9xVpE47k9Em0x9ZoOczUrIfCL.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "1分半たったら、穴の空いたレードルなどでポーチドエッグを引き上げます。氷水にとらずに、バットや皿におきましょう。\n「温かいまま食べていただく料理なので、氷水にはおとさないでくださいね。予熱でいい感じの半熟卵になります」"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/8edJc7cLuA7tYXkI5zBSs0w3TXGS5S02D01shnpT.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "白身の余分な部分はスプーンやキッチンばさみなどを使って取り除き、深みのあるお皿に盛り付けます。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "4. 2にコーンスターチを加える"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/Mi5jt0Sn87Mjl115mQIuUMcLKrccjJ8eJelW42tW.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "赤ワインの分量が1/6ほどになり、フランス語でミロワールといわれるつやつやした状態になったら、水溶きコーンスターチを加えてとろみをつけます。塩で味を調えればソースのできあがりです。\n「本来は大量のバターを加えてとろみをつけますが、今回はヘルシーに水溶きコーンスターチで代用しましょう。ソースは酸っぱく感じるかもしれませんが、半熟卵と食べるとちょうどよいと思います。酸味のあるソースと濃厚な卵黄の相性は抜群にいいんです！」"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "5. 盛り付ける"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/hcBCYdLdQiGbc5RM7hlmWmfjDJwiZ6mrtM1Z0DZ2.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "ポーチドエッグを囲むようにムーレットソースをかけ、中央にクルトン、パセリのみじん切りを飾ります。ウフ・アン・ムーレットの完成です！\n食材のコクや旨味が凝縮。半熟卵のぜいたくなひと皿"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/WMb42PpekrPTWpccjRVYylU7fh6BvLQKrpWfLB9d.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "濃厚な卵黄の甘味やコクが、赤ワインベースのムーレットソースの酸味と絡み、なんとも大人な味わい。ワイン、そしてこんがり焼いた薄切りバゲットと一緒に、半熟卵の魅力を楽しめる料理です。卵のイエローとソースのワインレッドの色合いも美しい！\n「日本ではまだあまり知られていないお料理ですが、使っているのは基本の素材ばかりなので、ワインのお供としてご家庭でのパーティの前菜として、気軽に取り入れていただけたらと思います。冷たくてもおいしいので、作り置きにも適していますよ。\nマッシュルームはしめじやまいたけに変えてもOKです。このソースはお肉との相性も抜群。少し多めに作ってお肉に添えるのもおすすめです」\n教えてくれた人"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/109/109611/EMlBngOLcsPWUtKKbJmYv8hxopkiFYEj2V9Cy1Rw.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "「モルソー」オーナーシェフ／秋元さくらさん\n福井県出身。国際線のCAを経て料理人に。新宿「モンドカフェ」、白金「オーギャマンドトキオ」など都内フレンチの名店で腕を磨く。2009年にフレンチビストロ「モルソー」を目黒にオープンし、19年に東京ミッドタウン日比谷に移転。メディアからもひっぱりだこの人気女性シェフ\n取材協力\n取材・文／古川あや\n撮影／宮本信義\n関連レシピはこちら▼"},
			}},
		},
	},
	"https://macaro-ni.jp/90067": {
		Title:    "わさびがポイント！お豆腐の大葉肉巻き",
		Image:    "https://cdn.macaro-ni.jp/image/summary/90/90067/yw8njYPbCLPv97YaraNnNyHQQuliiNRYuGHAV9My.jpeg?p=1x1",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "木綿豆腐", Amount: "300g", Comment: ""},
			{Group: "", Name: "豚バラ肉", Amount: "200g", Comment: "薄切り"},
			{Group: "", Name: "大葉", Amount: "10枚", Comment: ""},
			{Group: "", Name: "片栗粉", Amount: "適量", Comment: ""},
			{Group: "a.", Name: "酒", Amount: "大さじ1杯", Comment: ""},
			{Group: "a.", Name: "みりん", Amount: "大さじ1杯", Comment: ""},
			{Group: "a.", Name: "砂糖", Amount: "大さじ1杯", Comment: ""},
			{Group: "a.", Name: "しょうゆ", Amount: "大さじ2杯", Comment: ""},
			{Group: "", Name: "わさび", Amount: "大さじ1/2杯", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "豆腐をキッチンペーパーでくるみ、レンジ600Wで3分加熱します。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/90/90067/ekhmUcHBt9BZjyjMkeQ0anUcPf4zDwiW50cJkbuj.jpeg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "重石をして10分ほどおき、しっかり水切りをして1cm幅に切ります。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/90/90067/xxt0MRAfIq3cBn552blhp3Z62sEyFiAkESYLBhlX.jpeg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "豚肉に大葉、豆腐をのせて巻き上げ、片栗粉をまぶします。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/90/90067/lxE9asgFlMtHHYFNYxXglLhX49x8ujwJqgkkQDVi.jpeg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンにサラダ油を引いて熱し、③を並べて焼き色がつくまで焼きます。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/90/90067/lHi9RkopTzOU57FWnoNN0JVPPpkCjpf4Gf9XQyqZ.jpeg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "(a) を加えて全体がからんだら火を弱め、わさびを加えて全体にからめたら完成です。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/90/90067/x3UpTsrLKH0C4o4VilxiiVJrABJdbC76jeDHm75h.jpeg"}}},
		},
	},
	"https://macaro-ni.jp/85520": {
		Title:    "メイン食材1つ。豆腐のスティックフライ",
		Image:    "https://cdn.macaro-ni.jp/image/summary/85/85520/488fa7c470eb630b98de98fc66548774.jpg?p=1x1",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "木綿豆腐", Amount: "200g", Comment: ""},
			{Group: "", Name: "片栗粉", Amount: "適量", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "適量", Comment: ""},
			{Group: "", Name: "塩", Amount: "少々", Comment: ""},
			{Group: "", Name: "粗挽き黒こしょう", Amount: "少々", Comment: ""},
			{Group: "", Name: "ガーリックパウダー", Amount: "少々", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "キッチンペーパーで豆腐を包み、耐熱皿にのせて電子レンジ500Ｗ3分加熱し、水切りをします。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "取り出したら1cm角の棒状に切ります。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "片栗粉を全体にまぶします。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンに170℃のサラダ油を熱し、③を加えて揚げ焼きにします。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "油をよく切り、塩、こしょう、ガーリックパウダーを振ったら完成です。"}}},
		},
	},
	"https://macaro-ni.jp/35774": {
		Title:    "必見！コツをおさえた『基本の鮭フライ＆タルタルソース』の作り方",
		Image:    "https://cdn.macaro-ni.jp/image/summary/35/35774/2a8f6aae712a66cdf5c8485a93a3f044.jpg?p=1x1",
		Servings: 2,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "鮭", Amount: "3〜4切れ", Comment: ""},
			{Group: "", Name: "薄力粉", Amount: "大さじ1", Comment: ""},
			{Group: "", Name: "卵", Amount: "1玉", Comment: ""},
			{Group: "", Name: "酒", Amount: "大さじ1", Comment: ""},
			{Group: "", Name: "揚げ油", Amount: "適量", Comment: ""},
			{Group: "Ａ", Name: "パン粉", Amount: "1/2カップ", Comment: ""},
			{Group: "Ａ", Name: "粉チーズ", Amount: "大さじ1", Comment: ""},
			{Group: "", Name: "レモン", Amount: "お好みで", Comment: ""},
			{Group: "＜タルタルソース＞", Name: "ゆで卵", Amount: "1個", Comment: ""},
			{Group: "＜タルタルソース＞", Name: "玉ねぎ", Amount: "1/4玉", Comment: ""},
			{Group: "Ｂ", Name: "酢", Amount: "小さじ1/2", Comment: ""},
			{Group: "Ｂ", Name: "マヨネーズ", Amount: "大さじ3", Comment: ""},
			{Group: "Ｂ", Name: "塩胡椒", Amount: "少々", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "下準備\n・溶き卵に酒を加えてよく混ぜておく。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/888df0a1adc2bccd77f08abe78b43cbe.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "衣の卵に酒を加えることで、卵のみよりも揚げた時に水分が早く抜けて、さっくりした食感に仕上がります。\n・パン粉に粉チーズを加え、混ぜる。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/502d8c84bdd7e156f9ff1ecba72fc92f.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "パン粉に粉チーズを加えることで、味にコクが生まれます。\nたっぷり入れなくても、ほんの少量で十分です。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "①鮭は食べやすいよう、1切れを2〜3切れに切る。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/567f5cb41b2286ae60d4d50922f023b9.jpg?p=medium"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "②塩胡椒(分量外)を全体にまんべんなく振り、5分ほど置く。水分が出るのでキッチンペーパー等で軽く拭きとる。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/c7e74110c599d13774d13970fdad4d06.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "塩胡椒を振って水分を出すことで生臭みを取るので、必ずこの行程は行いましょう。\nまた、塩を振ることで旨味も増します。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "③袋に鮭と薄力粉を入れ、薄力粉を全体にまぶす。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/6a9cde73befc66ba24908e8d83b39c63.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "空気を含んだ状態で口を閉じ、全体を揺するようにすると薄力粉がきれいに付きます。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "④袋から鮭を取り出し、一つずつ軽く叩いて余分な薄力粉を落とす。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/00bfda22dee251416d728d9c2c7cdded.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "薄力粉・卵・パン粉の3つとも、余分に付いていると揚げた時にムラができます。\nきれいな仕上がりにするために、軽く叩きましょう。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "⑤鮭を卵にくぐらせる。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/28d95776755897528b210f32c65f6e34.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "薄力粉と同じく、卵はくぐらせた後によく水気を切りましょう。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "⑥粉チーズを混ぜたパン粉に鮭をのせ、パン粉を全体によくまぶす。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/a12119be3895d3041acf3bf1898244cf.jpg?p=medium"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/8095dd34daaa3e9e013297131193770c.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "パン粉を付けたら、鮭は軽く振って余分なパン粉を落とし、別の容器に置きましょう。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "⑦170℃に熱した油で、きつね色になるまでじっくり揚げる。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/d7358658a5b263523fdb003fb0a4ee04.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "火加減の目安は中火です。油が十分に温まり、なおかつ高温すぎないのがポイント。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "⑧ふちがきつね色になったらひっくり返し、反対側も同じように揚げる。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/40a21368c402b5b0d98c0f35a0723550.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "ふちがきつね色になってからひっくり返すと、火加減が均等で色も良く仕上がります。\nひっくり返す時に、衣が硬くサクッとしていればOK。\n衣が柔らかい場合は火が弱いので、火加減を調整しましょう。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "⑨バットに上げて、余分な油を落とす。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/854e7a24b157f83b5bb4789e137cfa9c.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "油切れが悪いと、ベタついたり胃もたれの原因になります。\n必ず油を落としましょう。\n次にタルタルソースを作ります。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "⑩玉ねぎはみじん切りにして、水にさらす。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/b820fa7fcd527d067253ca6a4a0160af.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "玉ねぎの辛味を抜きます。時間がない場合は、みじん切りして塩をまぶし、軽く握りつぶしてから水にさらすと繊維が潰れて辛味が早く抜けます。"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "⑪茹で卵をフォーク等で潰す。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/8b82745ef832ea1a7c2dc8194821695f.jpg?p=medium"},
			}},
			{Elements: []rexch.InstructionElement{
				&rexch.TextInstructionElement{Text: "⑫茹で卵、水気を切った玉ねぎにBの材料を加え、よく混ぜ合わせる。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/489ae3463574a9953ff59d25120f4c70.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "ぜひ揚げたての食感を楽しんで！"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/58a32da160ee9b7f1b5d66241874a551.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "揚げたてのサクサクをお召し上がりください。\nお好みでレモンを絞ると、より爽やかに頂けます。"},
				&rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/35/35774/9590d4298b5153378bebb6584810ea9c.jpg?p=medium"},
				&rexch.TextInstructionElement{Text: "タルタルソースはたっぷりかけるのが美味しさの秘訣！\nタルタルソースのみでも、レモンをプラスしても、さらにソースをかけても美味しいです。\nビールと一緒に食べると最高！揚げた後に油をしっかり落とせば、油やけの心配もありません。\nタルタルソースはお好みでキュウリのみじん切りやピクルス等を入れて、食感の変化をお楽しみください。"},
			}},
		},
	},
	"https://macaro-ni.jp/45500": {
		Title:    "ごはん無限レベル！豚肉とれんこんの甘辛炒め【作り置き】",
		Image:    "https://cdn.macaro-ni.jp/image/summary/45/45500/7d761cc229abdcb5088e77942c7bc1ed.jpg?p=1x1",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "豚こま切れ肉", Amount: "300g", Comment: ""},
			{Group: "", Name: "れんこん", Amount: "200g", Comment: ""},
			{Group: "", Name: "片栗粉", Amount: "大さじ2杯", Comment: "豚肉用"},
			{Group: "", Name: "片栗粉", Amount: "大さじ2杯", Comment: "れんこん用"},
			{Group: "", Name: "塩こしょう", Amount: "少々", Comment: "豚肉下味用"},
			{Group: "", Name: "サラダ油", Amount: "適量", Comment: ""},
			{Group: "", Name: "☆砂糖", Amount: "大さじ3杯", Comment: ""},
			{Group: "", Name: "☆しょうゆ", Amount: "大さじ3杯", Comment: ""},
			{Group: "", Name: "☆酢", Amount: "大さじ3杯", Comment: ""},
			{Group: "", Name: "白ごま", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "豚こま切れ肉とれんこんに片栗粉を大さじ2杯ずつまぶします。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンにサラダ油を深さ3cm入れて170℃に熱し、れんこんを揚げ焼きにします。まぶした片栗粉がカリっとしてきたら油から上げます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "次に豚こま切れ肉を入れて揚げ焼きにし、カリカリになったら油を切って上げます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンの油をふき取り☆の調味料を入れてひと煮立ちさせます。とろみがついたら②のれんこんと③の豚こま切れ肉を加えて絡めます。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "保存容器に入れ、お好みで白ごまを振って完成です！保存期間は冷蔵5日間です♪"}}},
		},
	},
	"https://macaro-ni.jp/104118": {
		Title:    "ランチにも♪ 豚こまガーリックライス",
		Image:    "https://cdn.macaro-ni.jp/image/summary/104/104118/P55cvVVOQNXc4ze0Q967VN79JtjAPAOSrwNTSZAJ.jpg?p=1x1",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "豚こま肉", Amount: "120g", Comment: ""},
			{Group: "", Name: "ごはん", Amount: "400g", Comment: ""},
			{Group: "", Name: "にんにく", Amount: "2片", Comment: ""},
			{Group: "", Name: "塩", Amount: "少々", Comment: ""},
			{Group: "", Name: "黒こしょう", Amount: "少々", Comment: ""},
			{Group: "", Name: "にんにく", Amount: "小さじ1杯", Comment: "すりおろし"},
			{Group: "", Name: "しょうゆ", Amount: "小さじ2杯", Comment: ""},
			{Group: "", Name: "ウスターソース", Amount: "小さじ1杯", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "大さじ1杯", Comment: ""},
			{Group: "", Name: "トッピング", Amount: "", Comment: ""},
			{Group: "", Name: "黒こしょう", Amount: "適量", Comment: ""},
			{Group: "", Name: "ドライパセリ", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "にんにくの両端を切り落とします。爪楊枝で芯の小さい方から大きい方に向かって、爪楊枝の太い方を使って押し出します。薄切りします。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/104/104118/j4VGePFKJdYYNAIRCbhlmhcOkfuaN6XPDroqGyi3.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンにサラダ油、にんにくをいれて弱火で熱します。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/104/104118/bIo25QPs7WJVrvzJJ1nXiQx30um7m5Y82ajCWn5o.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "にんにくに色がついたら取り出し、豚こま肉、塩、黒こしょうを入れて中火で炒めます。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/104/104118/uF8b5bROiuYF0tMb06SUlICuMEbZW62yIjq9LN28.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "豚肉の色が変わってきたらごはんを加えて強火で炒めます。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/104/104118/GhzdktAdSNfsgB5ZYEMdUihEhbXbwYdS6JSwZoO6.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "すりおろしにんにく、しょうゆ、ウスターソースを加えて混ぜ合わせて完成です。器に盛り付けてお好みでパセリ、黒こしょうをトッピングして完成です。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/104/104118/RursIk492XUUqqXf0QbOtJrO9QlIb0kEcaN0z9tW.jpg"}}},
		},
	},
	"https://macaro-ni.jp/84903": {
		Title:    "油揚げで簡単。豚こまとんかつ",
		Image:    "https://cdn.macaro-ni.jp/image/summary/84/84903/2HcJACf7b2xi6FMYZqRyjxHRWUb7O6bUxTBcgO3r.jpeg?p=1x1",
		Servings: 0, // TODO
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "豚こま肉", Amount: "240g", Comment: ""},
			{Group: "", Name: "油揚げ", Amount: "2枚", Comment: ""},
			{Group: "", Name: "塩", Amount: "少々", Comment: ""},
			{Group: "", Name: "こしょう", Amount: "少々", Comment: ""},
			{Group: "", Name: "片栗粉", Amount: "大さじ1杯", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "適量", Comment: ""},
			{Group: "", Name: "トッピング", Amount: "", Comment: ""},
			{Group: "", Name: "中濃ソース", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "油揚げは開きやすいように菜箸を転がします。包丁で横に切り込みを入れて、裏返します。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "豚こま肉は塩、こしょうで下味をつけて、片栗粉をまぶします。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "①の中に②を入れて切り口を折り返します。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンに深さ1cmほどのサラダ油を入れて熱し、③を入れて中火で揚げ焼きにします。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "焼き色がついたら裏返し、5〜8分揚げ焼きして中までしっかりと火を通します。"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "食べやすい大きさに切って器に盛り付けたら完成です。お好みでソースをかけて召し上がれ。"}}},
		},
	},
	"https://macaro-ni.jp/104802": {
		Title:    "豚こまで作る♪白菜のとんぺい焼き",
		Image:    "https://cdn.macaro-ni.jp/image/summary/104/104802/TJjkBPAVvCIAKGFJ6tkvngGe2lbUuGCfzGlvygIs.jpg?p=1x1",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "白菜", Amount: "180g", Comment: ""},
			{Group: "", Name: "豚こま肉", Amount: "80g", Comment: ""},
			{Group: "", Name: "もやし", Amount: "40g", Comment: ""},
			{Group: "", Name: "めんつゆ", Amount: "小さじ2杯", Comment: "３倍濃縮"},
			{Group: "", Name: "味付塩こしょう", Amount: "少々", Comment: ""},
			{Group: "", Name: "とろけるチーズ", Amount: "20g", Comment: ""},
			{Group: "", Name: "卵", Amount: "3個", Comment: ""},
			{Group: "", Name: "水溶き片栗粉", Amount: "片栗粉:小さじ2杯、水:小さじ2杯", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "小さじ1杯×2", Comment: ""},
			{Group: "", Name: "トッピング", Amount: "", Comment: ""},
			{Group: "", Name: "お好み焼きソース", Amount: "適量", Comment: ""},
			{Group: "", Name: "マヨネーズ", Amount: "適量", Comment: ""},
			{Group: "", Name: "小口ねぎ", Amount: "適量", Comment: ""},
			{Group: "", Name: "七味唐辛子", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "白菜を1cm幅に切ります。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/104/104802/vUImcG2miFkNjbvn0ehEL4fMJQiaFhDQmBZQH7BF.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ボウルに卵、水溶き片栗粉を入れて混ぜ合わせます。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/104/104802/jW7SHFJ8WvnB0wmkTzCtOvVePPlJSGw0sbsBWcbk.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンにサラダ油を引き中火で熱し、豚肉を炒めます。肉の色が変わったら白菜、もやしを加えて強火で1分ほど炒めます。味付塩こしょう、めんつゆを加えさっと炒め合わせ、いったん取り出します。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/104/104802/KIv7kKx8FU9AAxE2gfym4E4GxxOcbdyNGfG2urfQ.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "フライパンにサラダ油を入れ中火で熱し、溶き卵を流し入れます。菜箸で軽く混ぜ、半熟状になったら、とろけるチーズと③を中央にのせます。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/104/104802/xSA9nnKkkgD0I3CeoVHUC4GCDR8S9kSXB5FKX1Tz.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "具を包むように卵の両端を折りたたみます。折りたたんだ部分が下になるように皿に盛りつけます。ソース、マヨネーズ、七味唐辛子、小口ねぎをかけて完成です。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/104/104802/8oCfEGVMk14LAoYZ4WliIM1uYQUaBVETh6TDCvAy.jpg"}}},
		},
	},
	"https://macaro-ni.jp/94837": {
		Title:    "食べ応え抜群♪生姜焼き豚つくね",
		Image:    "https://cdn.macaro-ni.jp/image/summary/94/94837/LxT9jnKHHTrKofxUxgGz9XEJMbbJcIuzRipox9IP.jpg?p=1x1",
		Servings: 0,
		Ingredients: []rexch.Ingredient{
			{Group: "", Name: "豚こま肉", Amount: "300g", Comment: ""},
			{Group: "", Name: "玉ねぎ", Amount: "1/4個", Comment: ""},
			{Group: "", Name: "しょうが", Amount: "1片(15g)", Comment: ""},
			{Group: "", Name: "片栗粉", Amount: "大さじ2杯", Comment: ""},
			{Group: "", Name: "酒", Amount: "大さじ1/2杯", Comment: ""},
			{Group: "", Name: "塩", Amount: "少々", Comment: ""},
			{Group: "", Name: "こしょう", Amount: "少々", Comment: ""},
			{Group: "a.", Name: "酒", Amount: "大さじ2杯", Comment: ""},
			{Group: "a.", Name: "みりん", Amount: "大さじ2杯", Comment: ""},
			{Group: "a.", Name: "砂糖", Amount: "小さじ1杯", Comment: ""},
			{Group: "a.", Name: "はちみつ", Amount: "小さじ1杯", Comment: ""},
			{Group: "a.", Name: "しょうゆ", Amount: "大さじ2杯", Comment: ""},
			{Group: "", Name: "サラダ油", Amount: "大さじ1/2杯", Comment: ""},
			{Group: "", Name: "トッピング", Amount: "", Comment: ""},
			{Group: "", Name: "白いりごま", Amount: "適量", Comment: ""},
			{Group: "", Name: "小口ねぎ", Amount: "適量", Comment: ""},
		},
		Instructions: []rexch.Instruction{
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "豚こま肉は包丁で叩き、粗めのみじん切りにします。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/94/94837/RA6phTGBbz9j2DUC3csGGoNJGTN5tP81FLQGwVhI.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ボウルに①、玉ねぎ、しょうが（半量）、塩、こしょう、酒、片栗粉を加えて粘りが出るまでよく混ぜ合わせます。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/94/94837/D7yU8wwASwViMQDBZ0XBHObJPziCzfJZMJtgf3g0.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "ひとつ分手に取り形を整えたら、サラダ油を熱したフライパンにのせて両面焼きます。焼き色がついたらフタをし、3〜4分蒸し焼きにして中までしっかり火を通します。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/94/94837/RpjdeV1LIZrRNXenOPvWJbklwQZQhSZFM7DtZX9e.jpg"}}},
			{Label: "", Elements: []rexch.InstructionElement{&rexch.TextInstructionElement{Text: "(a)の調味料、残りのしょうがを加えて煮絡めます。器に盛り付けたら完成です。お好みでいりごま、小口ねぎをトッピングして召し上がれ。"}, &rexch.ImageInstructionElement{URL: "https://cdn.macaro-ni.jp/image/summary/94/94837/QkkC4SN5xT9nAthgx9Qh7jfcO4CvP1ymU7VE7G5m.jpg"}}},
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
