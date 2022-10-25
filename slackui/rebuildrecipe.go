package slackui

import (
	"context"

	"github.com/psyark/recipebot/async"
	"github.com/slack-go/slack"
)

// rebuildRecipeButton は「レシピを再構築」ボタンです
type rebuildRecipeButton struct {
	ui *UI
}

func (b *rebuildRecipeButton) Render(pageID string) slack.BlockElement {
	return button("rebuildRecipe", pageID, "レシピを再構築")
}

func (b *rebuildRecipeButton) React(callback *slack.InteractionCallback, action *slack.BlockAction) (bool, error) {
	if action.ActionID != "rebuildRecipe" {
		return false, nil
	}

	ctx := context.Background()
	return true, async.RebuildRecipe(ctx, action.Value)
}
