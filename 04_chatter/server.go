package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

// listen to incoming udp packets
func updServer() {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{
		Port: udpPort,
		IP:   net.ParseIP("0.0.0.0"),
	})
	if err != nil {
		log.Fatalf("UDP server failed: %v", err)
	}

	defer conn.Close()

	ui.AddMessage(infoID, fmt.Sprintf("UDP server listening @ %s...", conn.LocalAddr().String()))

	for {
		data := make([]byte, 256)

		dlen, addr, err := conn.ReadFromUDP(data[:])
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
	rawID, ok := peers.Load(ip)
	if !ok {
		// Ignore if we don't know that peer
		return
	}

	id := rawID.(string)

	ui.AddMessage(id, msg)
}
