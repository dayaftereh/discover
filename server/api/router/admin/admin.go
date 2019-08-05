package admin

import (
	"context"

	"github.com/dayaftereh/discover/server/api"
	"github.com/dayaftereh/discover/server/api/server/router"
)

type adminRouter struct {
	backend Backend
	routes  []router.Route
}

func NewRouter(backend Backend) router.Router {
	router := &adminRouter{
		backend: backend,
	}

	router.initRoutes()

	return router
}

func (admin *adminRouter) Routes() []router.Route {
	return admin.routes
}

func (admin *adminRouter) Close() {

}

func (admin *adminRouter) initRoutes() {
	admin.routes = []router.Route{
		//GET
		router.NewGetRoute("/admin/star-systems", admin.allStarSystems),
		router.NewGetRoute("/admin/star-system/{name}", admin.starSystem),
	}
}

func (admin *adminRouter) IsAdmin(ctx context.Context) (bool, error) {
	// get the session id
	sessionID, err := api.SessionIdFromContext(ctx)
	if err != nil {
		return false, err
	}

	// check if the player exists
	player := admin.backend.GetPlayerSession(sessionID)
	// check if a player was found
	if player == nil {
		return false, nil
	}

	return player.Admin, nil
}
