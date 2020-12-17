package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// ACTIVE represented by '#' is an active square
const ACTIVE = '#'

// INACTIVE represented by '.' is an inactive square
const INACTIVE = '.'

func readFileToLines(path string) ([]string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

// Grid resresents an infinite hypercube matrix, represented by a sparse map
type Grid struct {
	grid map[int]map[int]map[int]map[int]rune
	wMax int
	wMin int
	zMax int
	zMin int
	yMax int
	yMin int
	xMax int
	xMin int
}

// NewGrid initializes the 4D grid.
func NewGrid() Grid {
	g := Grid{}
	g.grid = make(map[int]map[int]map[int]map[int]rune)
	g.wMax = 0
	g.wMin = 0
	g.zMax = 0
	g.zMin = 0
	g.yMax = 0
	g.yMin = 0
	g.xMax = 0
	g.xMin = 0

	return g
}

func (g Grid) setPosition(w int, z int, y int, x int, val rune) {
	if _, ok := g.grid[w]; !ok {
		g.grid[w] = make(map[int]map[int]map[int]rune)
	}
	if _, ok := g.grid[w][z]; !ok {
		g.grid[w][z] = make(map[int]map[int]rune)
	}
	if _, ok := g.grid[w][z][y]; !ok {
		g.grid[w][z][y] = make(map[int]rune)
	}
	g.grid[w][z][y][x] = val
}

func (g Grid) getPosition(w int, z int, y int, x int) rune {
	if _, ok := g.grid[w]; !ok {
		return INACTIVE
	}
	if _, ok := g.grid[w][z]; !ok {
		return INACTIVE
	}
	if _, ok := g.grid[w][z][y]; !ok {
		return INACTIVE
	}
	if _, ok := g.grid[w][z][y][x]; !ok {
		return INACTIVE
	}

	return g.grid[w][z][y][x]
}

// During a cycle, all cubes simultaneously change their state according to the following rules:
// If a cube is active and exactly 2 or 3 of its neighbors are also active, the cube remains
// active. Otherwise, the cube becomes inactive.
// If a cube is inactive but exactly 3 of its neighbors are active, the cube becomes active.
// Otherwise, the cube remains inactive.
func (g Grid) calculateSquareNextState(w int, z int, y int, x int) rune {
	// count neighbors
	activeNeighbors := 0
	for _, wOff := range []int{-1, 0, 1} {
		for _, zOff := range []int{-1, 0, 1} {
			for _, yOff := range []int{-1, 0, 1} {
				for _, xOff := range []int{-1, 0, 1} {
					// ignore the current coordinate
					if wOff != 0 || zOff != 0 || yOff != 0 || xOff != 0 {
						otherSquare := g.getPosition(w+wOff, z+zOff, y+yOff, x+xOff)
						if otherSquare == ACTIVE {
							activeNeighbors++
						}
					}
				}
			}
		}
	}

	// Evaluate rules
	// If a cube is active and exactly 2 or 3 of its neighbors are also active,
	// the cube remains active. Otherwise, the cube becomes inactive.
	// If a cube is inactive but exactly 3 of its neighbors are active, the
	// cube becomes active. Otherwise, the cube remains inactive.
	currentSquare := g.getPosition(w, z, y, x)
	if currentSquare == ACTIVE && (activeNeighbors == 2 || activeNeighbors == 3) {
		return ACTIVE
	} else if currentSquare == INACTIVE && activeNeighbors == 3 {
		return ACTIVE
	}
	return INACTIVE
}

func (g Grid) printGrid() {
	for w := g.wMin; w <= g.wMax; w++ {
		for z := g.zMin; z <= g.zMax; z++ {
			fmt.Println("z =", z, "w =", w)
			for y := g.yMin; y <= g.yMax; y++ {
				for x := g.xMin; x <= g.xMax; x++ {
					val := g.getPosition(z, y, x, w)
					if val == INACTIVE {
						fmt.Print(".")
					} else {
						fmt.Print("#")
					}
				}
				fmt.Println()
			}
		}
	}
}

func (g Grid) cycle() Grid {
	nextState := NewGrid()
	nextState.zMax = g.zMax
	nextState.zMin = g.zMin
	nextState.yMax = g.yMax
	nextState.yMin = g.yMin
	nextState.xMax = g.xMax
	nextState.xMin = g.xMin
	nextState.wMin = g.wMin
	nextState.wMax = g.wMax
	for w := nextState.wMin - 1; w <= nextState.wMax+1; w++ {
		for z := nextState.zMin - 1; z <= nextState.zMax+1; z++ {
			for y := nextState.yMin - 1; y <= nextState.yMax+1; y++ {
				for x := nextState.xMin - 1; x <= nextState.xMax+1; x++ {
					// look at the current state. compute the state, set its state
					// in the new grid
					nextSquareState := g.calculateSquareNextState(w, z, y, x)
					if nextSquareState == ACTIVE {
						nextState.setPosition(w, z, y, x, nextSquareState)

						// check if the bounds need updating also
						if w < nextState.wMin {
							nextState.wMin = w
						} else if w > nextState.wMax {
							nextState.wMax = w
						}

						if z < nextState.zMin {
							nextState.zMin = z
						} else if z > nextState.zMax {
							nextState.zMax = z
						}

						if y < nextState.yMin {
							nextState.yMin = y
						} else if y > nextState.yMax {
							nextState.yMax = y
						}

						if x < nextState.xMin {
							nextState.xMin = x
						} else if x > nextState.xMax {
							nextState.xMax = x
						}
					}
				}
			}
		}
	}

	return nextState
}

func (g Grid) getActiveCount() int {
	count := 0
	for _, w := range g.grid {
		for _, z := range w {
			for _, y := range z {
				for _, x := range y {
					if x == ACTIVE {
						count++
					}
				}
			}
		}
	}
	return count
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	lines, err := readFileToLines(filePath)
	if err != nil {
		panic(err)
	}

	grid := NewGrid()
	for y, v := range lines {
		grid.yMax = y

		for x, val := range v {
			grid.xMax = x
			grid.setPosition(0, 0, y, x, val)
		}
	}

	fmt.Println("🎄 Part 1 🎁: 207") // Answer: 207

	fmt.Println("🎄 Part 2 🎁: ") // Answer: 2308

	// six cycles and return count
	for i := 0; i < 6; i++ {
		grid = grid.cycle()
		//fmt.Println("After", i+1, "cycle:")
		//grid.printGrid()
	}
	fmt.Println(grid.getActiveCount())
}
