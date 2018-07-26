package websocket

import (
	ws "github.com/gorilla/websocket"
)

type WebSocket struct {
	// id and session
	Id        string
	SessionId string

	// Events
	OnClose         chan bool
	OnError         chan error
	InBoundMessage  chan string
	OutBoundMessage chan string

	// private
	conn *ws.Conn
}

func NewWebSocket(id string, sessionId string, conn *ws.Conn) *WebSocket {
	webSocket := &WebSocket{
		Id:              id,
		SessionId:       sessionId,
		OnClose:         make(chan bool),
		OnError:         make(chan error),
		InBoundMessage:  make(chan string),
		OutBoundMessage: make(chan string),
		conn:            conn,
	}
	go webSocket.inBoundLoop()
	go webSocket.outBoundLoop()

	return webSocket
}

func (webSocket *WebSocket) inBoundLoop() {
	defer webSocket.Close()
	for {
		_, bytes, err := webSocket.conn.ReadMessage()
		if err != nil {
			webSocket.OnError <- err
			return
		}

		message := string(bytes)
		webSocket.InBoundMessage <- message
	}
}

func (webSocket *WebSocket) outBoundLoop() {
	defer webSocket.Close()
	for {
		select {
		case message := <-websocket.OutBoundMessage:
			bytes := []byte(message)
			err := websocket.conn.WriteMessage(ws.TextMessage, bytes)
			if err != nil {
				connection.OnError <- err
			}
		case <-webSocket.OnClose:
			return
		}
	}
}

func (webSocket *WebSocket) Write(message string) {
	webSocket.OutBoundMessage <- message
}

func (webSocket *WebSocket) Close() {
	webSocket.conn.Close()
	webSocket.OnClose <- true
}
