package main

import (
	"strings"
	"testing"
)

func TestFollowInstructions1(t *testing.T) {
	resetVariableStates()
	followExampleInstructions("R2, L3", 5, t)
}

func TestFollowInstructions2(t *testing.T) {
	resetVariableStates()
	followExampleInstructions("R2, R2, R2", 2, t)
}

func TestFollowInstructions3(t *testing.T) {
	resetVariableStates()
	followExampleInstructions("R5, L5, R5, R3", 12, t)
}

func TestFollowInstructions4(t *testing.T) {
	resetVariableStates()
	followExampleInstructions("R1, R1, R1, R1", 0, t)
}

func TestFollowInstructions5(t *testing.T) {
	resetVariableStates()
	followExampleInstructions("L1, L1, L1, L1", 0, t)
}

func resetVariableStates() {
	x = 0
	y = 0
	cardinalDirection = 0
	distance = 0
}

func followExampleInstructions(testInput string, expectedDistance int, t *testing.T) {
	instructions := strings.Split(testInput, ", ")
	followInstructions(instructions)
	calculateDistance()
	if distance != expectedDistance {
		t.Error("Expected distance", expectedDistance, "is not equal to", distance)
	}
}
