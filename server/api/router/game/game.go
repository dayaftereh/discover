package game

import (
	"github.com/dayaftereh/discover/server/api/server/router"
)

type commonGame struct {
	backend Backend
	routes  []router.Route
}

func NewRouter(backend Backend) router.Router {
	router := &commonGame{
		backend: backend,
	}

	router.initRoutes()

	return router
}

func (game *commonGame) Routes() []router.Route {
	return game.routes
}

func (game *commonGame) Close() {

}

func (game *commonGame) initRoutes() {
	game.routes = []router.Route{
		// POST
		router.NewPostRoute("/game/ready", game.ready),
	}
}
