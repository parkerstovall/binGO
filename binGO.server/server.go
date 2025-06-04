package main

import (
	"fmt"

	game_manager "binGOserver/binGO.server.game"
)

func main() {
	gm := game_manager.NewGameManager()
	for {
		ball, error := gm.CallBingoBall()
		if error != nil {
			fmt.Printf("Error: %v\n", error)
			break
		}

		ballText := game_manager.GetBingoBallText(ball)
		fmt.Printf("Bingo ball called: %v\n", ballText)
		fmt.Scanln() // Wait for user input before calling the next ball
	}
}
