# CETI Wix Publisher

Command-line application written in Go for preparing CETI English blog
posts and creating unpublished drafts in Wix Blog.

## Normal workflow

``` text
Create post → Write article → Add hero image → Validate → Create Wix draft → Review → Publish
```

## Requirements

-   Go installed.
-   Wix API key.
-   Wix Site ID.
-   Wix Member ID used as the blog post owner.

## Configuration

``` bash
export WIX_API_KEY="YOUR_WIX_API_KEY"
export WIX_SITE_ID="YOUR_WIX_SITE_ID"
export WIX_MEMBER_ID="YOUR_WIX_MEMBER_ID"
```

Never commit the Wix API key to Git.

## Build

``` bash
go test ./...
go build -o ceti-publisher ./cmd/ceti-publisher
```

## Create a new post

``` bash
./ceti-publisher init trabajo-ingles-entrevistas
```

This creates:

``` text
posts/trabajo-ingles-entrevistas/
├── article.md
└── post.yaml
```

Add the hero image as:

``` text
posts/trabajo-ingles-entrevistas/hero.png
```

The completed package is:

``` text
posts/trabajo-ingles-entrevistas/
├── post.yaml
├── article.md
└── hero.png
```

## post.yaml

Example:

``` yaml
title: "Inglés para entrevistas de trabajo"
slug: "ingles-entrevistas-trabajo"
excerpt: "Consejos para prepararte para una entrevista profesional en inglés."
language: "es"
article: "article.md"
hero: "hero.png"
alt_text: "Profesional preparándose para una entrevista en inglés"
```

CETI posts are currently written in Spanish, so normally use
`language: "es"`.

## article.md

Example:

``` markdown
Prepararte para una entrevista en inglés requiere más que memorizar respuestas.

## Preguntas frecuentes

Practica situaciones que realmente pueden aparecer durante una entrevista.

- Presentar tu experiencia profesional.
- Explicar tus responsabilidades.
- Hablar de tus objetivos profesionales.

Conoce más en [CETI English](https://www.cetienglish.com/).
```

The current converter supports normal paragraphs, H2 headings (`##`),
bullet lists (`-`), and Markdown links.

## Validate

``` bash
./ceti-publisher validate posts/trabajo-ingles-entrevistas
```

A successful validation confirms the post metadata, article, and hero
image are available.

## Create the Wix draft

``` bash
./ceti-publisher create posts/trabajo-ingles-entrevistas
```

The application performs:

``` text
Validate package
      ↓
Read article.md
      ↓
Convert Markdown to Wix Rich Content
      ↓
Upload hero.png to Wix Media
      ↓
Create Wix Blog draft
      ↓
Return Wix Post ID
```

A successful result ends with `Status: UNPUBLISHED`.

The tool does not automatically publish the article. Review the draft in
Wix before publishing.

## Other commands

Inspect generated Wix Rich Content JSON:

``` bash
./ceti-publisher convert posts/trabajo-ingles-entrevistas/article.md
```

Upload an image directly:

``` bash
./ceti-publisher upload /path/to/image.png
```

Show available commands:

``` bash
./ceti-publisher
```

## Recommended CETI workflow

``` text
1. ./ceti-publisher init <post-name>
2. Edit posts/<post-name>/post.yaml
3. Write posts/<post-name>/article.md
4. Add posts/<post-name>/hero.png
5. ./ceti-publisher validate posts/<post-name>
6. ./ceti-publisher create posts/<post-name>
7. Review the draft in Wix
8. Publish from Wix when approved
```

## Repository structure

``` text
cmd/ceti-publisher/   Main CLI
internal/config/      Wix configuration
internal/content/     Post initialization, loading and validation
internal/ricos/       Markdown → Wix Rich Content
internal/wix/         Wix API, media upload and draft creation
posts/                Local blog post workspace
```

## Security

Keep credentials in environment variables. Never put `WIX_API_KEY` in
source code, `post.yaml`, `article.md`, or Git commits.

The repository `.gitignore` should exclude `.env`, the compiled binary,
and generated content under `posts/`.

## Current status

The publisher has been tested end-to-end with CETI English. It can
validate a local post package, convert Markdown, upload the hero image,
and create an unpublished Wix Blog draft successfully.
