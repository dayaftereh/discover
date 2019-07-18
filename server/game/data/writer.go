package data

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/dayaftereh/discover/server/utils"
)

// Write the game to the given data directory
func Write(dataPath string, game *Game) error {
	// write the players
	err := wirtePlayers(dataPath, game.Players)
	if err != nil {
		return err
	}

	// write the universe
	err = wirteUniverse(dataPath, game.Universe)
	return err
}

func wirtePlayers(directory string, players map[string]*Player) error {
	err := writeJSON(directory, playersFileName, players)
	return err
}

func wirteUniverse(directory string, universe *Universe) error {
	err := writeJSON(directory, universeFileName, universe)
	return err
}

func writeJSON(directory string, fileName string, v interface{}) error {
	// check if a data directory exists
	ok, err := utils.Exists(directory)
	if err != nil {
		return err
	}

	// create directory if needed
	if !ok {
		err = os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// get the path to the file
	fpath := path.Join(directory, fileName)

	// convert data to json
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	// write the json file
	err = ioutil.WriteFile(fpath, bytes, os.ModePerm)
	return err
}
