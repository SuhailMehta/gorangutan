package main

import (
	// "github.com/garyburd/redigo/redis"
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
		},
	}

	return routes
}
