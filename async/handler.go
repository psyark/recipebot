package async

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/psyark/recipebot/core"
)

type Handler struct {
	coreService *core.Service
}

func NewHandler(coreService *core.Service) *Handler {
	return &Handler{coreService: coreService}
}

func (h *Handler) HandleCloudTasksRequest(rw http.ResponseWriter, req *http.Request) error {
	// X-Cloudtasks-Queuename:[rebuild-recipe]
	// X-Cloudtasks-Tasketa:[1666670855.3296249]
	// X-Cloudtasks-Taskexecutioncount:[0]
	// X-Cloudtasks-Taskname:[05885295116503454631]
	// X-Cloudtasks-Taskretrycount:[0]

	pay := payload{}
	if err := json.NewDecoder(req.Body).Decode(&pay); err != nil {
		return err
	}

	switch pay.Type {
	case rebuildRecipe:
		return h.coreService.UpdateRecipe(context.Background(), pay.RecipeID)

	default:
		return fmt.Errorf("unknown type: %v", pay.Type)
	}
}
