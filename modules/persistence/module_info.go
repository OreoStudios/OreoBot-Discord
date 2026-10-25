package persistence

import "github.com/el211/GoModulith/module"

// PackageInfo describes persistence as an application module.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "persistence",
		ImportPath:  "discord-bot/modules/persistence",
		Description: "Database drivers, database lifecycle, and guild state storage",
		Version:     "0.1.0",
		DependsOn:   []string{"config"},
		Exports:     []string{"storage.InitDB", "storage.GetGuild", "storage.GetAllGuildStates"},
	}
}
