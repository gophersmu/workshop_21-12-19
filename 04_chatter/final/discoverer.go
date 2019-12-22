package main

import (
	"fmt"
	"log"
	"time"

	"github.com/schollz/peerdiscovery"
)

// Discover peers on same network
func discoverer() {
	// use "peerdiscovery" package to actively discover
	// new peers on the network
	// every time there's a new peer, "store" it in "peers"
	// and "add" the user on the UI (i.e use ui.AddUser)
	// if there's an error, log.Fatal ;)
	ui.AddMessage(infoID, "Peer discovery started...")

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

				_, loaded := peers.LoadOrStore(ip, id)
				if loaded {
					// User already exist
					// just update the id
					peers.Store(ip, id)
					return
				}

				ui.AddMessage(infoID, fmt.Sprintf("%s has joined ^_^", id))
				ui.AddUser(id)
			},
		},
	)

	if err != nil {
		log.Fatalf("Peer discovery failed: %v", err)
	}
}
