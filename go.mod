module github.com/psyark/recipebot

go 1.16

replace github.com/psyark/recipebot => ./

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.6.1
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/joho/godotenv v1.4.0
	github.com/microcosm-cc/bluemonday v1.0.21
	github.com/mvdan/xurls v1.1.0
	github.com/psyark/jsonld v0.0.0-20220825042757-0bc9534f3d66
	github.com/psyark/notionapi v0.0.0-20221004085326-30a515443734
	github.com/psyark/slackbot v0.0.0-20221001084604-53e16179df81
	github.com/slack-go/slack v0.11.3
	golang.org/x/net v0.0.0-20221004154528-8021a29435af
	golang.org/x/sync v0.0.0-20220929204114-8fcdb60fdcc0
	golang.org/x/text v0.3.7
)
