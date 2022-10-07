package rexch

import "github.com/psyark/recipebot/recipe"

func (r *Recipe) BackCompat() *recipe.Recipe {
	b := &recipe.Recipe{
		Title: r.Title,
		Image: r.Image,
	}
	for _, idg := range r.Ingredients {
		b.AddIngredient(idg.Group, recipe.Ingredient{Name: idg.Name, Amount: idg.Amount, Comment: idg.Comment})
	}
	for _, ist := range r.Instructions {
		stp := recipe.Step{}
		for _, e := range ist.Elements {
			switch e := e.(type) {
			case *TextInstructionElement:
				stp.Text += e.Text
			case *ImageInstructionElement:
				stp.Images = append(stp.Images, e.URL)
			}
		}
		b.Steps = append(b.Steps, stp)
	}
	return b
}
