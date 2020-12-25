package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/crwilcox/advent-of-code/2020/utils"
)

// The handshake used by the card and the door involves an operation that
// transforms a subject number. To transform a subject number, start with the
// value 1. Then, a number of times called the loop size, perform the following steps:
// Set the value to itself multiplied by the subject number.
// Set the value to the remainder after dividing the value by 20201227.
func getLoopSize(subjectNumber int, publicKey int) int {
	modulus := 20201227
	value := 1
	loopSize := 0
	for ; value != publicKey; loopSize++ {
		value *= subjectNumber
		value = value % modulus
	}

	return loopSize
}

func produceEncryptionKey(subjectNumber int, loopSize int) int {
	modulus := 20201227
	value := 1
	for i := 0; i < loopSize; i++ {
		value *= subjectNumber
		value = value % modulus
	}

	return value
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	lines, err := utils.ReadFileToLines(filePath)
	if err != nil {
		panic(err)
	}

	doorPk, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}
	cardPk, err := strconv.Atoi(lines[1])
	if err != nil {
		panic(err)
	}

	subjectNumber := 7

	fmt.Println("🎄 This trip is nearly over! 🎁: ") // Answer: 9714832
	// derive the encryption key by finding the door loopsize and using the
	// provided card pk.
	doorLoopSize := getLoopSize(subjectNumber, doorPk)
	encryptionKey := produceEncryptionKey(cardPk, doorLoopSize)
	fmt.Println(encryptionKey)

}
