package main

import (
	"strconv"
	"testing"
)

func TestCountTLSAddresses(t *testing.T) {
	input = "test-input.txt"
	correctAddresses := 2
	main()
	if addresses != correctAddresses {
		t.Error("Addresses: " + strconv.Itoa(addresses) + ". Expected:" + strconv.Itoa(correctAddresses) + ".")
	}
}

func TestCheckForTLS(t *testing.T) {
	addresses := []string{
		"abba[mnop]qrst",
		"ioxxoj[asdfgh]zxcvbn",
	}
	for _, address := range addresses {
		if !addressSupportsTLS(address) {
			t.Error("Address: " + address + " is marked incorrectly. It does support TLS.")
		}
	}
}

func TestCheckForNotTLS(t *testing.T) {
	addresses := []string{
		"abcd[bddb]xyyx",
		"aaaa[qwer]tyui",
	}
	for _, address := range addresses {
		if addressSupportsTLS(address) {
			t.Error("Address: " + address + " is marked incorrectly. It does NOT support TLS.")
		}
	}
}
