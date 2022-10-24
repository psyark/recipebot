package slackui

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

func (ui *UI) HandleHTTP(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		defer func() {
			if err := recover(); err != nil {
				ui.ShowError(fmt.Errorf("panic: %#v", err))
				rw.WriteHeader(http.StatusInternalServerError)
			}
		}()

		if err := ui.handlePostRequest(rw, req); err != nil {
			ui.ShowError(err)
			rw.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (ui *UI) handlePostRequest(rw http.ResponseWriter, req *http.Request) error {
	payload, err := ui.getPayload(req)
	if err != nil {
		return fmt.Errorf("getPayload(%#v): %#v", req, err)
	}

	// TODO:
	// if err := r.verifyHeader(req.Header); err != nil {
	// 	return err
	// }

	event, err := slackevents.ParseEvent(payload, slackevents.OptionNoVerifyToken())
	if err != nil {
		return fmt.Errorf("slackevents.ParseEvent(%#v): %w", payload, err)
	}

	switch event.Type {
	case slackevents.URLVerification:
		if uve, ok := event.Data.(*slackevents.EventsAPIURLVerificationEvent); ok {
			return ui.verifyURL(rw, uve)
		}
		return fmt.Errorf("event.Data is not *EventsAPIURLVerificationEvent: %#v", event.Data)

	case slackevents.CallbackEvent:
		if err := ui.handleCallback(req, &event); err != nil {
			return err
		}
		return nil

	case string(slack.InteractionTypeBlockActions):
		callback := slack.InteractionCallback{}
		if err := json.Unmarshal(payload, &callback); err != nil {
			return fmt.Errorf("json.Unmarshal(%#v): %w", payload, err)
		}

		for _, action := range callback.ActionCallback.BlockActions {
			if err := ui.handleBlockAction(req, &callback, action); err != nil {
				return err
			}
		}
		return nil

	case string(slack.InteractionTypeViewSubmission):
		callback := slack.InteractionCallback{}
		if err := json.Unmarshal(payload, &callback); err != nil {
			return fmt.Errorf("json.Unmarshal(%#v): %w", payload, err)
		}

		if res, err := ui.handleViewSubmission(req, &callback); err != nil {
			return err
		} else if res != nil {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			return json.NewEncoder(rw).Encode(res)
		} else {
			return nil
		}

	default:
		return fmt.Errorf("unknown type: %#v", event.Type)
	}

}

func (ui *UI) handleCallback(req *http.Request, event *slackevents.EventsAPIEvent) error {
	switch innerEvent := event.InnerEvent.Data.(type) {
	// case *slackevents.AppHomeOpenedEvent:
	// 	return ui.OnAppHomeOpened(req, innerEvent)
	case *slackevents.MessageEvent:
		return ui.handleCallbackMessage(req, innerEvent)

	default:
		return fmt.Errorf("unknown type: %#v/%#v", event.Type, event.InnerEvent.Type)
	}
}

func (ui *UI) getPayload(req *http.Request) ([]byte, error) {
	switch req.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		if err := req.ParseForm(); err != nil {
			return nil, err
		}
		return []byte(req.Form.Get("payload")), nil
	case "application/json":
		return io.ReadAll(req.Body)
	default:
		return nil, fmt.Errorf("unsupported content-type: %v", req.Header.Get("Content-Type"))
	}
}

func (ui *UI) verifyURL(rw http.ResponseWriter, uvEvent *slackevents.EventsAPIURLVerificationEvent) error {
	rw.Header().Set("Content-Type", "text/plain")
	_, err := rw.Write([]byte(uvEvent.Challenge))
	return err
}
