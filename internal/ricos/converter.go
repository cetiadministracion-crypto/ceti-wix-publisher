package ricos

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

type RichContent struct {
	Nodes []Node `json:"nodes"`
}

type Node struct {
	Type             string         `json:"type"`
	ID               string         `json:"id"`
	Nodes            []Node         `json:"nodes"`
	TextData         *TextData      `json:"textData,omitempty"`
	ParagraphData    map[string]any `json:"paragraphData,omitempty"`
	HeadingData      *HeadingData   `json:"headingData,omitempty"`
	BulletedListData map[string]any `json:"bulletedListData,omitempty"`
}

type TextData struct {
	Text        string       `json:"text"`
	Decorations []Decoration `json:"decorations"`
}

type Decoration struct {
	Type           string          `json:"type"`
	LinkData       *LinkData       `json:"linkData,omitempty"`
	DecorationData *DecorationData `json:"decorationData,omitempty"`
}

type LinkData struct {
	Link *Link `json:"link,omitempty"`
}

type Link struct {
	URL    string `json:"url"`
	Target string `json:"target,omitempty"`
}

type DecorationData struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type HeadingData struct {
	Level     int            `json:"level"`
	TextStyle map[string]any `json:"textStyle"`
}

func Convert(markdown string) RichContent {
	lines := strings.Split(markdown, "\n")
	var nodes []Node
	var paragraphLines []string
	var bulletItems []string

	flushParagraph := func() {
		if len(paragraphLines) == 0 {
			return
		}

		text := strings.Join(paragraphLines, " ")
		nodes = append(nodes, paragraph(text))
		paragraphLines = nil
	}

	flushBullets := func() {
		if len(bulletItems) == 0 {
			return
		}

		var items []Node
		for _, item := range bulletItems {
			items = append(items, Node{
				Type: "LIST_ITEM",
				ID:   newID(),
				Nodes: []Node{
					paragraph(item),
				},
			})
		}

		nodes = append(nodes, Node{
			Type:  "BULLETED_LIST",
			ID:    newID(),
			Nodes: items,
			BulletedListData: map[string]any{
				"indentation": 0,
			},
		})

		bulletItems = nil
	}

	for _, raw := range lines {
		line := strings.TrimSpace(raw)

		switch {
		case line == "":
			flushParagraph()
			flushBullets()

		case strings.HasPrefix(line, "## "):
			flushParagraph()
			flushBullets()

			nodes = append(nodes, heading(
				strings.TrimSpace(strings.TrimPrefix(line, "## ")),
			))

		case strings.HasPrefix(line, "- "):
			flushParagraph()
			bulletItems = append(
				bulletItems,
				strings.TrimSpace(strings.TrimPrefix(line, "- ")),
			)

		default:
			flushBullets()
			paragraphLines = append(paragraphLines, line)
		}
	}

	flushParagraph()
	flushBullets()

	return RichContent{Nodes: nodes}
}

func paragraph(text string) Node {
	return Node{
		Type:          "PARAGRAPH",
		ID:            newID(),
		Nodes:         parseInline(text),
		ParagraphData: map[string]any{},
	}
}

func heading(text string) Node {
	return Node{
		Type: "HEADING",
		ID:   newID(),
		Nodes: []Node{
			textNode(text),
		},
		HeadingData: &HeadingData{
			Level: 2,
			TextStyle: map[string]any{
				"textAlignment": "AUTO",
			},
		},
	}
}

func textNode(text string) Node {
	return Node{
		Type:  "TEXT",
		ID:    newID(),
		Nodes: []Node{},
		TextData: &TextData{
			Text:        text,
			Decorations: []Decoration{},
		},
	}
}

func newID() string {
	b := make([]byte, 4)

	if _, err := rand.Read(b); err != nil {
		panic(err)
	}

	return hex.EncodeToString(b)
}

func parseInline(text string) []Node {
	var nodes []Node

	for {
		open := strings.Index(text, "[")
		if open == -1 {
			if text != "" {
				nodes = append(nodes, textNode(text))
			}
			break
		}

		middle := strings.Index(text[open:], "](")
		if middle == -1 {
			nodes = append(nodes, textNode(text))
			break
		}
		middle += open

		close := strings.Index(text[middle+2:], ")")
		if close == -1 {
			nodes = append(nodes, textNode(text))
			break
		}
		close += middle + 2

		if open > 0 {
			nodes = append(nodes, textNode(text[:open]))
		}

		label := text[open+1 : middle]
		url := text[middle+2 : close]

		linkNode := textNode(label)
		linkNode.TextData.Decorations = []Decoration{
			{
				Type: "LINK",
				LinkData: &LinkData{
					Link: &Link{
						URL:    url,
						Target: "BLANK",
					},
				},
			},
		}

		nodes = append(nodes, linkNode)
		text = text[close+1:]
	}

	return nodes
}
