package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/utils"
)

type Grid struct {
	grid map[int]map[int]int
}

func (g *Grid) incrementPosition(x int, y int) {
	if _, ok := g.grid[x]; !ok {
		g.grid[x] = make(map[int]int)
	}
	g.grid[x][y] += 1
}

func (grid Grid) printMap() {
	// Grid on site is y/x, print it that way
	size := 9
	for x := 0; x <= size; x++ {
		for y := 0; y < size; y++ {
			v := grid.grid[y][x]
			if v == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println()
	}
}

func (grid Grid) countOverlappingPositions() int {
	sum := 0
	for _, v := range grid.grid {
		for _, w := range v {
			if w >= 2 {
				sum += 1
			}
		}
	}
	return sum
}

func part1(path string) int {
	lines, _ := utils.ReadFileToLines(path)

	grid := Grid{}
	grid.grid = make(map[int]map[int]int)

	for _, line := range lines {
		// fmt.Println(i, line)
		points := strings.Split(line, " -> ")
		start := strings.Split(points[0], ",")
		end := strings.Split(points[1], ",")
		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])

		// fmt.Println("x1,y1,x2,y2", x1, y1, x2, y2)
		if x1 == x2 { // vertical line
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			for i := y1; i <= y2; i++ {
				//fmt.Println("MarkingV:", x1, i)
				grid.incrementPosition(x1, i)
			}

		} else if y1 == y2 { //horizontal line
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for i := x1; i <= x2; i++ {
				//fmt.Println("MarkingH:", i, y1)
				grid.incrementPosition(i, y1)
			}
		} else {
			// later, for now is fine
			continue
		}
	}
	// grid.printMap()
	return grid.countOverlappingPositions()
}

func part2(path string) int {
	lines, _ := utils.ReadFileToLines(path)

	grid := Grid{}
	grid.grid = make(map[int]map[int]int)

	for _, line := range lines {
		// fmt.Println(line)
		points := strings.Split(line, " -> ")
		start := strings.Split(points[0], ",")
		end := strings.Split(points[1], ",")
		x1, _ := strconv.Atoi(start[0])
		y1, _ := strconv.Atoi(start[1])
		x2, _ := strconv.Atoi(end[0])
		y2, _ := strconv.Atoi(end[1])

		// fmt.Println("x1,y1,x2,y2", x1, y1, x2, y2)
		if x1 == x2 { // vertical line
			if y1 > y2 {
				y1, y2 = y2, y1
			}

			for i := y1; i <= y2; i++ {
				//fmt.Println("MarkingV:", x1, i)
				grid.incrementPosition(x1, i)
			}

		} else if y1 == y2 { //horizontal line
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for i := x1; i <= x2; i++ {
				//fmt.Println("MarkingH:", i, y1)
				grid.incrementPosition(i, y1)
			}
		} else {
			// diagonal line at exactly 45 degrees.
			if x1 > x2 {
				x1, x2 = x2, x1
				y1, y2 = y2, y1
			}

			y := y1
			for x := x1; x <= x2; x++ {
				//fmt.Println("MarkingD:", x, y)
				grid.incrementPosition(x, y)
				if y1 < y2 {
					y++
				} else {
					y--
				}
			}
		}
	}
	// grid.printMap()
	return grid.countOverlappingPositions()
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
