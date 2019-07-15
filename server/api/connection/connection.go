package connection

import (
	"github.com/dayaftereh/discover/server/game/player"
	"github.com/gorilla/websocket"
)

type Connection struct {
	ID     string
	Player *player.Player

	// Events
	OnClose         chan bool
	OnError         chan error
	InBoundMessage  chan string
	OutBoundMessage chan string

	conn *websocket.Conn
	open bool
}

func NewConnection(id string, player *player.Player, conn *websocket.Conn) *Connection {
	connection := &Connection{
		ID:     id,
		Player: player,
		// Events
		OnClose:         make(chan bool),
		OnError:         make(chan error),
		InBoundMessage:  make(chan string),
		OutBoundMessage: make(chan string),
		// private
		conn: conn,
		open: true,
	}

	connection.init()

	return connection
}

func (connection *Connection) Id() string {
	return connection.ID
}

func (connection *Connection) init() {
	go connection.inBoundLoop()
	go connection.outBoundLoop()
}

func (connection *Connection) inBoundLoop() {
	defer connection.Close()
	for {
		_, bytes, err := connection.conn.ReadMessage()
		if err != nil {
			connection.OnError <- err
			return
		}

		message := string(bytes)
		connection.InBoundMessage <- message
	}
}

func (connection *Connection) outBoundLoop() {
	defer connection.Close()
	for {
		select {
		case message := <-connection.OutBoundMessage:
			bytes := []byte(message)
			err := connection.conn.WriteMessage(websocket.TextMessage, bytes)
			if err != nil {
				connection.OnError <- err
			}
		case <-connection.OnClose:
			return
		}
	}
}

func (connection *Connection) Write(message string) {
	if connection.open {
		connection.OutBoundMessage <- message
	}
}

func (connection *Connection) Close() {
	connection.open = false
	connection.conn.Close()
	connection.OnClose <- true
}
