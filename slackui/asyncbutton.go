package slackui

import (
	"context"

	"github.com/psyark/recipebot/async"
	"github.com/slack-go/slack"
)

// asyncButton は「主な材料を更新」ボタンです
type asyncButton struct {
	actionID  string
	label     string
	asyncType string
}

func (b *asyncButton) Render(pageID string, active bool) slack.BlockElement {
	elem := button(b.actionID, pageID, b.label)
	if active {
		elem.Style = slack.StyleDanger
	}
	return elem
}

func (b *asyncButton) React(callback *slack.InteractionCallback, action *slack.BlockAction) (bool, error) {
	if action.ActionID != b.actionID {
		return false, nil
	}

	pay := async.Payload{
		Type:      b.asyncType,
		ChannelID: callback.Channel.ID,
		Timestamp: callback.Message.Timestamp,
		RecipeID:  action.Value,
	}

	return true, async.CallAsync(context.Background(), pay)
}
