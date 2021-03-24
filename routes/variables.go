package routes

import "html/template"

var (
	checkoutView = template.Must(template.ParseFiles("views/checkout.html"))
	successView  = template.Must(template.ParseFiles("views/success.html"))
)
