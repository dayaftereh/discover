package common

import "github.com/dayaftereh/discover/server/game/player"

type Backend interface {
	SessionByName(id string, name string) (*player.Player, error)
	GetPlayerSession(id string) *player.Player
	DropPlayerSession(id string)
}
