package connection

type Function func(connection *Connection, message string) error
