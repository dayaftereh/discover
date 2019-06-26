package backend

import (
	"github.com/dayaftereh/discover/server/http"
	"github.com/dayaftereh/discover/server/http/session"
	"github.com/gorilla/mux"
)

type Http struct {
	Server         *http.Server
	SessionManager *session.Manager
}

type Backend struct {
	HTTP *Http
}

func NewBackend(router *mux.Router) (*Backend, error) {

	sessionManager, err := session.NewSessionManager()

	backend := &Backend{
		HTTP: &Http{
			Server:         http.NewServer(router),
			SessionManager: sessionManager,
		},
	}

}
