module github.com/psyark/recipebot

go 1.16

replace github.com/psyark/recipebot => ./

require (
	cloud.google.com/go/cloudtasks v1.6.0
	github.com/GoogleCloudPlatform/functions-framework-go v1.6.1
	github.com/PuerkitoBio/goquery v1.8.0
	github.com/joho/godotenv v1.4.0
	github.com/kylelemons/godebug v1.1.0
	github.com/microcosm-cc/bluemonday v1.0.21
	github.com/mvdan/xurls v1.1.0
	github.com/psyark/jsonld v0.0.0-20221024015310-f6b6cba323ae
	github.com/psyark/notionapi v0.0.0-20221005024044-6d36d5bffc5a
	github.com/slack-go/slack v0.11.3
	golang.org/x/net v0.0.0-20221002022538-bcab6841153b
	golang.org/x/sync v0.1.0
	golang.org/x/text v0.3.7
	google.golang.org/genproto v0.0.0-20220920201722-2b89144ce006
)
