// Package api defines the anti-scam feature lifecycle used by Discord composition.
package api

import "github.com/bwmarrin/discordgo"

// Feature registers message scanning and phishing-feed refresh behavior.
type Feature interface {
	Register(session *discordgo.Session)
	StartPhishingFeed()
}
