package main

import "testing"

func TestGeneratePassword(t *testing.T) {
	input = "abc"
	main()
	correctPassword := "18f47a30"
	if password != correctPassword {
		t.Error("Password: " + password + " is not correct. Should be " + correctPassword)
	}
}
