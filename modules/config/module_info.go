package config

import "github.com/el211/GoModulith/module"

// PackageInfo describes config as an application module.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "config",
		ImportPath:  "discord-bot/modules/config",
		Description: "Application configuration, validation, and guild settings",
		Version:     "0.1.0",
		Exports:     []string{"Config", "LoadConfig", "SaveConfig", "GuildState"},
	}
}
