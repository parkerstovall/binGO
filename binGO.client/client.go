package binGO_client

import (
	"context"
	"fmt"
	"math/rand"
	"slices"
	"strconv"

	"github.com/coder/websocket"
	"github.com/fatih/color"
)

type binGOclient struct {
	Board       [][]int
	CalledBalls []int
	ctx         context.Context
}

func NewClient(ctx context.Context) *binGOclient {
	client := &binGOclient{
		Board:       make([][]int, 5),
		CalledBalls: make([]int, 0),
		ctx:         ctx,
	}

	client.initializeBoard()

	return client
}

func (c *binGOclient) PrintBoard() {
	fmt.Println()
	fmt.Println()
	fmt.Println()

	uncalled := color.New(color.FgRed).PrintFunc()
	called := color.New(color.FgGreen).PrintFunc()
	// Print the Bingo board in a formatted way
	for i := range c.Board {
		for j := range c.Board[i] {
			letters := []string{"B", "I", "N", "G", "O"}
			msg := fmt.Sprintf("%s%d", letters[j], c.Board[i][j])
			if c.Board[i][j] < 10 {
				msg += " "
			}

			if slices.Contains(c.CalledBalls, c.Board[i][j]) {
				called(msg + " ")
			} else {
				uncalled(msg + " ")
			}
		}

		fmt.Println()
	}
}

func (c *binGOclient) initializeBoard() {
	// Initialize the Bingo board
	for i := range c.Board {
		c.Board[i] = make([]int, 5)

		// Fill the board with random numbers
		for j := range c.Board[i] {
			rand_number := rand.Intn(15) + 1 + (j * 15)
			c.Board[i][j] = rand_number
		}
	}
}

func (c *binGOclient) CheckForBingo() int {
	for i := range 5 {
		// Horizontal
		if slices.Contains(c.CalledBalls, c.Board[i][0]) &&
			slices.Contains(c.CalledBalls, c.Board[i][1]) &&
			slices.Contains(c.CalledBalls, c.Board[i][2]) &&
			slices.Contains(c.CalledBalls, c.Board[i][3]) &&
			slices.Contains(c.CalledBalls, c.Board[i][4]) {
			return 1
		}

		// Vertical
		if slices.Contains(c.CalledBalls, c.Board[0][i]) &&
			slices.Contains(c.CalledBalls, c.Board[1][i]) &&
			slices.Contains(c.CalledBalls, c.Board[2][i]) &&
			slices.Contains(c.CalledBalls, c.Board[3][i]) &&
			slices.Contains(c.CalledBalls, c.Board[4][i]) {
			return 1
		}
	}

	// Diagonal (top-left to bottom-right)
	if slices.Contains(c.CalledBalls, c.Board[0][0]) &&
		slices.Contains(c.CalledBalls, c.Board[1][1]) &&
		slices.Contains(c.CalledBalls, c.Board[2][2]) &&
		slices.Contains(c.CalledBalls, c.Board[3][3]) &&
		slices.Contains(c.CalledBalls, c.Board[4][4]) {
		return 1
	}

	// Diagonal (top-right to bottom-left)
	if slices.Contains(c.CalledBalls, c.Board[0][4]) &&
		slices.Contains(c.CalledBalls, c.Board[1][3]) &&
		slices.Contains(c.CalledBalls, c.Board[2][2]) &&
		slices.Contains(c.CalledBalls, c.Board[3][1]) &&
		slices.Contains(c.CalledBalls, c.Board[4][0]) {
		return 1
	}

	return 0
}

func (c *binGOclient) ListenForCalledBalls(conn *websocket.Conn) {
	// Handle incoming messages from the server
	for {
		_, msg, err := conn.Read(c.ctx)

		if err != nil {
			err_msg := fmt.Sprintf("Failed to read message: %s\n", err.Error())
			color.Red(err_msg)
			continue
		}

		ball, err := strconv.Atoi(string(msg))
		if err != nil {
			color.Red("Received invalid message from server: %s", string(msg))
			continue
		}

		c.CalledBalls = append(c.CalledBalls, ball)
		c.PrintBoard()

		if c.CheckForBingo() == 1 {
			color.Green("BINGO! You won!")
			conn.Write(c.ctx, websocket.MessageText, []byte("BINGO!"))
		} else {
			color.Red("No Bingo yet, keep playing!")
		}
	}
}
