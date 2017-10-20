package main

import "net/http"

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
		Ping,
	},
	Route{
		"VeinOrderCreate",
		"POST",
		"/v1/vein/order-create",
		VeinOrderCreate,
	},
	Route{
		"VeinOrderStatus",
		"POST",
		"/v1/vein/order-status",
		VeinOrderStatus,
	},
	Route{
		"VeinOrderPayment",
		"POST",
		"/v1/vein/order-payment",
		VeinOrderPayment,
	},
	Route{
		"VeinRedirectCreate",
		"POST",
		"/v1/vein/redirect-create",
		VeinRedirectCreate,
	},
	Route{
		"VeinViewCreate",
		"POST",
		"/v1/vein/view-create",
		VeinViewCreate,
	},
}