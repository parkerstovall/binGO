package main

import (
	"fmt"
)

func main() {
	ballsCalled := []int{}
	for {
		ball := callBingoBall(ballsCalled)
		if ball == -1 {
			fmt.Println("All balls have been called.")
			break
		}
		ballsCalled = append(ballsCalled, ball)
		ballText := getBingoBallText(ball)
		fmt.Printf("Bingo ball called: %v\n", ballText)
		fmt.Scanln() // Wait for user input before calling the next ball
	}
}
