package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func readFileToInts(path string) ([]int, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	ints := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		ints = append(ints, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return ints, nil
}

// Given an array, find the first entry that does not satisfy the constraint
// in the given previous entries, where two values cannot be summed to the
// current entry.
func findFirstInvalidNumber(ints []int, preamble int, previous int) int {
	for i := preamble; i < len(ints); i++ {
		if !canSumTwoToTarget(ints[i-previous:i], ints[i]) {
			return ints[i]
		}
	}
	return -1
}

// Iterate over a given slice and return if it is possible to sum two values to
// a given target int.
func canSumTwoToTarget(ints []int, target int) bool {
	for k, v := range ints {
		for _, w := range ints[k:] {
			if v+w == target && v != w {
				return true
			}
		}
	}
	return false
}

// Given a slice and a target sum, find the inner slice that sums to that value.
func findContiguousSliceWithSum(ints []int, sum int) []int {
	for k := range ints {
		runningSum := 0
		for j, w := range ints[k:] {
			runningSum += w
			if runningSum == sum {
				return ints[k : j+k+1]
			} else if runningSum > sum {
				break
			}
		}
	}
	return []int{}
}

// Given a slice, find the min and max values and sum them.
func sumMinMaxInSlice(ints []int) int {
	min := ints[0]
	max := ints[0]
	for _, v := range ints {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return min + max
}

func main() {
	preamble := 25
	previous := 25

	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	ints, err := readFileToInts(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Print("Part 1: ")
	invalidNumber := findFirstInvalidNumber(ints, preamble, previous)
	fmt.Println(invalidNumber) // 57195069

	fmt.Print("Part 2: ")
	contiguousSlice := findContiguousSliceWithSum(ints, invalidNumber)
	minMaxSum := sumMinMaxInSlice(contiguousSlice)
	fmt.Println(minMaxSum) // 7409241
}
