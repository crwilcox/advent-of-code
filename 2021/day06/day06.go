package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/utils"
)

func simulateFish(path string, days int) int {
	lines, _ := utils.ReadFileToLines(path)

	// fish all have value 0 to 9, so represent them as a compressed slice
	// this isn't necessary for part 1, but part 2 makes it more or less req'd
	fish := make([]int, 9)
	for _, line := range lines {
		intervals := strings.Split(line, ",")
		for _, interval := range intervals {
			number, _ := strconv.Atoi(interval)
			fish[number]++
		}
	}

	for d := 0; d < days; d++ {
		// fmt.Println(fish)
		procreators := fish[0]

		// shift procreators off, each other fish decrements its counter
		fish = fish[1:9]
		fish = append(fish, procreators)
		fish[6] += procreators
	}

	sum := 0
	for _, v := range fish {
		sum += v
	}
	return sum
}

func part1(path string) int {
	return simulateFish(path, 80)
}

func part2(path string) int {
	return simulateFish(path, 256)
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
