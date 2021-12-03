package main

import (
	"fmt"
	"os"

	"github.com/crwilcox/advent-of-code/utils"
)

func zeroAtPosition(number uint64, position int) bool {
	return number>>(position)&1 == 0
}

func part1(path string) uint64 {
	strLines, _ := utils.ReadFileToLines(path)
	bits := len(strLines[0])
	lines, _ := utils.StringSliceToUintSlice(strLines)

	gammaRate := uint64(0)
	epsRate := uint64(0)
	for i := 0; i < bits; i++ {
		zeros := 0
		ones := 0
		for _, line := range lines {
			if zeroAtPosition(line, bits-i-1) {
				zeros++
			} else {
				ones++
			}
		}

		// gamma keeps the most common, epsilon keeps the least common bit.
		if zeros > ones {
			gammaRate = gammaRate << 1
			epsRate = epsRate<<1 | 1
		} else {
			gammaRate = gammaRate<<1 | 1
			epsRate = epsRate << 1
		}
	}

	return epsRate * gammaRate
}

func part2(path string) uint64 {
	strLines, _ := utils.ReadFileToLines(path)
	bits := len(strLines[0])
	lines, _ := utils.StringSliceToUintSlice(strLines)

	oGenRating := lines
	cGenRating := lines
	for i := bits; i > 0; i-- {
		if len(oGenRating) == 1 {
			break
		}
		zeros := 0
		ones := 0
		for _, line := range oGenRating {
			if zeroAtPosition(line, i-1) {
				zeros++
			} else {
				ones++
			}
		}

		// keep most common
		newRatings := make([]uint64, 0)
		for _, line := range oGenRating {
			if zeros > ones {
				// keep 0 for oxygen generator
				if zeroAtPosition(line, i-1) {
					newRatings = append(newRatings, line)
				}

			} else {
				// keep 1 for oxygen generator
				if !zeroAtPosition(line, i-1) {
					newRatings = append(newRatings, line)
				}
			}
			oGenRating = newRatings
		}
	}

	// keep least common
	for i := bits; i > 0; i-- {
		if len(cGenRating) == 1 {
			break
		}
		zeros := 0
		ones := 0
		for _, line := range cGenRating {
			if zeroAtPosition(line, i-1) {
				zeros++
			} else {
				ones++
			}
		}

		newRatings := make([]uint64, 0)
		for _, line := range cGenRating {
			if zeros > ones {
				// keep 1 for co2 generator
				if !zeroAtPosition(line, i-1) {
					newRatings = append(newRatings, line)
				}
			} else {
				// keep 0 for co2 generator
				if zeroAtPosition(line, i-1) {
					newRatings = append(newRatings, line)
				}
			}
			cGenRating = newRatings
		}
	}
	return oGenRating[0] * cGenRating[0]
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
