package handler

// TODO パッケージごと廃止

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/psyark/recipebot/async"
	"github.com/psyark/recipebot/slackui"
)

type Handler struct {
	slackUI *slackui.UI
}

func NewHandler(slackUI *slackui.UI) *Handler {
	return &Handler{
		slackUI: slackUI,
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

	ctx := context.Background()

	switch pay.Type {
	case async.TypeRebuildRecipe:
		return h.slackUI.UpdateRecipeWithInteraction(ctx, pay)

	case async.TypeUpdateIngredients:
		return h.slackUI.UpdateIngredientsWithInteraction(ctx, pay)

	default:
		return fmt.Errorf("unknown type: %v", pay.Type)
	}
}
