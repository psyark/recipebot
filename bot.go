package recipebot

import (
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/psyark/notionapi"
	slackservice "github.com/psyark/recipebot/service/slack"
	"github.com/psyark/slackbot"
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
	registry := slackbot.NewRegistry()

	slacksvc := slackservice.New(
		slack.New(os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN")),
		notionapi.NewClient(os.Getenv("NOTION_API_KEY")),
		registry,
	)

	opt := &slackbot.GetHandlerOption{
		Registry: registry,
		Message:  slacksvc.OnCallbackMessage,
		Error:    slacksvc.OnError,
	}
	functions.HTTP("main", slackbot.GetHandler(opt))
}
