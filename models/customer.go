package models

import (
	"fmt"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/paymentmethod"
)

func FindCustomer(secretApiKey, email string) (*stripe.Customer, error) {
	stripe.Key = secretApiKey

	params := &stripe.CustomerListParams{
		Email: stripe.String(email),
	}

	iter := customer.List(params)
	if iter.Next() {
		if err := iter.Err(); err != nil {
			return nil, err
		}
		return iter.Customer(), nil
	}
	return nil, nil
}

func CreateCustomer(secretApiKey string, c *Customer) (*stripe.Customer, error) {
	stripe.Key = secretApiKey

	foundCustomer, err := FindCustomer(secretApiKey, *c.Email)
	if err == nil && foundCustomer != nil {
		var updCustomer *stripe.Customer
		updCustomer, err = UpdateCustomer(secretApiKey, foundCustomer.ID, c)
		if err != nil {
			return nil, err
		}
		return updCustomer, err
	}

	var address *stripe.AddressParams
	if c.Address != nil {
		address = &stripe.AddressParams{
			City:       c.Address.City,
			Country:    c.Address.Country,
			Line1:      c.Address.Line1,
			Line2:      c.Address.Line2,
			PostalCode: c.Address.PostalCode,
			State:      c.Address.State,
		}
	}

	params := &stripe.CustomerParams{
		Name:    c.Name,
		Email:   c.Email,
		Address: address,
	}

	cus, err := customer.New(params)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n:: customer {id: %s}", cus.ID)

	return cus, nil
}

func GetCustomerPaymentMethods(secretApiKey string, cus *stripe.Customer) error {
	stripe.Key = secretApiKey

	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String(cus.ID),
		Type:     stripe.String(AliPayType),
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

func UpdateCustomer(secretApiKey, id string, c *Customer) (*stripe.Customer, error) {
	stripe.Key = secretApiKey
	var address *stripe.AddressParams
	if c.Address != nil {
		address = &stripe.AddressParams{
			City:       c.Address.City,
			Country:    c.Address.Country,
			Line1:      c.Address.Line1,
			Line2:      c.Address.Line2,
			PostalCode: c.Address.PostalCode,
			State:      c.Address.State,
		}
	}

	params := &stripe.CustomerParams{
		Name:    c.Name,
		Email:   c.Email,
		Address: address,
	}

	cus, err := customer.Update(id, params)
	if err != nil {
		return nil, err
	}

	fmt.Printf("\n:: customer {id: %s}", cus.ID)

	return cus, nil
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
