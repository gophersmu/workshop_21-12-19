package main

// listen to incoming udp packets
func updServer() {
	// starts a udp listener (connection) on port "udpPort"
	// use a "for loop" to keep on handling messages (i.e call "handleMessage")
	// remember connections are resources and need to be closed if opened ;)
	// if there's an error when starting server, log.Fatal ;)
}

// handleMessage handles a message from updServer
func handleMessage(ip string, msg string) {
	// use "ip" to get "id" from "peers"
	// if "ip" does not exist in "peers", ignore the "msg"
	// display "id" and "msg" on the UI (i.e use ui.AddMessage)
}
