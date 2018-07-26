package websocket

import (
	"../../utils"
	ws "github.com/gorilla/websocket"
	"log"
	"sync"
)

type WebSocketManager struct {
	lock            sync.RWMutex
	OnOpen          chan OpenEvent
	OnClose         chan CloseEvent
	OnError         chan ErrorEvent
	InBoundMessage  chan InBoundMessageEvent
	OutBoundMessage chan OutBoundMessageEvent
	webSockets      map[string]*WebSocket
}

func (webSocketManager *WebSocketManager) Open(sessionId string, conn *ws.Conn) error {
	// generate new connection id
	id, err := utils.RandString(64)
	if err != nil {
		return err
	}

	log.Println("websocket [", id, "] established for session [", sessionId, "]")

	// create a new web-socket
	webSocket := NewWebSocket(id, sessionId, con)

	// aquire lock
	webSocketManager.lock.Lock()
	// store the websocket
	webSocketManager.webSockets[id] = webSocket
	// release the lock
	webSocketManager.lock.Unlock()

	// start the forwarder
	go forwarder(webSocket)

	return nil
}

func (webSocketManager *WebSocketManager) forwarder(websocket *WebSocket) {
	// close the websocket on end
	defer webSocketManager.close(websocket)

	// notify about open
	webSocketManager.OnOpen <- OpenEvent{
		WebSocket: webSocket,
	}

	for {
		select {
		case <-websocket.OnClose:
			return
		case err := <-websocket.OnError:
			webSocketManager.OnError <- ErrorEvent{
				Error:     err,
				WebSocket: webSocket,
			}
		case message := <-websocket.InBoundMessage:
			webSocketManager.InBoundMessage <- InBoundMessageEvent{
				Message:   message,
				WebSocket: webSocket,
			}
		case message := <-websocket.OutBoundMessage:
			webSocketManager.OutBoundMessage <- OutBoundMessageEvent{
				Message:   message,
				WebSocket: webSocket,
			}
		}
	}

}

func (webSocketManager *WebSocketManager) close(websocket *WebSocket) {
	// close the underlying connection
	websocket.Close()

	// notify about close
	webSocketManager.OnClose <- CloseEvent{
		WebSocket: websocket,
	}

	// aquire lock
	webSocketManager.lock.Lock()
	// delete the websocket
	delete(webSocketManager.webSockets, websocket.Id)
	// release the lock
	webSocketManager.lock.Unlock()
}
