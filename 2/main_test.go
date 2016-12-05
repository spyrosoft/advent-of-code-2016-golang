package main

import "testing"

// 123
// 456
// 789

// 2 1 1
// 2 3 6 9 9
// 8 5 6 9 8
// 5 2 2 2 5

// XX1XX
// X234X
// 56789
// XABCX
// XXDXX

// 5 D B 3

func TestDoorCode(t *testing.T) {
	input = `ULL
RRDDD
LURDL
UUUUD`
	main()
	if squareDoorCode != "1985" {
		t.Error("Expected door code 1985 is not the result:", squareDoorCode)
	}
	if diamondDoorCode != "5DB3" {
		t.Error("Expected door code 5DB3 is not the result:", diamondDoorCode)
	}
}
