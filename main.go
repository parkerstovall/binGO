package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	binGO_client "binGO/binGO.client"
	binGO_server "binGO/binGO.server"
	game_manager "binGO/binGO.server/binGO.server.game"

	"github.com/coder/websocket"
	"github.com/fatih/color"
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

// This function creates a web socket
// connection to the server and handles incoming messages.
func startClient() {
	ctx := context.Background()

	conn, _, err := websocket.Dial(ctx, "ws://localhost:8080", &websocket.DialOptions{
		Subprotocols: []string{"binGO"},
	})

	if err != nil {
		error_msg := fmt.Sprintf("Failed to connect to server: %s", err.Error())
		color.White(error_msg)
		return
	}

	defer conn.Close(websocket.StatusInternalError, "closing")

	client := binGO_client.NewClient(ctx)
	client.PrintBoard()
	client.ListenForCalledBalls(conn)
}

func main() {
	// Ask if the user wants to start the server or run the client
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "server":
			startServer()
			return
		case "client":
			startClient()
			return
		}
	}

	var choice string
	color.White("Do you want to start the server or run the client? (server/client)")
	fmt.Scanln(&choice)

	switch choice {
	case "server":
		startServer()
	case "client":
		startClient()
	default:
		color.White("Invalid choice. Please enter 'server' or 'client'.")
	}
}
