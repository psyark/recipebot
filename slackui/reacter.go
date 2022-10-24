package slackui

import (
	"github.com/slack-go/slack"
)

type BlockActionReacter interface {
	React(*slack.InteractionCallback, *slack.BlockAction) (bool, error)
}

type ViewSubmissionReacter interface {
	React(*slack.InteractionCallback) (bool, *slack.ViewSubmissionResponse, error)
}

var (
	_ BlockActionReacter    = BlockActionReacters{}
	_ ViewSubmissionReacter = ViewSubmissionReacters{}
)

type BlockActionReacters []BlockActionReacter

func (r BlockActionReacters) React(callback *slack.InteractionCallback, action *slack.BlockAction) (bool, error) {
	for _, reacter := range r {
		if ok, err := reacter.React(callback, action); err != nil || ok {
			return ok, err
		}
	}
	return false, nil
}

type ViewSubmissionReacters []ViewSubmissionReacter

func (r ViewSubmissionReacters) React(callback *slack.InteractionCallback) (bool, *slack.ViewSubmissionResponse, error) {
	for _, reacter := range r {
		if ok, resp, err := reacter.React(callback); err != nil || ok {
			return ok, resp, err
		}
	}
	return false, nil, nil
}
