package main

import (
	"testing"
)

func compareByteArrays(expected [][]byte, actual [][]byte) bool {
	for x, v := range expected {
		for y, val := range v {
			if val != actual[x][y] {
				return false
			}
		}
	}
	return true
}
func TestGetLoopSize(t *testing.T) {
	pkCard := 5764801
	pkDoor := 17807724
	sn := 7

	cardLoopSize := getLoopSize(sn, pkCard)

	if cardLoopSize != 8 {
		t.Error("Card Loop size want 11, got:", cardLoopSize)

	}

	doorLoopSize := getLoopSize(sn, pkDoor)
	if doorLoopSize != 11 {
		t.Error("Door Loop size want 11, got:", doorLoopSize)
	}

}

func TestProduceEncryptionKey(t *testing.T) {
	// you can use either device's loop size with the other device's public key
	// to calculate the encryption key. Transforming the subject number of 17807724
	// (the door's public key) with a loop size of 8 (the card's loop size)
	// produces the encryption key, 14897079. (Transforming the subject number
	// of 5764801 (the card's public key) with a loop size of 11
	// (the door's loop size) produces the same encryption key: 14897079.)

	got := produceEncryptionKey(17807724, 8)
	if got != 14897079 {
		t.Error("Door: want: 14897079, got:", got)
	}

	got = produceEncryptionKey(5764801, 11)
	if got != 14897079 {
		t.Error("Card: want: 14897079, got:", got)
	}
}
