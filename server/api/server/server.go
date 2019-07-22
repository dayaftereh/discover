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
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  10 * time.Second,
	}

	return nil
}

// Serve start the serving of the server
func (server *Server) Serve() {
	go func() {
		log.Printf("Listening on %v\n", server.Server.Addr)
		err := server.Server.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				log.Printf("server listening thread closed\n")
			} else {
				log.Printf("server listening thread recived an error: %v", err)
			}
		}
	}()
}

// Destroy and shutdown the started server
func (server *Server) Shutdown() error {
	if server.Server == nil {
		return nil
	}

	log.Println("shutdown server...")

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
