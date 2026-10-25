package shared

import "github.com/el211/GoModulith/module"

// Descriptor provides GoModulith metadata for shared contracts.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "shared",
		ImportPath:  "discord-bot/modules/shared",
		Description: "Low-level value and Discord helpers shared across feature modules",
		Version:     "0.1.0",
		Exports:     []string{"domain.ParseDuration", "discord.SendTemp", "discord.FindURLs", "discord.FindDiscordInvite", "discord.HasDiscordInvite", "discord.ExtractHost", "discord.DownloadAttachments", "discord.AsDiscordFiles", "discord.ArchivedAttachment"},
	}
}
