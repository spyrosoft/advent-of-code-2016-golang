package main

import "testing"

func TestCheckForRealRoom(t *testing.T) {
	realRooms := [][]string{
		{"aaaaa-bbb-z-y-x", "abxyz"},
		{"a-b-c-d-e-f-g-h", "abcde"},
		{"not-a-real-room", "oarel"},
	}
	for _, realRoom := range realRooms {
		if !checkForRealRoom(realRoom[0], realRoom[1]) {
			t.Error("The room '" + realRoom[0] + "[" + realRoom[1] + "]' was erroneously marked. It is A REAL ROOM.")
		}
	}
	fakeRooms := [][]string{
		{"totally-real-room", "decoy"},
	}
	for _, fakeRoom := range fakeRooms {
		if checkForRealRoom(fakeRoom[0], fakeRoom[1]) {
			t.Error("The room '" + fakeRoom[0] + "[" + fakeRoom[1] + "]' was erroneously marked. It is NOT A REAL ROOM.")
		}
	}
}

func TestAddSectorIDsIfReal(t *testing.T) {
	inputFile = "test-input.txt"
	main()
}

func TestDecipherRoom(t *testing.T) {
	room := "qzmt-zixmtkozy-ivhz"
	sectorID := 343
	decipheredRoom := "very encrypted name"
	if decipherRoom(room, sectorID) != decipheredRoom {
		t.Error("The room: " + room + " with sector ID: " + string(sectorID) + " did not decipher to: " + decipheredRoom)
	}
}
