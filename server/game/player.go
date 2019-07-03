package game

import (
	"github.com/dayaftereh/discover/server/game/player"
	"github.com/dayaftereh/discover/server/mathf"
)

func (game *Game) SessionByName(id string, name string) *player.Player {
	// get or create the player session for the given name
	player := game.playerManager.SessionByName(id, name)
	return player
}

// GetPlayerSession returns the session for the player
func (game *Game) GetPlayerSession(id string) *player.Player {
	// get the player session
	player := game.playerManager.GetSession(id)

	return player
}

// DropPlayerSession for given id from Game
func (game *Game) DropPlayerSession(id string) {
	// remove the player
	player := game.playerManager.DropSession(id)

	// check for player found
	if player == nil {
		return
	}

	// get the player star system
	starSystem := game.universe.GetStarSystem(*player.StarSystem)
	// check if a star system exists
	if starSystem == nil {
		return
	}
	// drop the player from star system
	starSystem.DropPlayer(player)
}

func (game *Game) JoinPlayer(player *player.Player) {
	// get the player star system
	starSystem := game.universe.GetStarSystem(*player.StarSystem)
	// check if a star system exists
	if starSystem == nil {
		return
	}
	// let the player join the star system
	starSystem.JoinPlayer(player)
}

func (game *Game) Movement(player *player.Player, lookAt *mathf.Vec3) {
	// get the player star system
	starSystem := game.universe.GetStarSystem(*player.StarSystem)
	// check if a star system exists
	if starSystem == nil {
		return
	}

	starSystem.UpdatePlayer(player, lookAt)
}
