package common

import (
	"github.com/dayaftereh/discover/server/api/server/router"
)

type commonRouter struct {
	backend Backend
	routes  []router.Route
}

func NewRouter(backend Backend) router.Router {
	router := &commonRouter{
		backend: backend,
	}

	router.initRoutes()

	return router
}

func (common *commonRouter) Routes() []router.Route {
	return common.routes
}

func (common *commonRouter) Close() {

}

func (common *commonRouter) initRoutes() {
	common.routes = []router.Route{
		//GET
		router.NewGetRoute("/common/status", common.status),
		// POST
		router.NewPostRoute("/common/login", common.login),
		router.NewPostRoute("/common/logout", common.logout),
	}
}
