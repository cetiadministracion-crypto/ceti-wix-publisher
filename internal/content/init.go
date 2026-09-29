package content

import (
	"fmt"
	"os"
	"path/filepath"
)

const defaultPostYAML = `title: "Nuevo artículo CETI"
slug: "nuevo-articulo-ceti"
excerpt: "Descripción breve del artículo."
language: "es"
article: "article.md"
hero: "hero.png"
alt_text: "CETI English"
`

const defaultArticle = `Escribe aquí la introducción del artículo.

## Primer tema

Escribe aquí el contenido.

## Segundo tema

- Punto uno
- Punto dos
- Punto tres

## CETI English

Escribe aquí el cierre del artículo.
`

func Init(name string) (string, error) {
	folder := filepath.Join("posts", name)

	if _, err := os.Stat(folder); err == nil {
		return "", fmt.Errorf("post already exists: %s", folder)
	}

	if err := os.MkdirAll(folder, 0755); err != nil {
		return "", fmt.Errorf("create post folder: %w", err)
	}

	if err := os.WriteFile(
		filepath.Join(folder, "post.yaml"),
		[]byte(defaultPostYAML),
		0644,
	); err != nil {
		return "", fmt.Errorf("create post.yaml: %w", err)
	}

	if err := os.WriteFile(
		filepath.Join(folder, "article.md"),
		[]byte(defaultArticle),
		0644,
	); err != nil {
		return "", fmt.Errorf("create article.md: %w", err)
	}

	return folder, nil
}
