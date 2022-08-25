package recipebot

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/mvdan/xurls"
	"github.com/slack-go/slack"
	"github.com/slack-go/slack/slackevents"
)

// BotのSlack側の処理を書く

const (
	actionSetCategory = "set_category"
	actionCreateMenu  = "create_menu"
	botMemberID       = "U03SCN7MYEQ"
	botChannelID      = "D03SNU2C80H"
)

var (
	categories = []string{"主食", "主菜", "副菜", "汁物", "弁当", "おつまみ", "デザート"}
)

// POSTリクエストに応答する
func (r *Bot) RespondPostRequest(rw http.ResponseWriter, req *http.Request) error {
	payload, err := r.getPayload(req)
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
		return r.RespondURLVerification(rw, &event)

	case slackevents.CallbackEvent:
		if err := r.RespondCallback(req, &event); err != nil {
			return err
		}
		return nil

	case string(slack.InteractionTypeBlockActions):
		intCb := slack.InteractionCallback{}
		if err := json.Unmarshal(payload, &intCb); err != nil {
			return err
		}
		if err := r.RespondBlockActions(req, &intCb); err != nil {
			return err
		}
		return nil

	default:
		return fmt.Errorf("unknown type: %v", event.Type)
	}
}

func (r *Bot) verifyHeader(header http.Header) error {
	sv, err := slack.NewSecretsVerifier(header, os.Getenv("SLACK_SIGNING_SECRET"))
	if err != nil {
		return err
	}

	return sv.Ensure()
}

// SlackのURLVerificationへの応答
func (r *Bot) RespondURLVerification(rw http.ResponseWriter, event *slackevents.EventsAPIEvent) error {
	uvEvent := event.Data.(slackevents.EventsAPIURLVerificationEvent)
	rw.Header().Set("Content-Type", "text/plain")
	_, err := rw.Write([]byte(uvEvent.Challenge))
	return err
}

// SlackのCallbackイベントへの応答
func (r *Bot) RespondCallback(req *http.Request, event *slackevents.EventsAPIEvent) error {
	switch innerEvent := event.InnerEvent.Data.(type) {
	case *slackevents.MessageEvent:
		return r.RespondCallbackMessage(req, innerEvent)

	default:
		return fmt.Errorf("unknown type: %v/%v", event.Type, event.InnerEvent.Type)
	}
}

// SlackのCallbackMessageへの応答
func (r *Bot) RespondCallbackMessage(req *http.Request, event *slackevents.MessageEvent) error {
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

		if err := r.slack.AddReaction("thumbsup", ref); err != nil {
			return &FancyError{err}
		}

		page, err := r.autoUpdateRecipePage(ctx, url)
		if err != nil {
			return &FancyError{err}
		}

		rbi, err := r.GetRecipeBlocksInfo(ctx, page.ID)
		if err != nil {
			return &FancyError{err}
		}

		_, _, err = r.slack.PostMessage(event.Channel, slack.MsgOptionBlocks(CreateRecipeBlocks(rbi)...))
		if err != nil {
			return &FancyError{err}
		}

		return nil
	} else {
		return r.slack.AddReaction("thinking_face", ref)
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
		default:
			return fmt.Errorf("unsupported action: %v", ba.ActionID)
		}
	}
	return nil
}

func (b *Bot) RespondSetCategory(event *slack.InteractionCallback, selectedValue string) error {
	ctx := context.Background()

	pair := [2]string{}
	if err := json.Unmarshal([]byte(selectedValue), &pair); err != nil {
		return err
	}

	_, err := b.updateCategory(ctx, pair[0], pair[1])
	if err != nil {
		return err
	}

	rbi, err := b.GetRecipeBlocksInfo(ctx, pair[0])
	if err != nil {
		return err
	}

	_, _, _, err = b.slack.UpdateMessage(
		event.Channel.ID,
		event.Message.Timestamp,
		slack.MsgOptionBlocks(CreateRecipeBlocks(rbi)...),
	)
	return err
}

func (r *Bot) getPayload(req *http.Request) ([]byte, error) {
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
