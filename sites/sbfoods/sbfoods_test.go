package sbfoods

import (
	"context"
	"encoding/json"
	"fmt"
)

func ExampleParser() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://www.sbfoods.co.jp/recipe/detail/05912.html")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "蒸し鶏の怪味ソースがけ",
	//   "Image": "https://cdn.sbfoods.co.jp/recipes/05912_l.jpg",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "鶏もも肉",
	//           "Amount": "1枚(250g)",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "サニーレタス",
	//           "Amount": "2枚",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "長ねぎ",
	//           "Amount": "5cm",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "酒",
	//           "Amount": "大さじ1",
	//           "Comment": ""
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "【Ａ】",
	//       "Children": [
	//         {
	//           "Name": "醤油",
	//           "Amount": "大さじ1と1/4",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "酢",
	//           "Amount": "大さじ1/2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "砂糖",
	//           "Amount": "小さじ2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "豆板醤",
	//           "Amount": "小さじ1",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "練りごま",
	//           "Amount": "大さじ1と1/4",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "ラー油",
	//           "Amount": "大さじ1/4",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "おろしにんにく",
	//           "Amount": "小さじ1/3",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "おろししょうが",
	//           "Amount": "小さじ1/2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "花椒（パウダー）",
	//           "Amount": "小さじ1/4",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "鶏もも肉は肉の厚い部分に斜めに浅く包丁を入れて開き、肉の厚さを均等にします。酒を全体に振って手で押さえ、耐熱皿にのせラップをかけて電子レンジ(600W)で４分加熱します。長ねぎはせん切りにして水にさらした後キッチンペーパーに包み、軽くもんで白髪ねぎを作ります。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "【Ａ】を混ぜ合わせて怪味ソースを作ります。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "【１】の鶏肉をやけどに気を付けて食べやすく切り分け、サニーレタスを敷いた器に盛り、上に【２】の怪味ソースをかけ、お好みで花椒（分量外）を振り、白髪ねぎをのせます。",
	//       "Images": null
	//     }
	//   ]
	// }
}
