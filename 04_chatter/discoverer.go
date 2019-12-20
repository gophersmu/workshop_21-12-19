package main

import (
	"fmt"
	"log"
	"time"

	"github.com/schollz/peerdiscovery"
)

// Discover peers on same network
func discoverer() {
	ui.AddMessage("SERVER", "Peer discovery started...")

	_, err := peerdiscovery.Discover(
		peerdiscovery.Settings{
			Limit:     -1,              // Unlimited
			Payload:   []byte(id),      // Send a random string of length 10 as payload (used as ID)
			Delay:     1 * time.Second, // Discover every 1 second
			TimeLimit: -1,              // Keep on scanning undefinitely
			Notify: func(d peerdiscovery.Discovered) {
				// A new peer was discovered
				id := string(d.Payload)
				ip := d.Address
				peers.Store(ip, id)

				ui.AddMessage("SERVER", fmt.Sprintf("%s has joined ^_^", id))
				ui.AddUser(id)
			},
		},
	)

	if err != nil {
		log.Fatalf("Peer discovery failed: %v", err)
	}
}
