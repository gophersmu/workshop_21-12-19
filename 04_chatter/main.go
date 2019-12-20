package main

import (
	"sync"

	"github.com/gophersmu/workshop_21-12-19/chatter/rand"
)

const (
	// udp port used for our server and remote clients
	udpPort = 13337
)

var (
	// id represents our client id
	id string

	// peers represents a map[ip]id
	peers sync.Map
)

func main() {
	// Generate our ID using a random string of length 10
	id = rand.String(10)

	// Starts discoverer in non-blocking
	go discoverer()

	// Starts UDP Server in non-blocking
	go updServer()

	// Starts the terminal reader in blocking
	reader()
}
