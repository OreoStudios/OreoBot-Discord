package discord

import (
	"strings"

	"discord-bot/modules/config"
	"github.com/bwmarrin/discordgo"
)

// Feature is the Discord adapter for giveaway commands and entry buttons.
type Feature struct{}

func NewFeature() *Feature { return &Feature{} }

func (*Feature) Commands() []*discordgo.ApplicationCommand { return giveawayCommands() }

func (*Feature) HandleCommand(name string, s *discordgo.Session, i *discordgo.InteractionCreate, authorized bool) bool {
	if name != "giveaway" {
		return false
	}
	handleGiveawayCommand(s, i, authorized)
	return true
}

func (*Feature) HandleComponent(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	id := i.MessageComponentData().CustomID
	if strings.HasPrefix(id, "giveaway_enter:") {
		HandleGiveawayEnter(s, i)
		return true
	}
	return strings.HasPrefix(id, "giveaway_ended_")
}

func (*Feature) RestoreTimers(s *discordgo.Session, gs *config.GuildState) {
	RestoreGiveawayTimers(s, gs)
}
