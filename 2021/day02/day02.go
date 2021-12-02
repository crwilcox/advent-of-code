package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/utils"
)

func part1(path string) int {
	lines, _ := utils.ReadFileToLines(path)

	horizontal := 0
	depth := 0
	for _, line := range lines {
		stringParts := strings.Split(line, " ")
		direction := stringParts[0]
		amount, _ := strconv.Atoi(stringParts[1])
		switch direction {
		case "forward":
			horizontal += amount
		case "down":
			depth += amount
		case "up":
			depth -= amount
		}
	}

	return horizontal * depth
}

func part2(path string) int {
	lines, _ := utils.ReadFileToLines(path)

	horizontal := 0
	depth := 0
	aim := 0
	for _, line := range lines {
		stringParts := strings.Split(line, " ")
		direction := stringParts[0]
		amount, _ := strconv.Atoi(stringParts[1])
		switch direction {
		case "forward":
			horizontal += amount
			depth += aim * amount
		case "down":
			aim += amount
		case "up":
			aim -= amount
		}
	}

	return horizontal * depth
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
