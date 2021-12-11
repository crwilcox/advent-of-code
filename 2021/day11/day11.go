package main

import (
	"fmt"
	"os"

	"github.com/crwilcox/advent-of-code/utils"
)

// Returns number of octopi that flashed
func processStep(grid [][]int) int {
	// First, the energy level of each octopus increases by 1.
	for x, rowVals := range grid {
		for y, _ := range rowVals {
			grid[x][y]++
		}
	}

	type pair struct{ x, y int }

	// Any octopus with an energy level greater than 9 flashes. This increases the energy level of
	// all adjacent octopuses by 1, including octopuses that are diagonally adjacent. If this causes
	// an octopus to have an energy level greater than 9, it also flashes. This process continues as
	// long as new octopuses keep having their energy level increased beyond 9. (An octopus can only
	// flash at most once per step.)
	toVisit := make([]pair, 0)
	for x, row := range grid {
		for y, _ := range row {
			toVisit = append(toVisit, pair{x, y})
		}
	}
	flashed := make([][]bool, len(grid))
	for i := 0; i < len(flashed); i++ {
		flashed[i] = make([]bool, len(grid[i]))
	}
	for len(toVisit) > 0 {
		loc := toVisit[0]
		toVisit = toVisit[1:]
		x := loc.x
		y := loc.y

		if grid[x][y] > 9 && !flashed[x][y] {
			flashed[x][y] = true
			for xd := -1; xd <= 1; xd++ {
				for yd := -1; yd <= 1; yd++ {
					if x+xd >= 0 && x+xd < len(grid) && y+yd >= 0 && y+yd < len(grid[0]) {
						grid[x+xd][y+yd] += 1
						toVisit = append(toVisit, pair{x + xd, y + yd})
					}
				}
			}
		}
	}

	// Finally, any octopus that flashed during this step has its energy level set to 0, as it used all of its energy to flash.
	flashes := 0
	for x, rowVals := range grid {
		for y, _ := range rowVals {
			if grid[x][y] > 9 {
				grid[x][y] = 0
				flashes++
			}
		}
	}
	return flashes
}

func part1(path string) int {
	grid, _ := utils.ReadFileTo2DIntArray(path)
	flashCount := 0
	for i := 1; i <= 100; i++ {
		flashCount += processStep(grid)
		// printGrid(grid, i)
	}
	return flashCount
}

func part2(path string) int {
	grid, _ := utils.ReadFileTo2DIntArray(path)
	flashCount := 0
	for i := 1; ; i++ {
		flashCount = processStep(grid)
		// printGrid(grid, i)
		if flashCount == len(grid)*len(grid[0]) {
			return i
		}
	}
}

func printGrid(grid [][]int, i int) {
	if i < 10 || i%10 == 0 {
		fmt.Println("After step ", i, ":")
		for _, row := range grid {
			fmt.Println(row)
		}
	}
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
