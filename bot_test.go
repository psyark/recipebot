package recipebot

import (
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("test.env"); err != nil {
		panic(err)
	}
}

// func TestUpdate(t *testing.T) {
// 	ctx := context.Background()

// 	bot := NewBot(
// 		slack.New(os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN")),
// 		notionapi.NewClient(os.Getenv("NOTION_API_KEY")),
// 	)

// 	if _, err := bot.autoUpdateRecipePage(ctx, "https://dancyu.jp/recipe/2020_00003879.html"); err != nil {
// 		t.Fatal(err)
// 	}
// }
