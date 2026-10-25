// Package api defines the commission feature contract used by Discord composition.
package api

import "github.com/bwmarrin/discordgo"

// Feature exposes commission commands and Discord event handlers.
type Feature interface {
	Commands() []*discordgo.ApplicationCommand
	HandleCommand(name string, session *discordgo.Session, event *discordgo.InteractionCreate) bool
	HandleAutocomplete(session *discordgo.Session, event *discordgo.InteractionCreate) bool
	HandleComponent(session *discordgo.Session, event *discordgo.InteractionCreate) bool
	HandleModal(session *discordgo.Session, event *discordgo.InteractionCreate) bool
	HandleMessage(session *discordgo.Session, event *discordgo.MessageCreate)
}
