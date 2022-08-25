package recipebot

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/psyark/notionapi"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

// Slackからの呼び出しのルーティング関連

func init() {
	functions.HTTP("main", runMain)
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

// POSTリクエストに応答する
func (b *Bot) RespondPostRequest(rw http.ResponseWriter, req *http.Request) error {
	payload, err := b.getPayload(req)
	if err != nil {
		return err
	}

	// まだ動かない
	// if err := r.verifyHeader(req.Header); err != nil {
	// 	return err
	// }

	event, err := slackevents.ParseEvent(payload, slackevents.OptionNoVerifyToken())
	if err != nil {
		return err
	}

	switch event.Type {
	case slackevents.URLVerification:
		return b.RespondURLVerification(rw, &event)

	case slackevents.CallbackEvent:
		if err := b.RespondCallback(req, &event); err != nil {
			return err
		}
		return nil

	case string(slack.InteractionTypeBlockActions):
		intCb := slack.InteractionCallback{}
		if err := json.Unmarshal(payload, &intCb); err != nil {
			return err
		}
		if err := b.RespondBlockActions(req, &intCb); err != nil {
			return err
		}
		return nil

	default:
		return fmt.Errorf("unknown type: %v", event.Type)
	}
}

func (b *Bot) verifyHeader(header http.Header) error {
	sv, err := slack.NewSecretsVerifier(header, os.Getenv("SLACK_SIGNING_SECRET"))
	if err != nil {
		return err
	}

	return sv.Ensure()
}

// SlackのCallbackイベントへの応答
func (b *Bot) RespondCallback(req *http.Request, event *slackevents.EventsAPIEvent) error {
	switch innerEvent := event.InnerEvent.Data.(type) {
	case *slackevents.MessageEvent:
		return b.RespondCallbackMessage(req, innerEvent)

	default:
		return fmt.Errorf("unknown type: %v/%v", event.Type, event.InnerEvent.Type)
	}
}

// SlackのBlockActionsへの応答
func (b *Bot) RespondBlockActions(req *http.Request, event *slack.InteractionCallback) error {
	for _, ba := range event.ActionCallback.BlockActions {
		switch ba.ActionID {
		case actionCreateMenu:
			ctx := context.Background()
			page, err := b.notion.RetrievePage(ctx, ba.Value)
			if err != nil {
				return err
			}
			if _, _, err := b.slack.PostMessage(event.Channel.ID, slack.MsgOptionText(page.URL, true)); err != nil {
				return err
			}
		case actionSetCategory:
			return b.RespondSetCategory(event, ba.SelectedOption.Value)
		case actionRebuild:
			return b.RespondRebuild(ba.Value)
		default:
			return fmt.Errorf("unsupported action: %v", ba.ActionID)
		}
	}
	return nil
}

func (b *Bot) getPayload(req *http.Request) ([]byte, error) {
	switch req.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		if err := req.ParseForm(); err != nil {
			return nil, err
		}
		return []byte(req.Form.Get("payload")), nil
	case "application/json":
		return ioutil.ReadAll(req.Body)
	default:
		return nil, fmt.Errorf("unsupported content-type: %v", req.Header.Get("Content-Type"))
	}
}
