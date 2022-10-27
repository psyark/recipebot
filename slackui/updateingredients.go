package slackui

import (
	"context"
	"fmt"
	"sort"

	"github.com/psyark/recipebot/async"
)

func (ui *UI) UpdateIngredientsWithInteraction(ctx context.Context, pay async.Payload) error {
	page, err := ui.coreService.RetrievePage(ctx, pay.RecipeID)
	if err != nil {
		return nil
	}

	if err := ui.UpdateRecipeMessage(ctx, pay.ChannelID, pay.Timestamp, page, &RecipeBlocksOption{IsUpdateIngredientsButtonActive: true}); err != nil {
		return nil
	}

	stockMap, err := ui.coreService.GetStockMap(ctx)
	if err != nil {
		return err
	}

	result, err := ui.coreService.UpdateRecipeIngredients(ctx, pay.RecipeID, stockMap)
	if err != nil {
		return err
	}

	foundItems := []string{}
	notFoundItems := []string{}
	for name, found := range result {
		if found {
			foundItems = append(foundItems, name)
		} else {
			notFoundItems = append(notFoundItems, name)
		}
	}

	additonalText := ""
	if len(foundItems) != 0 {
		sort.Strings(foundItems)
		additonalText += fmt.Sprintf("材料を設定しました: %v\n", foundItems)
	}
	if len(notFoundItems) != 0 {
		sort.Strings(notFoundItems)
		additonalText += fmt.Sprintf("材料が見つかりませんでした: %v\n", notFoundItems)
	}

	if err := ui.UpdateRecipeMessage(ctx, pay.ChannelID, pay.Timestamp, page, &RecipeBlocksOption{AdditionalText: additonalText}); err != nil {
		return nil
	}

	return nil
}
