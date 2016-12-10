package main

import (
	"fmt"
	"testing"
)

func TestRect(t *testing.T) {
	width = 7
	height = 3
	input = "test-input.txt"
	main()
	correctScreen := [][]bool{
		{false, true, false, false, true, false, true},
		{true, false, true, false, false, false, false},
		{false, true, false, false, false, false, false},
	}
	if fmt.Sprint(screen) != fmt.Sprint(correctScreen) {
		t.Error("Computed screen does not match expected screen.")
		fmt.Println("Computed:")
		printScreen(screen)
		fmt.Println()
		fmt.Println("Expected:")
		printScreen(correctScreen)
	}
}

// rect 3x2
// rotate column x=1 by 1
// rotate row y=0 by 4
// rotate column x=1 by 1
