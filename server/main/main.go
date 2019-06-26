package main

import (
	"log"

	"github.com/pkg/errors"
)

var (
	// VERSION - the server version
	VERSION string
	// RELEASE - the server release date
	RELEASE string
)

func main() {

	log.Println(VERSION, RELEASE)

	if false {
		errors.New("Error")
	}
}
