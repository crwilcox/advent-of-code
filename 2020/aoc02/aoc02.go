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

type PasswordEntry struct {
	Minimum   int
	Maximum   int
	Character string
	Password  string
}

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

func parseLine(line string) PasswordEntry {
	r := regexp.MustCompile(`(?P<Min>\d+)-(?P<Max>\d+) (?P<Char>.): (?P<Password>.+)`)
	matches := r.FindStringSubmatch(line)
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

	return PasswordEntry{minimum, maximum, character, password}
}

func part1(input []string) int {
	valid := 0
	for _, v := range input {
		entry := parseLine(v)
		count := strings.Count(entry.Password, entry.Character)
		if count >= entry.Minimum && count <= entry.Maximum {
			valid += 1
		}
	}
	return valid
}

func part2(input []string) int {
	valid := 0
	for _, v := range input {
		entry := parseLine(v)
		position1 := entry.Minimum
		position2 := entry.Maximum
		character := entry.Character[0]
		password := entry.Password

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
	fmt.Println(part1(lines)) // 378
	fmt.Print("Part 2: ")
	fmt.Println(part2(lines)) // 280
}
