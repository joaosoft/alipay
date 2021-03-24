package models

type (
	Customer struct {
		Name    *string
		Email   *string
		Address *Address
	}

	Address struct {
		City       *string `form:"city"`
		Country    *string `form:"country"`
		Line1      *string `form:"line1"`
		Line2      *string `form:"line2"`
		PostalCode *string `form:"postal_code"`
		State      *string `form:"state"`
	}
)
