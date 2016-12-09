package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"strings"
)

var (
	input          = "ugkcyxxp"
	firstPassword  = ""
	secondPassword = ""
)

func main() {
	firstPassword = generateFirstPassword()
	fmt.Println("First Password:", firstPassword)
	secondPassword = generateSecondPassword()
	fmt.Println("Second Password:", secondPassword)
}

func generateFirstPassword() string {
	password := ""
	for i := 0; len(password) < 8 && i >= 0; i++ {
		currentInput := input + strconv.Itoa(i)
		hasher := md5.New()
		io.WriteString(hasher, currentInput)
		hexMD5 := fmt.Sprintf("%x", hasher.Sum(nil))
		if hexMD5[0:5] == "00000" {
			password += string(hexMD5[5])
		}
	}
	return password
}

func generateSecondPassword() string {
	password := "        "
	for i := 0; strings.Contains(password, " ") && i >= 0; i++ {
		currentInput := input + strconv.Itoa(i)
		hasher := md5.New()
		io.WriteString(hasher, currentInput)
		hexMD5 := fmt.Sprintf("%x", hasher.Sum(nil))
		if hexMD5[0:5] == "00000" {
			characterIndex, err := strconv.Atoi(string(hexMD5[5]))
			if err != nil || characterIndex > 7 || password[characterIndex] != ' ' {
				continue
			}
			//fmt.Println("Character Index:", characterIndex)
			password = password[0:characterIndex] + string(hexMD5[6]) + password[characterIndex+1:8]
		}
	}
	return password
}
