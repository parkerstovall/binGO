package game_manager

import (
	"slices"
	"testing"
)

func TestCallBingoBall(t *testing.T) {
	gm := NewGameManager()

	ballsCalled := []int{}
	for range 75 {
		ball, error := gm.CallBingoBall()
		if error != nil {
			t.Errorf("Unexpected error: %v", error)
			return
		}

		if slices.Contains(ballsCalled, ball) {
			t.Errorf("Ball %d has already been called. Previous balls: %v", ball, ballsCalled)
			return
		}

		ballsCalled = append(ballsCalled, ball)
	}

	// Check if all balls have been called
	_, error := gm.CallBingoBall()
	if error == nil || error.Error() != "all balls have been called" {
		t.Errorf("Expected error 'all balls have been called', but got: %v", error)
	}
}

func TestGetBingoBallText(t *testing.T) {
	tests := []struct {
		ball     int
		expected string
	}{
		{1, "B1"},
		{15, "B15"},
		{16, "I16"},
		{30, "I30"},
		{31, "N31"},
		{45, "N45"},
		{46, "G46"},
		{60, "G60"},
		{61, "O61"},
		{75, "O75"},
	}

	for _, test := range tests {
		result := GetBingoBallText(test.ball)
		if result != test.expected {
			t.Errorf("Expected %s for ball %d, but got %s", test.expected, test.ball, result)
		}
	}
}
