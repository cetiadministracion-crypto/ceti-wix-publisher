package content

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func Load(folder string) (*Post, error) {
	configPath := filepath.Join(folder, "post.yaml")

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("read post.yaml: %w", err)
	}

	var post Post
	if err := yaml.Unmarshal(data, &post); err != nil {
		return nil, fmt.Errorf("parse post.yaml: %w", err)
	}

	if post.Title == "" {
		return nil, fmt.Errorf("title is required")
	}
	if post.Slug == "" {
		return nil, fmt.Errorf("slug is required")
	}
	if post.Excerpt == "" {
		return nil, fmt.Errorf("excerpt is required")
	}
	if post.Article == "" {
		return nil, fmt.Errorf("article is required")
	}
	if post.Hero == "" {
		return nil, fmt.Errorf("hero is required")
	}

	if post.Language == "" {
		post.Language = "es"
	}

	post.Article = filepath.Join(folder, post.Article)
	post.Hero = filepath.Join(folder, post.Hero)

	return &post, nil
}
