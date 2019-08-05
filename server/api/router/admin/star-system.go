package admin

import (
	"context"
	"net/http"

	"github.com/dayaftereh/discover/server/api"
)

func (admin *adminRouter) starSystem(ctx context.Context, response http.ResponseWriter, request *http.Request, variables map[string]string) error {
	// check if the user is admin
	isAdmin, err := admin.IsAdmin(ctx)
	if err != nil {
		return err
	}

	// response Forbidden because not a admin
	if !isAdmin {
		return api.Forbidden(response)
	}

	// get the name of the star system
	name, ok := variables["name"]

	// return bad request because name missing
	if !ok {
		return api.BadRequest(response)
	}

	// get the star system for the given name
	starSystem := admin.backend.GetStarSystem(name)
	// check if a star system with the name exists
	if starSystem == nil {
		return api.NotFound(response)
	}

	// response all star systems
	return api.WriteJSON(response, http.StatusOK, starSystem.Data)
}
