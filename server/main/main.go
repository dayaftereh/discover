package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	// VERSION - the server version
	VERSION string
	// RELEASE - the server release date
	RELEASE string
)

func main() {
	log.Printf("Version: %s\n", VERSION)
	log.Printf("Release: %s\n", RELEASE)

	// start discovery
	insatnce, err := start()
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	// wait for shutdown signals
	wait()

	// shutdown
	err = shutdown(insatnce)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	log.Println("Done.")

	//buf := make([]byte, 1<<16)
	//runtime.Stack(buf, true)
	//fmt.Printf("%s", buf)
}

func wait() {
	latch := make(chan bool, 1)
	signals := make(chan os.Signal, 1)

	// register signals
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// forward function
	go func() {
		<-signals
		latch <- true
	}()

	// wati for signals
	<-latch

	signal.Stop(signals)

	log.Println("Received shutdown signal")
}
