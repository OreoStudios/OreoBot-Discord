package antiscam

import "github.com/el211/GoModulith/module"

// Descriptor provides GoModulith metadata for anti-scam protection.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "antiscam",
		ImportPath:  "discord-bot/modules/antiscam",
		Description: "Message risk scoring, attachment OCR, phishing feeds, and moderation actions",
		Version:     "0.1.0",
		DependsOn:   []string{"config", "localization", "persistence", "shared"},
		Exports:     []string{"api.Feature", "infrastructure/discord.NewFeature"},
	}
}
