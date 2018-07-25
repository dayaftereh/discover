package http

import (
	"github.com/gorilla/mux"
)

type RouterFactory func(*HttpServer) (*mux.Router, error)
