package persistence

import (
	"image"
	"image/png"
	"os"
	"path/filepath"

	"github.com/dayaftereh/discover/server/utils"
)

func writeImage(filename string, img *image.RGBA) error {
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

	// create the file
	fPath, err := os.Create(filename)
	if err != nil {
		return err
	}

	// write the image as png
	err = png.Encode(fPath, img)
	return err
}
