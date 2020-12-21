package main

import (
	"fmt"
	"os"

	"github.com/crwilcox/advent-of-code/2020/utils"
)

// Determine the number of trees you would encounter if you traversed it with
// the slope or right and down.
func treesOnSlope(input [][]byte, right int, down int) int {
	TREE := byte("#"[0])

	maxRowIndex := len(input[0])

	treeCount := 0
	currCol := 0
	for currRow := 0; currRow < len(input); currRow += down {
		if input[currRow][currCol] == TREE {
			treeCount++
		}
		currCol = (currCol + right) % (maxRowIndex)
	}

	return treeCount
}

// Determine the number of trees you would encounter if you traversed with a
// slope of  Right 3, down 1.
func part1(input [][]byte) int {
	return treesOnSlope(input, 3, 1)
}

// Determine the number of trees you would encounter if, for each of the
// following slopes, you start at the top-left corner and traverse the map all
// the way to the bottom:
// Right 1, down 1.
// Right 3, down 1. (This is the slope you already checked.)
// Right 5, down 1.
// Right 7, down 1.
// Right 1, down 2.
func part2(input [][]byte) int {
	s1 := treesOnSlope(input, 1, 1)
	s2 := treesOnSlope(input, 3, 1)
	s3 := treesOnSlope(input, 5, 1)
	s4 := treesOnSlope(input, 7, 1)
	s5 := treesOnSlope(input, 1, 2)
	return s1 * s2 * s3 * s4 * s5
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	lines, err := utils.ReadFileTo2DByteArray(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Print("Part 1: ")
	fmt.Println(part1(lines)) // 272
	fmt.Print("Part 2: ")
	fmt.Println(part2(lines)) // 3898725600
}
