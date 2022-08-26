package recipebot

import (
	"context"
	"encoding/json"

	"github.com/slack-go/slack"
)

// レシピのカテゴリー変更

func (b *Bot) RespondSetCategory(event *slack.InteractionCallback, selectedValue string) error {
	ctx := context.Background()

	pair := [2]string{}
	if err := json.Unmarshal([]byte(selectedValue), &pair); err != nil {
		return err
	}

	if err := b.SetRecipeCategory(ctx, pair[0], pair[1]); err != nil {
		return err
	}

	return b.UpdateRecipeBlocks(ctx, event.Channel.ID, event.Message.Timestamp, pair[0])
}
