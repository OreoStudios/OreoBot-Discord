package discord

import (
	"discord-bot/modules/moderation/api"

	"github.com/bwmarrin/discordgo"
)

type feature struct {
	muteSync api.MuteSynchronizer
}

// NewFeature constructs the Discord adapter for the moderation module.
func NewFeature() api.Feature { return &feature{} }

func (f *feature) Commands() []*discordgo.ApplicationCommand { return moderationCommands() }

func (f *feature) SetMuteSynchronizer(syncer api.MuteSynchronizer) { f.muteSync = syncer }

func (f *feature) Handle(name string, s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	switch name {
	case "ban":
		handleBan(s, i)
	case "unban":
		handleUnban(s, i)
	case "kick":
		handleKick(s, i)
	case "mute":
		handleMute(s, i, f.muteSync)
	case "unmute":
		handleUnmute(s, i, f.muteSync)
	case "warn":
		handleWarn(s, i)
	case "warnings":
		handleWarnings(s, i)
	case "clearwarnings":
		handleClearWarnings(s, i)
	case "purge", "clear":
		handlePurge(s, i)
	case "slowmode":
		handleSlowmode(s, i)
	case "lock":
		handleLock(s, i)
	case "unlock":
		handleUnlock(s, i)
	case "modlog":
		handleModlog(s, i)
	case "userinfo":
		handleUserinfo(s, i)
	default:
		return false
	}
	return true
}
