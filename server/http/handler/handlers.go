package handler

import (
	"github.com/dayaftereh/discover/server/backend"
	"github.com/dayaftereh/discover/server/http/handler/common"
	"github.com/gorilla/mux"
)

func CreateMuxRouter(backend *backend.Backend) *mux.Router {
	m := mux.NewRouter()

	// common
	m.HandleFunc("/login", common.LoginHandler(backend)).Methods("POST")
	m.HandleFunc("/status", common.StatusHandler(backend)).Methods("GET")
	m.HandleFunc("/logout", common.LogoutHandler(backend)).Methods("POST")

	return m
}
