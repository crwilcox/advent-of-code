package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/crwilcox/advent-of-code/utils"
)

// findMissingCharacter either returns a missing character or "", and
// potentially remaining to match characters
func findMissingCharacter(line string) (rune, []rune) {
	match := map[rune]rune{'[': ']', '(': ')', '{': '}', '<': '>'}
	closures := make([]rune, 0)
	for _, r := range line {
		switch r {
		case '(', '[', '{', '<':
			closures = append(closures, r)
		case ')', ']', '}', '>':
			open := closures[len(closures)-1]
			closures = closures[:len(closures)-1]
			if match[open] != r {
				return r, nil
			}
		}
	}
	return rune(0), closures
}

func completeLine(remaining []rune) int {
	score := 0
	for i := len(remaining) - 1; i >= 0; i-- {
		switch remaining[i] {
		case '(':
			score = score*5 + 1
		case '[':
			score = score*5 + 2
		case '{':
			score = score*5 + 3
		case '<':
			score = score*5 + 4
		}
	}
	return score
}

func part1(path string) int {
	values := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}
	score := 0
	lines, _ := utils.ReadFileToLines(path)
	for _, line := range lines {
		missing, _ := findMissingCharacter(line)
		score += values[missing]
	}
	return score
}

func part2(path string) int {
	scores := make([]int, 0)
	lines, _ := utils.ReadFileToLines(path)
	for _, line := range lines {
		missing, remaining := findMissingCharacter(line)
		if missing == rune(0) {
			score := completeLine(remaining)
			scores = append(scores, score)
		}
	}

	sort.Ints(scores)
	return scores[len(scores)/2]
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
