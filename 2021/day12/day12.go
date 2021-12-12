package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/crwilcox/advent-of-code/utils"
)

func findPaths(caveMap map[string][]string, path []string) int {
	lastPlace := path[len(path)-1]

	if lastPlace == "end" {
		// fmt.Println(path)
		return 1
	}

	found := 0
	for _, v := range caveMap[lastPlace] {
		// if lowercae we don't visit again
		visited := false
		if unicode.IsLower(rune(v[0])) {
			visited = utils.StringSliceContains(path, v)
		}
		if !visited {
			found += findPaths(caveMap, append(path, v))
		}
	}
	return found
}

func findPathsVisitTwice(caveMap map[string][]string, path []string, smallVisitedTwice bool) int {
	lastPlace := path[len(path)-1]

	if lastPlace == "end" {
		//fmt.Println(path)
		return 1
	}

	found := 0
	for _, v := range caveMap[lastPlace] {
		// if lowercae we don't visit again
		visited := false
		if v != "start" && unicode.IsLower(rune(v[0])) {
			visited = utils.StringSliceContains(path, v)
		}
		if visited && !smallVisitedTwice {
			// visited, but we haven't visited a node twice, so we can do this once.
			found += findPathsVisitTwice(caveMap, append(path, v), true)
		} else if !visited && v != "start" {
			found += findPathsVisitTwice(caveMap, append(path, v), smallVisitedTwice)
		}
	}
	return found
}

func generateCaveMap(path string) map[string][]string {
	caveMap := make(map[string][]string)
	edges, _ := utils.ReadFileToLines(path)

	for _, edge := range edges {
		points := strings.Split(edge, "-")
		orig := points[0]
		dest := points[1]

		if !utils.StringSliceContains(caveMap[orig], dest) {
			caveMap[orig] = append(caveMap[orig], dest)
		}
		if !utils.StringSliceContains(caveMap[dest], orig) {
			caveMap[dest] = append(caveMap[dest], orig)
		}
	}
	return caveMap
}
func part1(path string) int {
	caveMap := generateCaveMap(path)
	pathCount := findPaths(caveMap, []string{"start"})
	return pathCount
}

func part2(path string) int {
	caveMap := generateCaveMap(path)
	pathCount := findPathsVisitTwice(caveMap, []string{"start"}, false)
	return pathCount
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
