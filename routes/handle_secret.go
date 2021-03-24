package routes

import (
	"alipay/models"
	"encoding/json"
	"github.com/stripe/stripe-go"
	"net/http"
	"strconv"
)

func handleSecret(publicApiKey, secretApiKey, returnUrl string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		client := r.URL.Query().Get("client")
		amount, _ := strconv.ParseInt(r.URL.Query().Get("amount"), 10, 64)
		confirm, _ := strconv.ParseBool(r.URL.Query().Get("confirm"))
		cancel, _ := strconv.ParseBool(r.URL.Query().Get("cancel"))

		var err error
		var cus *stripe.Customer

		if client != "" {
			customer := &models.Customer{
				Name:  stripe.String(client),
				Email: stripe.String("joaosoft@gmail.com"),
				Address: &models.Address{
					City:       stripe.String("Porto"),
					Country:    stripe.String("Portugal"),
					Line1:      stripe.String("Rua dos testes"),
					Line2:      stripe.String("Número 69"),
					PostalCode: stripe.String("4400-201"),
					State:      stripe.String("Portugal"),
				},
			}
			if cus, err = models.CreateCustomer(secretApiKey, customer); err != nil {
				handleError(err, w)
				return
			}
		}

		var pi *stripe.PaymentIntent
		if pi, err = models.CreatePaymentIntent(secretApiKey, cus, amount); err != nil {
			handleError(err, w)
			return
		}

		var data interface{}
		if confirm {
			if pi, err = models.ConfirmPaymentIntent(publicApiKey, returnUrl, pi); err != nil {
				handleError(err, w)
				return
			}

			data = struct {
				Url string `json:"url"`
			}{
				Url: pi.NextAction.AlipayHandleRedirect.URL,
			}
		} else {
			data = struct {
				ClientSecret string `json:"client_secret"`
			}{
				ClientSecret: pi.ClientSecret,
			}
		}

		if cancel {
			if pi, err = models.CancelPaymentIntent(secretApiKey, pi.ID, models.CanceledReasonRequestedByCustomer); err != nil {
				handleError(err, w)
				return
			}
		}

		if err = json.NewEncoder(w).Encode(data); err != nil {
			handleError(err, w)
			return
		}
	}
}
