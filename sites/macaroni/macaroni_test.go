package macaroni

import (
	"context"
	"encoding/json"
	"fmt"
)

func ExampleParser() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://macaro-ni.jp/109611")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "とろ〜り半熟卵で作る「ウフ・アン・ムーレット」【秋元さくらシェフ直伝】",
	//   "Image": "https://cdn.macaro-ni.jp/image/summary/109/109611/BN1C2JEIn3tfIUOSJZUGolyDUYOXsDN6BY3HjDBp.jpg?p=medium",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "玉ねぎ",
	//           "Amount": "1/6個",
	//           "Comment": "みじん切り"
	//         },
	//         {
	//           "Name": "マッシュルーム",
	//           "Amount": "6個",
	//           "Comment": "薄切り"
	//         },
	//         {
	//           "Name": "ベーコン",
	//           "Amount": "40g",
	//           "Comment": "5mm幅・拍子切り"
	//         },
	//         {
	//           "Name": "にんにく",
	//           "Amount": "少々",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "赤ワイン",
	//           "Amount": "400cc",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "水溶きコーンスターチ",
	//           "Amount": "小さじ1杯（コーンスターチ：小さじ1/2杯、水：小さじ1/2杯）",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "オリーブオイル",
	//           "Amount": "大さじ1杯",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "バター",
	//           "Amount": "20g",
	//           "Comment": "有塩"
	//         },
	//         {
	//           "Name": "塩",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "クルトン",
	//           "Amount": "適宜",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "パセリ",
	//           "Amount": "適宜",
	//           "Comment": "みじん切り"
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "〈ポーチドエッグ 〉",
	//       "Children": [
	//         {
	//           "Name": "卵",
	//           "Amount": "2個",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "水",
	//           "Amount": "2000cc",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "酢",
	//           "Amount": "大さじ1杯",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "1. 具材を炒める\n鍋にオリーブオイルを入れ、玉ねぎ、にんにくを加えて炒めます。にんにくの香りを引き出しながら、玉ねぎがくったりするまで炒めます。\n\n「鍋肌の温度が上がるまでは強火にしましょう。ジリジリっとした音が聞こえてきたら中火にします。今回は家庭料理なので玉ねぎを使用しますが、本来はエシャロットを使います。\n\n炒めるときの注意点は、あまりぐるぐるかきまぜないこと。触れば触るほど鍋中の温度は下がるので、かき混ぜるのは時々にしてくださいね」\nバター、マッシュルーム、ベーコンを加えて、さらに炒めます。\n\n「素材の水分を出すために通常は塩をふりますが、今回はなし。ソース作りの過程でワインをかなり煮詰めていくので、ここで塩分を入れると塩味が強くなりがちなんです」\n「マッシュルームから水分が出てきます。この水分を生かしながら、蒸すようにじっくり炒めて、それぞれの素材の甘みを引き出しましょう」\n「甘味は旨味でもあります。きのこのグルタミン酸、ベーコンのアミノ酸の旨味をじっくり濃縮するのがポイントです。1/3量になるまで炒めてくださいね」",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/2tsacMfe0uGbQGt9gxXOW3T3IvU9NEmH1JkEYSgA.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/fAtKszCFK5BW2is92bvWKEdQWUtUcHojthu1RdYm.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/TEPWPU31mmym2HBQSIZm6lmrGrdl3eJQy0Dl5wVa.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/dw3TGr0oexZd1sSUDzPOgu5cWQC0QNgAx2WvgSNy.jpg?p=medium"
	//       ]
	//     },
	//     {
	//       "Text": "2. 赤ワインを加えて煮詰める\n赤ワインを加えて、中火でコトコト煮詰めていきます。赤ワインが1/6量になるぐらいが目安です。\n\n「このお料理はブルゴーニュ地方のスペシャリテなので、同じ土地で生まれたブルゴーニュのワインだといいですね！高いワインを使うともちろんおいしいソースができますが、赤ワインならどんなものでもOKです」",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/sYfFTga1KGecbXr8uIGb3ka79k95SO3MgkmSpDUK.jpg?p=medium"
	//       ]
	//     },
	//     {
	//       "Text": "3. ポーチドエッグを作る\n鍋に2リットルの熱湯を沸かし、大さじ1杯の酢を加え、おたまを入れてぐるぐる混ぜて渦を作ります。ポーチドエッグは1個ずつ作るので、小ぶりな容器に卵を1個割り入れます。\n\n「塩も加えるレシピもありますが、シンプルに酢だけを使いましょう。いろいろな組み合わせで試しましたが、大きな違いはありませんでした。それなら簡単なほうがいいですよね」\n鍋の中央にできた渦の中に卵を落とし、1分半ほどゆでます。\n\n「成功の秘訣は、よく冷えた新しい卵を使うことです。お鍋の大きさによりますが、卵の高さの3倍量のお湯は必須ですね」\n1分半たったら、穴の空いたレードルなどでポーチドエッグを引き上げます。氷水にとらずに、バットや皿におきましょう。\n\n「温かいまま食べていただく料理なので、氷水にはおとさないでくださいね。予熱でいい感じの半熟卵になります」\n白身の余分な部分はスプーンやキッチンばさみなどを使って取り除き、深みのあるお皿に盛り付けます。",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/L8YTIVGx5JJyRLInfLsJd70agBynPRiMUYHk258J.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/rzRaq7X2XjKCD2cSBTkenp347VSjqMzGYYQfEzze.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/EVuQOquPuyDyI5r9xVpE47k9Em0x9ZoOczUrIfCL.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/8edJc7cLuA7tYXkI5zBSs0w3TXGS5S02D01shnpT.jpg?p=medium"
	//       ]
	//     },
	//     {
	//       "Text": "4. 2にコーンスターチを加える\n赤ワインの分量が1/6ほどになり、フランス語でミロワールといわれるつやつやした状態になったら、水溶きコーンスターチを加えてとろみをつけます。塩で味を調えればソースのできあがりです。\n\n「本来は大量のバターを加えてとろみをつけますが、今回はヘルシーに水溶きコーンスターチで代用しましょう。ソースは酸っぱく感じるかもしれませんが、半熟卵と食べるとちょうどよいと思います。酸味のあるソースと濃厚な卵黄の相性は抜群にいいんです！」",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/Mi5jt0Sn87Mjl115mQIuUMcLKrccjJ8eJelW42tW.jpg?p=medium"
	//       ]
	//     },
	//     {
	//       "Text": "5. 盛り付ける\nポーチドエッグを囲むようにムーレットソースをかけ、中央にクルトン、パセリのみじん切りを飾ります。ウフ・アン・ムーレットの完成です！",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/hcBCYdLdQiGbc5RM7hlmWmfjDJwiZ6mrtM1Z0DZ2.jpg?p=medium"
	//       ]
	//     },
	//     {
	//       "Text": "食材のコクや旨味が凝縮。半熟卵のぜいたくなひと皿\n濃厚な卵黄の甘味やコクが、赤ワインベースのムーレットソースの酸味と絡み、なんとも大人な味わい。ワイン、そしてこんがり焼いた薄切りバゲットと一緒に、半熟卵の魅力を楽しめる料理です。卵のイエローとソースのワインレッドの色合いも美しい！\n\n「日本ではまだあまり知られていないお料理ですが、使っているのは基本の素材ばかりなので、ワインのお供としてご家庭でのパーティの前菜として、気軽に取り入れていただけたらと思います。冷たくてもおいしいので、作り置きにも適していますよ。\n\nマッシュルームはしめじやまいたけに変えてもOKです。このソースはお肉との相性も抜群。少し多めに作ってお肉に添えるのもおすすめです」\n\nポーチドエッグは作ったことがない！という方でも、秋元シェフ紹介の作り方なら失敗知らずです。エッグベネディクト、オープンサンド、サラダといろいろな料理に活かせるので、この機会に作り方を習得してはいかがでしょうか。",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/WMb42PpekrPTWpccjRVYylU7fh6BvLQKrpWfLB9d.jpg?p=medium"
	//       ]
	//     },
	//     {
	//       "Text": "教えてくれた人\n取材協力\n取材・文／古川あや\n撮影／宮本信義\n\n\n関連レシピはこちら▼",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/109/109611/EMlBngOLcsPWUtKKbJmYv8hxopkiFYEj2V9Cy1Rw.jpg?p=medium"
	//       ]
	//     }
	//   ]
	// }
}

