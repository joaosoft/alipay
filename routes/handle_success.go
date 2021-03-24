package routes

import (
	"net/http"
)

func handleSuccess(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		if err := successView.Execute(w, nil); err != nil {
			handleError(err, w)
			return
		}
}
