package game

import (
	"github.com/dayaftereh/discover/server/game/player"
	"github.com/dayaftereh/discover/server/mathf"
	"github.com/pkg/errors"
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

func (game *Game) Ready(player *player.Player) error {
	// check if player has star system
	if player.StarSystem == nil {
		// get the initial star system
		initialStarSystem, err := game.universe.GetInitialStarSystem()
		if err != nil {
			return err
		}
		// let the player join the star system
		initialStarSystem.JoinPlayer(player)
		return nil
	}

	// get the player star system
	starSystem := game.universe.GetStarSystem(*player.StarSystem)
	// check if a star system exists
	if starSystem == nil {
		return errors.Errorf("unable to join player [ %s ] into star-system [ %d ], because star-system not found", player.Name, starSystem.ID)
	}
	// let the player join the star system
	starSystem.JoinPlayer(player)

	return nil
}

func (game *Game) Movement(player *player.Player, move *mathf.Vec3, rotation *mathf.Vec3) {
	// get the player star system
	starSystem := game.universe.GetStarSystem(*player.StarSystem)
	// check if a star system exists
	if starSystem == nil {
		return
	}

	starSystem.UpdatePlayer(player, move, rotation)
}
