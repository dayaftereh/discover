package middleware

import "github.com/dayaftereh/discover/server/api"

type Middleware interface {
	WrapHandler(handler api.Function) api.Function
}
