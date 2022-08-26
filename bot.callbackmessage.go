package recipebot

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/mvdan/xurls"
	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/recipe"
	"github.com/psyark/recipebot/sites/united"
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
	rcp, err := united.Parsers.Parse(ctx, recipeURL)
	if err != nil {
		return nil, err
	}

	opt := &notionapi.QueryDatabaseOptions{Filter: notionapi.PropertyFilter{
		Property: RECIPE_ORIGINAL,
		URL:      &notionapi.TextFilterCondition{Equals: recipeURL},
	}}

	pagi, err := b.notion.QueryDatabase(ctx, RECIPE_DB_ID, opt)
	if err != nil {
		return nil, err
	}

	if len(pagi.Results) == 0 {
		opt := &notionapi.CreatePageOptions{
			Parent: &notionapi.Parent{Type: "database_id", DatabaseID: RECIPE_DB_ID},
			Properties: map[string]notionapi.PropertyValue{
				"title":         {Type: "title", Title: toRichTextArray(rcp.Title)},
				RECIPE_ORIGINAL: {Type: "url", URL: recipeURL},
				RECIPE_EVAL:     {Type: "select", Select: notionapi.SelectPropertyValueData{Name: "👀次作る"}},
			},
		}
		if rcp.GetEmoji() != "" {
			opt.Icon = &notionapi.FileOrEmoji{Type: "emoji", Emoji: rcp.GetEmoji()}
		}
		if rcp.Image != "" {
			opt.Cover = &notionapi.File{Type: "external", External: notionapi.ExternalFileData{URL: rcp.Image}}
		}

		page, err := b.notion.CreatePage(ctx, opt)
		if err != nil {
			return nil, err
		}

		return page, b.autoUpdateRecipePageContent(ctx, page.ID, rcp)
	} else {
		return &pagi.Results[0], nil
	}
}

func (b *MyBot) autoUpdateRecipePageContent(ctx context.Context, pageID string, rcp *recipe.Recipe) error {
	// 以前のブロックを削除
	pagi, err := b.notion.RetrieveBlockChildren(ctx, pageID)
	if err != nil {
		return err
	}

	if pagi.HasMore {
		return fmt.Errorf("autoUpdateRecipePageContent: Not implemented")
	}

	for _, block := range pagi.Results {
		_, err := b.notion.DeleteBlock(ctx, block.ID)
		if err != nil {
			return err
		}
	}

	opt := &notionapi.AppendBlockChildrenOptions{Children: ToNotionBlocks(rcp)}
	_, err = b.notion.AppendBlockChildren(ctx, pageID, opt)
	return err
}

func (b *MyBot) updateCategory(ctx context.Context, pageID string, category string) (*notionapi.Page, error) {
	return b.notion.UpdatePage(ctx, pageID, &notionapi.UpdatePageOptions{
		Properties: map[string]notionapi.PropertyValue{
			"分類": {Type: "select", Select: notionapi.SelectPropertyValueData{Name: category}},
		},
	})
}
