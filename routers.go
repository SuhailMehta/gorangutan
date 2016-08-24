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
			"/androidPush/", "POST", "/androidPush/", client.AndroidPushNotification,
		}, {
			"/register/", "POST", "/register/", client.RegisterDevice,
		},
	}

	return routes
}
