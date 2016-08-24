package main

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

func (client *DbController) NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.NotFoundHandler = http.HandlerFunc(MethodNotFound)
	routes := client.GetRoutes()
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = context.ClearHandler(handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
