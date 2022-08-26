package recipebot

import (
	"net/http"

	"github.com/slack-go/slack/slackevents"
)

type URLVerifier interface {
	URLVerify(http.ResponseWriter, slackevents.EventsAPIURLVerificationEvent) error
}

var _ URLVerifier = urlVerifier{}

type urlVerifier struct{}

func (s urlVerifier) URLVerify(rw http.ResponseWriter, uvEvent slackevents.EventsAPIURLVerificationEvent) error {
	rw.Header().Set("Content-Type", "text/plain")
	_, err := rw.Write([]byte(uvEvent.Challenge))
	return err
}
