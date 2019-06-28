package router

import "github.com/dayaftereh/discover/server/api"

// Route defines an individual API route in the server.
type Route interface {
	// Handler returns the raw function to create the http handler.
	Handler() api.Function
	// Method returns the http method that the route responds to.
	Method() string
	// Path returns the subpath where the route responds to.
	Path() string
}
