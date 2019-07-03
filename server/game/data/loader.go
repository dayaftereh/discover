package data

import (
	"encoding/json"
	"io/ioutil"
	"path"

	"github.com/dayaftereh/discover/server/utils"
)

const dataFileName = "data.json"

// Load the Game data from the given directory
func Load(dataPath string) (*Game, error) {
	// get the path to the data file
	dataFile := path.Join(dataPath, dataFileName)

	// check if a the data file exists
	ok, err := utils.Exists(dataFile)
	if err != nil {
		return nil, err
	}

	// if file not exists return empty
	if !ok {
		return NewGame(), nil
	}

	// read the file
	bytes, err := ioutil.ReadFile(dataPath)

	if err != nil {
		return nil, err
	}

	// parse the json
	var game Game
	err = json.Unmarshal(bytes, &game)
	// check for success
	if err != nil {
		return nil, err
	}

	// return game
	return &game, nil
}

func NewGame() *Game {
	return &Game{
		Players:     make(map[string]*Player),
		StarSystems: make(map[int64]*StarSystem),
	}
}
