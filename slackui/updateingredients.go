package slackui

import (
	"context"

	"github.com/psyark/recipebot/async"
	"github.com/slack-go/slack"
)

// updateIngredientsButton は「主な材料を更新」ボタンです
type updateIngredientsButton struct {
	ui *UI
}

func (b *updateIngredientsButton) Render(pageID string) slack.BlockElement {
	return button("updateIngredients", pageID, "主な材料を更新")
}

func (b *updateIngredientsButton) React(callback *slack.InteractionCallback, action *slack.BlockAction) (bool, error) {
	if action.ActionID != "updateIngredients" {
		return false, nil
	}

	return true, async.UpdateIngredients(context.Background(), action.Value)
}
