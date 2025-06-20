package binGO_client

import (
	"context"
	"fmt"
	"time"

	"github.com/coder/websocket"
)

// This function creates a web socket
// connection to the server and handles incoming messages.
func StartClient() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, _, err := websocket.Dial(ctx, "ws://localhost:8080", &websocket.DialOptions{
		Subprotocols: []string{"binGO"},
	})

	if err != nil {
		fmt.Printf("Failed to connect to server: %s\n", err.Error())
		return
	}

	defer conn.Close(websocket.StatusInternalError, "closing")

	// Handle incoming messages from the server
	for {
		readCtx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
		_, msg, err := conn.Read(readCtx)
		cancel()

		if err != nil {
			fmt.Printf("Failed to read message: %s\n", err.Error())
			continue
		}

		// Process the received message (e.g., display it)
		fmt.Println("Received message:", string(msg))
	}
}
