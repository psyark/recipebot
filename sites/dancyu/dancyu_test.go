package dancyu

import (
	"context"
	"encoding/json"
	"fmt"
)

func ExampleParser_a() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://dancyu.jp/recipe/2022_00005935.html")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))

	// output:
	// {
	//   "Title": "サーモンのクリームパスタ",
	//   "Image": "https://dancyu.jp/images/m5935.jpg",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "パスタ",
	//           "Amount": "160g",
	//           "Comment": "モリサーナ　1.45mm"
	//         },
	//         {
	//           "Name": "サーモン",
	//           "Amount": "100g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "ルッコラ",
	//           "Amount": "20g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "生クリーム",
	//           "Amount": "50g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "ケイパー",
	//           "Amount": "6g",
	//           "Comment": "酢漬け"
	//         },
	//         {
	//           "Name": "パン粉",
	//           "Amount": "3g",
	//           "Comment": "細かいタイプ"
	//         },
	//         {
	//           "Name": "塩",
	//           "Amount": "9g",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "パスタをゆでる\n鍋に3Lの湯を沸かし、塩とパスタを入れる。時々混ぜながら袋の表示通りにゆでる。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "下準備\nサーモンは1.5cm角に、ルッコラは軸を3cm幅に、葉を1cm幅に切る。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "ソースをつくる\nフライパンを中火で熱し、生クリーム、ルッコラの軸、パン粉を入れ、①の湯を70ml加えて伸ばす。",
	//       "Images": [
	//         "https://dancyu.jp/images/5935a.jpg",
	//         "https://dancyu.jp/images/5935b.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "具材を温める\nケイパーとサーモンを入れ温める。魚介には酢漬けのケイパーがよく合います。",
	//       "Images": [
	//         "https://dancyu.jp/images/5935d.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "ルッコラの葉を加える\nルッコラの葉の3/4を加えさっと和える。",
	//       "Images": [
	//         "https://dancyu.jp/images/5935e.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "合わせる\nゆであがったパスタを⑤のフライパンに入れ、混ぜ合わせる。",
	//       "Images": [
	//         "https://dancyu.jp/images/5935f.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "盛りつける\n器に盛り、飾り用に残しておいた1/4のルッコラの葉を飾り完成。",
	//       "Images": [
	//         "https://dancyu.jp/images/5935g.jpg"
	//       ]
	//     }
	//   ]
	// }
}

func ExampleParser_b() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://dancyu.jp/recipe/2021_00005129.html")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))

	// output:
	// {
	//   "Title": "豚肉と卵のココナッツジュース煮",
	//   "Image": "https://dancyu.jp/images/m5129.jpg",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "豚肩ロース肉＊1",
	//           "Amount": "300g",
	//           "Comment": "塊"
	//         },
	//         {
	//           "Name": "ゆで卵",
	//           "Amount": "2個",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "ココナッツジュース＊2",
	//           "Amount": "150ml",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "赤唐辛子",
	//           "Amount": "1～2本",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "水",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "ヌクマム",
	//           "Amount": "大さじ1／2～",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "砂糖",
	//           "Amount": "大さじ1／2～",
	//           "Comment": ""
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "A",
	//       "Children": [
	//         {
	//           "Name": "砂糖",
	//           "Amount": "大さじ2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "湯",
	//           "Amount": "大さじ4",
	//           "Comment": ""
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "B",
	//       "Children": [
	//         {
	//           "Name": "ヌクマム",
	//           "Amount": "大さじ2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "砂糖",
	//           "Amount": "大さじ2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "黒胡椒",
	//           "Amount": "小さじ1／2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "にんにく",
	//           "Amount": "2片分",
	//           "Comment": "つぶしたもの"
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "下ごしらえ\n豚肉はキッチンペーパーで水気を拭き取り、5cm角に切る。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "調味料を火にかける\n鍋にAの砂糖と半量の湯を入れて中火にかける。全体がカラメル色になったら火を止めて、残りの湯を入れ、熱いうちにBを入れて混ぜ合わせる。1を入れて全体にからめたら、ときどき上下を返しながら30分ほど漬ける。",
	//       "Images": [
	//         "https://dancyu.jp/images/5129a.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "ココナッツジュースを加える\n2の鍋を中火にかけ、煮立ってきたら火を弱める。煮詰めながら、ときどき豚肉の上下を返し、煮汁をからめていく。煮汁がしっかりと煮詰まり、豚肉の表面が熱で固まってきたら、ココナッツジュースと赤唐辛子を加え、豚肉がかぶるほどの水を加える。",
	//       "Images": [
	//         "https://dancyu.jp/images/5129b.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "煮る\n強火にかけ、沸騰したら弱火にしてアクを取る。蓋をして豚肉がやわらかくなるまで30～40分煮る。しばらく煮ると煮汁の表面に脂が浮いてくるので、気になるようならすくって取り除く。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "仕上げ\n煮汁の味見をし、ヌクマムと砂糖を加える。2でつくったカラメルの焦がし具合で味が変わるので、表記の量を目安に、ここで塩気と甘味のバランスをととのえるとよい。ゆで卵を加え、蓋はせずに中火で煮汁を煮詰めながら10分ほど煮て仕上げる。",
	//       "Images": [
	//         "https://dancyu.jp/images/5129c.jpg"
	//       ]
	//     }
	//   ]
	// }
}

