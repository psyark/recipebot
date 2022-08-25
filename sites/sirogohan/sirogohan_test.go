package sirogohan

import (
	"context"
	"encoding/json"
	"fmt"
)

func ExampleParser() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://www.sirogohan.com/recipe/nikudouhu/")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "肉豆腐",
	//   "Image": "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig6717.JPG",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "牛こま切れ肉",
	//           "Amount": "200ｇ",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "木綿豆腐",
	//           "Amount": "1と1/2丁（計450ｇほど）",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "玉ねぎ",
	//           "Amount": "1個",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "椎茸",
	//           "Amount": "4枚",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "えのき茸",
	//           "Amount": "100ｇ",
	//           "Comment": ""
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "A",
	//       "Children": [
	//         {
	//           "Name": "醤油",
	//           "Amount": "大さじ5",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "酒",
	//           "Amount": "大さじ4",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "みりん",
	//           "Amount": "大さじ4",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "砂糖",
	//           "Amount": "大さじ2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "水",
	//           "Amount": "100ml",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "肉豆腐の下ごしらえ\n\nこの肉豆腐は26㎝ほどのフライパンで作る工程で紹介しています。フライパンの体積を考えると木綿豆腐１と1/2丁くらいがちょうどよいので、半端ですが一度上記分量で作ってみてください。\n\nはじめに木綿豆腐は1丁を8等分ほどに切って、キッチンペーパーの上にしばらく置いて水を吸わせます。\n\n水気を切ることで煮汁が水っぽくなったりせず、味もしみ込みやすくなります。置く時間は10分ほどでOKです。\n\n\t\n\t\t\n\t\t\n\t\n\n※木綿豆腐の水切りをする時間がないときは、より水気が少ない「焼き豆腐」を買ってきて作っても。\n\n玉ねぎは10～12等分ほどのくし切りにします。椎茸は軸を切り落として半分に、えのき茸は石づきを切り落として食べやすい大きさに手で割いておきます。",
	//       "Images": [
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig1.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig4.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig2.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig3.JPG"
	//       ]
	//     },
	//     {
	//       "Text": "肉豆腐の味付け/レシピ\n\n肉豆腐の味付けは牛肉や野菜からの出汁が出るので、特にだし汁は必要ありません。Aの調味料（醬油大さじ5、みりんと酒が各大さじ4、砂糖大さじ2、水100ml）を合わせて箸で混ぜて軽く溶かし、フライパンを火にかける前に、切った玉ねぎを入れておきます。\n\n\t\n\t\t\n\t\t\n\t\n\nフライパンを中火にかけて沸くのを待ち、煮汁が沸いたら玉ねぎを端に寄せて牛肉を入れるスペースを作ります。\n\n\t\n\t\t\n\t\t\n\t\n\n煮汁が出た空いたスペースに、牛肉を入れて箸でほぐしながら全体的に色が変わるまで火を通します。火が通った牛肉は箸ですくい上げて取り出します。\n\n\t\n\t\t\n\t\t\n\t\n\n※肉豆腐は豆腐や玉ねぎなどやわらかい食感の具材が多いので、牛肉が固くならないよう、はじめにさっと火を通して取り出し、最後に戻し入れて温めるという火の入れ方にしています。\n\n牛肉を取り出したら、玉ねぎを端に寄せたまま豆腐をそっと入れます（玉ねぎは重なってもよいので、豆腐が重ならないように）。豆腐がすべて入れば、上にきのこを広げるようにのせます。\n\n\t\n\t\t\n\t\t\n\t\n\nきのこを上にのせて少し煮るとしんなりしてくるので、きのこをフライパンの端や、豆腐と豆腐の隙間など、入れ込むことができる部分に入れます。\n\n\t\n\t\t\n\t\t\n\t\n\nきのこがある程度煮汁に浸かれば落し蓋をします（落し蓋をすれば自然とカサは減るのですべてが浸かる必要はありません）。蓋の下の煮汁がぐつぐつと沸く火加減で、10～12分煮ます。\n\n\t\n\t\t\n\t\t\n\t\n\n10分ほど煮れば、玉ねぎやきのこもしっかりとやわらかく煮上がり、煮汁も適度に煮詰まってきます。\nきのこや玉ねぎを豆腐の上に移動させ、先に取り出しておいた牛肉を煮汁に戻し入れ、1～2分煮てしっかり温めます。\n\nこれで出来上がり、煮汁をたっぷりかけて盛り付けてください。",
	//       "Images": [
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig5.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig6.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig7.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig8.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig9.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig10.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig0.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig12.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig13.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig14.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig15.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig16.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig17.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig18.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig6717small.JPG"
	//       ]
	//     },
	//     {
	//       "Text": "肉豆腐を丼ぶりにしても！\n\nご飯の上に豆腐、牛肉、野菜をのせて、煮汁をたっぷりかけると牛丼のような美味しい丼ものにもなります。\n好みで一味唐辛子をふりかけても美味しいです。2日目などの食べ方としてぜひ！",
	//       "Images": [
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig19.JPG",
	//         "https://www.sirogohan.com/_files/recipe/images/nikudouhu/nikudouhubig20.JPG"
	//       ]
	//     }
	//   ]
	// }
}
