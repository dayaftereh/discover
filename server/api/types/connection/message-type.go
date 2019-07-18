package connection

type MessageType string

const (
	Move        MessageType = "move"
	WorldUpdate MessageType = "world-update"
)
