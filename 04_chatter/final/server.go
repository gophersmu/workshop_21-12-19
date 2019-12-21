package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// updServer listens to incoming udp packets
func updServer() {
	// starts a udp listener (connection) on port "udpPort"
	// use a "for loop" to keep on handling messages (i.e call "handleMessage")
	// remember connections are resources and need to be closed (aka 'freed') if opened ;)
	// if there's an error when starting server, log.Fatal ;)
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: udpPort,
	})
	if err != nil {
		log.Fatalf("UDP server failed: %v", err)
	}

	defer conn.Close()

	ui.AddMessage(infoID, fmt.Sprintf("UDP server listening @ %s...", conn.LocalAddr().String()))

	for {
		data := make([]byte, 256)

		dlen, addr, err := conn.ReadFromUDP(data)
		if err != nil {
			continue
		}

		ip := addr.IP.String()
		msg := strings.TrimSpace(string(data[:dlen]))

		go handleMessage(ip, msg)
	}
}

// handleMessage handles a message from updServer
func handleMessage(ip string, msg string) {
	// use "ip" to get "id" from "peers"
	// if "ip" does not exist in "peers", ignore the "msg"
	// display "id" and "msg" on the UI (i.e use ui.AddMessage)
	rawID, ok := peers.Load(ip)
	if !ok {
		return
	}

	id := rawID.(string)

	ui.AddMessage(id, msg)
}
