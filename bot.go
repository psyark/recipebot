package recipebot

import (
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/slackui"
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
	ui := slackui.New(
		slack.New(os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN")),
		notionapi.NewClient(os.Getenv("NOTION_API_KEY")),
	)

	functions.HTTP("main", ui.HandleHTTP)
}
