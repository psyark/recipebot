module github.com/psyark/recipebot

go 1.16

replace github.com/psyark/recipebot => ./

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.6.1
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/cloudevents/sdk-go/v2 v2.11.0 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/joho/godotenv v1.4.0
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/microcosm-cc/bluemonday v1.0.19
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/mvdan/xurls v1.1.0
	github.com/psyark/jsonld v0.0.0-20220823063026-e08edf883a0d
	github.com/psyark/notionapi v0.0.0-20220822092621-3b048662c38b
	github.com/psyark/slackbot v0.0.0-20220826063954-684bcd5f8acc
	github.com/slack-go/slack v0.11.2
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.22.0 // indirect
	golang.org/x/net v0.0.0-20220812174116-3211cb980234
	golang.org/x/text v0.3.7
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f
)
