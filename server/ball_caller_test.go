package main

import (
	"slices"
	"testing"
)

func TestCallBingoBall(t *testing.T) {
	ballsCalled := []int{}
	for range 75 {
		ball := callBingoBall(ballsCalled)
		if ball == -1 {
			t.Error("Expected a valid ball number, but got -1 indicating all balls have been called.")
			return
		}

		if slices.Contains(ballsCalled, ball) {
			t.Errorf("Ball %d has already been called. Previous balls: %v", ball, ballsCalled)
			return
		}

		ballsCalled = append(ballsCalled, ball)
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
		result := getBingoBallText(test.ball)
		if result != test.expected {
			t.Errorf("Expected %s for ball %d, but got %s", test.expected, test.ball, result)
		}
	}
}
