package router

import "github.com/dayaftereh/discover/server/api"

func NewRoute(method string, path string, handler api.Function) Route {
	var route Route = localRoute{method, path, handler}
	return route
}

func NewGetRoute(path string, handler api.Function) Route {
	return NewRoute("GET", path, handler)
}

func NewPostRoute(path string, handler api.Function) Route {
	return NewRoute("POST", path, handler)
}
