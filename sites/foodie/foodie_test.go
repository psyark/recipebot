package foodie

import (
	"context"
	"encoding/json"
	"fmt"
)

func ExampleParser() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://mi-journey.jp/foodie/59727/")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "鮭のムニエルのレシピ～洋食店のように仕上げる焼き方のコツ 【シェフ直伝】",
	//   "Image": "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere12.jpg",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "生鮭",
	//           "Amount": "2切れ",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "塩",
	//           "Amount": "少々",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "小麦粉",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "オリーブオイル",
	//           "Amount": "大さじ1",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "バター",
	//           "Amount": "15g",
	//           "Comment": ""
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "ソース",
	//       "Children": [
	//         {
	//           "Name": "バター",
	//           "Amount": "60g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "しょうゆ",
	//           "Amount": "小さじ2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "にんにく",
	//           "Amount": "1/2かけ分",
	//           "Comment": "みじん切り"
	//         },
	//         {
	//           "Name": "トマト",
	//           "Amount": "1/2個分（約30g）",
	//           "Comment": "湯むきして種を除き、角切りにしたもの"
	//         },
	//         {
	//           "Name": "レモン",
	//           "Amount": "1/4個分（約7g）",
	//           "Comment": "種と薄皮を除き角切りにしたもの"
	//         },
	//         {
	//           "Name": "ケイパー",
	//           "Amount": "大さじ2",
	//           "Comment": "酢漬け"
	//         },
	//         {
	//           "Name": "パセリ",
	//           "Amount": "各適量",
	//           "Comment": "みじん切り）、レモンの皮（好みで"
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "鮭の両面に塩をふり、10分ほどおく\n\n「ソースの味つけがしっかりしている場合でも、素材にはきちんと下味をつけるようにしましょう。塩は鮭に対し、1％の量を目安にふるとちょうどいい加減になります。今回は120gの鮭を使用したので、1.2gの塩をふりました」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere4.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "ペーパータオルで水けをふき取る\n\n10分ほどたって表面に水分が出てきたら、ペーパータオルでしっかりふき取ります。\n「魚は加熱する直前ではなく、10〜20分前に塩をふるのが基本。余分な水分を出すことで、臭みを取り除きます」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere5.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "鮭に小麦粉をまんべんなくまぶす\n\n「小麦粉をまぶす際、身の部分にはムラなくしっかりとつけますが、皮目は焦げやすいのであえてふるう必要はありません。身全体に小麦粉をまぶしたら、はたいて余分な粉をきちんと落としましょう。粉が余分についていると加熱中にはがれやすく、油や衣が汚れる原因に」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere6.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "冷たいフライパンにオリーブオイルを入れ、鮭を皮目から入れる\n\n「必ずフライパンが冷たい状態で油と鮭を入れてください。バターは焦げやすいので、最初はオリーブオイルを焼き油として使用し、途中から加えることで風味よく仕上げます。\n鮭は2切れをくっつけて並べ、皮目だけフライパンの底に当たるようにしてください。ひと切れずつ焼く場合は、アルミ箔などをクシャクシャと丸めて支えにしてあげると、立てて焼くことができます」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere7.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "弱火にかけ、鮭をじっくり焼く\n\n鮭を並べ入れたタイミングで火にかけます。しばらくすると鮭から脂が出てくるので、ペーパータオルでこまめにふき取ります。\n「皮は意外に厚いので、時間をかけてじっくり焼きます。皮の表面だけを焼きたいわけではないので、弱火でゆっくり火を入れましょう」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere8.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "皮目に焼き色がついたら身を焼く\n\n皮目を焼きはじめて5〜6分たち、皮に焼き色がついたら身を寝かせます。\n「仕上がりが美しくなるよう、器に盛りつける際に表になる面から焼きます」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere9.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "バターを加える\n\n身を焼くタイミングでバターを加えます。\n「バターは焦げやすいので、フライパンをときどきゆすって油脂の温度を一定に保ち、細かく泡が立っている状態をなるべく長くキープしましょう」\n",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere10.jpg",
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere11.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "裏返して、バターをすくってかけながら焼く\n\n2〜3分たったら裏返してもう片面を焼きます。焼き油をスプーンですくって身にかけながら加熱します。\n「身が厚い部分を重点的にアロゼしながら焼きます。火を均一に通すと同時に、バターの香りを鮭にまとわせ、香りよく仕上げましょう」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere12.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "取り出してペーパータオルで油をきる\n\n程よく焼き色がついたらペーパータオルの上に取り出し、余分な油をきります。",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere13.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "焦がしバターソースソースを作る\n\n9のフライパンの汚れをふき取り、バターを入れて中火で溶かします。",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere14.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "フライパンをときどきゆすりながら均一に色づける\n\n「状態の変化に気をつけながらバターを焦がします。最初に大きな泡が立ち、徐々に細かい泡に変化します。その泡が引いてきて茶色く色づきはじめたら、一気にソースを仕上げましょう」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere15.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "細かい泡が消えてきたら火を止め、しょうゆを加える\n\n香ばしい香りが立ち、薄く色づきはじめたら火から外してしょうゆを加えます。\n「しょうゆを加えることで、これ以上バターが焦げるのを止めることができます」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere16.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "残りの材料を加え、混ぜれば完成\n\nにんにく、トマト、レモン、ケイパー、パセリを手早く加え、フライパンをゆすって全体が混ざったらソースのできあがりです。\n「器に鮭を盛り、好みの野菜を添えて焦がしバターソースをたっぷりかけます。仕上げにレモンの皮をすりおろしてかけて、爽やかな風味をプラスしてもよいでしょう」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2019/10/191007salmonmeuniere17.jpg"
	//       ]
	//     }
	//   ]
	// }
}

