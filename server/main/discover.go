package main

import (
	"log"

	"github.com/dayaftereh/discover/server/utils"

	"github.com/dayaftereh/discover/server/api/server"
	"github.com/dayaftereh/discover/server/game"
)

type discover struct {
	Game   *game.Game
	Server *server.Server
}

func start() (*discover, error) {
	t := utils.SystemSeconds()
	log.Println("starting discover...")

	// create the game
	game, err := initGame()
	if err != nil {
		return nil, err
	}

	// create the server
	server, err := initAPI(game)
	if err != nil {
		return nil, err
	}

	startupTime := utils.SystemSeconds() - t
	log.Printf("discover started in [ %f s ]\n", startupTime)

	return &discover{
		Game:   game,
		Server: server,
	}, nil
}

func shutdown(instance *discover) error {
	log.Println("shutdown discover...")

	// shutdown the api
	err := shutdownAPI(instance.Server)
	if err != nil {
		return err
	}

	// shutdown the game
	err = shutdownGame(instance.Game)
	if err != nil {
		return err
	}

	return nil
}
