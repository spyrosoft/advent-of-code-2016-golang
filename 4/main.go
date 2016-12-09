package main

import (
	"bufio"
	"fmt"
	"math"
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
		processRoom(scanner.Text())
	}
	fmt.Println("Sector IDs Sum:", sectorIDsSum)
}

func processRoom(room string) {
	lettersSectorIDChecksumRegex := regexp.MustCompile("^([^\\d]+)(\\d+)\\[(.+)\\]$")
	regexMatches := lettersSectorIDChecksumRegex.FindAllStringSubmatch(room, -1)
	//fmt.Println(regexMatches)
	letters := regexMatches[0][1]
	sectorID, _ := strconv.Atoi(regexMatches[0][2])
	checksum := regexMatches[0][3]
	addSectorIDIfRealRoom(letters, sectorID, checksum)
	decipherRoom(letters, sectorID)
}

func addSectorIDIfRealRoom(letters string, sectorID int, checksum string) {
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
	correctChecksum := ""
	for _, letterCount := range letterCounts {
		if len(correctChecksum) == 5 || letterCount.Count == 0 {
			break
		}
		//fmt.Println(string(letterCount.Letter), letterCount.Count)
		correctChecksum += string(letterCount.Letter)
	}
	return correctChecksum == checksum
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
	letterCounts := make(letterCounts, len(lettersFrequencyMap)/2)
	for letter, count := range lettersFrequencyMap {
		//fmt.Println("letter:", string(letter))
		letterCount := letterCount{letter, count}
		letterCounts = append(letterCounts, letterCount)
	}
	return letterCounts
}

// Part 2
func decipherRoom(room string, sectorID int) string {
	// a: 97, z: 122
	maxLetterValue := 122
	rotation := int(math.Mod(float64(sectorID), float64(26)))
	//fmt.Println("Rotation: ", rotation)
	decipheredRoom := ""
	for i := 0; i < len(room); i++ {
		if room[i] == '-' {
			decipheredRoom += " "
			continue
		}
		letter := int(room[i])
		letter += rotation
		if letter > maxLetterValue {
			letter -= 26 // 26 letters in the alphabet
		}
		decipheredRoom += string(rune(letter))
	}
	fmt.Println("Deciphered Room:", decipheredRoom, " - Sector ID:", sectorID)
	return decipheredRoom
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