func ExampleParser_a() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://macaro-ni.jp/90067")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "わさびがポイント！お豆腐の大葉肉巻き",
	//   "Image": "https://cdn.macaro-ni.jp/image/summary/90/90067/yw8njYPbCLPv97YaraNnNyHQQuliiNRYuGHAV9My.jpeg?p=1x1",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "木綿豆腐",
	//           "Amount": "300g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "豚バラ肉",
	//           "Amount": "200g",
	//           "Comment": "薄切り"
	//         },
	//         {
	//           "Name": "大葉",
	//           "Amount": "10枚",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "片栗粉",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "わさび",
	//           "Amount": "大さじ1/2杯",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "サラダ油",
	//           "Amount": "適量",
	//           "Comment": ""
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "a.",
	//       "Children": [
	//         {
	//           "Name": "酒",
	//           "Amount": "大さじ1杯",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "みりん",
	//           "Amount": "大さじ1杯",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "砂糖",
	//           "Amount": "大さじ1杯",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "しょうゆ",
	//           "Amount": "大さじ2杯",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "豆腐をキッチンペーパーでくるみ、レンジ600Wで3分加熱します。",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/90/90067/ekhmUcHBt9BZjyjMkeQ0anUcPf4zDwiW50cJkbuj.jpeg"
	//       ]
	//     },
	//     {
	//       "Text": "重石をして10分ほどおき、しっかり水切りをして1cm幅に切ります。",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/90/90067/xxt0MRAfIq3cBn552blhp3Z62sEyFiAkESYLBhlX.jpeg"
	//       ]
	//     },
	//     {
	//       "Text": "豚肉に大葉、豆腐をのせて巻き上げ、片栗粉をまぶします。",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/90/90067/lxE9asgFlMtHHYFNYxXglLhX49x8ujwJqgkkQDVi.jpeg"
	//       ]
	//     },
	//     {
	//       "Text": "フライパンにサラダ油を引いて熱し、③を並べて焼き色がつくまで焼きます。",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/90/90067/lHi9RkopTzOU57FWnoNN0JVPPpkCjpf4Gf9XQyqZ.jpeg"
	//       ]
	//     },
	//     {
	//       "Text": "(a) を加えて全体がからんだら火を弱め、わさびを加えて全体にからめたら完成です。",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/90/90067/x3UpTsrLKH0C4o4VilxiiVJrABJdbC76jeDHm75h.jpeg"
	//       ]
	//     }
	//   ]
	// }
}

