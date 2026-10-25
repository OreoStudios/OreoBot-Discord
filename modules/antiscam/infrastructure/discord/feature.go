package discord

import (
	"discord-bot/modules/antiscam/api"
	"discord-bot/modules/config"
)

type feature struct {
	cfg *config.Config
}

// NewFeature creates the Discord anti-scam adapter.
func NewFeature(cfg *config.Config) api.Feature { return &feature{cfg: cfg} }
