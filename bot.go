package recipebot

import (
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

const (
	actionSetCategory = "set_category"
	actionCreateMenu  = "create_menu"
	actionRebuild     = "rebuild"
	botMemberID       = "U03SCN7MYEQ"
	botChannelID      = "D03SNU2C80H"
	RECIPE_DB_ID      = "ff24a40498c94ac3ac2fa8894ac0d489"
	RECIPE_ORIGINAL   = "%5CiX%60"
	RECIPE_EVAL       = "Ha%3Ba"
	RECIPE_CATEGORY   = "gmv%3A"
)

// Bot はGoogle Cloud Functionsへの応答を行うクラスです
type Bot struct {
	urlVerifier
	recipeService
	chatService
	slack  *slack.Client
	notion *notionapi.Client
}

func NewBot(slackClient *slack.Client, notionClient *notionapi.Client) *Bot {
	return &Bot{
		recipeService: recipeService{
			client: notionClient,
		},
		chatService: chatService{
			notion: notionClient,
			slack:  slackClient,
		},
		slack:  slackClient,
		notion: notionClient,
	}
}
