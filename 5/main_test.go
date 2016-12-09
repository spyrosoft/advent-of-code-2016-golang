package main

import "testing"

func TestGeneratePasswords(t *testing.T) {
	input = "abc"
	main()
	correctFirstPassword := "18f47a30"
	if firstPassword != correctFirstPassword {
		t.Error("First Password: " + firstPassword + " is not correct. Should be " + correctFirstPassword)
	}
	correctSecondPassword := "05ace8e3"
	//secondPassword = generateSecondPassword()
	if secondPassword != correctSecondPassword {
		t.Error("Second Password: " + secondPassword + " is not correct. Should be " + correctSecondPassword)
	}
}
