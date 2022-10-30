package rexch

import "fmt"

type ingredients []Ingredient
type instructions []Instruction

// TODO: 実装
func (r Recipe) GoString() string {
	return fmt.Sprintf(`{
	Title: %q,
	Image: %q,
	Servings: %v,
	Ingredients: %#v,
	Instructions: %#v,
}`, r.Title, r.Image, r.Servings, ingredients(r.Ingredients), instructions(r.Instructions))
}

func (i ingredients) GoString() string {
	s := "\n"
	for _, igr := range i {
		s += igr.GoString() + ",\n"
	}
	return fmt.Sprintf(`[]rexch.Ingredient{%s}`, s)
}

func (i instructions) GoString() string {
	s := "\n"
	for _, ist := range i {
		s += ist.GoString() + ",\n"
	}
	return fmt.Sprintf(`[]rexch.Instruction{%s}`, s)
}

func (i Ingredient) GoString() string {
	return fmt.Sprintf(`{Group: %q, Name: %q, Amount: %q, Comment: %q}`, i.Group, i.Name, i.Amount, i.Comment)
}

func (i Instruction) GoString() string {
	return fmt.Sprintf(`{Label: %q, Elements: %#v}`, i.Label, i.Elements)
}

func (e *TextInstructionElement) GoString() string {
	return fmt.Sprintf(`&rexch.TextInstructionElement{Text: %q}`, e.Text)
}

func (e *ImageInstructionElement) GoString() string {
	return fmt.Sprintf(`&rexch.ImageInstructionElement{URL: %q}`, e.URL)
}
