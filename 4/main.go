package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	inputFile    = "input.txt"
	sectorIDsSum = 0
)

func main() {
	file, err := os.Open(inputFile)
	panicOnError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		addSectorIDIfRealRoom(scanner.Text())
	}
}

func addSectorIDIfRealRoom(room string) {
	//Throw away dashes; they don't seem to matter
	room = strings.Replace(room, "-", "", -1)
	lettersSectorIDChecksumRegex := regexp.MustCompile("^([^\\d]+)(\\d+)\\[(.+)\\]$")
	regexMatches := lettersSectorIDChecksumRegex.FindAllStringSubmatch(room, -1)
	fmt.Println(regexMatches)
	letters := regexMatches[0][1]
	sectorID, _ := strconv.Atoi(regexMatches[0][2])
	checksum := regexMatches[0][3]
	if checkForRealRoom(letters, checksum) {
		sectorIDsSum += sectorID
	}
}

func checkForRealRoom(letters string, checksum string) bool {
	lettersCounter := make(map[rune]int)
	for _, letter := range letters {
		lettersCounter[letter]++
	}
	return true
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
