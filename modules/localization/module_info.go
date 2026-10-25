package localization

import "github.com/el211/GoModulith/module"

// PackageInfo describes localization as an application module.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "localization",
		ImportPath:  "discord-bot/modules/localization",
		Description: "Language file loading and translation lookup",
		Version:     "0.1.0",
		Exports:     []string{"lang.Load", "lang.Reload", "lang.T", "lang.TL"},
	}
}
