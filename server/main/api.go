package main

import (
	"github.com/dayaftereh/discover/server/api/connection/dispatch"
	"github.com/dayaftereh/discover/server/api/connection/dispatch/handler"
	"github.com/dayaftereh/discover/server/api/connection/handler/movement"
	"github.com/dayaftereh/discover/server/api/router/common"
	"github.com/dayaftereh/discover/server/api/router/connection"
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
	// create the connection dispatcher
	dispatcher := dispatch.NewDispatcher()

	// initialize the dispatcher
	err := initDispatcher(game, dispatcher)
	if err != nil {
		return err
	}

	// create the routers for the server
	routers := []router.Router{
		// common
		common.NewRouter(game),
		// connection
		connection.NewRouter(game, dispatcher),
	}

	// register the routers
	server.UseRouter(routers...)

	return nil
}

func initDispatcher(game *game.Game, dispatcher *dispatch.Dispatcher) error {

	// create the handler for dispatcher
	handlers := []handler.Handler{
		// movement
		movement.NewHandler(game),
	}

	// register handlers
	dispatcher.UseHandlers(handlers...)

	return nil
}
