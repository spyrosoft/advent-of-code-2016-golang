package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	input  = "input.txt"
	width  = 50
	height = 6
	screen [][]bool
)

func main() {
	screen = createScreen(width, height)
	file, err := os.Open(input)
	panicOnError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		followInstruction(scanner.Text())
	}
	fmt.Println()
	fmt.Println("Result:")
	printScreen(screen)
	fmt.Println("Total pixels lit", totalPixelsLit(screen))
	fmt.Println()
}

func followInstruction(instruction string) {
	parts := strings.Split(instruction, " ")
	if parts[0] == "rect" {
		widthAndHeight := strings.Split(parts[1], "x")
		width, _ := strconv.Atoi(widthAndHeight[0])
		height, _ := strconv.Atoi(widthAndHeight[1])
		screenRect(width, height)
	} else if parts[0] == "rotate" {
		rotation := strings.Split(parts[2], "=")
		if parts[1] == "column" {
			column, _ := strconv.Atoi(rotation[1])
			amount, _ := strconv.Atoi(parts[4])
			screenRotate(column, amount, "column")
		} else if parts[1] == "row" {
			row, _ := strconv.Atoi(rotation[1])
			amount, _ := strconv.Atoi(parts[4])
			screenRotate(row, amount, "row")
		} else {
			panic("Unknown instruction: " + instruction)
		}
	} else {
		panic("Unknown instruction: " + instruction)
	}
}

func screenRect(width int, height int) {
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			screen[y][x] = true
		}
	}
}

func screenRotate(index int, amount int, columnOrRow string) {
	column, row, x, y := false, false, 0, 0
	var newColumnOrRow []bool
	if columnOrRow == "row" {
		row = true
		x = 0
		y = index
		newColumnOrRow = make([]bool, width)
	} else {
		column = true
		x = index
		y = 0
		newColumnOrRow = make([]bool, height)
	}

	i := 0
	for column && i < height || row && i < width {
		if column {
			rotation := y + amount
			if rotation >= height {
				rotation -= height
			}
			newColumnOrRow[rotation] = screen[i][x]
			y++
		} else {
			rotation := x + amount
			if rotation >= width {
				rotation -= width
			}
			newColumnOrRow[rotation] = screen[y][i]
			x++
		}
		i++
	}
	for i, onOrOff := range newColumnOrRow {
		if column {
			screen[i][x] = onOrOff
		} else {
			screen[y][i] = onOrOff
		}
	}
}

func printScreen(whichScreen [][]bool) {
	for _, y := range whichScreen {
		for _, cell := range y {
			if cell {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func totalPixelsLit(whichScreen [][]bool) int {
	total := 0
	for _, y := range whichScreen {
		for _, cell := range y {
			if cell {
				total++
			}
		}
	}
	return total
}

func createScreen(width int, height int) [][]bool {
	newScreen := make([][]bool, height)
	for i, _ := range newScreen {
		newScreen[i] = make([]bool, width)
	}
	return newScreen
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
