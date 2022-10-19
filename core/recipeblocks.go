package core

import (
	"strings"

	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/recipe"
)

func toBlocks(rcp recipe.Recipe) []notionapi.Block {
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
	for _, group := range rcp.IngredientGroups {
		if group.Name == "" {
			width := group.LongestNameWidth() + 1
			for _, igd := range group.Children {
				blocks = append(blocks, toIngredient(igd, width))
			}
		}
	}
	for _, group := range rcp.IngredientGroups {
		if group.Name != "" {
			blocks = append(blocks, toHeading3(group.Name))
			width := group.LongestNameWidth() + 1
			for _, igd := range group.Children {
				blocks = append(blocks, toIngredient(igd, width))
			}
		}
	}
	blocks = append(blocks, toHeading1("手順"))
	for idx, stp := range rcp.Steps {
		emoji := "🔢"
		if idx < len(indices) {
			emoji = indices[idx]
		}
		blocks = append(blocks, toCallout(stp.Text, emoji, stp.Images))
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

func toIngredient(igd recipe.Ingredient, width int) notionapi.Block {
	todo := toToDo(igd.Name + strings.Repeat("　", width-igd.NameWidth()) + igd.Amount)
	if igd.Comment != "" {
		comment := toRichTextArray(" （" + igd.Comment + "）")
		comment[0].Annotations = &notionapi.Annotations{Color: "green"}
		todo.ToDo.RichText = append(todo.ToDo.RichText, comment...)
	}
	return todo
}

func toCallout(str string, emoji string, images []string) notionapi.Block {
	block := notionapi.Block{
		Object: "block",
		Type:   "callout",
		Callout: notionapi.CalloutBlockData{
			RichText: toRichTextArray(str),
			Icon:     &notionapi.Emoji{Type: "emoji", Emoji: emoji},
			Color:    "gray_background",
		},
	}
	for _, url := range images {
		block.Callout.Children = append(block.Callout.Children, toImage(url))
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
