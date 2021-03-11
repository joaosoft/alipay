package models

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/paymentmethod"
)

func CreateCustomer(name string) (*stripe.Customer, error) {
	params := &stripe.CustomerParams{
		Name: stripe.String(name),
	}

	cus, err := customer.New(params)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n:: customer {id: %s}", cus.ID)

	return cus, nil
}

func GetCustomerPaymentMethods(cus *stripe.Customer) error {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(cus.ID),
		Type:     stripe.String(aliPayType),
	}

	pmList := paymentmethod.List(params)

	if pmList.Err() != nil {
		return pmList.Err()
	}

	fmt.Print(":: payment methods:")
	for {
		fmt.Printf("\n:: payment method: %+v", pmList.Current())
		if !pmList.Next() {
			break
		}
	}

	return nil
}

func AddCustomerPaymentMethod(cus *stripe.Customer, paymentMethodID string) error {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String(cus.ID),
	}

	pm, err := paymentmethod.Attach(paymentMethodID, params)

	if err != nil {
		return err
	}

	fmt.Printf("\n:: payment method: %+v", pm.ID)

	return nil
}