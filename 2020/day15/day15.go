package main

import (
	"fmt"
)

// Given the array ints, seed a game. for each iteration process the following
// rules, execute to provided number of iterations providing result.
// Rules:
// Each turn consists of considering the most recently spoken number:
// If that was the first time the number has been spoken, result is 0
// Otherwise, the number had been spoken before; result is how many turns apart
// the number is from when it was previously spoken.
// for example, if 0,3,6 is used to seed => 0,3,3,1,0,4,0,2,0,2,2,1,8,...
func utter(ints []int, iterations int) int {
	uttered := make(map[int]int, 0)
	lastNumber := -1
	for i := 0; i < iterations; i++ {
		//fmt.Println("i:", i, "Number:", lastNumber)
		if i < len(ints) {
			// seed the game from provided array
			uttered[ints[i]] = i
			lastNumber = ints[i]
		} else {
			if lastUttered, ok := uttered[lastNumber]; ok {
				// number has been uttered before
				timeSince := i - 1 - lastUttered
				uttered[lastNumber] = i - 1
				lastNumber = timeSince
			} else {
				// if it hasn't been uttered, player says 0
				uttered[lastNumber] = i - 1
				lastNumber = 0
			}
		}
	}
	return lastNumber
}
func main() {
	input := []int{1, 0, 16, 5, 17, 4}

	fmt.Println("🎄 Part 1 🎁: ") // Answer: 1294
	fmt.Println(utter(input, 2020))

	fmt.Println("🎄 Part 2 🎁: ") // Answer: 573522
	fmt.Println(utter(input, 30000000))
}
