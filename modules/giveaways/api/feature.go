// Package api defines the giveaway feature contract used by Discord composition.
package api

import (
	"discord-bot/modules/config"

	"github.com/bwmarrin/discordgo"
)

// Feature exposes giveaway commands, buttons, and timer restoration.
type Feature interface {
	Commands() []*discordgo.ApplicationCommand
	HandleCommand(name string, session *discordgo.Session, event *discordgo.InteractionCreate, authorized bool) bool
	HandleComponent(session *discordgo.Session, event *discordgo.InteractionCreate) bool
	RestoreTimers(session *discordgo.Session, guild *config.GuildState)
}
