package http

import (
	"github.com/gorilla/mux"
)

type RouterFactory func(server *Server) (*mux.Router, error)
