package server

import (
	"github.com/dayaftereh/discover/server/api/server/router"
	"github.com/gorilla/mux"
)

// UseRouter registered given routers to the server
func (server *Server) UseRouter(routers ...router.Router) {
	server.routers = append(server.routers, routers...)
}

func (server *Server) createMux() *mux.Router {
	// create the router mux
	m := mux.NewRouter()

	// create a sub router for the api
	apiRouter := m.PathPrefix("api")

	// get all registered routers
	for _, router := range server.routers {
		// get the routes from the router
		for _, route := range router.Routes() {
			// make the route handler to a handlerFunc
			handlerFunc := server.makeHTTPHandler(route.Handler())
			// register the handler for the given route
			apiRouter.Path(route.Path()).Methods(route.Method()).Handler(handlerFunc)
		}
	}

	return m
}
