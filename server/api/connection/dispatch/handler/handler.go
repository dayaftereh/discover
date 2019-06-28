package handler

import (
	"github.com/dayaftereh/discover/server/api/connection"
	types "github.com/dayaftereh/discover/server/api/types/connection"
)

type Handler interface {
	MessageType() types.MessageType
	Handler() connection.Function
}
