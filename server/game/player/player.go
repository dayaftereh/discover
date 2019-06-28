package player

type Player struct {
	Id   string
	Name *string
}

func NewPlayer(id string) *Player {
	return &Player{
		Id: id,
	}
}
