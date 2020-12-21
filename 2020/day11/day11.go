package main

import (
	"fmt"
	"os"

	"github.com/crwilcox/advent-of-code/2020/utils"
)

const (
	floor    = byte('.')
	empty    = byte('L')
	occupied = byte('#')
)

// helper function to display the grid
func printGrid(grid [][]byte) {
	for _, v := range grid {
		fmt.Println(string(v))
	}
}

func gridsEqual(before [][]byte, after [][]byte) bool {
	for rowIndex, rowValue := range before {
		for colIndex := range rowValue {
			if before[rowIndex][colIndex] != after[rowIndex][colIndex] {
				return false
			}
		}
	}
	return true
}

func countOccupiedSeats(grid [][]byte) int {
	count := 0
	for _, rowValue := range grid {
		for _, square := range rowValue {
			if square == occupied {
				count++
			}
		}
	}
	return count
}

func countOccupiedSeatsAdjacent(grid [][]byte, row int, col int) int {
	rowLength := len(grid)
	colLength := len(grid[0])

	count := 0
	for _, rowShift := range []int{-1, 0, 1} {
		for _, colShift := range []int{-1, 0, 1} {
			if (rowShift != 0 || colShift != 0) &&
				row+rowShift >= 0 && row+rowShift < rowLength &&
				col+colShift >= 0 && col+colShift < colLength &&
				grid[row+rowShift][col+colShift] == occupied {
				count++
			}
		}
	}
	return count
}

// Looks at the next square in a particular direction. Returns true if an
// occupied seat is seen, False if an empty seat (or board edge) seen.
func occupiedInDirection(grid [][]byte, row int, col int, rowChange int, colChange int) bool {
	for {
		row += rowChange
		col += colChange
		if row < 0 || row >= len(grid) {
			return false // out of rows!
		} else if col < 0 || col >= len(grid[0]) {
			return false // out of columns!
		}
		if grid[row][col] == occupied {
			// seat is occupied in direction
			return true
		} else if grid[row][col] == empty {
			// seat is empty in direcion
			return false
		}
		// need to look at the next square...
	}
}

// Count occupied seats in all directions from view of seat. Skip over floor.
func countOccupiedSeatsInDirections(grid [][]byte, row int, col int) int {
	count := 0
	for _, rowShift := range []int{-1, 0, 1} {
		for _, colShift := range []int{-1, 0, 1} {
			if (rowShift != 0 || colShift != 0) &&
				occupiedInDirection(grid, row, col, rowShift, colShift) {
				count++
			}
		}
	}
	return count
}

// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
// If a seat is occupied (#) and four or more seats adjacent to it are also occupied, the seat becomes empty.
// Otherwise, the seat's state does not change.
func applySeatingRulesPart1(grid [][]byte) [][]byte {
	retGrid := make([][]byte, 0)
	for rowIndex, rowValue := range grid {
		retRow := make([]byte, 0)
		for colIndex, square := range rowValue {
			adjacentOccupied := countOccupiedSeatsAdjacent(grid, rowIndex, colIndex)
			newSquare := square
			if square == empty && adjacentOccupied == 0 {
				newSquare = occupied
			} else if square == occupied && adjacentOccupied >= 4 {
				newSquare = empty
			}
			retRow = append(retRow, newSquare)
		}
		retGrid = append(retGrid, retRow)
	}
	return retGrid
}

// If a seat is empty (L) and there are no occupied seats adjacent to it, the seat becomes occupied.
// If a seat is occupied (#) and five or more seats adjacent to it are also occupied, the seat becomes empty.
// Otherwise, the seat's state does not change.
func applySeatingRulesPart2(grid [][]byte) [][]byte {
	retGrid := make([][]byte, 0)
	for rowIndex, rowValue := range grid {
		retRow := make([]byte, 0)
		for colIndex, square := range rowValue {
			adjacentOccupied := countOccupiedSeatsInDirections(grid, rowIndex, colIndex)
			newSquare := square
			if square == empty && adjacentOccupied == 0 {
				newSquare = occupied
			} else if square == occupied && adjacentOccupied >= 5 {
				newSquare = empty
			}
			retRow = append(retRow, newSquare)
		}
		retGrid = append(retGrid, retRow)
	}
	return retGrid
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	grid, err := utils.ReadFileTo2DByteArray(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Print("Part 1: ") // Answer: 2108
	for gridsNotEqual := true; gridsNotEqual; {
		newGrid := applySeatingRulesPart1(grid)
		gridsNotEqual = !gridsEqual(grid, newGrid)
		grid = newGrid
	}
	fmt.Println(countOccupiedSeats(grid))

	fmt.Print("Part 2: ") // Answer: 1897
	grid, err = utils.ReadFileTo2DByteArray(filePath)
	if err != nil {
		panic(err)
	}

	for gridsNotEqual := true; gridsNotEqual; {
		newGrid := applySeatingRulesPart2(grid)
		gridsNotEqual = !gridsEqual(grid, newGrid)
		grid = newGrid
	}
	fmt.Println(countOccupiedSeats(grid))
}
