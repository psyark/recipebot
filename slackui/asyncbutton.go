package slackui

import (
	"context"

	"github.com/psyark/recipebot/async"
	"github.com/slack-go/slack"
)

// asyncButton は「主な材料を更新」ボタンです
type asyncButton struct {
	ui        *UI
	actionID  string
	label     string
	asyncType string
}

func (b *asyncButton) Render(pageID string) slack.BlockElement {
	return button(b.actionID, pageID, b.label)
}

func (b *asyncButton) React(callback *slack.InteractionCallback, action *slack.BlockAction) (bool, error) {
	if action.ActionID != b.actionID {
		return false, nil
	}

	// 自分自身を探してスタイル変更
	for _, block := range callback.Message.Blocks.BlockSet {
		switch block := block.(type) {
		case *slack.ActionBlock:
			for _, elem := range block.Elements.ElementSet {
				switch elem := elem.(type) {
				case *slack.ButtonBlockElement:
					if elem.ActionID == b.actionID {
						elem.Text.Text = b.label + "⏳"
					}
				}
			}
		}
	}

	_, _, _, err := b.ui.slackClient.UpdateMessage(callback.Channel.ID, callback.Message.Timestamp, slack.MsgOptionBlocks(callback.Message.Blocks.BlockSet...))
	if err != nil {
		return true, err
	}

	pay := async.Payload{
		Type:      b.asyncType,
		ChannelID: callback.Channel.ID,
		Timestamp: callback.Message.Timestamp,
		RecipeID:  action.Value,
	}

	return true, async.CallAsync(context.Background(), pay)
}
