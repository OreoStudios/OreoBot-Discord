// Package api defines the Discord playback adapter contract for music.
package api

import (
	"discord-bot/modules/music"
	"github.com/bwmarrin/discordgo"
)

// Feature exposes music slash commands without coupling the core player to Discord.
type Feature interface {
	Commands() []*discordgo.ApplicationCommand
	SetManager(*music.Manager)
	SetDJAuthorizer(func(*discordgo.Session, *discordgo.InteractionCreate) bool)
	HandleCommand(name string, session *discordgo.Session, event *discordgo.InteractionCreate) bool
}
