package dispatch

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/dayaftereh/discover/server/api/connection"
	"github.com/dayaftereh/discover/server/api/connection/dispatch/handler"
	types "github.com/dayaftereh/discover/server/api/types/connection"
	"github.com/pkg/errors"
)

type Dispatcher struct {
	lock        sync.RWMutex
	handlers    map[types.MessageType]handler.Handler
	connections map[string]*connection.Connection
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		handlers:    make(map[types.MessageType]handler.Handler),
		connections: make(map[string]*connection.Connection),
	}
}

func (dispatcher *Dispatcher) EmitOpen(connection *connection.Connection) {
	// lock for adding connection
	dispatcher.lock.Lock()
	defer dispatcher.lock.Unlock()

	// store the connection
	dispatcher.connections[connection.ID] = connection

	// start dispatch loop for connection
	go dispatcher.dispatchLoop(connection)
}

func (dispatcher *Dispatcher) dispatchLoop(connection *connection.Connection) {
	defer dispatcher.drop(connection)
	for {
		select {
		case message := <-connection.InBoundMessage:
			// handle received message
			dispatcher.handle(connection, message)
		case <-connection.OnClose:
			return
		}
	}
}

func (dispatcher *Dispatcher) drop(connection *connection.Connection) {
	// lock for removeing of the connection
	dispatcher.lock.Lock()
	defer dispatcher.lock.Unlock()
	// remove the connection
	delete(dispatcher.connections, connection.ID)

	// drop the connection from player
	connection.Player.DropConnection(connection)
}

func (dispatcher *Dispatcher) handle(connection *connection.Connection, message string) {
	err := dispatcher.dispatch(connection, message)
	if err != nil {
		log.Println(err)
	}
}

func (dispatcher *Dispatcher) dispatch(connection *connection.Connection, content string) error {
	// convert content to bytes
	data := []byte(content)
	// decode for default message
	var message types.Message
	// unmarshal message
	err := json.Unmarshal(data, &message)
	if err != nil {
		return errors.Wrapf(err, "fail to unmarshal received message")
	}

	if message.Type == nil {
		return errors.Wrapf(err, "received message has now message type")
	}

	// execute the handler for the received message
	err = dispatcher.execute(connection, *message.Type, content)
	return errors.Wrapf(err, "executed message handler returnes with an error")
}

func (dispatcher *Dispatcher) Close() {
	dispatcher.lock.RLock()
	defer dispatcher.lock.RUnlock()

	for _, connection := range dispatcher.connections {
		connection.Close()
	}
}
