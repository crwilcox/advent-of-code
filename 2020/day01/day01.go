package main

import (
	"fmt"
	"os"

	"github.com/crwilcox/advent-of-code/utils"
)

// Given an input array, find the product of the two values
// within that sum to 2020
func part1(input []int) int {
	// for each subsequent iteration, you only need to look at a subset of the
	// array, as the previous bits have already been compared
	for i, v := range input {
		for _, w := range input[i+1:] {
			if v+w == 2020 {
				return v * w
			}
		}
	}
	return -1
}

// Given an input array, find the product of the three values
// within that sum to 2020
func part2(input []int) int {
	// for each subsequent iteration, you only need to look at a subset of the
	// array, as the previous bits have already been compared
	for i, v := range input {
		for j, w := range input[i+1:] {
			for _, y := range input[i+j+1:] {
				if v+w+y == 2020 {
					return v * w * y
				}
			}
		}
	}
	return -1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	lines, err := utils.ReadFileToIntArray(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Print("Part 1: ")
	fmt.Println(part1(lines)) // 63616
	fmt.Print("Part 2: ")
	fmt.Println(part2(lines)) // 67877784
}
