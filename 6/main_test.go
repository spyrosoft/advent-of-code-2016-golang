package main

import "testing"

func TestGeneratePasswords(t *testing.T) {
	input = "test-input.txt"
	main()
	correctMessage := "easter"
	if ungarbledMessage != correctMessage {
		t.Error("Ungarbled Message: " + ungarbledMessage + " is not correct. Should be " + correctMessage)
	}
}
