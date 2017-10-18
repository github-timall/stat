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
		"VeinOrderCreate",
		"POST",
		"/v1/vein/order-create",
		VeinOrderCreate,
	},
	Route{
		"LeadIndex",
		"GET",
		"/lead",
		LeadIndex,
	},
	Route{
		"LeadView",
		"GET",
		"/lead/{leadId}",
		TodoView,
	},
}