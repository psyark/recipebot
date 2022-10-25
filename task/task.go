package task

import (
	"context"
	"encoding/json"

	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"google.golang.org/genproto/googleapis/cloud/tasks/v2"
)

const (
	rebuildRecipe = "rebuildRecipe"
)

type payload struct {
	Type     string `json:"type"`
	RecipeID string `json:"recipeID"`
}

func CreateRebuildRecipeTask(ctx context.Context, recipeID string) error {
	return createTask(ctx, payload{Type: rebuildRecipe, RecipeID: recipeID})
}

func createTask(ctx context.Context, pay payload) error {
	ctCli, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return err
	}

	defer ctCli.Close()

	body, err := json.Marshal(pay)
	if err != nil {
		return err
	}

	req := &tasks.CreateTaskRequest{
		Parent: `projects/notion-recipe-importer/locations/asia-northeast1/queues/recipebot`,
		Task: &tasks.Task{
			MessageType: &tasks.Task_HttpRequest{HttpRequest: &tasks.HttpRequest{Url: "https://recipebot2-n2nmszkvha-an.a.run.app", Body: body}},
		},
	}

	_, err = ctCli.CreateTask(ctx, req)
	return err
}
