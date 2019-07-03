package main

import (
	"log"
	"os"
	"path"

	"github.com/dayaftereh/discover/server/game"
)

func gameDirectory() (string, error) {
	// get current directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// path to the game data
	gameDirectory := path.Join(cwd, "data")

	return gameDirectory, nil
}

func initGame() (*game.Game, error) {
	// create a new game
	game := game.NewGame()

	// get the game directory
	gameDirectory, err := gameDirectory()
	if err != nil {
		return nil, err
	}

	log.Printf("loading game data from [ %s ]\n", gameDirectory)

	// load the game data
	err = game.Init(gameDirectory)
	if err != nil {
		return nil, err
	}

	log.Println("game successful initialized")

	return game, nil
}

func shutdownGame(game *game.Game) error {
	log.Println("shutdown game...")

	// get the game directory
	gameDirectory, err := gameDirectory()
	if err != nil {
		return err
	}

	// shutdown the game
	err = game.Shutdown(gameDirectory)
	return err
}
