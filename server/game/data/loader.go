package data

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"github.com/dayaftereh/discover/server/utils"
)

const playersFileName = "players.json"
const universeFileName = "universe.json"

// Load the Game data from the given directory
func Load(dataPath string) (*Game, error) {

	// load the universe
	universe, err := loadUniverse(dataPath)
	if err != nil {
		return nil, err
	}

	// load the players
	players, err := loadPlayers(dataPath)
	if err != nil {
		return nil, err
	}

	// return a new game
	return &Game{
		Players:  players,
		Universe: universe,
	}, nil
}

func loadUniverse(directory string) (*Universe, error) {
	var universe Universe
	// read the universe
	err := readJSON(directory, universeFileName, &universe)
	if err != nil {
		return nil, err
	}

	return &universe, nil
}

func loadPlayers(directory string) (map[string]*Player, error) {
	var players map[string]*Player
	// read the players
	err := readJSON(directory, playersFileName, &players)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func readJSON(directory string, fileName string, v interface{}) error {
	// get the path to the file
	fpath := path.Join(directory, fileName)

	ok, err := utils.Exists(fpath)
	// check for error
	if err != nil {
		return err
	}

	// check if file exists
	if !ok {
		return nil
	}

	// read the file
	bytes, err := ioutil.ReadFile(fpath)

	// check if read success
	if err != nil {
		return err
	}

	// decode json
	err = json.Unmarshal(bytes, v)
	return err

}
