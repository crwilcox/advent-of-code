package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func readFileToArray(path string) ([]int, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		i, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, err
		}

		lines = append(lines, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// Given an input array, find the product of the two values
// within that sum to 2020
func part1(input []int) int {
	// for each subsequent iteration, you only need to look at a subset of the
	// array, as the previous bits have already been compared
	for i, v := range input {
		for _, w := range input[i+1:] {
			if v+w == 2020 {
				return v * w
			}
		}
	}
	return -1
}

// Given an input array, find the product of the three values
// within that sum to 2020
func part2(input []int) int {
	// for each subsequent iteration, you only need to look at a subset of the
	// array, as the previous bits have already been compared
	for i, v := range input {
		for j, w := range input[i+1:] {
			for _, y := range input[i+j+1:] {
				if v+w+y == 2020 {
					return v * w * y
				}
			}
		}
	}
	return -1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	lines, err := readFileToArray(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Print("Part 1: ")
	fmt.Println(part1(lines)) // 63616
	fmt.Print("Part 2: ")
	fmt.Println(part2(lines)) // 67877784
}
