package recipe

type Recipe struct {
	Title            string
	Image            string
	IngredientGroups []IngredientGroup
	Steps            []Step
}

type IngredientGroup struct {
	Name     string
	Children []Ingredient
}

type Ingredient struct {
	Name    string
	Amount  string
	Comment string
}

type Step struct {
	Text   string
	Images []string
}

func (r *Recipe) AddIngredient(group string, ingredient Ingredient) {
	for i := range r.IngredientGroups {
		g := &r.IngredientGroups[i]
		if g.Name == group {
			g.Children = append(g.Children, ingredient)
			return
		}
	}
	r.IngredientGroups = append(r.IngredientGroups, IngredientGroup{
		Name:     group,
		Children: []Ingredient{ingredient},
	})
}

func (g IngredientGroup) LongestNameWidth() int {
	longest := 0
	for _, idg := range g.Children {
		if longest < idg.NameWidth() {
			longest = idg.NameWidth()
		}
	}
	return longest
}

func (idg Ingredient) NameWidth() int {
	return len([]rune(idg.Name))
}
