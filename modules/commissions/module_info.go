package commissions

import "github.com/el211/GoModulith/module"

// Descriptor provides GoModulith metadata for commissions.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "commissions",
		ImportPath:  "discord-bot/modules/commissions",
		Description: "Commission requests, client and freelancer quotes, and invoices",
		Version:     "0.1.0",
		DependsOn:   []string{"config", "localization", "payments", "persistence", "tickets"},
		Exports:     []string{"api.Feature", "infrastructure/discord.NewFeature"},
	}
}
