package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func readFileToJoltages(path string) ([]int, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	joltages := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		i, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		joltages = append(joltages, i)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return joltages, nil
}

// Initial Approach to Part 2, while this will *work* it takes far too
// long for a real dataset.
func validChains(remainingArray []int) int {
	if len(remainingArray) <= 1 {
		return 1
	}

	choices := 0
	for k, v := range remainingArray[1:] {
		if v-remainingArray[0] <= 3 {
			choices += validChains(remainingArray[k+1:])
		}
	}

	return choices
}

// Given joltages, calculate how many different joltages are reachable
// given a restrcition of 1-3 jolatge increase
// from there, we can go backwards, figuring out the total number of paths
// through the adapters.
func determinePotentialJoltageSequences(joltages []int) int {
	// Creates a gap list. This tells us how many other adapters are reachable
	// use this to control the next loop adds.
	gaps := make([]int, len(joltages))
	gaps[len(joltages)-1] = 1
	for k1, v1 := range joltages {
		for _, v := range joltages[k1+1:] {
			if v-v1 <= 3 {
				gaps[k1]++
			}
		}
	}
	// [1 1 3 2 1 1 2 1 1 1 1 1 1]
	// Now we can figure out the path to the end by adding the possible path
	// changes, from the end.
	for i := len(gaps) - 2; i >= 0; i-- {
		acc := 0
		for j := 1; j <= gaps[i]; j++ {
			acc += gaps[i+j]
		}
		gaps[i] = acc
	}
	// [8 8 8 4 2 2 2 1 1 1 1 1 1]
	return gaps[0]
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	joltages, err := readFileToJoltages(filePath)
	if err != nil {
		panic(err)
	}

	// Sort the joltages, add the zeroeth element and the +3 element.
	sort.Ints(joltages)
	joltages = append([]int{0}, joltages...)
	joltages = append(joltages, joltages[len(joltages)-1]+3)

	fmt.Print("Part 1: ")
	counts := make([]int, 4)
	for i := 1; i < len(joltages); i++ {
		difference := joltages[i] - joltages[i-1]
		counts[difference]++
	}
	fmt.Println(counts[1] * counts[3]) // 1980

	fmt.Print("Part 2: ")
	// Initial approach, turned out to be far too slow
	// fmt.Println(validChains(joltages))
	// Different approach, don't actually explore the 'graph'
	fmt.Println(determinePotentialJoltageSequences(joltages))
}
