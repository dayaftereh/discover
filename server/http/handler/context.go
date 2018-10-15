package handler

import (
	server "../http"
	"../session"
	"net/http"
)

type HandlerContext struct {
	Request  *http.Request
	Response http.ResponseWriter
	Server   *server.HttpServer
	Session  *session.HttpSession
}
