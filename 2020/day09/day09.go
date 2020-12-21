package main

import (
	"fmt"
	"os"

	"github.com/crwilcox/advent-of-code/2020/utils"
)

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
	ints, err := utils.ReadFileToIntArray(filePath)
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
