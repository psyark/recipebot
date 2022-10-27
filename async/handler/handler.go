package handler

// TODO パッケージごと廃止

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/psyark/recipebot/async"
	"github.com/psyark/recipebot/core"
	"github.com/psyark/recipebot/slackui"
	"github.com/slack-go/slack"
)

type Handler struct {
	coreService *core.Service
	slackClient *slack.Client
	slackUI     *slackui.UI
}

func NewHandler(coreService *core.Service, slackClient *slack.Client, slackUI *slackui.UI) *Handler {
	return &Handler{
		coreService: coreService,
		slackClient: slackClient,
		slackUI:     slackUI,
	}
}

func (h *Handler) HandleCloudTasksRequest(rw http.ResponseWriter, req *http.Request) error {
	// X-Cloudtasks-Queuename:[rebuild-recipe]
	// X-Cloudtasks-Tasketa:[1666670855.3296249]
	// X-Cloudtasks-Taskexecutioncount:[0]
	// X-Cloudtasks-Taskname:[05885295116503454631]
	// X-Cloudtasks-Taskretrycount:[0]

	pay := async.Payload{}
	if err := json.NewDecoder(req.Body).Decode(&pay); err != nil {
		return err
	}

	switch pay.Type {
	case async.TypeRebuildRecipe:
		return h.coreService.UpdateRecipe(context.Background(), pay.RecipeID)

	case async.TypeUpdateIngredients:
		ctx := context.Background()
		return h.slackUI.UpdateIngredientsWithInteraction(ctx, pay)

	default:
		return fmt.Errorf("unknown type: %v", pay.Type)
	}
}
