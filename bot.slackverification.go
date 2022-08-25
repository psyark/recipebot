package recipebot

import (
	"net/http"

	"github.com/slack-go/slack/slackevents"
)

// SlackのURLVerificationへの応答
func (b *Bot) RespondURLVerification(rw http.ResponseWriter, event *slackevents.EventsAPIEvent) error {
	uvEvent := event.Data.(slackevents.EventsAPIURLVerificationEvent)
	rw.Header().Set("Content-Type", "text/plain")
	_, err := rw.Write([]byte(uvEvent.Challenge))
	return err
}
