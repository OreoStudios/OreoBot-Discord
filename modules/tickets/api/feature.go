// Package api defines the ticket feature contract used by Discord composition.
package api

import "github.com/bwmarrin/discordgo"

// Feature exposes the ticket command and component handlers.
type Feature interface {
	Commands() []*discordgo.ApplicationCommand
	HandleCommand(name string, session *discordgo.Session, event *discordgo.InteractionCreate) bool
	HandleComponent(session *discordgo.Session, event *discordgo.InteractionCreate) bool
}
