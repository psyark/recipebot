package recipe

import (
	"strings"
)

type emojiRule struct {
	Emoji     string
	Substring string
}

var emojiRules = []emojiRule{}

func init() {
	rule := func(emoji string, names ...string) {
		for _, name := range names {
			emojiRules = append(emojiRules, emojiRule{emoji, name})
		}
	}

	// 特徴的な料理名は一番上
	rule("🥘", "パエリア")
	rule("🍲", "すき焼き")
	rule("🍳", "目玉焼き")
	rule("🥟", "シュウマイ", "焼売")
	// 特徴的な素材は上へ
	rule("🥥", "ココナッツ")
	rule("🍓", "いちご")
	rule("🥦", "ブロッコリ")
	rule("🍅", "トマト")
	rule("🐟", "サーモン", "鮭")
	rule("🥓", "ベーコン")
	rule("🌰", "栗", "マロン")
	rule("🥫", "コンビーフ")
	// 一般的な料理名
	rule("🥗", "サラダ")
	rule("🥪", "サンド")
	rule("🍝", "パスタ", "スパゲティ", "スパゲッティ")
	// 一般的な素材は下へ
	rule("🍤", "海老")
	rule("🍄", "きのこ", "茸", "えのき")
	rule("🐖", "豚")
	rule("🐔", "鶏", "鳥")
	rule("🥚", "卵", "エッグ")
	rule("🍆", "なす", "ナス", "茄子")
	rule("🫑", "ピーマン")
	rule("🧅", "玉ねぎ", "たまねぎ", "タマネギ")
	rule("🧀", "チーズ")
	rule("🐟", "ツナ")
	rule("🍖", "肉")
}

func (r Recipe) GetEmoji() string {
	for _, rule := range emojiRules {
		if strings.Contains(r.Title, rule.Substring) {
			return rule.Emoji
		}
	}
	return ""
}
