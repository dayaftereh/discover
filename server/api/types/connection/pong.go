package connection

type PongMessage struct {
	Type              MessageType `json:"type"`
	ClientSendTime    float64     `json:"clientSendTime"`
	ServerReceiveTime float64     `json:"serverReceiveTime"`
	ServerSendTime    float64     `json:"serverSendTime"`
}
