package main

import (
	"strconv"
	"testing"
)

func TestDecompressedLength(t *testing.T) {
	inputs := map[string]int{
		"ADVENT":            6,
		"A(1x5)BC":          7,
		"(3x3)XYZ":          9,
		"A(2x2)BCD(2x2)EFG": 11,
		"(6x1)(1x3)A":       6,
		"X(8x2)(3x3)ABCY":   18,
		"(3x20)ZFPXXX":      63,
	}
	for input, expectedLength := range inputs {
		data = ""
		decompress(input)
		if len(data) != expectedLength {
			t.Error("Expected length for: " + input + ": " + strconv.Itoa(expectedLength) + " does not match decompressed length: " + strconv.Itoa(len(data)))
		}
	}
}
