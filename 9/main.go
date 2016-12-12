package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	input = "input.txt"
	data  = ""
)

func main() {
	file, err := os.Open(input)
	panicOnError(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		decompress(scanner.Text())
	}
	fmt.Println("Length of Decompressed Data:", len(data))
}

func decompress(compressed string) {
	//fmt.Println("Instructions:", compressed)
	for i := 0; i < len(compressed); i++ {
		if compressed[i] == '(' {
			newData, fastForward := parseInstruction(compressed, i+1)
			data += newData
			//fmt.Println("Fast Forward:", fastForward-i)
			i = fastForward
		} else {
			data += string(compressed[i])
		}
	}
	//fmt.Println("Data:", data)
}

func parseInstruction(compressed string, i int) (newData string, fastForward int) {
	distanceString := ""
	for ; compressed[i] != 'x'; i++ {
		distanceString += string(compressed[i])
	}
	distance, _ := strconv.Atoi(distanceString)
	i++

	timesString := ""
	for ; compressed[i] != ')'; i++ {
		timesString += string(compressed[i])
	}
	times, _ := strconv.Atoi(timesString)
	i++

	stringToRepeat := ""
	for repeat := 0; repeat < distance; repeat++ {
		stringToRepeat += string(compressed[i])
		i++
	}

	for repitition := 0; repitition < times; repitition++ {
		newData += stringToRepeat
	}

	//fmt.Println("D:", distance, "T:", times, "I:", i, "Repeat String:", stringToRepeat)

	fastForward = i - 1
	return
}

// func parseInstruction(compressed string, i int) (newData string, fastForward int) {
// 	distanceString := ""
// 	for ; compressed[i] != 'x'; i++ {
// 		distanceString += string(compressed[i])
// 	}
// 	distance, _ := strconv.Atoi(distanceString)
// 	i++

// 	timesString := ""
// 	for ; compressed[i] != ')'; i++ {
// 		timesString += string(compressed[i])
// 	}
// 	times, _ := strconv.Atoi(timesString)
// 	i++

// 	if compressed[i] == '(' {
// 		fastForward = i - 1
// 		return
// 	}

// 	stringToRepeat := ""
// 	for repeat := 0; repeat < distance; repeat++ {
// 		stringToRepeat += string(compressed[i])
// 		i++
// 	}
// 	i -= distance

// 	for repitition := 0; repitition < times; repitition++ {
// 		newData += stringToRepeat
// 		i++
// 	}

// 	fastForward = i
// 	return
// }

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
