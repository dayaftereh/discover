package persistence

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/dayaftereh/discover/server/utils"
)

func readJSON(filename string, v interface{}) error {
	ok, err := utils.Exists(filename)
	// check for error
	if err != nil {
		return err
	}

	// check if file exists
	if !ok {
		return nil
	}

	// read the file
	bytes, err := ioutil.ReadFile(filename)

	// check if read success
	if err != nil {
		return err
	}

	// decode json
	err = json.Unmarshal(bytes, v)
	return err

}

func writeJSON(filename string, v interface{}) error {
	directory := filepath.Dir(filename)
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

	// convert data to json
	bytes, err := json.Marshal(v)
	if err != nil {
		return err
	}

	// write the json file
	err = ioutil.WriteFile(filename, bytes, os.ModePerm)
	return err
}
