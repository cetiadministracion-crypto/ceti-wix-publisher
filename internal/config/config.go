package config

import (
	"fmt"
	"os"
)

type Config struct {
	APIKey string
	SiteID string
}

func Load() (Config, error) {
	cfg := Config{
		APIKey: os.Getenv("WIX_API_KEY"),
		SiteID: os.Getenv("WIX_SITE_ID"),
	}

	if cfg.APIKey == "" {
		return Config{}, fmt.Errorf("WIX_API_KEY is not set")
	}

	if cfg.SiteID == "" {
		return Config{}, fmt.Errorf("WIX_SITE_ID is not set")
	}

	return cfg, nil
}
