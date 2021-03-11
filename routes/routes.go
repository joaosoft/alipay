package routes

import (
	"net/http"
)

func InitializeRouters(publicApiKey string) {
	http.HandleFunc("/secret", handleSecret)
	http.HandleFunc("/checkout", handleCheckout(publicApiKey))
	http.Handle("/", http.FileServer(http.Dir("./static")))
}
