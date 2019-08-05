package game

func (game *Game) Shutdown(directory string) error {
	// shutdown the universe
	game.universe.Shutdown()

	return nil
}
