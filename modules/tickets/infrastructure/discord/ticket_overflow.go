package discord

import (
	"fmt"
	"strings"
	"sync"

	"discord-bot/modules/config"

	"github.com/bwmarrin/discordgo"
)

const discordCategoryChannelLimit = 50
const maxTicketOverflowAttempts = 6

// Per-guild locks avoid two ticket requests claiming the last category slot
// (or the same ticket number) in a single bot process.
var ticketCreationLocks sync.Map

func ticketCreationLock(guildID string) *sync.Mutex {
	lock, _ := ticketCreationLocks.LoadOrStore(guildID, &sync.Mutex{})
	return lock.(*sync.Mutex)
}

// ticketOverflowCategoryName stays within Discord's 100-character name limit.
func ticketOverflowCategoryName(primaryName string, number int) string {
	suffix := fmt.Sprintf(" - Tickets %d", number)
	prefix := []rune(strings.TrimSpace(primaryName))
	maxPrefix := 100 - len([]rune(suffix))
	if len(prefix) > maxPrefix {
		prefix = prefix[:maxPrefix]
	}
	return string(prefix) + suffix
}

// selectTicketParent prefers the primary Discord category, then previously
// created overflow categories in order. Only IDs recorded by this bot are
// candidates: an unrelated category with a similar name is never reused.
// Returning a nil target means a new overflow category is required.
func selectTicketParent(channels []*discordgo.Channel, primaryID string, overflowIDs []string) (primary, target *discordgo.Channel, err error) {
	byID := make(map[string]*discordgo.Channel, len(channels))
	childCounts := make(map[string]int)
	for _, channel := range channels {
		if channel == nil {
			continue
		}
		byID[channel.ID] = channel
		if channel.ParentID != "" {
			childCounts[channel.ParentID]++
		}
	}

	primary = byID[primaryID]
	if primary == nil || primary.Type != discordgo.ChannelTypeGuildCategory {
		return nil, nil, fmt.Errorf("configured ticket category %s does not exist or is not a Discord category; run /ticket setup", primaryID)
	}
	if childCounts[primaryID] < discordCategoryChannelLimit {
		return primary, primary, nil
	}
	for _, id := range overflowIDs {
		if candidate := byID[id]; candidate != nil &&
			candidate.Type == discordgo.ChannelTypeGuildCategory &&
			childCounts[id] < discordCategoryChannelLimit {
			return primary, candidate, nil
		}
	}
	return primary, nil, nil
}

// createTicketChannel creates ticket channels, adding new Discord categories
// when the configured parent and its tracked overflow categories are full.
// It is called with the per-guild ticketCreationLock held.
func createTicketChannel(
	s *discordgo.Session,
	guildID, primaryCategoryID, channelName string,
	overwrites []*discordgo.PermissionOverwrite,
	gs *config.GuildState,
) (*discordgo.Channel, error) {
	channelData := discordgo.GuildChannelCreateData{
		Name:                 channelName,
		Type:                 discordgo.ChannelTypeGuildText,
		PermissionOverwrites: overwrites,
	}
	if primaryCategoryID == "" {
		// Maintain support for servers that explicitly leave tickets uncategorized.
		return s.GuildChannelCreateComplex(guildID, channelData)
	}

	for attempt := 0; attempt < maxTicketOverflowAttempts; attempt++ {
		channels, err := s.GuildChannels(guildID)
		if err != nil {
			return nil, fmt.Errorf("cannot load Discord categories: %w", err)
		}

		gs.Lock()
		tracked := append([]string(nil), gs.TicketRuntime.OverflowCategories[primaryCategoryID]...)
		gs.Unlock()
		primary, target, err := selectTicketParent(channels, primaryCategoryID, tracked)
		if err != nil {
			return nil, err
		}

		if target == nil {
			// Retain the original Discord category permissions. Individual tickets
			// still receive their own private owner/staff permission overwrites.
			overflow, createErr := s.GuildChannelCreateComplex(guildID, discordgo.GuildChannelCreateData{
				Name:                 ticketOverflowCategoryName(primary.Name, len(tracked)+2),
				Type:                 discordgo.ChannelTypeGuildCategory,
				PermissionOverwrites: primary.PermissionOverwrites,
			})
			if createErr != nil {
				return nil, fmt.Errorf("cannot create ticket overflow category (Manage Channels permission required): %w", createErr)
			}

			gs.Lock()
			if gs.TicketRuntime.OverflowCategories == nil {
				gs.TicketRuntime.OverflowCategories = make(map[string][]string)
			}
			gs.TicketRuntime.OverflowCategories[primaryCategoryID] = append(gs.TicketRuntime.OverflowCategories[primaryCategoryID], overflow.ID)
			gs.Unlock()
			if saveErr := gs.Save(); saveErr != nil {
				// Don't orphan an untracked overflow category if persistence fails.
				gs.Lock()
				ids := gs.TicketRuntime.OverflowCategories[primaryCategoryID]
				gs.TicketRuntime.OverflowCategories[primaryCategoryID] = ids[:len(ids)-1]
				gs.Unlock()
				_, _ = s.ChannelDelete(overflow.ID)
				return nil, fmt.Errorf("cannot save ticket overflow category: %w", saveErr)
			}
			target = overflow
		}

		channelData.ParentID = target.ID
		channel, createErr := s.GuildChannelCreateComplex(guildID, channelData)
		if createErr == nil {
			return channel, nil
		}
		if !strings.Contains(createErr.Error(), "CHANNEL_PARENT_MAX_CHANNELS") {
			return nil, fmt.Errorf("cannot create ticket channel: %w", createErr)
		}
		// Another writer may have filled the selected category. Reload Discord's
		// current channel list and select the next available parent.
	}

	return nil, fmt.Errorf("ticket categories filled during creation; please retry")
}
