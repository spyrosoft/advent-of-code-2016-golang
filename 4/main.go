package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
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

type letterCount struct {
	Letter rune
	Count  int
}

type letterCounts []letterCount

func (l letterCounts) Len() int      { return len(l) }
func (l letterCounts) Swap(i, j int) { l[i], l[j] = l[j], l[i] }
func (l letterCounts) Less(i, j int) bool {
	if l[i].Count == l[j].Count {
		return int(l[i].Letter) < int(l[j].Letter)
	}
	return int(l[i].Count) > int(l[j].Count)
}

func checkForRealRoom(letters string, checksum string) bool {
	lettersFrequencyMap := populateLettersFrequencyMap(letters)
	letterCounts := populateLetterCounts(lettersFrequencyMap)
	sort.Sort(letterCounts)
	for _, letterCount := range letterCounts {
		fmt.Println(string(letterCount.Letter), letterCount.Count)
	}
	return true
}

func populateLettersFrequencyMap(letters string) map[rune]int {
	lettersFrequencyMap := make(map[rune]int)
	for _, letter := range letters {
		if letter == '-' {
			continue
		}
		lettersFrequencyMap[letter]++
	}
	return lettersFrequencyMap
}

func populateLetterCounts(lettersFrequencyMap map[rune]int) letterCounts {
	letterCounts := make(letterCounts, len(lettersFrequencyMap))
	for letter, count := range lettersFrequencyMap {
		fmt.Println("letter:", string(letter))
		letterCount := letterCount{letter, count}
		letterCounts = append(letterCounts, letterCount)
	}
	return letterCounts
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
