package minecraft

import "github.com/el211/GoModulith/module"

// PackageInfo describes Minecraft integration as an application module.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "minecraft",
		ImportPath:  "discord-bot/modules/minecraft",
		Description: "Minecraft RCON client and command transport",
		Version:     "0.1.0",
		Exports:     []string{"NewClient", "Client"},
	}
}
