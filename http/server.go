package http

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type HttpServer struct {
	Server        *http.Server
	Config        *Config
	Errors        chan error
	RouterFactory *RouterFactory
}

func NewHttpServer(config *Config, routerFactory *RouterFactory) *HttpServer {
	return &HttpServer{
		Config:        config,
		RouterFactory: routerFactory,
	}
}

func (server *HttpServer) Init() error {
	// check if config is given
	if server.Config == nil {
		server.Config = DefaultConfig()
	}

	// create the root-router
	router := server.createRouter()
	// build bind address
	address := fmt.Sprintf(":%d", server.Config.Port)

	// create the http server
	server.Server = &http.Server{
		Addr:         address,
		Handler:      router,
		WriteTimeout: server.Config.WriteTimeout,
		ReadTimeout:  server.Config.ReadTimeout,
		IdleTimeout:  server.Config.IdleTimeout,
	}

	return nil
}

func (server *HttpServer) createRouter() (error, *Router) {
	if server.RouterFactory == nil {
		return mux.NewRouter(), nil
	}
	router, err := server.RouterFactory(server)
	return router, err
}

func (server *HttpServer) Serve() {
	if server.Server == nil {
		server.Init()
	}

	go func() {
		log.Printf("Listening on %v", server.Server.Addr)
		err := server.Server.ListenAndServe()
		if err != nil {
			server.Errors <- err
		}
	}()
}

func (server *HttpServer) Destroy() error {
	if server.Server == nil {
		return nil
	}

	var wait time.Duration = time.Second * 5
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	err := server.Server.Shutdown(ctx)
	return err
}
