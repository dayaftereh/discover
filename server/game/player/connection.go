package player

import (
	"encoding/json"

	"github.com/pkg/errors"
)

type Connection interface {
	Id() string
	Write(message string)
	Close()
}

func (player *Player) AddConnection(connection Connection) {
	player.connections[connection.Id()] = connection
}

func (player *Player) DropConnection(connection Connection) {
	delete(player.connections, connection.Id())
}

func (player *Player) Push(v interface{}) error {
	// convert given object to jsoin
	bytes, err := json.Marshal(v)
	if err != nil {
		return errors.Wrapf(err, "fail to marshal outbound data")
	}
	// make a string
	message := string(bytes)

	// write to all connections
	for _, connection := range player.connections {
		connection.Write(message)
	}

	return nil
}
