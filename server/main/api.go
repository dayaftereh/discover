package main

import (
	"github.com/dayaftereh/discover/server/api/router/common"
	"github.com/dayaftereh/discover/server/api/server"
	"github.com/dayaftereh/discover/server/api/server/middleware"
	"github.com/dayaftereh/discover/server/api/server/router"
	"github.com/dayaftereh/discover/server/api/session"
	"github.com/dayaftereh/discover/server/game"
	"github.com/pkg/errors"
)

func initAPI(game *game.Game) (*server.Server, error) {
	server := server.NewServer()

	// init middleware for the server
	err := initMiddleware(server)
	if err != nil {
		return nil, err
	}

	// init routers
	err = initRouters(game, server)
	if err != nil {
		return nil, err
	}

	return server, nil
}

func initMiddleware(server *server.Server) error {
	// register the debug request middleware
	debugRequestMiddleware := middleware.NewDebugRequestMiddleware()
	server.UseMiddleware(debugRequestMiddleware)

	// register session middleware
	sessionManager, err := createSessionManager()
	if err != nil {
		return err
	}
	// create the session middleware for the session manager
	sessionMiddleware := middleware.NewSessionMiddleware(sessionManager)
	server.UseMiddleware(sessionMiddleware)

	return nil
}

func createSessionManager() (*session.Manager, error) {
	manager, err := session.NewSessionManager()
	if err != nil {
		return nil, errors.Wrapf(err, "fail to create session-manager")
	}
	return manager, nil
}

func initRouters(game *game.Game, server *server.Server) error {
	// create the routers for the server
	routers := []router.Router{
		// common
		common.NewRouter(game),
	}

	// register the routers
	server.UseRouter(routers...)

	return nil
}
