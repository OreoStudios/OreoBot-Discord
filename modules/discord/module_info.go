package discord

import "github.com/el211/GoModulith/module"

// PackageInfo describes Discord integration as an application module.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "discord",
		ImportPath:  "discord-bot/modules/discord",
		Description: "Discord session lifecycle, commands, and feature handlers",
		Version:     "0.1.0",
		DependsOn:   []string{"antiscam", "commissions", "config", "giveaways", "localization", "minecraft", "moderation", "music", "payments", "persistence", "shared", "tickets"},
		Exports:     []string{"bot.New", "handlers.NewHandler"},
	}
}
