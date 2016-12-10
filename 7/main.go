package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	input        = "input.txt"
	tlsAddresses = 0
	sslAddresses = 0
)

func main() {
	file, err := os.Open(input)
	panicOnError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processIPAddress(scanner.Text())
	}
	fmt.Println("Addresses Supporting TLS:", tlsAddresses)
	fmt.Println("Addresses Supporting SSL:", sslAddresses)
}

func processIPAddress(address string) {
	if addressSupportsTLS(address) {
		tlsAddresses++
	}
	if addressSupportsSSL(address) {
		sslAddresses++
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

func addressSupportsSSL(address string) bool {
	brackets, noBrackets := parseBracketsAndNoBrackets(address)
	for _, chunk := range brackets {
		if abaExists(chunk, noBrackets) {
			return true
		}
	}
	return false
}

func abaExists(chunk string, noBrackets []string) bool {
	for i := 0; i < len(chunk)-2; i++ {
		if chunk[i] == chunk[i+2] {
			bab := chunk[i+1:i+3] + string(chunk[i+1])
			if babExists(bab, noBrackets) {
				return true
			}
		}
	}
	return false
}

func babExists(bab string, noBrackets []string) bool {
	for _, chunk := range noBrackets {
		if strings.Contains(chunk, bab) {
			return true
		}
	}
	return false
}

func parseBracketsAndNoBrackets(address string) (brackets []string, noBrackets []string) {
	brackets = make([]string, 0)
	noBrackets = make([]string, 0)
	bracketsRegex := regexp.MustCompile("[\\[\\]]")
	bracketsSplit := bracketsRegex.Split(address, -1)
	inBrackets := false
	for _, bracketSplit := range bracketsSplit {
		if inBrackets {
			brackets = append(brackets, bracketSplit)
		} else {
			noBrackets = append(noBrackets, bracketSplit)
		}
		inBrackets = !inBrackets
	}
	return
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