func ExampleParser_c() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://dancyu.jp/recipe/2022_00005778.html")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))

	// output:
	// {
	//   "Title": "海老ときのこのパスタ",
	//   "Image": "https://dancyu.jp/images/m5778.jpg",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "ショートパスタ",
	//           "Amount": "180g",
	//           "Comment": "リガトーニ"
	//         },
	//         {
	//           "Name": "有頭えび",
	//           "Amount": "4尾",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "マッシュルーム",
	//           "Amount": "100g",
	//           "Comment": "しいたけでも"
	//         },
	//         {
	//           "Name": "パッサータ",
	//           "Amount": "200ml",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "オリーブオイル",
	//           "Amount": "大さじ1",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "にんにく",
	//           "Amount": "1片",
	//           "Comment": "芯を抜いて潰す"
	//         },
	//         {
	//           "Name": "白ワイン",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "イタリアンパセリ",
	//           "Amount": "適量",
	//           "Comment": "刻む"
	//         },
	//         {
	//           "Name": "塩",
	//           "Amount": "9g",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "パスタをゆで始める\n鍋に3Lの湯を沸かし、塩とショートパスタを入れる。時々混ぜながら袋の表示通りにゆでる。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "マッシュルームの下ごしらえ\nマッシュルームの軸を外し、軸は1／4個に、かさは薄切りにする。",
	//       "Images": [
	//         "https://dancyu.jp/images/5778b.jpg",
	//         "https://dancyu.jp/images/5778c.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "海老の下準備\n海老の頭を外し、殻を剥き、尾っぽも取る。身は一口大に切る。",
	//       "Images": [
	//         "https://dancyu.jp/images/5778d.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "具材を焼く\nフライパンを中火で熱し、オリーブオイルとにんにくとマッシュルームの軸を入れ、香りが出てきたら、海老の頭とマッシュルームのかさも加え、にんにくは取り出す。ポイントはあまり動かさず焼きつけること。ある程度焼けたら、白ワインを回しかける。",
	//       "Images": [
	//         "https://dancyu.jp/images/5778e.jpg",
	//         "https://dancyu.jp/images/5778f.jpg",
	//         "https://dancyu.jp/images/5778g.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "パッサータを加える\nマッシュルームに火が入り、全体に水分が飛んできたらパッサータを加える。",
	//       "Images": [
	//         "https://dancyu.jp/images/5778h.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "海老の頭を潰す\nマッシャーなどで海老の頭を押して、味噌や旨味をソースに出すように炒める。",
	//       "Images": [
	//         "https://dancyu.jp/images/5778i.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "パスタと合わせる\n⑥のフライパンに、①のリガトーニとゆで汁を入れさっと合わせたら、③の海老の身も加える。ソースが煮詰まったら完成。器に盛り、イタリアンパセリを散らす。",
	//       "Images": [
	//         "https://dancyu.jp/images/5778j.jpg",
	//         "https://dancyu.jp/images/5778k.jpg",
	//         "https://dancyu.jp/images/5778l.jpg"
	//       ]
	//     }
	//   ]
	// }
}

