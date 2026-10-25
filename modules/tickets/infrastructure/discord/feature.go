package discord

import "github.com/bwmarrin/discordgo"

// Feature is the Discord adapter for ticket commands and components.
type Feature struct{}

func NewFeature() *Feature { return &Feature{} }

func (*Feature) Commands() []*discordgo.ApplicationCommand { return ticketCommands() }

func (*Feature) HandleCommand(name string, s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	switch name {
	case "ticket":
		handleTicketCommand(s, i)
	case "close":
		handleCloseCommand(s, i)
	case "add":
		handleAddUser(s, i)
	case "remove":
		handleRemoveUser(s, i)
	default:
		return false
	}
	return true
}

func (*Feature) HandleComponent(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	switch i.MessageComponentData().CustomID {
	case "ticket_category_select":
		handleTicketCategorySelect(s, i)
	case "ticket_subcategory_select":
		handleTicketSubcategorySelect(s, i)
	case "ticket_close_btn":
		handleCloseButton(s, i)
	case "ticket_close_confirm":
		handleCloseConfirm(s, i)
	case "ticket_close_cancel":
		handleCloseCancel(s, i)
	default:
		return false
	}
	return true
}
