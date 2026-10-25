package moderation

import "github.com/el211/GoModulith/module"

// Descriptor provides GoModulith metadata for moderation.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "moderation",
		ImportPath:  "discord-bot/modules/moderation",
		Description: "Member moderation commands, warnings, and moderation logs",
		Version:     "0.1.0",
		DependsOn:   []string{"config", "localization", "persistence", "shared"},
		Exports:     []string{"api.Feature", "api.MuteSynchronizer", "infrastructure/discord.NewFeature"},
	}
}
