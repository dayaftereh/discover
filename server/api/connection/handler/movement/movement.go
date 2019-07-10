package movement

import (
	"encoding/json"

	"github.com/dayaftereh/discover/server/api/connection"
	"github.com/dayaftereh/discover/server/api/connection/dispatch/handler"
	types "github.com/dayaftereh/discover/server/api/types/connection"
	"github.com/pkg/errors"
)

type movementHandler struct {
	backend Backend
}

func NewHandler(backend Backend) handler.Handler {
	return &movementHandler{
		backend: backend,
	}
}

func (handler *movementHandler) MessageType() types.MessageType {
	return types.Move
}

func (handler *movementHandler) Handler() connection.Function {
	return func(connection *connection.Connection, content string) error {
		// convert to bytes
		data := []byte(content)

		// Unmarshal movement message
		var movement types.Movement
		err := json.Unmarshal(data, &movement)
		if err != nil {
			return errors.Wrapf(err, "fail to unmarshal movement message")
		}

		// emit via backend
		handler.backend.Movement(connection.Player, movement.Move, movement.Rotation)

		return nil
	}
}
