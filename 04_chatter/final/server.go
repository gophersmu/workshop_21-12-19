package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// listen to incoming udp packets
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
	// starts a udp listener (connection) on port "udpPort"
	// remember connections are resources and need to be closed (aka 'freed') if opened ;)
	// use a "for loop" to continuously "handle" incoming messages (i.e use "handleMessage")
	// if there's an error when starting server, log.Fatal ;)
	rawID, ok := peers.Load(ip)
	if !ok {
		return
	}

	id := rawID.(string)

	ui.AddMessage(id, msg)
}
