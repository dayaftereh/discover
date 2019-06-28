package router

import "github.com/dayaftereh/discover/server/api"

type localRoute struct {
	method  string
	path    string
	handler api.Function
}

func (route localRoute) Method() string {
	return route.method
}

func (route localRoute) Path() string {
	return route.path
}

func (route localRoute) Handler() api.Function {
	return route.handler
}
