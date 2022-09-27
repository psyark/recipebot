module github.com/psyark/recipebot

go 1.16

replace github.com/psyark/recipebot => ./

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.6.1
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/joho/godotenv v1.4.0
	github.com/microcosm-cc/bluemonday v1.0.20
	github.com/mvdan/xurls v1.1.0
	github.com/psyark/jsonld v0.0.0-20220825042757-0bc9534f3d66
	github.com/psyark/notionapi v0.0.0-20220927054835-47e2ac37872e
	github.com/psyark/slackbot v0.0.0-20220919132108-aa6f92ad32db
	github.com/slack-go/slack v0.11.3
	golang.org/x/net v0.0.0-20220926192436-02166a98028e
	golang.org/x/sync v0.0.0-20220923202941-7f9b1623fab7
	golang.org/x/text v0.3.7
)
