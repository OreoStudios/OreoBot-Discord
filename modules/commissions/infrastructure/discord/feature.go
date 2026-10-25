package discord

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Feature is the Discord adapter for commission requests, quotes, and invoices.
type Feature struct{}

func NewFeature() *Feature { return &Feature{} }

func (*Feature) Commands() []*discordgo.ApplicationCommand { return commissionCommands() }

func (*Feature) HandleCommand(name string, s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	switch name {
	case "commission":
		handleCommissionCommand(s, i)
	case "invoice":
		handleInvoiceCommand(s, i)
	default:
		return false
	}
	return true
}

func (*Feature) HandleAutocomplete(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	if i.ApplicationCommandData().Name != "invoice" {
		return false
	}
	handleInvoiceCurrencyAutocomplete(s, i)
	return true
}

func (*Feature) HandleComponent(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	id := i.MessageComponentData().CustomID
	switch {
	case strings.HasPrefix(id, "commission_invoice_btn:"):
		handleCommissionInvoiceButton(s, i)
	case strings.HasPrefix(id, "invoice_crypto:"):
		handleInvoiceCryptoButton(s, i)
	case strings.HasPrefix(id, "invoice_lang:"):
		handleInvoiceLangButton(s, i)
	case strings.HasPrefix(id, "commission_panel_lang:"):
		handleCommissionPanelLang(s, i)
	case strings.HasPrefix(id, "commission_ticket_lang:"):
		handleCommissionTicketLang(s, i)
	case strings.HasPrefix(id, "commission_quote_btn:"):
		handleCommissionQuoteButton(s, i)
	case strings.HasPrefix(id, "commission_quote_accept:"):
		handleCommissionQuoteAccept(s, i)
	case strings.HasPrefix(id, "commission_quote_decline:"):
		handleCommissionQuoteDecline(s, i)
	case id == "commission_order":
		handleCommissionOrder(s, i)
	case id == "commission_service_select":
		handleCommissionServiceSelect(s, i)
	case id == "commission_close_btn":
		handleCommissionCloseButton(s, i)
	case id == "commission_close_confirm":
		handleCommissionCloseConfirm(s, i)
	case id == "commission_close_cancel":
		handleCommissionCloseCancel(s, i)
	default:
		return false
	}
	return true
}

func (*Feature) HandleModal(s *discordgo.Session, i *discordgo.InteractionCreate) bool {
	id := i.ModalSubmitData().CustomID
	switch {
	case strings.HasPrefix(id, "commission_form:"):
		handleCommissionFormSubmit(s, i)
	case strings.HasPrefix(id, "commission_invoice_modal:"):
		handleCommissionInvoiceModalSubmit(s, i)
	case strings.HasPrefix(id, "commission_quote_modal:"):
		handleCommissionQuoteModalSubmit(s, i)
	case strings.HasPrefix(id, "commission_quote_decline_modal:"):
		handleCommissionQuoteDeclineModalSubmit(s, i)
	default:
		return false
	}
	return true
}

func (*Feature) HandleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	handleCommissionMessageCreate(s, m)
}
