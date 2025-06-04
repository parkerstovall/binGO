package game_manager

import (
	"fmt"
	"math/rand"
	"slices"
)

type GameManager struct {
	CalledBalls []int
}

func NewGameManager() *GameManager {
	return &GameManager{
		CalledBalls: []int{},
	}
}

func (gm *GameManager) CallBingoBall() (int, error) {
	min := 1
	max := 76

	if len(gm.CalledBalls) >= (max - min) {
		return -1, fmt.Errorf("all balls have been called")
	}

	for {
		ball := rand.Intn(max-min) + min

		// Check if the ball has already been called
		if !slices.Contains(gm.CalledBalls, ball) {
			// If not, return the ball
			gm.CalledBalls = append(gm.CalledBalls, ball)
			return ball, nil
		}
	}
}

// Determine the column based on the ball number
func GetBingoBallText(ball int) string {
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
