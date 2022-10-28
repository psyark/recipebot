package core

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/psyark/notionapi"
)

var client *notionapi.Client

func init() {
	if err := godotenv.Load("../test.env"); err != nil {
		panic(err)
	}

	client = notionapi.NewClient(os.Getenv("NOTION_API_KEY"))
}

// func TestGetStockMap(t *testing.T) {
// 	ctx := context.Background()
// 	service := New(client)
// 	stockMap, err := service.GetStockMap(ctx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Println(len(stockMap))
// 	fmt.Println(stockMap)
// }

// func TestUpdateRecipe(t *testing.T) {
// 	ctx := context.Background()
// 	service := New(client)

// 	if err := service.UpdateRecipe(ctx, "956851db33774257a4a4b4d987d853cd"); err != nil {
// 		t.Fatal(err)
// 	}
// }

// func TestService(t *testing.T) {
// 	ctx := context.Background()
// 	service := New(client)

// 	stockMap, err := service.GetStockMap(ctx)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Printf("%v件のストック\n", len(stockMap))

// 	result, err := service.UpdateRecipeIngredients(ctx, "2aed1d7818ad463bb2b894ca1812571d", stockMap)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	fmt.Println(result)
// }
