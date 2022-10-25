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

	return true, async.CallAsync(context.Background(), async.Payload{Type: b.asyncType, RecipeID: action.Value})
}
