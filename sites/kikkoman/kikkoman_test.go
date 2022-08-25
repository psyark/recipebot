package kikkoman

import (
	"context"
	"encoding/json"
	"fmt"
)

func ExampleParser() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://www.kikkoman.co.jp/homecook/search/recipe/00004691/index.html")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))
	// output:
	// {
	//   "Title": "基本の肉じゃが",
	//   "Image": "https://www.kikkoman.co.jp/homecook/search/recipe/img/00004691.jpg",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "じゃがいも",
	//           "Amount": "３個",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "玉ねぎ",
	//           "Amount": "1/2個",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "にんじん",
	//           "Amount": "1/2本",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "牛肩肉(薄切り・切り落とし)",
	//           "Amount": "１００ｇ",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "しらたき",
	//           "Amount": "１００ｇ",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "サラダ油",
	//           "Amount": "小さじ２",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "かつおだし",
	//           "Amount": "１と1/2カップ",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "醤油",
	//           "Amount": "大さじ２",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "みりん",
	//           "Amount": "大さじ３",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "じゃがいもはひと口大に切って水にさらし、水気をきる。玉ねぎはくし形切り、にんじんは乱切りにする。牛肉はひと口大に切る。しらたきはゆでて食べやすく切る。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "鍋にサラダ油を熱して（１）の玉ねぎを炒め、牛肉を加えてさらに炒める。にんじん、じゃがいも、しらたきも入れて炒め合わせる。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "かつおだしを注ぎ、沸騰したらアクを取り、しょうゆ、みりんを加えて落しぶたをする。沸騰したら弱火で１５分くらい煮る。",
	//       "Images": null
	//     }
	//   ]
	// }
}