func ExampleParser_d() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://dancyu.jp/recipe/2022_00005958.html")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))

	// output:
	// {
	//   "Title": "ピーマンと豚肉のからし酢醤油炒め",
	//   "Image": "https://dancyu.jp/images/m5958.jpg",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "豚バラ肉",
	//           "Amount": "250g",
	//           "Comment": "しゃぶしゃぶ用"
	//         },
	//         {
	//           "Name": "ピーマン",
	//           "Amount": "3個",
	//           "Comment": "約150g"
	//         },
	//         {
	//           "Name": "玉ねぎ",
	//           "Amount": "1/2個",
	//           "Comment": "約50g"
	//         },
	//         {
	//           "Name": "生姜",
	//           "Amount": "40g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "片栗粉",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "塩",
	//           "Amount": "ひとつまみ",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "黒胡椒",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "酒",
	//           "Amount": "大さじ1と1/2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "胡麻油",
	//           "Amount": "大さじ1/2",
	//           "Comment": ""
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "A（※混ぜておく）",
	//       "Children": [
	//         {
	//           "Name": "醤油",
	//           "Amount": "大さじ1と1/2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "米酢",
	//           "Amount": "大さじ1と1/2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "練り辛子",
	//           "Amount": "小さじ1と1/2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "砂糖",
	//           "Amount": "小さじ1/2",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "材料を切る\nピーマンは縦半分に切ってから斜めに細切り、玉ねぎはくし切り、生姜は太めのせん切りにする。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "豚肉に片栗粉をはたく\n豚肉は広げて片栗粉を薄くまぶし、6～7cm幅に切る。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "肉を炒める\nフライパンに胡麻油を入れて中火で熱し、2を入れてさっと炒める。色が変わり始めたら塩、黒胡椒をふり、酒を加えてほぐしながら炒める。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "仕上げる\n1を加えてさっと炒め、さらにAを加えて炒め合わせる。器に盛り、好みで黒胡椒を振る。",
	//       "Images": [
	//         "https://dancyu.jp/images/5958a.jpg"
	//       ]
	//     }
	//   ]
	// }
}

func ExampleParser_e() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://dancyu.jp/recipe/2022_00005756.html")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))

	// output:
	// {
	//   "Title": "カオマンガイ風ツナサンド",
	//   "Image": "https://dancyu.jp/images/m5756.jpg",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "角食パン",
	//           "Amount": "2枚",
	//           "Comment": ""
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "★ 具",
	//       "Children": [
	//         {
	//           "Name": "ツナ",
	//           "Amount": "",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "きゅうり",
	//           "Amount": "",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "生姜",
	//           "Amount": "",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "パクチー",
	//           "Amount": "",
	//           "Comment": ""
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "★ 調味料",
	//       "Children": [
	//         {
	//           "Name": "スイートチリソース",
	//           "Amount": "",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "パンにスイートチリソースを塗る\nベースのパンとフタのパンの片面にスイートチリソースを塗る。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "具を盛る\nきゅうりと生姜をみじん切りにし、ベースのパンに軽く水気をきったツナをのせる。その上にきゅうり、生姜をのせる。",
	//       "Images": [
	//         "https://dancyu.jp/images/5756a.jpg",
	//         "https://dancyu.jp/images/5756b.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "焼く\nフタのパンにちぎったパクチーを散らし、1にかぶせて焼く。",
	//       "Images": [
	//         "https://dancyu.jp/images/5756c.jpg"
	//       ]
	//     }
	//   ]
	// }
}

