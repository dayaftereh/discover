package object

import (
	"github.com/dayaftereh/discover/server/mathf"
)

type Player struct {
	id   int64
	body *Body
}

func NewPlayer(id int64, position *mathf.Vec3) *Player {
	return &Player{
		id:   id,
		body: NewBody(position, 1.0),
	}
}

func (player *Player) ID() int64 {
	return player.id
}

func (player *Player) Body() *Body {
	return player.body
}

func (player *Player) Update(move *mathf.Vec3, rotation *mathf.Vec3) {

}
