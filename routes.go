package main

import(
	"net/http"
	"github.com/github-timall/stat/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Ping",
		"GET",
		"/v1/ping",
		handlers.Ping,
	},
	Route{
		"VeinOrderCreate",
		"POST",
		"/v1/vein/order-create",
		handlers.VeinOrderCreate,
	},
	Route{
		"VeinOrderStatus",
		"POST",
		"/v1/vein/order-status",
		handlers.VeinOrderStatus,
	},
	Route{
		"VeinOrderPayment",
		"POST",
		"/v1/vein/order-payment",
		handlers.VeinOrderPayment,
	},
	Route{
		"VeinRedirectCreate",
		"POST",
		"/v1/vein/redirect-create",
		handlers.VeinRedirectCreate,
	},
}