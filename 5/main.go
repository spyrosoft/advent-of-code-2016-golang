package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
)

var (
	input    = "ugkcyxxp"
	password = ""
)

func main() {
	for i := 0; len(password) < 8 && i >= 0; i++ {
		currentInput := input + strconv.Itoa(i)
		hasher := md5.New()
		io.WriteString(hasher, currentInput)
		hexMD5 := fmt.Sprintf("%x", hasher.Sum(nil))
		if hexMD5[0:5] == "00000" {
			password += string(hexMD5[5])
		}
	}
	fmt.Println("Password:", password)
}
