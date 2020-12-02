package main

import (
	"bufio"
	"fmt"
	"io"
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
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		i, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, err
		}

		lines = append(lines, i)
	}

	return lines, nil
}

// Given an input array, find the product of the two values
// within that sum to 2020
func part1(input []int) int {
	for i, v := range input {
		for j, w := range input {
			if i != j {
				if v+w == 2020 {
					return v * w
				}
			}
		}
	}
	return -1
}

// Given an input array, find the product of the three values
// within that sum to 2020
func part2(input []int) int {
	for i, v := range input {
		for j, w := range input {
			for k, y := range input {
				if i != j && j != k {
					if v+w+y == 2020 {
						return v * w * y
					}
				}
			}
		}
	}
	return -1
}

func main() {
	lines, err := readFileToArray("/2020/aoc01/input")
	if err != nil {
		panic(err)
	}
	fmt.Print("Part 1: ")
	fmt.Println(part1(lines)) // 63616
	fmt.Print("Part 2: ")
	fmt.Println(part2(lines)) // 67877784
}
