package main

import (
	"log"
	"sync"

	"github.com/gophersmu/workshop_21-12-19/chatter/rand"
)

const (
	// udp port used for our server and remote clients
	udpPort = 13337

	// id used to display INFO messages
	infoID = "INFO"
)

var (
	// id represents our client id
	id string

	// peers represents a map[ip]id
	peers sync.Map

	// ui represents the terminal chat ui
	ui *UI
)

func main() {
	var err error

	// Generate our ID using a random string of length 10
	id = rand.String(10)

	// Create a new terminal UI
	ui, err = NewUI()
	if err != nil {
		log.Fatalf("failed to create ui: %v", err)
	}

	// Starts discoverer in non-blocking
	go discoverer()

	// Starts UDP Server in non-blocking
	go updServer()

	// Starts the terminal ui in blocking
	if err := ui.Run(); err != nil {
		log.Fatalf("failed to start ui: %v", err)
	}
}
