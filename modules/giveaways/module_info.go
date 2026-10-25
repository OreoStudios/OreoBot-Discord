package giveaways

import "github.com/el211/GoModulith/module"

// Descriptor provides GoModulith metadata for giveaways.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "giveaways",
		ImportPath:  "discord-bot/modules/giveaways",
		Description: "Giveaway creation, entries, timed endings, and winner rerolls",
		Version:     "0.1.0",
		DependsOn:   []string{"config", "localization", "persistence", "shared"},
		Exports:     []string{"api.Feature", "infrastructure/discord.NewFeature"},
	}
}
