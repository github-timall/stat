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
		"/v1/vein/redirect-create",
		VeinTrackingCreate,
	},
}