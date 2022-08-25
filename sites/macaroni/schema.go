package macaroni

type Recipe struct {
	Type               string              `json:"@type"`
	Name               string              `json:"name"`
	Image              []string            `json:"image"`
	RecipeIngredient   []string            `json:"recipeIngredient"`
	RecipeInstructions []RecipeInstruction `json:"recipeInstructions"`
}

type RecipeInstruction struct {
	Image string `json:"image"`
	Text  string `json:"text"`
}
