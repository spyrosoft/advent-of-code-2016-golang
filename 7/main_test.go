package main

import (
	"strconv"
	"testing"
)

func TestCountTLSAddresses(t *testing.T) {
	input = "test-input.txt"
	correctAddresses := 2
	main()
	if tlsAddresses != correctAddresses {
		t.Error("Addresses: " + strconv.Itoa(tlsAddresses) + ". Expected:" + strconv.Itoa(correctAddresses) + ".")
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

func TestCheckForSSL(t *testing.T) {
	addresses := []string{
		"aba[bab]xyz",
		"aaa[kek]eke",
		"zazbz[bzb]cdb",
		"nfvkyixiivbqfuqpgae[jfnytsqtgdvuspfj]wfziiabvnejeswdfu[mvulpnuojuhdbljoj]hteozxzmoyyozjgmvi[lqytlkseljwqsthg]nwnvpacojowhtywdcgue",
	}
	for _, address := range addresses {
		if !addressSupportsSSL(address) {
			t.Error("Address: " + address + " is marked incorrectly. It does supports SSL.")
		}
	}
}

func TestCheckForNotSSL(t *testing.T) {
	addresses := []string{
		"xyx[xyx]xyx",
	}
	for _, address := range addresses {
		if addressSupportsSSL(address) {
			t.Error("Address: " + address + " is marked incorrectly. It does NOT support SSL.")
		}
	}
}
