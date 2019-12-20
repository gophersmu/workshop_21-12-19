package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// reader reads indefintely from the terminal
func reader() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message: ")

		msg, _ := reader.ReadString('\n')
		msg = strings.Trim(msg, "\n")
		if msg == "" {
			continue
		}

		go sendMessage(msg)
	}
}

// sendMessage sends a message to all peers including ourself
func sendMessage(msg string) {
	// range over peers
	peers.Range(func(key interface{}, value interface{}) bool {
		peerIP := value.(string)

		conn, err := net.DialUDP("udp", nil, &net.UDPAddr{
			Port: udpPort,
			IP:   net.ParseIP(peerIP),
		})
		if err != nil {
			return true
		}
		defer conn.Close()

		// Write on conn
		conn.Write([]byte(msg))

		return true
	})

	// print msg on our side
	log.Printf("<%s> %s", id, msg)
}
