package slackui

import (
	"fmt"

	"github.com/slack-go/slack"
)

// ShowError は発生したエラーを適切な場所（現在は #検針 チャンネル）に出力します
func (ui *UI) ShowError(err error) {
	input := slack.NewPlainTextInputBlockElement(
		nil, "",
	)
	input.InitialValue = fmt.Sprintf("%#v", err)

	block := slack.NewInputBlock(
		"",
		plain(err.Error()),
		nil,
		input,
	)
	ui.slackClient.PostMessage(cookingChannelID, slack.MsgOptionBlocks(block))
}
