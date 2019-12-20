package main

import "net"

// sendMessage sends a message to all peers
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
}
