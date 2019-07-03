package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dayaftereh/discover/server/api/server/middleware"
	"github.com/dayaftereh/discover/server/api/server/router"
)

type Server struct {
	Server *http.Server
	Errors chan error

	routers     []router.Router
	middlewares []middleware.Middleware
}

// NewServer creates a new Server
func NewServer() *Server {
	return &Server{}
}

// Init initlize the server with the given configuration
func (server *Server) Init() error {
	// build bind address
	address := fmt.Sprintf(":%d", 4000)

	router := server.createMux()

	// create the http server
	server.Server = &http.Server{
		Addr:         address,
		Handler:      router,
		WriteTimeout: 10,
		ReadTimeout:  10,
		IdleTimeout:  10,
	}

	return nil
}

// Serve start the serving of the server
func (server *Server) Serve() {
	go func() {
		log.Printf("Listening on %v", server.Server.Addr)
		err := server.Server.ListenAndServe()
		if err != nil {
			server.Errors <- err
		}
	}()
}

// Destroy and shutdown the started server
func (server *Server) Shutdown() error {
	if server.Server == nil {
		return nil
	}

	// close all routers
	for _, router := range server.routers {
		router.Close()
	}

	var wait = time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	err := server.Server.Shutdown(ctx)
	return err
}
