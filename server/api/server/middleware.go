package server

import (
	"github.com/dayaftereh/discover/server/api"
	"github.com/dayaftereh/discover/server/api/server/middleware"
)

// UseMiddleware registered a middelware wrapper function for all registered server routes
func (server *Server) UseMiddleware(m middleware.Middleware) {
	server.middlewares = append(server.middlewares, m)
}

func (server *Server) handlerWithGlobalMiddlewares(handler api.Function) api.Function {
	next := handler

	for _, m := range server.middlewares {
		next = m.WrapHandler(next)
	}

	return next
}
