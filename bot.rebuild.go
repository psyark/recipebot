package recipebot

import (
	"context"
)

func (b *MyBot) RespondRebuild(pageID string) error {
	ctx := context.Background()

	return b.UpdateRecipe(ctx, pageID)
}
