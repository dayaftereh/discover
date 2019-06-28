package common

import "github.com/dayaftereh/discover/server/game/player"

type Backend interface {
	GetPlayer(id string) (*player.Player, error)
	DropPlayer(id string) error
	HasPlayer(id string) bool
}
