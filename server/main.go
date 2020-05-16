package main

import (
	"fmt"
	"time"

	"./database"
	"./routing"
)

type someData struct {
	id        int
	message   string
	isChanged bool
}

func testMe(message string, f func() error) {
	start := time.Now()
	defer func() {
		duration := time.Since(start)

		// Formatted string, such as "2h3m0.5s" or "4.503μs"
		fmt.Println(message, duration)
	}()
	f()
}

func main() {
	//enable multi core and shit
	//runtime.GOMAXPROCS(runtime.NumCPU())

	go database.StartZoneStatusUpdate()

	server := routing.SetupRouter()
	// running
	server.Run()
}
