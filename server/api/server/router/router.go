package router

// Router defines an interface to specify a group of routes to add to the server.
type Router interface {
	// Routes returns the list of routes to add to the server.
	Routes() []Route
}