func ExampleParser_foodie38465() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://mi-journey.jp/foodie/38465/")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "ミラノ風カツレツのレシピ・作り方。「揚げない」のがコツ。サクッとジューシー",
	//   "Image": "https://mi-journey.jp/foodie/wp-content/uploads/2017/05/1704_90_katsuretsu_01.jpg",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "とんかつ用豚肉",
	//           "Amount": "2枚（1枚120g程度）",
	//           "Comment": "ロースでもヒレ肉でもOK"
	//         },
	//         {
	//           "Name": "小麦粉",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "卵黄",
	//           "Amount": "1個分",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "パン粉",
	//           "Amount": "適量",
	//           "Comment": "粒が細かいもの"
	//         },
	//         {
	//           "Name": "パルミジャーノレッジャーノ",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "バターとオリーブオイル",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "塩、こしょう",
	//           "Amount": "各適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "レモン",
	//           "Amount": "適量",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "①豚肉を叩いて約1.5倍の面積に伸ばす。\n\n「ラップなどを被せると肉が動きません。肉叩きがなければ片手鍋などでもOKです」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2017/05/1704_90_katsuretsu_03.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "②塩、こしょうを両面にまぶしたら、全体に小麦粉を薄くまぶす。\n\n「揚げ衣を薄めにとどめるのがカリッとジューシーに仕上げるコツ。小麦粉を全体にまぶしたら、肉同士を軽く叩いて余計な粉は落としましょう」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2017/05/1704_90_katsuretsu_04-.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "③溶いた卵黄に②を付け、パン粉を広げたバットに肉を置きます。片面にパルミジャーノレッジャーノを、両面にパン粉をまんべんなくまぶす。\n\n「パルミジャーノレッジャーノ、パン粉ともに同じく薄化粧にとどめてください。パン粉は粒の細かいものを使うと衣がサックリ仕上がります。まぶしたら、ギュッと軽く押し付けます」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2017/05/1704_90_katsuretsu_05.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "④フライパンにバターとオリーブオイルを同量入れ、中火にかけてムース状（乳化した状態）になったら③を入れる。\n\n「火加減は油がシュワシュワとムース状態をずっと保つように。ときどき少し揺らして、肉の表面の凹凸になっている面にもまんべんなく油があたるようにしましょう」",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2017/05/1704_90_katsuretsu_06.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "⑤周囲に浮いているパン粉がキツネ色になってきたら裏返す。全体がキツネ色になったらできあがり。食べる直前にレモンを絞る。\n",
	//       "Images": [
	//         "https://mi-journey.jp/foodie/wp-content/uploads/2017/05/1704_90_katsuretsu_07.jpg"
	//       ]
	//     }
	//   ]
	// }
}
