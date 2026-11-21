package discord

import (
	"fmt"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/bwmarrin/discordgo"
)

func testCategory(id, name string) *discordgo.Channel {
	return &discordgo.Channel{ID: id, Name: name, Type: discordgo.ChannelTypeGuildCategory}
}

func testChildren(parent string, count int) []*discordgo.Channel {
	children := make([]*discordgo.Channel, 0, count)
	for n := 0; n < count; n++ {
		children = append(children, &discordgo.Channel{
			ID:       fmt.Sprintf("%s-%d", parent, n),
			Type:     discordgo.ChannelTypeGuildText,
			ParentID: parent,
		})
	}
	return children
}

func TestSelectTicketParentPrefersPrimary(t *testing.T) {
	primary := testCategory("primary", "Support Tickets")
	overflow := testCategory("overflow", "Support Tickets - Tickets 2")
	channels := append([]*discordgo.Channel{primary, overflow}, testChildren("primary", 49)...)

	gotPrimary, gotTarget, err := selectTicketParent(channels, "primary", []string{"overflow"})
	if err != nil || gotPrimary != primary || gotTarget != primary {
		t.Fatalf("expected available primary; primary=%v target=%v err=%v", gotPrimary, gotTarget, err)
	}
}

func TestSelectTicketParentUsesOverflowWhenPrimaryFull(t *testing.T) {
	primary := testCategory("primary", "Support Tickets")
	overflow := testCategory("overflow", "Support Tickets - Tickets 2")
	channels := append([]*discordgo.Channel{primary, overflow}, testChildren("primary", 50)...)
	channels = append(channels, testChildren("overflow", 49)...)

	_, target, err := selectTicketParent(channels, "primary", []string{"overflow"})
	if err != nil || target != overflow {
		t.Fatalf("expected overflow category; target=%v err=%v", target, err)
	}
}

func TestSelectTicketParentRequestsNewOverflowWhenAllFull(t *testing.T) {
	primary := testCategory("primary", "Support Tickets")
	overflow := testCategory("overflow", "Support Tickets - Tickets 2")
	channels := append([]*discordgo.Channel{primary, overflow}, testChildren("primary", 50)...)
	channels = append(channels, testChildren("overflow", 50)...)

	gotPrimary, target, err := selectTicketParent(channels, "primary", []string{"overflow"})
	if err != nil || gotPrimary != primary || target != nil {
		t.Fatalf("expected new overflow category; primary=%v target=%v err=%v", gotPrimary, target, err)
	}
}

func TestSelectTicketParentNeverReusesUntrackedCategory(t *testing.T) {
	primary := testCategory("primary", "Support Tickets")
	unrelated := testCategory("other", "Support Tickets - Tickets 2")
	channels := append([]*discordgo.Channel{primary, unrelated}, testChildren("primary", 50)...)

	_, target, err := selectTicketParent(channels, "primary", nil)
	if err != nil || target != nil {
		t.Fatalf("an unrelated category must not be used; target=%v err=%v", target, err)
	}
}

func TestSelectTicketParentIgnoresDeletedOverflow(t *testing.T) {
	primary := testCategory("primary", "Support Tickets")
	channels := append([]*discordgo.Channel{primary}, testChildren("primary", 50)...)

	_, target, err := selectTicketParent(channels, "primary", []string{"deleted"})
	if err != nil || target != nil {
		t.Fatalf("deleted overflow needs recreation; target=%v err=%v", target, err)
	}
}

func TestSelectTicketParentRejectsInvalidPrimary(t *testing.T) {
	for _, channels := range [][]*discordgo.Channel{
		nil,
		{{ID: "primary", Type: discordgo.ChannelTypeGuildText}},
	} {
		_, _, err := selectTicketParent(channels, "primary", nil)
		if err == nil || !strings.Contains(err.Error(), "/ticket setup") {
			t.Fatalf("expected actionable invalid category error, got %v", err)
		}
	}
}

func TestSelectTicketParentCountsAllChildChannelTypes(t *testing.T) {
	primary := testCategory("primary", "Support Tickets")
	channels := append([]*discordgo.Channel{primary}, testChildren("primary", 49)...)
	channels = append(channels, &discordgo.Channel{
		ID: "voice", Type: discordgo.ChannelTypeGuildVoice, ParentID: "primary",
	})

	_, target, err := selectTicketParent(channels, "primary", nil)
	if err != nil || target != nil {
		t.Fatalf("voice channels also count toward the category limit; target=%v err=%v", target, err)
	}
}

func TestTicketOverflowCategoryNameLimit(t *testing.T) {
	got := ticketOverflowCategoryName(strings.Repeat("🎫", 110), 12)
	if utf8.RuneCountInString(got) > 100 {
		t.Fatalf("Discord category name exceeds 100 characters: %d", utf8.RuneCountInString(got))
	}
	if !strings.HasSuffix(got, " - Tickets 12") {
		t.Fatalf("unexpected category name: %q", got)
	}
}
