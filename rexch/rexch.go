// Package rexch はレシピの交換用のデータ構造を提供します
package rexch

type Recipe struct {
	Title        string        `json:"title"`
	Image        string        `json:"image"`
	Servings     int           `json:"servings,omitempty"`
	Ingredients  []Ingredient  `json:"ingredients"`
	Instructions []Instruction `json:"instructions"`
}

type Ingredient struct {
	Group   string `json:"group,omitempty"`
	Name    string `json:"name"`
	Amount  string `json:"amount,omitempty"`
	Comment string `json:"comment,omitempty"`
}

type Instruction struct {
	Label    string               `json:"label,omitempty"` // 「準備」など
	Elements []InstructionElement `json:"elements"`
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
