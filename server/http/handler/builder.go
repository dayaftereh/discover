package handler

import (
	"../http"
	"../session"
	"github.com/gorilla/mux"
)

type Builder struct {
	Router         *mux.Router
	Server         *http.Server
	SessionManager *session.SessionManager
}

func NewBuilder(sessionManager *session.SessionManager) *Builder {
	return &Builder{
		SessionManager: sessionManager,
	}
}

func (builder *Builder) ToRouterFactory() *RouterFactory {
	return func(server *http.Server) {
		builder.Server = server
		return builder.router
	}
}

func (builder *Builder) Register(path string, handlerFunc HandlerFunc) *mux.Router {
	router := bind(builder, path, handlerFunc)
	return router
}
