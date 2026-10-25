// Package api defines the provider contracts consumed by payment orchestration.
package api

import "discord-bot/modules/config"

// PayPalGateway creates, cancels, and checks PayPal invoices.
type PayPalGateway interface {
	CreateInvoice(*config.CommissionInvoice) (string, string, error)
	CancelInvoice(invoiceID string) error
	GetInvoiceStatus(invoiceID string) (string, error)
}

// StripeGateway creates, expires, and checks Stripe checkout sessions.
type StripeGateway interface {
	CreateCheckoutSession(*config.CommissionInvoice) (string, string, error)
	ExpireSession(sessionID string) error
	GetSessionStatus(sessionID string) (string, error)
}

// CoinbaseGateway manages Coinbase Commerce charges.
type CoinbaseGateway interface {
	CreateCharge(*config.CommissionInvoice) (string, string, error)
	GetChargeStatus(chargeID string) (string, error)
}

// CoinbaseCDPGateway manages receive addresses and checks incoming transfers.
type CoinbaseCDPGateway interface {
	GetSpotPrice(asset, fiat string) (float64, error)
	GetAccountID(asset string) (string, error)
	CreateAddress(accountID, name, network string) (address, addressID string, err error)
	AddressReceived(accountID, addressID string) (float64, error)
}
