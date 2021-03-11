package routes

import (
	"encoding/json"
	"net/http"
)

func handleError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	if err = json.NewEncoder(w).Encode(Error{
		Error: err,
	}); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}