package models

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/paymentmethod"
)

func GetPaymentMethod(paymentMethodID string) error {
	pm, err := paymentmethod.Get(paymentMethodID, nil)
	if err != nil {
		return err
	}

	fmt.Printf("\n:: payment method: {id: %s}", pm.Type)

	return nil
}

func CreatePaymentIntent(cus *stripe.Customer, amount int64) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(amount * 100),
		Currency: stripe.String(string(stripe.CurrencyEUR)),
		PaymentMethodTypes: []*string{
			stripe.String(aliPayType),
		},
	}

	if cus != nil {
		params.Customer = stripe.String(cus.ID)
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n:: payment intent: {id: %s, client_secret: %s}", pi.ID, pi.ClientSecret)

	return pi, nil
}
