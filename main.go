package main

import (
	"fmt"
	"net/http"

	binGO_client "binGO/binGO.client"
	binGO_server "binGO/binGO.server"
	game_manager "binGO/binGO.server/binGO.server.game"
)

func startServer() {
	gm := game_manager.NewGameManager()
	server := binGO_server.NewServer(gm)

	go server.StartBroadcaster()
	go server.StartBallCaller()

	http.Handle("/", server)
	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func main() {
	// Ask if the user wants to start the server or run the client
	var choice string
	fmt.Println("Do you want to start the server or run the client? (server/client)")
	fmt.Scanln(&choice)
	switch choice {
	case "server":
		startServer()
	case "client":
		binGO_client.StartClient()
	default:
		fmt.Println("Invalid choice. Please enter 'server' or 'client'.")
	}
}
