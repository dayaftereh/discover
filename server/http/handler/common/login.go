package common

import (
	"net/http"

	"github.com/dayaftereh/discover/server/backend"
)

func LoginHandler(backend *backend.Backend) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {}
}
