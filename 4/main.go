package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"regexp"
	"sort"
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
	lettersFrequency := populateLettersFrequency(letters)
	lettersList := populateLettersList(lettersFrequency)
	lettersList = sortLettersList(lettersList)
	return true
}

func populateLettersFrequency(letters string) map[rune]int {
	sort.Sort([]rune(letters))
	lettersFrequency := make(map[rune]int)
	for _, letter := range letters {
		lettersFrequency[letter]++
	}
	return lettersFrequency
}

type letterCount struct {
	Letter rune
	Count  int
}

func populateLettersList(lettersFrequency map[rune]int) *list.List {
	lettersList := list.New()
	for letter, count := range lettersFrequency {
		fmt.Println(string(letter))
		lettersList.PushFront(letterCount{letter, count})
	}
	return lettersList
}

func sortLettersList(lettersList *list.List) *list.List {
	sortedLettersList := list.New()
	for letterCount := lettersList.Front(); letterCount != nil; letterCount = letterCount.Next() {

	}
	return sortedLettersList
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
