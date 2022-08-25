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

	_, err := b.updateCategory(ctx, pair[0], pair[1])
	if err != nil {
		return err
	}

	rbi, err := b.GetRecipeBlocksInfo(ctx, pair[0])
	if err != nil {
		return err
	}

	_, _, _, err = b.slack.UpdateMessage(
		event.Channel.ID,
		event.Message.Timestamp,
		slack.MsgOptionBlocks(rbi.ToSlackBlocks()...),
	)
	return err
}
