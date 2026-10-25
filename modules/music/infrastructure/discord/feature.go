package discord

import (
	"discord-bot/modules/config"
	"discord-bot/modules/music"
	musicapi "discord-bot/modules/music/api"
	"github.com/bwmarrin/discordgo"
)

type feature struct {
	cfg     *config.MusicConfig
	manager *music.Manager
	canDJ   func(*discordgo.Session, *discordgo.InteractionCreate) bool
}

// NewFeature creates the Discord command adapter for the music module.
func NewFeature(cfg *config.MusicConfig) musicapi.Feature { return &feature{cfg: cfg} }

func (f *feature) SetManager(manager *music.Manager) { f.manager = manager }

func (f *feature) SetDJAuthorizer(authorize func(*discordgo.Session, *discordgo.InteractionCreate) bool) {
	f.canDJ = authorize
}

func (f *feature) Commands() []*discordgo.ApplicationCommand { return musicCommands() }

func (f *feature) HandleCommand(name string, s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	switch name {
	case "play", "skip", "stop", "queue", "volume", "nowplaying", "pause", "resume":
		f.handleMusicCommand(s, i, name)
		return true
	default:
		return false
	}
}

func (f *feature) authorizeDJ(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	return f.canDJ != nil && f.canDJ(s, i)
}