func ExampleParser_r85520() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://macaro-ni.jp/85520")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "メイン食材1つ。豆腐のスティックフライ",
	//   "Image": "https://cdn.macaro-ni.jp/image/summary/85/85520/488fa7c470eb630b98de98fc66548774.jpg?p=1x1",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "木綿豆腐",
	//           "Amount": "200g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "片栗粉",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "サラダ油",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "塩",
	//           "Amount": "少々",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "粗挽き黒こしょう",
	//           "Amount": "少々",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "ガーリックパウダー",
	//           "Amount": "少々",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "キッチンペーパーで豆腐を包み、耐熱皿にのせて電子レンジ500Ｗ3分加熱し、水切りをします。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "取り出したら1cm角の棒状に切ります。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "片栗粉を全体にまぶします。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "フライパンに170℃のサラダ油を熱し、③を加えて揚げ焼きにします。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "油をよく切り、塩、こしょう、ガーリックパウダーを振ったら完成です。",
	//       "Images": null
	//     }
	//   ]
	// }
}
func ExampleParser_r35774() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://macaro-ni.jp/35774")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "必見！コツをおさえた『基本の鮭フライ＆タルタルソース』の作り方",
	//   "Image": "https://cdn.macaro-ni.jp/image/summary/35/35774/2a8f6aae712a66cdf5c8485a93a3f044.jpg?p=medium",
	//   "IngredientGroups": null,
	//   "Steps": [
	//     {
	//       "Text": "\n①鮭は食べやすいよう、1切れを2〜3切れに切る。\n②塩胡椒(分量外)を全体にまんべんなく振り、5分ほど置く。水分が出るのでキッチンペーパー等で軽く拭きとる。\n塩胡椒を振って水分を出すことで生臭みを取るので、必ずこの行程は行いましょう。\nまた、塩を振ることで旨味も増します。\n③袋に鮭と薄力粉を入れ、薄力粉を全体にまぶす。\n空気を含んだ状態で口を閉じ、全体を揺するようにすると薄力粉がきれいに付きます。\n④袋から鮭を取り出し、一つずつ軽く叩いて余分な薄力粉を落とす。\n薄力粉・卵・パン粉の3つとも、余分に付いていると揚げた時にムラができます。\nきれいな仕上がりにするために、軽く叩きましょう。\n⑤鮭を卵にくぐらせる。\n薄力粉と同じく、卵はくぐらせた後によく水気を切りましょう。\n⑥粉チーズを混ぜたパン粉に鮭をのせ、パン粉を全体によくまぶす。\nパン粉を付けたら、鮭は軽く振って余分なパン粉を落とし、別の容器に置きましょう。\n⑦170℃に熱した油で、きつね色になるまでじっくり揚げる。\n火加減の目安は中火です。油が十分に温まり、なおかつ高温すぎないのがポイント。\n⑧ふちがきつね色になったらひっくり返し、反対側も同じように揚げる。\nふちがきつね色になってからひっくり返すと、火加減が均等で色も良く仕上がります。\nひっくり返す時に、衣が硬くサクッとしていればOK。\n衣が柔らかい場合は火が弱いので、火加減を調整しましょう。\n⑨バットに上げて、余分な油を落とす。\n油切れが悪いと、ベタついたり胃もたれの原因になります。\n必ず油を落としましょう。\n次にタルタルソースを作ります。\n⑩玉ねぎはみじん切りにして、水にさらす。\n玉ねぎの辛味を抜きます。時間がない場合は、みじん切りして塩をまぶし、軽く握りつぶしてから水にさらすと繊維が潰れて辛味が早く抜けます。\n⑪茹で卵をフォーク等で潰す。\n⑫茹で卵、水気を切った玉ねぎにBの材料を加え、よく混ぜ合わせる。",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/567f5cb41b2286ae60d4d50922f023b9.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/c7e74110c599d13774d13970fdad4d06.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/6a9cde73befc66ba24908e8d83b39c63.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/00bfda22dee251416d728d9c2c7cdded.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/28d95776755897528b210f32c65f6e34.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/a12119be3895d3041acf3bf1898244cf.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/8095dd34daaa3e9e013297131193770c.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/d7358658a5b263523fdb003fb0a4ee04.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/40a21368c402b5b0d98c0f35a0723550.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/854e7a24b157f83b5bb4789e137cfa9c.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/b820fa7fcd527d067253ca6a4a0160af.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/8b82745ef832ea1a7c2dc8194821695f.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/489ae3463574a9953ff59d25120f4c70.jpg?p=medium"
	//       ]
	//     },
	//     {
	//       "Text": "ぜひ揚げたての食感を楽しんで！\n揚げたてのサクサクをお召し上がりください。\nお好みでレモンを絞ると、より爽やかに頂けます。\nタルタルソースはたっぷりかけるのが美味しさの秘訣！\nタルタルソースのみでも、レモンをプラスしても、さらにソースをかけても美味しいです。\n\nビールと一緒に食べると最高！揚げた後に油をしっかり落とせば、油やけの心配もありません。\n\nタルタルソースはお好みでキュウリのみじん切りやピクルス等を入れて、食感の変化をお楽しみください。",
	//       "Images": [
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/58a32da160ee9b7f1b5d66241874a551.jpg?p=medium",
	//         "https://cdn.macaro-ni.jp/image/summary/35/35774/9590d4298b5153378bebb6584810ea9c.jpg?p=medium"
	//       ]
	//     }
	//   ]
	// }
}

