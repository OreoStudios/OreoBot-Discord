// Package domain contains ticket state rules shared by feature adapters.
package domain

import "discord-bot/modules/config"

// IsOpenSupportChannel reports whether a channel belongs to an open ticket or commission.
func IsOpenSupportChannel(gs *config.GuildState, channelID string) bool {
	gs.Lock()
	defer gs.Unlock()
	if _, ok := gs.TicketRuntime.OpenTickets[channelID]; ok {
		return true
	}
	if _, ok := gs.CommissionsRuntime.OpenCommissions[channelID]; ok {
		return true
	}
	return false
}
