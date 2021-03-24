package routes

import (
	"net/http"
)

func InitializeRouters(publicApiKey, secretApiKey, returnUrl string) {
	http.HandleFunc("/secret", handleSecret(publicApiKey, secretApiKey, returnUrl))
	http.HandleFunc("/checkout", handleCheckout(publicApiKey))
	http.HandleFunc("/success", handleSuccess)
	http.HandleFunc("/event", handleEvent)
	http.Handle("/", http.FileServer(http.Dir("./static")))
}