func ExampleParser_f() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://dancyu.jp/recipe/2019_00001391.html")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))

	// output:
	// {
	//   "Title": "マトンプラオ",
	//   "Image": "https://dancyu.jp/images/m1391.jpg",
	//   "IngredientGroups": [
	//     {
	//       "Name": "",
	//       "Children": [
	//         {
	//           "Name": "バスマティライス",
	//           "Amount": "1kg",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "マトン",
	//           "Amount": "1kg",
	//           "Comment": "5cm角の骨付き"
	//         },
	//         {
	//           "Name": "塩",
	//           "Amount": "小さじ1",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "サラダ油",
	//           "Amount": "100ml",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "にんにく",
	//           "Amount": "小さじ1",
	//           "Comment": "すりおろす"
	//         },
	//         {
	//           "Name": "生姜",
	//           "Amount": "小さじ1",
	//           "Comment": "すりおろす"
	//         },
	//         {
	//           "Name": "プレーンヨーグルト",
	//           "Amount": "50g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "塩",
	//           "Amount": "大さじ3",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "牛乳",
	//           "Amount": "大さじ2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "生姜",
	//           "Amount": "適量",
	//           "Comment": "細切り／仕上げ用"
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "A",
	//       "Children": [
	//         {
	//           "Name": "クローブ",
	//           "Amount": "10個",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "シナモンスティック",
	//           "Amount": "2本",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "ブラックカルダモン",
	//           "Amount": "4個",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "クミンシード",
	//           "Amount": "大さじ1",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "ベイリーフ",
	//           "Amount": "2枚",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "八角",
	//           "Amount": "2個",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "黒胡椒",
	//           "Amount": "小さじ1／2",
	//           "Comment": "粒"
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "B",
	//       "Children": [
	//         {
	//           "Name": "トマト",
	//           "Amount": "1／3個",
	//           "Comment": "薄切り"
	//         },
	//         {
	//           "Name": "生姜",
	//           "Amount": "20g",
	//           "Comment": "細切り"
	//         },
	//         {
	//           "Name": "フライドオニオン",
	//           "Amount": "5g",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "青唐辛子",
	//           "Amount": "7本",
	//           "Comment": "半分の長さに切る"
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "米を洗う\n\n米は3回ほど水で洗い、30分ほど常温の水に浸ける。ざるに上げて水気を切っておく。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "マトンのスープをつくる\n\nマトン、塩、水2l（分量外）を圧力鍋に入れ、20分加圧して火を止める。圧力鍋の蒸気が抜けたら蓋を取る。穴あきお玉などで肉を取り出し、スープと分けておく。スープの量が減っていたら、全体が1.5lになるよう水を足す。圧力鍋を使わない場合は、マトン、塩、水2lを鍋に入れて強火にかける。沸騰したらアクを取り、弱火にして2時間煮る。肉が柔らかくなったら火を止め、穴あきお玉などで肉を取り出し、スープと分けておく。スープの粗熱が取れたら、全体が1.5lになるよう水を足す。",
	//       "Images": [
	//         "https://dancyu.jp/images/1391d.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "スパイスを炒める\n\n鍋にサラダ油をひいて中火で熱し、にんにく、生姜を入れて炒める。茶色く色づいたらAのスパイスを加え、香りが立つまで炒める。",
	//       "Images": [
	//         "https://dancyu.jp/images/1391e.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "野菜を炒める\n\nBを加えて炒める。トマトが崩れてきたらヨーグルト、塩を加えてさらに炒める。",
	//       "Images": [
	//         "https://dancyu.jp/images/1391f.jpg",
	//         "https://dancyu.jp/images/1391g.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "マトンを炒める\n\n2の肉を加え、スパイスが全体になじむまで炒める。",
	//       "Images": [
	//         "https://dancyu.jp/images/1391h.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "スープを加えて煮る\n\n2のスープと牛乳を加えて混ぜ、蓋をして強火にする。沸騰したら弱めの中火で5分煮る。",
	//       "Images": [
	//         "https://dancyu.jp/images/1391i.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "米を加えて煮る\n\n1の米を加えて混ぜる。蓋をしないで、ときどき混ぜながら7分煮て汁けをとばす。",
	//       "Images": [
	//         "https://dancyu.jp/images/1391j.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "蓋と重石をして炊く\n\n鍋にホイルをかぶせ、蓋をして重石をのせる。弱火で20分炊く。蒸気を逃がすため、蓋と鍋の間に少し隙間をあけておく。",
	//       "Images": [
	//         "https://dancyu.jp/images/1391k.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "蒸らす\n\n火を止めて、そのまま10分置いて蒸らす。全体を混ぜ合わせたら、でき上がり。",
	//       "Images": [
	//         "https://dancyu.jp/images/1391l.jpg",
	//         "https://dancyu.jp/images/1391c.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "混ぜる\n\n材料をすべてボウルに入れて混ぜ合わせるだけ！",
	//       "Images": null
	//     }
	//   ]
	// }
}

