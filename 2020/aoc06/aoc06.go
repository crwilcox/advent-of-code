package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type group struct {
	size    int
	answers map[rune]int
}

// Given custom declaration forms, take groups, collapse to a set with counts
// of responses per group tallied.
func readFileToSets(path string) ([]group, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var groupSets []group
	currGroup := group{}
	currGroup.answers = make(map[rune]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		// if there is a blankline, that should be considered the end of a group
		if len(line) <= 0 {
			groupSets = append(groupSets, currGroup)
			currGroup = group{}
			currGroup.answers = make(map[rune]int)
		} else {
			currGroup.size++
			// parse an individual response adding responses to the group set
			for _, v := range line[:] {
				if val, ok := currGroup.answers[v]; ok {
					currGroup.answers[v] = val + 1
				} else {
					currGroup.answers[v] = 1
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// On the final iteration, we will not get an empty entry, commit here.
	groupSets = append(groupSets, currGroup)

	return groupSets, nil
}

// For each group, count the number of questions to which anyone answered
// "yes". What is the sum of those counts?
func part1(groupSets []group) int {
	sum := 0
	for _, v := range groupSets {
		sum += len(v.answers)
	}
	return sum
}

// For each group, count the number of questions to which everyone answered
//"yes". What is the sum of those counts?
func part2(groupSets []group) int {
	sum := 0

	for _, group := range groupSets {
		groupSize := group.size
		for _, question := range group.answers {
			if question == groupSize {
				sum++
			}
		}
	}
	return sum
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	groupSets, err := readFileToSets(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Print("Part 1: ")
	fmt.Println(part1(groupSets)) // 6809

	fmt.Print("Part 2: ")
	fmt.Println(part2(groupSets)) // 3394
}
