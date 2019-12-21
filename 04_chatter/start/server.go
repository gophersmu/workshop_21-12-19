package main

// updServer listens to incoming udp packets
func updServer() {
	// starts a udp listener (connection) on port "udpPort"
	// remember connections are resources and need to be closed (aka 'freed') if opened ;)
	// use a "for loop" to continuously "handle" incoming messages (i.e use "handleMessage")
	// if there's an error when starting server, log.Fatal ;)
}

// handleMessage handles a message from updServer
func handleMessage(ip string, msg string) {
	// use "ip" to get "id" from "peers"
	// if "ip" does not exist in "peers", ignore the "msg"
	// display "id" and "msg" on the UI (i.e use ui.AddMessage)
}
