package main_test

import (
	"testing"

	antiscammodule "discord-bot/modules/antiscam"
	commissionsmodule "discord-bot/modules/commissions"
	configmodule "discord-bot/modules/config"
	discordmodule "discord-bot/modules/discord"
	giveawaysmodule "discord-bot/modules/giveaways"
	localizationmodule "discord-bot/modules/localization"
	minecraftmodule "discord-bot/modules/minecraft"
	moderationmodule "discord-bot/modules/moderation"
	musicmodule "discord-bot/modules/music"
	paymentsmodule "discord-bot/modules/payments"
	persistencemodule "discord-bot/modules/persistence"
	sharedmodule "discord-bot/modules/shared"
	ticketsmodule "discord-bot/modules/tickets"
	"github.com/el211/GoModulith/architecture"
	"github.com/el211/GoModulith/module"
)

// TestModuleBoundaries keeps the application module graph explicit. Add a
// dependency here only when the owning module genuinely needs that capability.
func TestModuleBoundaries(t *testing.T) {
	metadata := []module.Provider{
		configmodule.Descriptor{},
		antiscammodule.Descriptor{},
		commissionsmodule.Descriptor{},
		discordmodule.Descriptor{},
		giveawaysmodule.Descriptor{},
		localizationmodule.Descriptor{},
		minecraftmodule.Descriptor{},
		moderationmodule.Descriptor{},
		musicmodule.Descriptor{},
		paymentsmodule.Descriptor{},
		persistencemodule.Descriptor{},
		sharedmodule.Descriptor{},
		ticketsmodule.Descriptor{},
	}
	seen := make(map[string]bool, len(metadata))
	for _, provider := range metadata {
		info := provider.PackageInfo()
		if info.Name == "" || info.ImportPath == "" || info.Description == "" || len(info.Exports) == 0 {
			t.Errorf("incomplete GoModulith metadata: %+v", info)
		}
		if seen[info.Name] {
			t.Errorf("duplicate module metadata name %q", info.Name)
		}
		seen[info.Name] = true
	}

	report, err := architecture.Analyze(architecture.Config{
		Root:       ".",
		ModulesDir: "modules",
		AllowedDependencies: map[string][]string{
			"discord":      {"antiscam", "commissions", "config", "giveaways", "localization", "minecraft", "moderation", "music", "payments", "persistence", "shared", "tickets"},
			"antiscam":     {"config", "localization", "persistence", "shared"},
			"commissions":  {"config", "localization", "payments", "persistence", "tickets"},
			"payments":     {"config", "persistence"},
			"giveaways":    {"config", "localization", "persistence", "shared"},
			"moderation":   {"config", "localization", "persistence", "shared"},
			"shared":       {},
			"tickets":      {"config", "localization", "persistence"},
			"music":        {"config"},
			"persistence":  {"config"},
			"config":       {},
			"localization": {},
			"minecraft":    {},
		},
	})
	if err != nil {
		t.Fatalf("analyze module architecture: %v", err)
	}
	if err := report.Verify(); err != nil {
		t.Fatalf("module architecture violation: %v", err)
	}
}
