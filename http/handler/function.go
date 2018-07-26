package handler

import (
	"net/http"
)

type HandlerFunc func(context *Context) error

func bind(builder *Builder, path string, handlerFunc HandlerFunc) *mux.Router {

	// register the handler func
	builder.Router.HandleFunc(path, func(response http.ResponseWriter, request *http.Request) {

		context := &HandlerContext{
			Server:   builder.Server,
			Request:  request,
			Response: response,
		}

		// check for session- manager
		if builder.SessionManager != nil {
			// get the session
			session, err := builder.SessionManager.Get(response, request)
			if err != nil {
				handleError(response, err)
				return
			}
			// set the session to context
			context.Session = session
		}

		// execute the handler function
		err := handlerFunc(context)
		if err != nil {
			handleError(response, err)
		}
	})
}

func handleError(response http.ResponseWriter, err error) {
	http.Error(response, e.Error(), 501)
}
