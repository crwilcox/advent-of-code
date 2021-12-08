package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/utils"
)

func part1(path string) int {
	lines, _ := utils.ReadFileToLines(path)
	uniqueDigit := 0
	for _, line := range lines {
		// 1 => 2 segments
		// 4 => 4 segments
		// 7 => 3 segments
		// 8 => 7 segments

		split := strings.Split(line, " | ")
		// inputEnc := split[0]
		outputEnc := split[1]

		for _, v := range strings.Split(outputEnc, " ") {
			switch len(v) {
			case 2:
				uniqueDigit++
			case 3:
				uniqueDigit++
			case 4:
				uniqueDigit++
			case 7:
				uniqueDigit++
			}
		}
	}
	return uniqueDigit
}

func valuesInCommon(a []byte, b []byte) int {
	inCommon := 0

	for _, c := range a {
		for _, d := range b {
			if c == d {
				inCommon++
				break
			}
		}
	}
	return inCommon
}

func decode(line string) int {
	knownValues := make(map[int][]byte)

	// Note: initially I thought you may need to make a second 'pass' in case'
	// 1 and 4, which are ultimately used to figure out the other digits,
	// came later, but since we look at all input (has 1, 4) before output,
	// this can be done in a single loop, provided we also look at the outputs
	// after.
	split := strings.Split(line, " | ")
	inputEnc := split[0]
	outputEnc := split[1]

	entries := strings.Split(inputEnc+" "+outputEnc, " ")
	for _, s := range entries {
		v := []byte(s)
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})

		switch len(v) {
		case 2: // 1
			knownValues[1] = v
		case 3: // 7
			knownValues[7] = v
		case 4: // 4
			knownValues[4] = v
		case 5: // 2, 3, 5
			if len(knownValues[4]) > 0 && len(knownValues[1]) > 0 {
				if valuesInCommon(knownValues[4], v) == 3 && valuesInCommon(knownValues[1], v) == 2 {
					// 3: shares 4 segments with 4
					knownValues[3] = v
				} else if valuesInCommon(knownValues[4], v) == 3 {
					// 5 shares 3 segments with 4
					knownValues[5] = v
				} else if valuesInCommon(knownValues[1], v) == 1 {
					// 2: shares one segment with 1,
					knownValues[2] = v
				}
			}
		case 6: // 0, 6, or 9
			if len(knownValues[4]) > 0 && len(knownValues[1]) > 0 {
				if valuesInCommon(knownValues[4], v) == 4 {
					// 9:  has all of 4's segments (4)
					knownValues[9] = v
				} else if valuesInCommon(knownValues[1], v) == 2 {
					// 0: contains all of one.
					knownValues[0] = v
				} else {
					// 6: otherwise
					knownValues[6] = v
				}
			}
		case 7: // 8
			knownValues[8] = v
		}
	}

	reverseMap := make(map[string]string)
	for i := 0; i < 10; i++ {
		if entry, ok := knownValues[i]; ok {
			b := []byte(entry)
			reverseMap[string(b)] = fmt.Sprint(i)
		}
	}

	outVal := ""
	for _, position := range strings.Split(outputEnc, " ") {
		b := []byte(position)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})

		outVal += reverseMap[string(b)]
	}
	n, _ := strconv.Atoi(outVal)
	return n
}

func part2(path string) int {
	sum := 0
	lines, _ := utils.ReadFileToLines(path)
	for _, line := range lines {
		n := decode(line)
		sum += n

	}
	return sum
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
