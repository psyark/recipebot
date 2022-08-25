package kurashiru

import (
	"context"
	"encoding/json"
	"fmt"
)

func ExampleParser() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://www.kurashiru.com/recipes/9d2cec44-130f-49e4-880c-b1ed15eec39b")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "オレンジチキン",
	//   "Image": "https://video.kurashiru.com/production/videos/9d2cec44-130f-49e4-880c-b1ed15eec39b/compressed_thumbnail_square_large.jpg?1639367889",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "鶏もも肉",
	//           "Amount": "300g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "塩こしょう",
	//           "Amount": "小さじ1/4",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "片栗粉",
	//           "Amount": "大さじ2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "揚げ油",
	//           "Amount": "適量",
	//           "Comment": ""
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "ソース",
	//       "Children": [
	//         {
	//           "Name": "100%オレンジジュース",
	//           "Amount": "30ml",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "しょうゆ",
	//           "Amount": "大さじ2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "酢",
	//           "Amount": "大さじ2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "はちみつ",
	//           "Amount": "大さじ1",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "すりおろしニンニク",
	//           "Amount": "小さじ1",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "すりおろし生姜",
	//           "Amount": "小さじ1",
	//           "Comment": ""
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "付け合せ",
	//       "Children": [
	//         {
	//           "Name": "ベビーリーフ",
	//           "Amount": "30g",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "鶏もも肉は一口大に切ります。塩こしょうをふり、片栗粉をまぶします。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "鍋底から3cm程の高さまで揚げ油を注ぎ、180℃に加熱します。1を入れ、鶏もも肉に火が通るまで5分ほど揚げたら油切りをします。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "ソースを作ります。フライパンにソースの材料を入れて中火で熱します。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "沸騰してから1分ほど中火で加熱し、2を入れます。よく絡め、全体に味がなじんだら火から下ろします。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "ベビーリーフと共にお皿に盛り付けてできあがりです。",
	//       "Images": null
	//     }
	//   ]
	// }
}
