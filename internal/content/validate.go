package content

import (
	"fmt"
	"os"
)

func Validate(post *Post) error {
	if _, err := os.Stat(post.Article); err != nil {
		return fmt.Errorf("article not found: %w", err)
	}

	if _, err := os.Stat(post.Hero); err != nil {
		return fmt.Errorf("hero image not found: %w", err)
	}

	return nil
}
