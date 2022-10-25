package async

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/psyark/recipebot/core"
	"github.com/slack-go/slack"
)

const cookingChannelID = "C03SNSP9HNV" // #料理 チャンネル

type Handler struct {
	coreService *core.Service
	slackClient *slack.Client
}

func NewHandler(coreService *core.Service, slackClient *slack.Client) *Handler {
	return &Handler{
		coreService: coreService,
		slackClient: slackClient,
	}
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

	case updateIngredients:
		ctx := context.Background()

		_, _, err := h.slackClient.PostMessage(cookingChannelID, slack.MsgOptionText("材料を設定します", true))
		if err != nil {
			return err
		}

		stockMap, err := h.coreService.GetStockMap(ctx)
		if err != nil {
			return err
		}

		result, err := h.coreService.UpdateRecipeIngredients(ctx, pay.RecipeID, stockMap)
		if err != nil {
			return err
		}

		foundItems := []string{}
		notFoundItems := []string{}
		for name, found := range result {
			if found {
				foundItems = append(foundItems, name)
			} else {
				notFoundItems = append(notFoundItems, name)
			}
		}

		if len(foundItems) != 0 {
			sort.Strings(foundItems)
			_, _, err := h.slackClient.PostMessage(cookingChannelID, slack.MsgOptionText(fmt.Sprintf("材料を設定しました: %v", foundItems), true))
			if err != nil {
				return err
			}
		}
		if len(notFoundItems) != 0 {
			sort.Strings(notFoundItems)
			_, _, err := h.slackClient.PostMessage(cookingChannelID, slack.MsgOptionText(fmt.Sprintf("材料が見つかりませんでした: %v", notFoundItems), true))
			if err != nil {
				return err
			}
		}

		return nil

	default:
		return fmt.Errorf("unknown type: %v", pay.Type)
	}
}
