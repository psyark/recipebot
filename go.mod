module github.com/psyark/recipebot

go 1.16

replace github.com/psyark/recipebot => ./

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.6.1
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/joho/godotenv v1.4.0
	github.com/microcosm-cc/bluemonday v1.0.21
	github.com/mvdan/xurls v1.1.0
	github.com/psyark/jsonld v0.0.0-20221024015310-f6b6cba323ae
	github.com/psyark/notionapi v0.0.0-20221005024044-6d36d5bffc5a
	github.com/psyark/slackbot v0.0.0-20221001084604-53e16179df81
	github.com/sergi/go-diff v1.2.0
	github.com/slack-go/slack v0.11.3
	golang.org/x/net v0.1.0
	golang.org/x/sync v0.1.0
	golang.org/x/text v0.4.0
)
