package payment

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	var paymentSources []types.PaymentSource
	for _, card := range cards {
		if card.Active && card.Balance > 0 {
			paymentSources = append(paymentSources, types.PaymentSource{
				Type:    "card",
				Number:  card.PAN,
				Balance: card.Balance,
			})
		}
	}
	return paymentSources
}
