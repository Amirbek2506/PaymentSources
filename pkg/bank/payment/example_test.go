package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources() {
	cards := []types.Card{
		{
			Balance: 20_00,
			Active:  true,
			PAN:     "4444 4545 2525 0101",
		},
		{
			Balance: 250_00,
			Active:  true,
			PAN:     "8085 2702 8080 8544",
		},
		{
			Balance: -5_00,
			Active:  true,
			PAN:     "8085 2702 8080 8544",
		},
		{
			Balance: 100_00,
			Active:  false,
			PAN:     "8085 2702 8080 8544",
		},
	}
	for _, paymentSource := range  PaymentSources(cards){
		fmt.Println(paymentSource.Number)
	}
	// Output: 
	//4444 4545 2525 0101
	//8085 2702 8080 8544
}
