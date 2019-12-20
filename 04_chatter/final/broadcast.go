package main

import "net"

// broadcastMessage sends a message to all peers
func broadcastMessage(msg string) {
	// range over "peers" and dial a udp connection for each peer
	// use the connection to send "msg" (e.g conn.Write(msg))
	// remember connections are resources and need to be closed (aka 'freed') if opened ;)
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
