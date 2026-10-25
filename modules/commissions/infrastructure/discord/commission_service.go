package discord

import (
	"discord-bot/modules/config"
	"discord-bot/modules/payments"
	"discord-bot/modules/persistence/storage"
)

type CommissionService struct {
	cfg *config.Config
	svc *payments.Service
}

func newCommissionService(cfg *config.Config, svc *payments.Service) *CommissionService {
	return &CommissionService{cfg: cfg, svc: svc}
}

func (cs *CommissionService) FindServiceByName(name string) (config.CommissionService, bool) {
	for _, svc := range cs.cfg.Commissions.Services {
		if svc.Name == name {
			return svc, true
		}
	}
	return config.CommissionService{}, false
}

func (cs *CommissionService) IsCommissionChannel(guildID, channelID string) bool {
	gs := storage.GetGuild(guildID)
	gs.Lock()
	_, ok := gs.CommissionsRuntime.OpenCommissions[channelID]
	gs.Unlock()
	return ok
}
