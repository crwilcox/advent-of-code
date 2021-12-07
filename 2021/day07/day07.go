package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/utils"
)

func getInput(path string) []int {
	crabs := make([]int, 0)
	lines, _ := utils.ReadFileToLines(path)
	for _, line := range lines {
		for _, v := range strings.Split(line, ",") {
			crab, _ := strconv.Atoi(v)
			crabs = append(crabs, crab)
		}
	}
	crabs = sort.IntSlice(crabs)
	return crabs
}

func part1(path string) int {
	crabs := getInput(path)
	sort.Ints(crabs)

	value := crabs[len(crabs)/2]

	fuel := 0
	for _, crab := range crabs {
		fuel += int(math.Abs(float64(crab - value)))
	}
	return fuel
}

func part2(path string) int {
	crabs := getInput(path)
	bestFuel := math.MaxInt
	for i := 0; i <= crabs[len(crabs)-1]; i++ {
		fuel := 0
		for _, crab := range crabs {
			// fuel rate is the sumation of 1-n for the distance
			distance := int(math.Abs(float64(crab - i)))
			sum := 0
			for i := 1; i <= distance; i++ {
				sum += i
			}
			fuel += sum
		}
		if fuel < bestFuel {
			bestFuel = fuel
		}
	}

	return bestFuel
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
