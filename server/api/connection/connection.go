package connection

import (
	"log"

	"github.com/dayaftereh/discover/server/game/player"
	"github.com/dayaftereh/discover/server/utils/atomic"
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
	open *atomic.AtomicBool
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
		open: atomic.NewAtomicBool(false),
	}

	connection.init()

	return connection
}

func (connection *Connection) Id() string {
	return connection.ID
}

func (connection *Connection) init() {
	// set connection to open
	connection.open.Set(true)
	// start in/out bound loop
	go connection.inBoundLoop()
	go connection.outBoundLoop()
}

func (connection *Connection) inBoundLoop() {
	defer connection.Close()
	for {

		if !connection.open.Get() {
			log.Printf("closing inbound thread for connection [ %s ]\n", connection.ID)
			return
		}

		_, bytes, err := connection.conn.ReadMessage()
		if err != nil {
			connection.OnError <- err
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
			log.Printf("closing outbound thread for connection [ %s ]\n", connection.ID)
			return
		}
	}
}

func (connection *Connection) Write(message string) {
	// check if connection is open
	if connection.open.Get() {
		connection.OutBoundMessage <- message
	}
}

func (connection *Connection) Close() {
	// check if connection is open and set to false
	if !connection.open.GetAndSet(false) {
		return
	}

	log.Printf("closing connection [ %s ]\n", connection.ID)

	// close the underlying connection
	connection.conn.Close()
	// notify in/out bound loop about close
	connection.OnClose <- true
}
