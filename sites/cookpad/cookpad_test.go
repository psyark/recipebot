package cookpad

import (
	"context"
	"encoding/json"
	"fmt"
)

func ExampleParser_a() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://cookpad.com/recipe/1885344")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "マーマレードですっきり甘☆豚のスペアリブ",
	//   "Image": "https://img.cpcdn.com/recipes/1885344/m/d7fdaff65e9e0694d16801432cc6ea89?u=529143\u0026p=1545602123",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "豚のスペアリブ（骨付き肉）",
	//           "Amount": "400~500g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "★オレンジマーマレード",
	//           "Amount": "80g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "★しょうゆ（こいくち）",
	//           "Amount": "40~50cc",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "★砂糖",
	//           "Amount": "大さじ1~",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "※水",
	//           "Amount": "150~200cc",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "☆サラダ油",
	//           "Amount": "適宜",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "厚手の鍋にサラダ油を熱し、豚の骨付き肉の表面に焼き目をつけます。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "※の水を加えて、煮立ったらアクをとります。★のマーマレードとしょうゆを加えて中火で煮ます。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "煮汁が半分くらいになったら、味をみて、砂糖を加えてください。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "煮汁がほとんどなくなるまで、煮からめながら、照りがでるように仕上げます。",
	//       "Images": null
	//     }
	//   ]
	// }
}

func ExampleParser_b() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://cookpad.com/recipe/1948575")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "ビーフシチュー・イタリアン",
	//   "Image": "https://img.cpcdn.com/recipes/1948575/m/a76fc44f26fe07abfdf08902f966df60?u=1252112\u0026p=1347192368",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "牛肉（バラ）有ればスネ肉",
	//           "Amount": "1キロ",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "玉ねぎ（みじん切り）",
	//           "Amount": "1個",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "人参（みじん切り）",
	//           "Amount": "2本",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "セロリ（みじん切り）",
	//           "Amount": "1本",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "赤ワイン",
	//           "Amount": "500cc",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "トマト缶",
	//           "Amount": "1個",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "塩コショウ",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "小麦粉（薄力粉）",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "粉チーズ",
	//           "Amount": "お好みで",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "まずはソフリットを作ります。フライパンにみじん切りにした玉ねぎ、人参、セロリをオリーブオイル（大３）で炒めます。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "水分が無くなるまで炒めます。本格的なソフリットは形が無くなるまで炒めますが、煮込んだら一緒なのでこの程度でもＯＫです。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "多めに作ってタッパーで冷凍すれば他の料理でも使えます。今回は半分を使いました。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "牛肉を大きめにカットして塩コショウし小麦粉を振ります。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "オリーブオイルで全面に香ばしく焼き色を付けます。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "別の鍋に肉を移し肉汁が残ったフライパンに赤ワインを入れアルコールを飛ばします。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "肉にソフリットと赤ワインを入れ強火で鍋を揺すりながら煮詰めていきます。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "これくらいまで煮詰めます。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "水をヒタヒタになるまで入れ（約１０００cc）トマト缶を入れて強火で熱し沸騰したら弱火で蓋をせずにコトコト煮ます。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "2時間以上煮たらこれくらいまで水分が無くなります。塩コショウで味を整えて好きな柔らかさになったら出来上がり。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "器に移してお好みで粉チーズを振りかけて下さい。",
	//       "Images": null
	//     }
	//   ]
	// }
}
