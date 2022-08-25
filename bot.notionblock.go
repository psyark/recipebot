package recipebot

import (
	"strings"

	"github.com/psyark/notionapi"
	"github.com/psyark/recipebot/recipe"
)

func ToNotionBlocks(rcp *recipe.Recipe) []notionapi.Block {
	indices := []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "🔟", "🔢"}

	blocks := []notionapi.Block{}
	blocks = append(blocks, toHeading1("材料"))
	for _, group := range rcp.IngredientGroups {
		if group.Name != "" {
			blocks = append(blocks, toHeading3(group.Name))
		}

		width := group.LongestNameWidth() + 1
		for _, idg := range group.Children {
			todo := toToDo(idg.Name + strings.Repeat("　", width-idg.NameWidth()) + idg.Amount)
			if idg.Comment != "" {
				comment := toRichTextArray(" （" + idg.Comment + "）")
				comment[0].Annotations = &notionapi.Annotations{Color: "green"}
				todo.ToDo.RichText = append(todo.ToDo.RichText, comment...)
			}
			blocks = append(blocks, todo)
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

func toCallout(str string, emoji string, images []string) notionapi.Block {
	block := notionapi.Block{
		Object: "block",
		Type:   "callout",
		Callout: notionapi.CalloutBlockData{
			RichText: toRichTextArray(str),
			Icon:     &notionapi.FileOrEmoji{Type: "emoji", Emoji: emoji},
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
		Image: notionapi.File{
			Type:     "external",
			External: notionapi.ExternalFileData{URL: url},
		},
	}
}

func toRichTextArray(text string) []notionapi.RichText {
	return []notionapi.RichText{{Type: "text", Text: notionapi.Text{Content: text}}}
}
