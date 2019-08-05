package game

func (game *Game) Init() error {
	// initialize the player manager
	err := game.playerManager.Init()
	if err != nil {
		return err
	}

	// initialize the universe
	err = game.universe.Init()

	return err
}
