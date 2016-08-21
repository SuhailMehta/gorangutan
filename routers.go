package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type routes []Route

func (client *DbController) GetRoutes() []Route {

	var routes = routes{
		Route{
			"android",
			"GET",
			"/android",
			client.AndroidGCM,
		}, {
			"Registration",
			"POST",
			"/register",
			client.RegisterDevice,
		},
	}

	return routes
}
