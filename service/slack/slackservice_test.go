package slack

// func TestXxx(t *testing.T) {
// 	ctx := context.Background()

// 	if err := godotenv.Load("../../test.env"); err != nil {
// 		t.Fatal(err)
// 	}

// 	bot := &Service{
// 		notionService: notion.New(notionapi.NewClient(os.Getenv("NOTION_API_KEY"))),
// 		client:        slack.New(os.Getenv("SLACK_BOT_USER_OAUTH_TOKEN")),
// 	}

// 	page, err := bot.autoUpdateRecipePage(ctx, "https://macaro-ni.jp/35774")
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	if err := bot.PostRecipeBlocks(ctx, cookingChannelID, page.ID); err != nil {
// 		t.Fatal(err)
// 	}
// 	fmt.Println(page.URL)
// }
