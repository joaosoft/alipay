package main

import (
	"github.com/joaosoft/alipay"
)

func main() {
	ap, err := alipay.NewAliPay()
	if err != nil {
		panic(err)
	}

	if err = ap.Start(); err != nil {
		panic(err)
	}
}
