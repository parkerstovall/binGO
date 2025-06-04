package main

import (
	"fmt"
	"math/rand"
	"slices"
)

func callBingoBall(previousBalls []int) int {
	min := 1
	max := 76

	if len(previousBalls) >= (max - min) {
		return -1 // All balls have been called
	}

	for {
		ball := rand.Intn(max-min) + min

		// Check if the ball has already been called
		if !slices.Contains(previousBalls, ball) {
			// If not, return the ball
			return ball
		}
	}
}

// Determine the column based on the ball number
func getBingoBallText(ball int) string {
	if ball < 1 || ball > 75 {
		return "Invalid ball number"
	} else if ball <= 15 {
		return fmt.Sprintf("B%d", ball)
	} else if ball <= 30 {
		return fmt.Sprintf("I%d", ball)
	} else if ball <= 45 {
		return fmt.Sprintf("N%d", ball)
	} else if ball <= 60 {
		return fmt.Sprintf("G%d", ball)
	}

	return fmt.Sprintf("O%d", ball)
}
