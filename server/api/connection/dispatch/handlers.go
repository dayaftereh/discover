package dispatch

import (
	"github.com/dayaftereh/discover/server/api/connection"
	"github.com/dayaftereh/discover/server/api/connection/dispatch/handler"
	types "github.com/dayaftereh/discover/server/api/types/connection"
	"github.com/pkg/errors"
)

func (dispatcher *Dispatcher) UseHandler(handler handler.Handler) {
	dispatcher.handlers[handler.MessageType()] = handler
}

func (dispatcher *Dispatcher) UseHandlers(handlers ...handler.Handler) {
	for _, handler := range handlers {
		dispatcher.UseHandler(handler)
	}
}

func (dispatcher *Dispatcher) execute(connection *connection.Connection, messageType types.MessageType, message string) error {
	handler, ok := dispatcher.handlers[messageType]
	if !ok {
		return errors.Errorf("unable to dispatch message with type [ %s ], because no handler found handler", messageType)
	}
	// get the function
	function := handler.Handler()
	// execute the handler
	return function(connection, message)
}
