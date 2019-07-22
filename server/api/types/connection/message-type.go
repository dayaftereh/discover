package connection

type MessageType string

const (
	Ping        MessageType = "ping"
	Pong        MessageType = "pong"
	Move        MessageType = "move"
	WorldUpdate MessageType = "world-update"
)