func ExampleParser_r45500() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://macaro-ni.jp/45500")
	if err != nil {
		panic(err)
	}
	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "ごはん無限レベル！豚肉とれんこんの甘辛炒め【作り置き】",
	//   "Image": "https://cdn.macaro-ni.jp/image/summary/45/45500/7d761cc229abdcb5088e77942c7bc1ed.jpg?p=1x1",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "豚こま切れ肉",
	//           "Amount": "300g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "れんこん",
	//           "Amount": "200g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "片栗粉",
	//           "Amount": "大さじ2杯",
	//           "Comment": "豚肉用"
	//         },
	//         {
	//           "Name": "片栗粉",
	//           "Amount": "大さじ2杯",
	//           "Comment": "れんこん用"
	//         },
	//         {
	//           "Name": "塩こしょう",
	//           "Amount": "少々",
	//           "Comment": "豚肉下味用"
	//         },
	//         {
	//           "Name": "サラダ油",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "☆砂糖",
	//           "Amount": "大さじ3杯",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "☆しょうゆ",
	//           "Amount": "大さじ3杯",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "☆酢",
	//           "Amount": "大さじ3杯",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "白ごま",
	//           "Amount": "適量",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "豚こま切れ肉とれんこんに片栗粉を大さじ2杯ずつまぶします。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "フライパンにサラダ油を深さ3cm入れて170℃に熱し、れんこんを揚げ焼きにします。まぶした片栗粉がカリっとしてきたら油から上げます。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "次に豚こま切れ肉を入れて揚げ焼きにし、カリカリになったら油を切って上げます。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "フライパンの油をふき取り☆の調味料を入れてひと煮立ちさせます。とろみがついたら②のれんこんと③の豚こま切れ肉を加えて絡めます。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "保存容器に入れ、お好みで白ごまを振って完成です！保存期間は冷蔵5日間です♪",
	//       "Images": null
	//     }
	//   ]
	// }
}
