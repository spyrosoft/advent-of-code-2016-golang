package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var (
	instructionsInput = "L1, L5, R1, R3, L4, L5, R5, R1, L2, L2, L3, R4, L2, R3, R1, L2, R5, R3, L4, R4, L3, R3, R3, L2, R1, L3, R2, L1, R4, L2, R4, L4, R5, L3, R1, R1, L1, L3, L2, R1, R3, R2, L1, R4, L4, R2, L189, L4, R5, R3, L1, R47, R4, R1, R3, L3, L3, L2, R70, L1, R4, R185, R5, L4, L5, R4, L1, L4, R5, L3, R2, R3, L5, L3, R5, L1, R5, L4, R1, R2, L2, L5, L2, R4, L3, R5, R1, L5, L4, L3, R4, L3, L4, L1, L5, L5, R5, L5, L2, L1, L2, L4, L1, L2, R3, R1, R1, L2, L5, R2, L3, L5, L4, L2, L1, L2, R3, L1, L4, R3, R3, L2, R5, L1, L3, L3, L3, L5, R5, R1, R2, L3, L2, R4, R1, R1, R3, R4, R3, L3, R3, L5, R2, L2, R4, R5, L4, L3, L1, L5, L1, R1, R2, L1, R3, R4, R5, R2, R3, L2, L1, L5"
	x                 = 0
	y                 = 0
	// Directions are 0-3 clockwise
	// 0 = north; 1 = east; 2 = south; 3 = west
	cardinalDirection = 0
	distance          = 0
)

func main() {
	instructions := strings.Split(instructionsInput, ", ")
	followInstructions(instructions)
	calculateDistance()
	fmt.Println("Distance: ", distance)
}

func followInstructions(instructions []string) {
	for _, i := range instructions {
		if i[0] == 'R' {
			turnRight()
		} else {
			turnLeft()
		}
		blocks, _ := strconv.Atoi(i[1:])
		move(blocks)
	}
}

func turnRight() {
	if cardinalDirection == 3 {
		cardinalDirection = 0
	} else {
		cardinalDirection++
	}
}

func turnLeft() {
	if cardinalDirection == 0 {
		cardinalDirection = 3
	} else {
		cardinalDirection--
	}
}

func move(blocks int) {
	if cardinalDirection == 0 {
		y += blocks
	} else if cardinalDirection == 1 {
		x += blocks
	} else if cardinalDirection == 2 {
		y -= blocks
	} else {
		x -= blocks
	}
}

func calculateDistance() {
	distance = int(math.Abs(float64(x)) + math.Abs(float64(y)))
}
