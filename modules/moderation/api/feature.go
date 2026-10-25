// Package api defines the Discord-facing moderation feature contract.
package api

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// MuteSynchronizer optionally mirrors Discord timeouts to a linked game server.
type MuteSynchronizer interface {
	SyncMuteToMC(discordUserID string, until time.Time, reason, byUsername string)
	SyncUnmuteToMC(discordUserID string)
}

// Feature exposes moderation commands to the Discord composition layer.
type Feature interface {
	Commands() []*discordgo.ApplicationCommand
	Handle(name string, session *discordgo.Session, event *discordgo.InteractionCreate) bool
	SetMuteSynchronizer(MuteSynchronizer)
}
