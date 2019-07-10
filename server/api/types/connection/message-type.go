package connection

type MessageType string

const (
	Move   MessageType = "move"
	Update MessageType = "update"
)
