package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/crwilcox/advent-of-code/utils"
)

func isLowSpot(grid [][]int, x int, y int) bool {
	point := grid[x][y]

	if x > 0 && grid[x-1][y] <= point {
		return false
	}
	if x+1 < len(grid) && grid[x+1][y] <= point {
		return false
	}
	if y > 0 && grid[x][y-1] <= point {
		return false
	}
	if y+1 < len(grid[0]) && grid[x][y+1] <= point {
		return false
	}

	return true
}

type pair struct {
	x int
	y int
}

func part1(path string) int {
	grid, _ := utils.ReadFileTo2DIntArray(path)
	summedRiskLevels := 0
	for x, row := range grid {
		for y, position := range row {
			// risk level is 1 plus height.
			if isLowSpot(grid, x, y) {
				summedRiskLevels += position + 1
			}
		}
	}
	return summedRiskLevels

}

// returns count of spots found in basin
func discoverBasin(grid [][]int, x int, y int) int {
	if grid[x][y] < 0 || grid[x][y] >= 9 {
		// visited
		return 0
	}
	visitedCount := 1
	grid[x][y] = -1

	if x > 0 {
		visitedCount += discoverBasin(grid, x-1, y)
	}
	if x+1 < len(grid) {
		visitedCount += discoverBasin(grid, x+1, y)
	}
	if y > 0 {
		visitedCount += discoverBasin(grid, x, y-1)
	}
	if y+1 < len(grid[0]) {
		visitedCount += discoverBasin(grid, x, y+1)
	}
	return visitedCount
}

func part2(path string) int {
	grid, _ := utils.ReadFileTo2DIntArray(path)
	summedRiskLevels := 0
	lowSpots := make([]pair, 0)
	for x, row := range grid {
		for y, position := range row {
			// risk level is 1 plus height.
			if isLowSpot(grid, x, y) {
				summedRiskLevels += position + 1
				lowSpots = append(lowSpots, pair{x: x, y: y})
			}
		}
	}

	// starting from each low spot, count locations to get basins and sizes
	basins := make([]int, 0)
	for _, lowSpot := range lowSpots {
		basin := discoverBasin(grid, lowSpot.x, lowSpot.y)
		if basin != 0 {
			basins = append(basins, basin)
		}
	}

	// get top 3 basins, multiply
	m := 1
	sort.Ints(basins)
	for _, v := range basins[len(basins)-3:] {
		m *= v
	}
	return m
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
