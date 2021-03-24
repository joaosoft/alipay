package models

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
	"github.com/stripe/stripe-go/paymentmethod"
	"net/url"
)

func GetPaymentMethod(paymentMethodID string) error {
	pm, err := paymentmethod.Get(paymentMethodID, nil)
	if err != nil {
		return err
	}

	fmt.Printf("\n:: payment method: {id: %s}", pm.Type)

	return nil
}

func CreatePaymentIntent(secretApiKey string, cus *stripe.Customer, amount int64) (*stripe.PaymentIntent, error) {
	stripe.Key = secretApiKey

	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(amount * 100),
		Currency:      stripe.String(string(stripe.CurrencyEUR)),
		CaptureMethod: stripe.String(CaptureMethodAutomatic),
		PaymentMethodTypes: []*string{
			stripe.String(AliPayType),
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

func CancelPaymentIntent(secretApiKey, id, cancelReason string) (*stripe.PaymentIntent, error) {
	stripe.Key = secretApiKey

	params := &stripe.PaymentIntentCancelParams{
		CancellationReason: stripe.String(cancelReason),
	}

	pi, err := paymentintent.Cancel(id, params)
	if err != nil {
		return nil, err
	}

	return pi, nil
}

func ConfirmPaymentIntent(publicApiKey, returnUrl string, pi *stripe.PaymentIntent) (*stripe.PaymentIntent, error) {
	stripe.Key = publicApiKey
	params := &stripe.PaymentIntentConfirmParams{
		ReturnURL: stripe.String(returnUrl),
		Params: stripe.Params{
			Extra: &stripe.ExtraValues{
				Values: url.Values{
					"expected_payment_method_type": []string{
						AliPayType,
					},
					"client_secret": []string{pi.ClientSecret},
				},
			},
		},
	}

	pi, err := paymentintent.Confirm(pi.ID, params)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n:: confirmed payment intent: {id: %s, status: %s}", pi.ID, pi.Status)

	return pi, nil
}
