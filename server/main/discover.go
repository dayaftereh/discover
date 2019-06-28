package main

import (
	"github.com/dayaftereh/discover/server/game"
)

func start() error {
	// create a new game
	game := game.NewGame()

	// create the server
	server, err := initAPI(game)
	if err != nil {
		return err
	}

	// inilize the server
	err = server.Init()
	if err != nil {
		return err
	}

	return nil
}
