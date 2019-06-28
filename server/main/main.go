package main

import (
	"log"
)

var (
	// VERSION - the server version
	VERSION string
	// RELEASE - the server release date
	RELEASE string
)

func main() {

	log.Println(VERSION, RELEASE)

	// start discovery
	err := start()
	if err != nil {
		log.Fatalf("%s\n", err)
	}
	
}
