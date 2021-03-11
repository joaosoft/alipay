package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func handleEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	bodyBuffer, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(bodyBuffer))

	fmt.Println("handled event!")
}
