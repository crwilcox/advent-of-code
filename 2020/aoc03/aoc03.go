package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func readFileToArray(path string) ([][]byte, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]byte
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)

		var runeSlice = []byte(line)
		if len(line) > 0 {
			lines = append(lines, runeSlice)
		}
	}

	return lines, nil
}

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
	lines, err := readFileToArray("/2020/aoc03/input")
	if err != nil {
		panic(err)
	}
	fmt.Print("Part 1: ")
	fmt.Println(part1(lines)) // 272
	fmt.Print("Part 2: ")
	fmt.Println(part2(lines)) // 3898725600
}
