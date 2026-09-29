package ricos

import "testing"

func TestConvertBasicContent(t *testing.T) {
	md := `Introducción.

## Título

- Punto uno
- Punto dos`

	content := Convert(md)

	if len(content.Nodes) != 3 {
		t.Fatalf("expected 3 top-level nodes, got %d", len(content.Nodes))
	}

	if content.Nodes[0].Type != "PARAGRAPH" {
		t.Errorf("expected PARAGRAPH, got %s", content.Nodes[0].Type)
	}

	if content.Nodes[1].Type != "HEADING" {
		t.Errorf("expected HEADING, got %s", content.Nodes[1].Type)
	}

	if content.Nodes[2].Type != "BULLETED_LIST" {
		t.Errorf("expected BULLETED_LIST, got %s", content.Nodes[2].Type)
	}
}

func TestConvertMarkdownLink(t *testing.T) {
	md := `Consulta nuestras [evaluaciones de nivel](https://www.cetienglish.com/general-8).`

	content := Convert(md)

	if len(content.Nodes) != 1 {
		t.Fatalf("expected 1 paragraph, got %d", len(content.Nodes))
	}

	nodes := content.Nodes[0].Nodes

	if len(nodes) != 3 {
		t.Fatalf("expected 3 text nodes, got %d", len(nodes))
	}

	link := nodes[1]

	if link.TextData == nil || link.TextData.Text != "evaluaciones de nivel" {
		t.Fatal("link text was not parsed correctly")
	}

	if len(link.TextData.Decorations) != 1 {
		t.Fatal("expected link decoration")
	}

	decoration := link.TextData.Decorations[0]

	if decoration.LinkData == nil ||
		decoration.LinkData.Link == nil ||
		decoration.LinkData.Link.URL != "https://www.cetienglish.com/general-8" {
		t.Fatal("link URL was not parsed correctly")
	}
}
