package connection

import (
	"github.com/dayaftereh/discover/server/api/connection/dispatch"
	"github.com/dayaftereh/discover/server/api/server/router"
	"github.com/gorilla/websocket"
)

type connectionRouter struct {
	backend    Backend
	routes     []router.Route
	dispatcher *dispatch.Dispatcher
	upgrader   websocket.Upgrader
}

func NewRouter(backend Backend, dispatcher *dispatch.Dispatcher) router.Router {
	router := &connectionRouter{
		backend:    backend,
		dispatcher: dispatcher,
	}

	router.initRoutes()

	return router
}

func (connection *connectionRouter) Routes() []router.Route {
	return connection.routes
}

func (connection *connectionRouter) Close() {
	connection.dispatcher.Close()
}

func (connection *connectionRouter) initRoutes() {
	connection.routes = []router.Route{
		router.NewGetRoute("/ws", connection.websocket),
	}
}
