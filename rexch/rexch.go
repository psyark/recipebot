// Package rexch はレシピの交換用のデータ構造を提供します
package rexch

import (
	"regexp"
	"strings"

	"golang.org/x/text/width"
)

var commentRegex = regexp.MustCompile(`　*（(.+?)）$`)

type Recipe struct {
	Title        string        `json:"title"`
	Image        string        `json:"image"`
	Servings     int           `json:"servings,omitempty"`
	Ingredients  []Ingredient  `json:"ingredients"`
	Instructions []Instruction `json:"instructions"`
}

func (r *Recipe) IngredientGroups() []string {
	groups := []string{}
	exists := map[string]bool{}
	for _, i := range r.Ingredients {
		if !exists[i.Group] {
			exists[i.Group] = true
			groups = append(groups, i.Group)
		}
	}
	return groups
}

func (r *Recipe) IngredientsByGroup(group string) []Ingredient {
	igds := []Ingredient{}
	for _, i := range r.Ingredients {
		if i.Group == group {
			igds = append(igds, i)
		}
	}
	return igds
}

type Ingredient struct {
	Group   string `json:"group,omitempty"`
	Name    string `json:"name"`
	Amount  string `json:"amount,omitempty"`
	Comment string `json:"comment,omitempty"`
}

func NewIngredient(nameAndComment string, amount string) *Ingredient {
	idg := &Ingredient{
		Name:   width.Widen.String(strings.TrimSpace(nameAndComment)),
		Amount: width.Fold.String(strings.TrimSpace(amount)),
	}
	if match := commentRegex.FindStringSubmatch(idg.Name); len(match) != 0 {
		idg.Name = strings.TrimSuffix(idg.Name, match[0])
		idg.Comment = match[1]
	}
	return idg
}

type Instruction struct {
	Label    string               `json:"label,omitempty"` // 「準備」など
	Elements []InstructionElement `json:"elements"`
}

func (i *Instruction) AddText(text string) {
	if len(i.Elements) != 0 {
		if elem, ok := i.Elements[len(i.Elements)-1].(*TextInstructionElement); ok {
			if elem.Text != "" {
				elem.Text += "\n"
			}
			elem.Text += text
			return
		}
	}
	i.Elements = append(i.Elements, &TextInstructionElement{Text: text})
}

func (i *Instruction) AddImage(url string) {
	i.Elements = append(i.Elements, &ImageInstructionElement{URL: url})
}

type InstructionElement interface {
	instructionElement()
}

type TextInstructionElement struct {
	Text string `json:"text"`
}

type ImageInstructionElement struct {
	URL string `json:"url"`
}

func (e *TextInstructionElement) instructionElement()  {}
func (e *ImageInstructionElement) instructionElement() {}

var (
	_ InstructionElement = &TextInstructionElement{}
	_ InstructionElement = &ImageInstructionElement{}
)
