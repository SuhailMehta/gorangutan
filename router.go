package main

import (
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"net/http"
)

func (client *DbController) NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
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
	r := OGRouter{router: router}
	router = r.router
	router.NotFoundHandler = http.HandlerFunc(r.MethodNotAllowed)
	return router
}

/* Handled mux router drawback of giving 404 in every case i.e if either URI or METHOD
*  is not matched mux router will give 404. So, this case is handled her, if Method will
*  not match we will through 405
 */
func (router *OGRouter) MethodNotAllowed(rw http.ResponseWriter, req *http.Request) {

	r := *router.router
	route := r.Get(req.RequestURI)

	if route != nil {
		if req.RequestURI == route.GetName() {

			routeMatch := mux.RouteMatch{Route: route}

			if !route.Match(req, &routeMatch) {
				rw.WriteHeader(http.StatusMethodNotAllowed)
			}
		} else {
			rw.WriteHeader(http.StatusNotFound)
		}
	} else {
		rw.WriteHeader(http.StatusNotFound)
	}
}

type OGRouter struct {
	router *mux.Router
}
