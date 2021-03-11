package routes

import (
	"alipay/models"
	"encoding/json"
	"github.com/stripe/stripe-go"
	"net/http"
	"strconv"
)

func handleSecret(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	client := r.URL.Query().Get("client")
	amount, _ := strconv.ParseInt(r.URL.Query().Get("amount"), 10, 64)

	var err error
	var cus *stripe.Customer

	if client != "" {
		if cus, err = models.CreateCustomer(client); err != nil {
			handleError(err, w)
			return
		}
	}

	var pi *stripe.PaymentIntent
	if pi, err = models.CreatePaymentIntent(cus, amount); err != nil {
		handleError(err, w)
		return
	}

	data := struct {
		ClientSecret string `json:"client_secret"`
	}{
		ClientSecret: pi.ClientSecret,
	}

	if err = json.NewEncoder(w).Encode(data); err != nil {
		handleError(err, w)
		return
	}
}
