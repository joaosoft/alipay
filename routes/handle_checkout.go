package routes

import (
	"net/http"
)

func handleCheckout(publicApiKey string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		data := struct {
			PublicApiKey string
		}{
			PublicApiKey: publicApiKey,
		}

		if err := checkoutTmpl.Execute(w, data); err != nil {
			handleError(err, w)
			return
		}
	}
}