func ExampleParser_g() {
	ctx := context.Background()
	rcp, err := NewParser().Parse(ctx, "https://dancyu.jp/recipe/2020_00003873.html")
	if err != nil {
		panic(err)
	}

	data, _ := json.MarshalIndent(rcp, "", "  ")
	fmt.Println(string(data))

	// output:
	// {
	//   "Title": "鮭と旬野菜のソテー　生姜醤油ソース",
	//   "Image": "https://dancyu.jp/images/m3873.jpg",
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
	//           "Name": "じゃがいも",
	//           "Amount": "中2個",
	//           "Comment": "約300g"
	//         },
	//         {
	//           "Name": "小松菜",
	//           "Amount": "3株",
	//           "Comment": "約100g"
	//         },
	//         {
	//           "Name": "「純正ごま油 濃口」",
	//           "Amount": "大さじ2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "塩",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "胡椒",
	//           "Amount": "適量",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "薄力粉",
	//           "Amount": "適量",
	//           "Comment": ""
	//         }
	//       ]
	//     },
	//     {
	//       "Name": "A 生姜醤油ソース",
	//       "Children": [
	//         {
	//           "Name": "生姜",
	//           "Amount": "1かけ",
	//           "Comment": "約15g"
	//         },
	//         {
	//           "Name": "「純正ごま油 濃口」",
	//           "Amount": "大さじ2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "醤油",
	//           "Amount": "大さじ1",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "味醂",
	//           "Amount": "大さじ1と1／2",
	//           "Comment": ""
	//         },
	//         {
	//           "Name": "梅肉",
	//           "Amount": "小さじ1／2",
	//           "Comment": ""
	//         }
	//       ]
	//     }
	//   ],
	//   "Steps": [
	//     {
	//       "Text": "「純正ごま油 濃口」が香る、生姜醤油ソースをつくる\nAの生姜1かけは皮をむいてせん切りにし、小さめのボウルに入れる。残りのA（「純正ごま油 濃口」大さじ2、醤油大さじ1、味醂大さじ1と1／2、梅肉小さじ1／2）を加え、約20分置いて生姜をしんなりさせる。ここで約20分置くことが、味を落ち着かせるポイント（つくりたては生姜の辛味が強い）。また、味醂のアルコール分が気になる場合は、味醂を耐熱容器に入れて電子レンジにかけ、アルコール分をとばすといい（600Wで1分目安）。",
	//       "Images": [
	//         "https://dancyu.jp/images/3873d.jpg",
	//         "https://dancyu.jp/images/3873e.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "鮭に塩をふって約20分置く\n生鮭2切れは、ザルなどにのせて塩少々をふり、約20分置く。このひと手間が、生魚からくさみを抜くための大切な作業。",
	//       "Images": [
	//         "https://dancyu.jp/images/3873f.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "小松菜を切る\n1と2を20分置いておく間に、野菜の下ごしらえをする。まず、小松菜3株はよく洗って水気をきり、根元に十字の切れ込みを入れて、長さ4cmに切る。",
	//       "Images": [
	//         "https://dancyu.jp/images/3873g.jpg",
	//         "https://dancyu.jp/images/3873h.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "じゃがいもは電子レンジにかける\nじゃがいも中2個はよく洗い、水気がついたまま1個ずつラップに包み、電子レンジにかける（600Wで5分目安）。粗熱が取れたらキッチンペーパーなどで包んで皮をむき、それぞれ3等分に切る。",
	//       "Images": [
	//         "https://dancyu.jp/images/3873i.jpg",
	//         "https://dancyu.jp/images/3873j.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "鮭の水分を拭き取り、胡椒をふって薄力粉をまぶす\n2の鮭は約20分置くと、身から水分が出ているはず（この水分がくさみのもと）。水分をキッチンペーパーで拭き取ってから、胡椒を軽くふり、薄力粉を茶こしなどで薄く全体にふり、余計な粉をはたく。",
	//       "Images": null
	//     },
	//     {
	//       "Text": "「純正ごま油 濃口」で、小松菜を炒める\nフライパンに「純正ごま油 濃口」大さじ1を熱し、3の小松菜を入れて塩少々をふり、中火で約1分炒める。塩をふることで小松菜から水分が出てきて、小松菜が短時間でしゃきっと炒め上がる。",
	//       "Images": [
	//         "https://dancyu.jp/images/3873r.jpg",
	//         "https://dancyu.jp/images/3873o.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "小松菜を取り出し、同じフライパンでじゃがいもを焼きつける\n6の小松菜は一度取り出し、続いて同じフライパンに4のじゃがいもを入れ、両面を2分ずつ焼く。じゃがいもは炒めるというより、表面を焼きつけてごま油の香りを移すイメージ。適度な焼き色がついたら取り出し、器に小松菜とともに盛りつける。",
	//       "Images": [
	//         "https://dancyu.jp/images/3873p.jpg",
	//         "https://dancyu.jp/images/3873q.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "「純正ごま油 濃口」で、鮭を焼く\n7のフライパンをキッチンペーパーで拭き、再度「純正ごま油 濃口」大さじ1を入れたら中火にし、5の鮭を皮側から入れる。",
	//       "Images": [
	//         "https://dancyu.jp/images/3873n.jpg",
	//         "https://dancyu.jp/images/3873s.jpg"
	//       ]
	//     },
	//     {
	//       "Text": "鮭が焼き上がったら盛りつけ、生姜醤油ソースをかける\n鮭の両面を2分ずつ焼き、ふっくらとして香ばしい焼き色がついたら、最後に皮目を焼きつける。7の小松菜とじゃがいもをのせた器に盛りつけ、1の生姜醤油ソースを好みの量かける。",
	//       "Images": [
	//         "https://dancyu.jp/images/3873t.jpg",
	//         "https://dancyu.jp/images/3873u.jpg"
	//       ]
	//     }
	//   ]
	// }
}
