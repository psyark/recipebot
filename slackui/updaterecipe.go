package slackui

import (
	"context"

	"github.com/psyark/recipebot/async"
)

func (ui *UI) UpdateRecipeWithInteraction(ctx context.Context, pay async.Payload) error {
	if err := ui.coreService.UpdateRecipe(ctx, pay.RecipeID); err != nil {
		return err
	}

	page, err := ui.coreService.RetrievePage(ctx, pay.RecipeID)
	if err != nil {
		return nil
	}

	return ui.UpdateRecipeMessage(ctx, pay.ChannelID, pay.Timestamp, page, nil)
}
