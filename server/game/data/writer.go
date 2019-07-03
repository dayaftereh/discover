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
	// check if a the data directory exists
	ok, err := utils.Exists(dataPath)
	if err != nil {
		return err
	}

	// create directory if needed
	if !ok {
		err = os.MkdirAll(dataPath, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// get the path to the data file
	dataFile := path.Join(dataPath, dataFileName)

	// convert game to json
	bytes, err := json.Marshal(game)
	if err != nil {
		return err
	}

	// write the json file
	err = ioutil.WriteFile(dataFile, bytes, os.ModePerm)
	return err
}
