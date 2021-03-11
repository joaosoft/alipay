package routes

import "html/template"

var(
	checkoutTmpl = template.Must(template.ParseFiles("views/checkout.html"))
)
