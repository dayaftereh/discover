package server

import (
	"net/http"

	"github.com/dayaftereh/discover/server/api"
	"github.com/gorilla/mux"
)

func (server *Server) makeHTTPHandler(handler api.Function) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// get the context for the request
		ctx := r.Context()

		// wrap the handler with the given middlewares
		handlerFunc := server.handlerWithGlobalMiddlewares(handler)

		// extract the variables
		vars := mux.Vars(r)
		if vars == nil {
			vars = make(map[string]string)
		}

		// execute the handler function
		err := handlerFunc(ctx, w, r, vars)

		// check if the handler function has an error
		if err != nil {
			// response the error
			api.MakeErrorHandler(err)(w, r)
		}
	}
}
