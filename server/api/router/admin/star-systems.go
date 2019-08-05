package admin

import (
	"context"
	"net/http"

	"github.com/dayaftereh/discover/server/api"
)

func (admin *adminRouter) allStarSystems(ctx context.Context, response http.ResponseWriter, request *http.Request, variables map[string]string) error {
	// check if the user is admin
	isAdmin, err := admin.IsAdmin(ctx)
	if err != nil {
		return err
	}

	// response Forbidden because not a admin
	if !isAdmin {
		return api.Forbidden(response)
	}

	// get all star systems
	starSystems := admin.backend.GetStarSystems()

	// response all star systems
	return api.WriteJSON(response, http.StatusOK, starSystems)
}
