package binGO_server

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	game_manager "binGO/binGO.server/binGO.server.game"

	"github.com/coder/websocket"
)

type binGOServer struct {
	gm        *game_manager.GameManager
	clients   map[*websocket.Conn]bool
	broadcast chan int
	mu        sync.Mutex
}

func NewServer(gm *game_manager.GameManager) *binGOServer {
	return &binGOServer{
		gm:        gm,
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan int),
	}
}

func (s *binGOServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols: []string{"binGO"},
	})

	if err != nil {
		http.Error(w, "WebSocket upgrade failed", http.StatusInternalServerError)
		return
	}

	if conn.Subprotocol() != "binGO" {
		http.Error(w, "Invalid subprotocol", http.StatusBadRequest)
		return
	}

	s.mu.Lock()
	s.clients[conn] = true
	s.mu.Unlock()

	fmt.Println("New client connected")

	go s.handleClient(conn)
}

func (s *binGOServer) handleClient(conn *websocket.Conn) {
	defer func() {
		s.mu.Lock()
		delete(s.clients, conn)
		s.mu.Unlock()
		conn.CloseNow()
		fmt.Println("Client disconnected")
	}()

	// Optionally read messages from client here
	for {
		_, msg, err := conn.Read(context.Background())
		if err != nil {
			fmt.Printf("Error reading from client: %v\n", err)
			break
		}
	}
}

func (s *binGOServer) StartBallCaller() {
	for {
		fmt.Println("Press Enter to call a bingo ball")
		fmt.Scanln()

		if string(msg) == "BINGO!" {
			fmt.Println("Bingo received from client")
		}
	}
}

func (s *binGOServer) StartBallCaller() {
	for {
		fmt.Println("Press Enter to call a bingo ball")
		fmt.Scanln()

		s.broadcast <- ball
	}
}

func (s *binGOServer) StartBroadcaster() {
	for msg := range s.broadcast {
		s.mu.Lock()
		for conn := range s.clients {
			go func(c *websocket.Conn) {
				ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				writer, err := c.Writer(ctx, websocket.MessageText)
				if err == nil {
					writer.Write([]byte(strconv.Itoa(msg)))
					writer.Close()
				}
			}(conn)
		}
		s.mu.Unlock()
	}
}
