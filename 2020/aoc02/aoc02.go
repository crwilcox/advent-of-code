package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func readFileToArray(path string) ([]string, error) {
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
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func part1(input []string) int {
	valid := 0
	for _, v := range input {
		r := regexp.MustCompile(`(?P<Min>\d+)-(?P<Max>\d+) (?P<Char>.): (?P<Password>.+)`)
		matches := r.FindStringSubmatch(v)
		minimum, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		maximum, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		character := matches[3]
		password := matches[4]

		count := strings.Count(password, character)
		if count >= minimum && count <= maximum {
			valid += 1
		}
	}
	return valid
}

func part2(input []string) int {
	valid := 0
	for _, v := range input {
		// 3-4 h: hhxnh
		r := regexp.MustCompile(`(?P<Pos1>\d+)-(?P<Pos2>\d+) (?P<Char>.): (?P<Password>.+)`)
		matches := r.FindStringSubmatch(v)
		position1, err := strconv.Atoi(matches[1])
		if err != nil {
			panic(err)
		}
		position2, err := strconv.Atoi(matches[2])
		if err != nil {
			panic(err)
		}
		character := matches[3][0]
		password := matches[4]

		if (password[position1-1] == character && password[position2-1] != character) ||
			(password[position1-1] != character && password[position2-1] == character) {
			valid += 1
		}
	}
	return valid
}

func main() {
	lines, err := readFileToArray("/2020/aoc02/input")
	if err != nil {
		panic(err)
	}
	fmt.Print("Part 1: ")
	fmt.Println(part1(lines)) // 63616
	fmt.Print("Part 2: ")
	fmt.Println(part2(lines)) // 67877784
}
