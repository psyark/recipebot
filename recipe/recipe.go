package recipe

import (
	"regexp"
	"strings"

	"golang.org/x/text/width"
)

var commentRegex = regexp.MustCompile(`　*（(.+?)）$`)

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

func GetIngredient(nameAndComment string, amount string) Ingredient {
	igd := Ingredient{
		Name:   width.Widen.String(strings.TrimSpace(nameAndComment)),
		Amount: width.Fold.String(strings.TrimSpace(amount)),
	}
	if match := commentRegex.FindStringSubmatch(igd.Name); len(match) != 0 {
		igd.Name = strings.TrimSuffix(igd.Name, match[0])
		igd.Comment = match[1]
	}
	return igd
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
