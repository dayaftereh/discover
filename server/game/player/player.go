package player

type Player struct {
	ID string
	// private
	name        *string
	connections map[string]Connection
}

func NewPlayer(id string) *Player {
	return &Player{
		ID:          id,
		connections: make(map[string]Connection),
	}
}

func (player *Player) GetName() *string {
	return player.name
}

func (player *Player) SetName(name *string) {
	player.name = name
}
