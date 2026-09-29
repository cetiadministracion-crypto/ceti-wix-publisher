package config

import (
	"fmt"
	"os"
)

type Config struct {
	APIKey   string
	SiteID   string
	MemberID string
}

func Load() (Config, error) {
	cfg := Config{
		APIKey:   os.Getenv("WIX_API_KEY"),
		SiteID:   os.Getenv("WIX_SITE_ID"),
		MemberID: os.Getenv("WIX_MEMBER_ID"),
	}

	if cfg.APIKey == "" {
		return Config{}, fmt.Errorf("WIX_API_KEY is not set")
	}

	if cfg.SiteID == "" {
		return Config{}, fmt.Errorf("WIX_SITE_ID is not set")
	}

	if cfg.MemberID == "" {
		return Config{}, fmt.Errorf("WIX_MEMBER_ID is not set")
	}

	return cfg, nil
}
