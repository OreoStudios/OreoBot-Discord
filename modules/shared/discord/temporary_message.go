// Package discord contains low-level Discord interaction helpers shared by features.
package discord

import (
	"time"

	"github.com/bwmarrin/discordgo"
)

// SendTemp sends a channel message and deletes it after the requested number of seconds.
func SendTemp(s *discordgo.Session, channelID, content string, seconds int) {
	msg, err := s.ChannelMessageSend(channelID, content)
	if err != nil {
		return
	}
	go func() {
		time.Sleep(time.Duration(seconds) * time.Second)
		_ = s.ChannelMessageDelete(channelID, msg.ID)
	}()
}
