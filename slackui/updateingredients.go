package slackui

import (
	"context"
	"fmt"
	"sort"

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

	ctx := context.Background()

	stockMap, err := b.ui.coreService.GetStockMap(ctx)
	if err != nil {
		return true, err
	}

	result, err := b.ui.coreService.UpdateRecipeIngredients(ctx, action.Value, stockMap)
	if err != nil {
		return true, err
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

	if len(foundItems) != 0 {
		sort.Strings(foundItems)
		_, _, err := b.ui.slackClient.PostMessage(callback.Channel.ID, slack.MsgOptionText(fmt.Sprintf("材料を設定しました: %v", foundItems), true))
		if err != nil {
			return true, err
		}
	}
	if len(notFoundItems) != 0 {
		sort.Strings(notFoundItems)
		_, _, err := b.ui.slackClient.PostMessage(callback.Channel.ID, slack.MsgOptionText(fmt.Sprintf("材料が見つかりませんでした: %v", notFoundItems), true))
		if err != nil {
			return true, err
		}
	}

	return true, nil
}
