package handlers

import (
	"sync"

	antiscamapi "discord-bot/modules/antiscam/api"
	antiscamdiscord "discord-bot/modules/antiscam/infrastructure/discord"
	commissionsapi "discord-bot/modules/commissions/api"
	commissionsdiscord "discord-bot/modules/commissions/infrastructure/discord"
	"discord-bot/modules/config"
	giveawaysapi "discord-bot/modules/giveaways/api"
	giveawaysdiscord "discord-bot/modules/giveaways/infrastructure/discord"
	"discord-bot/modules/minecraft"
	moderationapi "discord-bot/modules/moderation/api"
	moderationdiscord "discord-bot/modules/moderation/infrastructure/discord"
	"discord-bot/modules/music"
	musicapi "discord-bot/modules/music/api"
	musicdiscord "discord-bot/modules/music/infrastructure/discord"
	"discord-bot/modules/persistence/storage"
	ticketsapi "discord-bot/modules/tickets/api"
	ticketsdiscord "discord-bot/modules/tickets/infrastructure/discord"

	"github.com/bwmarrin/discordgo"
)

type Handler struct {
	cfg          *config.Config
	db           storage.Database
	musicFeature musicapi.Feature
	rcon         *minecraft.Client
	mcStore      MCLinkStore
	moderation   moderationapi.Feature
	antiscam     antiscamapi.Feature
	commissions  commissionsapi.Feature
	giveaways    giveawaysapi.Feature
	tickets      ticketsapi.Feature
	customCmds   map[string]config.CustomCommandConfig
	customCmdsMu sync.RWMutex
}

func NewHandler(cfg *config.Config, db storage.Database) *Handler {
	h := &Handler{
		cfg:          cfg,
		db:           db,
		antiscam:     antiscamdiscord.NewFeature(cfg),
		moderation:   moderationdiscord.NewFeature(),
		commissions:  commissionsdiscord.NewFeature(),
		giveaways:    giveawaysdiscord.NewFeature(),
		musicFeature: musicdiscord.NewFeature(&cfg.Music),
		tickets:      ticketsdiscord.NewFeature(),
	}
	h.musicFeature.SetDJAuthorizer(h.isDJ)
	return h
}

func (h *Handler) SetMusic(m *music.Manager)   { h.musicFeature.SetManager(m) }
func (h *Handler) SetRCON(r *minecraft.Client) { h.rcon = r }
func (h *Handler) SetMCStore(s MCLinkStore)    { h.mcStore = s }
func (h *Handler) GetRCON() *minecraft.Client  { return h.rcon }
func (h *Handler) SetMuteSynchronizer(s moderationapi.MuteSynchronizer) {
	h.moderation.SetMuteSynchronizer(s)
}
func (h *Handler) RegisterAntiScam(s *discordgo.Session) { h.antiscam.Register(s) }
func (h *Handler) StartPhishingFeed()                    { h.antiscam.StartPhishingFeed() }
func (h *Handler) RestoreGiveawayTimers(s *discordgo.Session, gs *config.GuildState) {
	h.giveaways.RestoreTimers(s, gs)
}

func (h *Handler) lookupCustomCommand(name string) (config.CustomCommandConfig, bool) {
	h.customCmdsMu.RLock()
	defer h.customCmdsMu.RUnlock()
	cc, ok := h.customCmds[name]
	return cc, ok
}
