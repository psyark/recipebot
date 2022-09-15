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
	github.com/psyark/notionapi v0.0.0-20220822092621-3b048662c38b
	github.com/psyark/slackbot v0.0.0-20220906030041-86efe33d1eec
	github.com/slack-go/slack v0.11.3
	golang.org/x/net v0.0.0-20220909164309-bea034e7d591
	golang.org/x/text v0.3.7
)
