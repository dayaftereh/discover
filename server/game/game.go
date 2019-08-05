package game

import (
	"github.com/dayaftereh/discover/server/game/persistence"
	"github.com/dayaftereh/discover/server/game/player"
	"github.com/dayaftereh/discover/server/game/universe"
)

type Game struct {
	// player
	playerManager *player.Manager
	// Universe
	universe *universe.Universe
	// persistence
	persistence *persistence.PersistenceManager
}

func NewGame(directory string) *Game {
	persistenceManager := persistence.NewPersistenceManager(directory)
	return &Game{
		universe:      universe.NewUniverse(persistenceManager),
		playerManager: player.NewPlayerManager(persistenceManager),
		persistence:   persistenceManager,
	}
}
