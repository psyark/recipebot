package async

import (
	"context"
	"encoding/json"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
)

const (
	TypeRebuildRecipe     = "rebuildRecipe"
	TypeUpdateIngredients = "updateIngredients"
)

type Payload struct {
	Type      string `json:"type"`
	ChannelID string `json:"channelID"`
	Timestamp string `json:"timestamp"`
	RecipeID  string `json:"recipeID"`
}

func CallAsync(ctx context.Context, pay Payload) error {
	ctCli, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return err
	}

	defer ctCli.Close()

	body, err := json.Marshal(pay)
	if err != nil {
		return err
	}

	req := &taskspb.CreateTaskRequest{
		Parent: `projects/notion-recipe-importer/locations/asia-northeast1/queues/recipebot`,
		Task: &taskspb.Task{
			MessageType: &taskspb.Task_HttpRequest{HttpRequest: &taskspb.HttpRequest{Url: "https://recipebot2-n2nmszkvha-an.a.run.app", Body: body}},
		},
	}

	_, err = ctCli.CreateTask(ctx, req)
	return err
}
