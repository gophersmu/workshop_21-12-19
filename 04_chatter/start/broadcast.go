package main

// broadcastMessage sends a message to all peers
func broadcastMessage(msg string) {
	// range over "peers" and dial a udp connection for each peer
	// use the connection to send "msg" (e.g conn.Write(msg))
	// remember connections are resources and need to be closed (aka 'freed') if opened ;)
}
