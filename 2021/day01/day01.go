package main

import (
	"fmt"
	"os"

	"github.com/crwilcox/advent-of-code/utils"
)

func part1(path string) int {
	lines, _ := utils.ReadFileToIntArray(path)

	previous := 0
	increased := 0
	for i, line := range lines {
		if i > 0 && line > previous {
			increased++
		}
		previous = line
	}

	return increased
}

func part2(path string) int {
	lines, _ := utils.ReadFileToIntArray(path)

	increased := 0
	for i := range lines {
		firstWindow := lines[i : i+3]
		secondWindow := lines[i+1 : i+4]
		if compareWindows(firstWindow, secondWindow) {
			increased++
		}
	}

	return increased
}

func compareWindows(first []int, second []int) bool {
	firstSum := 0
	for _, v := range first {
		firstSum += v
	}
	secondSum := 0
	for _, v := range second {
		secondSum += v
	}

	return secondSum > firstSum
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]

	fmt.Println("🎄 Part 1 🎁: ")
	part1 := part1(filePath)
	fmt.Println(part1)

	fmt.Println("🎄 Part 2 🎁: ")
	part2 := part2(filePath)
	fmt.Println(part2)
}
