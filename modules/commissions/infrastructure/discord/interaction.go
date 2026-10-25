package discord

import (
	"log/slog"

	"github.com/bwmarrin/discordgo"
)

var adminPermission int64 = discordgo.PermissionAdministrator

func respond(s *discordgo.Session, i *discordgo.InteractionCreate, content string, ephemeral bool) {
	flags := discordgo.MessageFlags(0)
	if ephemeral {
		flags = discordgo.MessageFlagsEphemeral
	}
	if err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{Content: content, Flags: flags},
	}); err != nil {
		slog.Error("failed to respond", "error", err)
	}
}

func subOptMap(options []*discordgo.ApplicationCommandInteractionDataOption) map[string]*discordgo.ApplicationCommandInteractionDataOption {
	result := make(map[string]*discordgo.ApplicationCommandInteractionDataOption)
	for _, option := range options {
		result[option.Name] = option
	}
	return result
}

func optStr(options map[string]*discordgo.ApplicationCommandInteractionDataOption, key, fallback string) string {
	if option, ok := options[key]; ok {
		return option.StringValue()
	}
	return fallback
}
