package ping

import (
	"encoding/json"

	"github.com/dayaftereh/discover/server/utils"

	"github.com/dayaftereh/discover/server/api/connection"
	"github.com/dayaftereh/discover/server/api/connection/dispatch/handler"
	types "github.com/dayaftereh/discover/server/api/types/connection"
	"github.com/pkg/errors"
)

type pingHandler struct {
	backend Backend
}

func NewHandler(backend Backend) handler.Handler {
	return &pingHandler{
		backend: backend,
	}
}

func (handler *pingHandler) MessageType() types.MessageType {
	return types.Ping
}

func (handler *pingHandler) Handler() connection.Function {
	return func(connection *connection.Connection, content string) error {
		// get the receive time
		receiveTime := utils.SystemMillis()
		// convert to bytes
		data := []byte(content)

		// Unmarshal ping message
		var ping types.PingMessage
		err := json.Unmarshal(data, &ping)
		if err != nil {
			return errors.Wrapf(err, "fail to unmarshal ping message")
		}

		// create the pong message
		pong := &types.PongMessage{
			Type:              types.Pong,
			ClientSendTime:    ping.ClientTime,
			ServerReceiveTime: receiveTime,
			ServerSendTime:    utils.SystemMillis(),
		}

		// convert pong back to json
		bytes, err := json.Marshal(pong)
		if err != nil {
			return errors.Wrapf(err, "fail to marshal pong message")
		}
		// make a string
		message := string(bytes)

		// write pong back
		connection.Write(message)

		return nil
	}
}
