package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	input             = "input.txt"
	ungarbledMessage  = ""
	letterFrequencies []map[rune]int
)

func main() {
	file, err := os.Open(input)
	panicOnError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processGarbledMessages(scanner.Text())
	}
	ungarbleMessage()
	fmt.Println("Ungarbled Message:", ungarbledMessage)
}

func processGarbledMessages(message string) {
	if letterFrequencies == nil {
		letterFrequencies = make([]map[rune]int, len(message))
		for i := 0; i < len(message); i++ {
			letterFrequencies[i] = make(map[rune]int)
		}
	}
	for i := 0; i < len(message); i++ {
		letterFrequencies[i][rune(message[i])]++
	}
}

func ungarbleMessage() {
	ungarbledMessage = ""
	for _, letterFrequency := range letterFrequencies {
		maxLetterFrequency := 0
		mostCommonLetter := '!'
		for letter, frequency := range letterFrequency {
			if frequency > maxLetterFrequency {
				mostCommonLetter = letter
				maxLetterFrequency = frequency
			}
		}
		ungarbledMessage += string(mostCommonLetter)
	}
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
