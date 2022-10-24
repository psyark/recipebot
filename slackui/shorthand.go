package slackui

import (
	"github.com/slack-go/slack"
)

// plain は type="plain_text" のテキストブロックを作成します
func plain(text string) *slack.TextBlockObject {
	return slack.NewTextBlockObject(slack.PlainTextType, text, true, false)
}

// mrkdwn は type="mrkdwn" のテキストブロックを作成します
func mrkdwn(text string) *slack.TextBlockObject {
	return slack.NewTextBlockObject(slack.MarkdownType, text, false, true)
}

// // button はButtonブロックエレメントを作成します
func button(actionID string, value string, text string) *slack.ButtonBlockElement {
	return slack.NewButtonBlockElement(actionID, value, plain(text))
}

func section(text *slack.TextBlockObject, accElem slack.BlockElement) *slack.SectionBlock {
	var acc *slack.Accessory
	if accElem != nil {
		acc = slack.NewAccessory(accElem)
	}
	return slack.NewSectionBlock(text, nil, acc)
}
