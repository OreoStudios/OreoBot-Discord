package music

import "github.com/el211/GoModulith/module"

// PackageInfo describes music as an application module.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "music",
		ImportPath:  "discord-bot/modules/music",
		Description: "Discord voice playback and Lavalink backends",
		Version:     "0.1.0",
		DependsOn:   []string{"config"},
		Exports:     []string{"NewManager", "Manager", "Backend", "api.Feature", "infrastructure/discord.NewFeature"},
	}
}
