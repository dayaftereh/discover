package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Server *http.Server
	Errors chan error
	Router *mux.Router
}

func NewServer(router *mux.Router) *Server {
	return &Server{
		Router: router,
	}
}

func (server *Server) Init() error {
	// build bind address
	address := fmt.Sprintf(":%d", 400)

	// create the http server
	server.Server = &http.Server{
		Addr:         address,
		Handler:      server.Router,
		WriteTimeout: 10,
		ReadTimeout:  10,
		IdleTimeout:  10,
	}

	return nil
}

func (server *Server) Serve() {
	go func() {
		log.Printf("Listening on %v", server.Server.Addr)
		err := server.Server.ListenAndServe()
		if err != nil {
			server.Errors <- err
		}
	}()
}

func (server *Server) Destroy() error {
	if server.Server == nil {
		return nil
	}

	var wait = time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	err := server.Server.Shutdown(ctx)
	return err
}
