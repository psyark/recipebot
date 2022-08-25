package recipebot

import (
	"fmt"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/psyark/notionapi"
	"github.com/slack-go/slack"
)

/*
Google Cloud Console
https://console.cloud.google.com/functions/list?project=notion-recipe-importer

Slack #recipe
https://app.slack.com/client/T03S6UY399B/C03RVMW9S3Z

recipebot
https://api.slack.com/apps/A03SNSS0S81
*/

func init() {
	functions.HTTP("main", runMain)
}

// Bot はGoogle Cloud Functionsへの応答を行うクラスです
type Bot struct {
	slack  *slack.Client
	notion *notionapi.Client
}

func NewBot(slackClient *slack.Client, notionClient *notionapi.Client) *Bot {
	return &Bot{
		slack:  slackClient,
		notion: notionClient,
	}
}

func runMain(w http.ResponseWriter, r *http.Request) {
	bot := NewBot(
		slack.New(os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN")),
		notionapi.NewClient(os.Getenv("NOTION_API_KEY")),
	)

	switch r.Method {
	case http.MethodPost:
		err := bot.RespondPostRequest(w, r)
		if err != nil {
			if err, ok := err.(*FancyError); ok {
				bot.slack.PostMessage(botChannelID, slack.MsgOptionText(err.Error(), true))
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintln(w, err.Error())
			}
		}
	}
}
