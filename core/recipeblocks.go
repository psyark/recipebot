package core

import (
	"strings"

	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/rexch"
)

func toBlocks(rex *rexch.Recipe) []notionapi.Block {
	indices := []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "🔟", "🔢"}

	blocks := []notionapi.Block{
		{
			Object: "block",
			Type:   "synced_block",
			SyncedBlock: &notionapi.SyncedBlockBlocks{
				SyncedFrom: &notionapi.SyncedFrom{Type: "block_id", BlockID: recipe_shared_header_id},
			},
		},
		toHeading1("材料"),
	}
	for _, group := range rex.IngredientGroups() {
		if group == "" {
			items := rex.IngredientsByGroup(group)
			width := getLongestNameWidth(items) + 1
			for _, igd := range items {
				blocks = append(blocks, toIngredient(igd, width))
			}
		}
	}
	for _, group := range rex.IngredientGroups() {
		if group != "" {
			blocks = append(blocks, toHeading3(group))
			items := rex.IngredientsByGroup(group)
			width := getLongestNameWidth(items) + 1
			for _, igd := range items {
				blocks = append(blocks, toIngredient(igd, width))
			}
		}
	}
	blocks = append(blocks, toHeading1("手順"))
	for idx, ist := range rex.Instructions {
		emoji := "🔢"
		if idx < len(indices) {
			emoji = indices[idx]
		}
		blocks = append(blocks, toCallout(emoji, ist.Elements))
	}
	return blocks
}

func toHeading1(str string) notionapi.Block {
	return notionapi.Block{
		Object:   "block",
		Type:     "heading_1",
		Heading1: notionapi.HeadingBlockData{RichText: toRichTextArray(str), Color: "default"},
	}
}

func toHeading3(str string) notionapi.Block {
	return notionapi.Block{
		Object:   "block",
		Type:     "heading_3",
		Heading3: notionapi.HeadingBlockData{RichText: toRichTextArray(str), Color: "default"},
	}
}

func toToDo(str string) notionapi.Block {
	return notionapi.Block{
		Object: "block",
		Type:   "to_do",
		ToDo:   notionapi.ToDoBlockData{RichText: toRichTextArray(str), Color: "default"},
	}
}

func toIngredient(igd rexch.Ingredient, width int) notionapi.Block {
	todo := toToDo(igd.Name + strings.Repeat("　", width-len([]rune(igd.Name))) + igd.Amount)
	if igd.Comment != "" {
		comment := toRichTextArray(" （" + igd.Comment + "）")
		comment[0].Annotations = &notionapi.Annotations{Color: "green"}
		todo.ToDo.RichText = append(todo.ToDo.RichText, comment...)
	}
	return todo
}

func toCallout(emoji string, elements []rexch.InstructionElement) notionapi.Block {
	block := notionapi.Block{
		Object: "block",
		Type:   "callout",
		Callout: notionapi.CalloutBlockData{
			Icon:  &notionapi.Emoji{Type: "emoji", Emoji: emoji},
			Color: "gray_background",
		},
	}
	for i, elem := range elements {
		switch elem := elem.(type) {
		case *rexch.TextInstructionElement:
			if i == 0 {
				block.Callout.RichText = toRichTextArray(elem.Text)
			} else {
				block.Callout.Children = append(block.Callout.Children, toParagraph(elem.Text))
			}
		case *rexch.ImageInstructionElement:
			block.Callout.Children = append(block.Callout.Children, toImage(elem.URL))
		}
	}
	return block
}

func toImage(url string) notionapi.Block {
	return notionapi.Block{
		Object: "block",
		Type:   "image",
		Image:  &notionapi.File{Type: "external", External: &notionapi.ExternalFileData{URL: url}},
	}
}

func toRichTextArray(text string) []notionapi.RichText {
	return []notionapi.RichText{{Type: "text", Text: &notionapi.Text{Content: text}}}
}

func toParagraph(str string) notionapi.Block {
	return notionapi.Block{
		Object: "block",
		Type:   "paragraph",
		Paragraph: notionapi.ParagraphBlockData{
			RichText: toRichTextArray(str),
		},
	}
}

func getLongestNameWidth(ingredients []rexch.Ingredient) int {
	longest := 0
	for _, igd := range ingredients {
		if longest < len([]rune(igd.Name)) {
			longest = len([]rune(igd.Name))
		}
	}
	return longest
}
