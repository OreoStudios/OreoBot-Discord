package payments

import "github.com/el211/GoModulith/module"

// PackageInfo describes payments as an application module.
type Descriptor struct{}

var _ module.Provider = Descriptor{}

func (Descriptor) PackageInfo() module.Info {
	return module.Info{
		Name:        "payments",
		ImportPath:  "discord-bot/modules/payments",
		Description: "Payment orchestration, provider gateways, invoice polling, and webhooks",
		Version:     "0.1.0",
		DependsOn:   []string{"config", "persistence"},
		Exports:     []string{"Start", "PaymentButton", "api.PayPalGateway", "api.StripeGateway", "api.CoinbaseGateway", "api.CoinbaseCDPGateway", "infrastructure/providers"},
	}
}
