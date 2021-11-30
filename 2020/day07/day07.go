package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/utils"
)

type bag struct {
	name  string
	count int
}

func readFileToSets(path string) (map[string][]bag, error) {
	lines, err := utils.ReadFileToLines(path)
	if err != nil {
		return nil, err
	}

	bags := make(map[string][]bag)
	for _, line := range lines {
		// Example Lines:
		// vibrant plum bags contain 5 faded blue bags, 1 dotted black bag.
		// faded blue bags contain no other bags.
		// split on 'contain', left side is bag names and counts
		split := strings.Split(line, " contain ")
		bagName := strings.TrimSuffix(split[0], " bags")
		for _, v := range strings.Split(split[1], ", ") {
			v = strings.TrimSuffix(v, ".")
			v = strings.TrimSuffix(v, " bags")
			v = strings.TrimSuffix(v, " bag")

			if v == "no other" {
				// do nothing, this indicates zero bags
			} else {
				b := bag{}
				b.count, _ = strconv.Atoi(v[0:1])
				if err != nil {
					return nil, err
				}
				b.name = v[2:]
				bags[bagName] = append(bags[bagName], b)
			}
		}
	}

	return bags, nil
}

func arrayContainsBag(arr []bag, bagName string) bool {
	for _, v := range arr {
		if v.name == bagName {
			return true
		}
	}
	return false
}

func canBagHoldShinyGoldBag(bag string, allBags map[string][]bag) bool {
	if arrayContainsBag(allBags[bag], "shiny gold") {
		// this bag contains a shiny gold bag directly.
		return true
	}

	// then for each bag in the bag, check if one of them can hold a shiny gold bag.
	for _, v := range allBags[bag] {
		if canBagHoldShinyGoldBag(v.name, allBags) {
			// this bag contains a shiny gold bag indirectly.
			return true
		}
	}

	// Seems this bag can't hold a shiny gold bag
	return false
}

// Given a bag and the ruleset of bags, return the number of bags within a bag
func countInnerBags(bag string, allBags map[string][]bag) int {
	sum := 0
	for _, v := range allBags[bag] {
		// Add each bag listed and it's inner bags
		sum += v.count * (1 + countInnerBags(v.name, allBags))
	}
	return sum
}

// Evaluate how many bags can eventually contain at least one shiny gold bag
func part1(allBags map[string][]bag) int {
	sum := 0
	for bagName := range allBags {
		// evaluate this bag.
		canIt := canBagHoldShinyGoldBag(bagName, allBags)
		if canIt {
			sum++
		}
	}
	return sum
}

// Evaluate how many individual bags are required inside a shiny gold bag
func part2(bags map[string][]bag) int {
	return countInnerBags("shiny gold", bags)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	bags, err := readFileToSets(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Print("Part 1: ")
	fmt.Println(part1(bags)) // 208

	fmt.Print("Part 2: ")
	fmt.Println(part2(bags)) // 1664
}
