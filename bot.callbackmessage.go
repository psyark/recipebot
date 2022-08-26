package recipebot

import (
	"context"
	"net/http"
	"strings"

	"github.com/mvdan/xurls"
	"github.com/psyark/notionapi"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

// SlackのCallbackMessageへの応答
func (b *MyBot) RespondCallbackMessage(req *http.Request, event *slackevents.MessageEvent) error {
	if req.Header.Get("X-Slack-Retry-Num") != "" {
		return nil // リトライは無視
	} else if event.User == botMemberID {
		return nil // 自分のメッセージは無視
	}

	ctx := context.Background()
	ref := slack.NewRefToMessage(event.Channel, event.TimeStamp)
	if url := xurls.Strict.FindString(event.Text); url != "" {
		if strings.Contains(url, "|") {
			url = strings.Split(url, "|")[0]
		}

		if err := b.slack.AddReaction("thumbsup", ref); err != nil {
			return &FancyError{err}
		}

		page, err := b.autoUpdateRecipePage(ctx, url)
		if err != nil {
			return &FancyError{err}
		}

		rbi, err := b.GetRecipeBlocksInfo(ctx, page.ID)
		if err != nil {
			return &FancyError{err}
		}

		_, _, err = b.slack.PostMessage(event.Channel, slack.MsgOptionBlocks(rbi.ToSlackBlocks()...))
		if err != nil {
			return &FancyError{err}
		}

		return nil
	} else {
		return b.slack.AddReaction("thinking_face", ref)
	}
}

func (b *MyBot) autoUpdateRecipePage(ctx context.Context, recipeURL string) (*notionapi.Page, error) {
	// レシピページを取得
	page, err := b.GetRecipeByURL(ctx, recipeURL)
	if err != nil {
		return nil, err
	}

	if page != nil {
		return page, nil
	}

	// レシピページがなければ作成
	return b.CreateRecipe(ctx, recipeURL)
}
