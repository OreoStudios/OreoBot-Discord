package tickets

import "github.com/el211/GoModulith/module"

// Descriptor provides GoModulith metadata for tickets.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "tickets",
		ImportPath:  "discord-bot/modules/tickets",
		Description: "Support ticket commands, panels, channels, and transcripts",
		Version:     "0.1.0",
		DependsOn:   []string{"config", "localization", "persistence"},
		Exports:     []string{"api.Feature", "domain.IsOpenSupportChannel", "infrastructure/discord.NewFeature", "infrastructure/discord.GenerateTranscript", "infrastructure/discord.ParseComponentEmoji"},
	}
}
