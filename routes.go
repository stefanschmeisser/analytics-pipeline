package main

import (
	"go-websocket-server/handlers"
	"net/http"
)

// Route struct defines name fields
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes type defines to be an array of route structs
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		handlers.Index,
	},
	Route{
		"Index",
		"GET",
		"/analyticsstream",
		handlers.GetAnalyticsStream,
	},
}
