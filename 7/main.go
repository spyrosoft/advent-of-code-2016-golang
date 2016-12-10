package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	input     = "input.txt"
	addresses = 0
)

func main() {
	file, err := os.Open(input)
	panicOnError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processIPAddress(scanner.Text())
	}
	fmt.Println(addresses)
}

func processIPAddress(address string) {
	if addressSupportsTLS(address) {
		addresses++
	}
}

func addressSupportsTLS(address string) bool {
	supportsTLS := false
	insideSquareBrackets := false
	for i := 0; i < len(address)-3; i++ {
		if address[i+3] == '[' {
			insideSquareBrackets = true
			i += 3
			continue
		}
		if address[i+3] == ']' {
			insideSquareBrackets = false
			i += 3
			continue
		}
		if address[i] == address[i+3] && address[i+1] == address[i+2] && address[i] != address[i+1] {
			if insideSquareBrackets {
				return false
			} else {
				supportsTLS = true
			}
		}
	}
	return supportsTLS
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
