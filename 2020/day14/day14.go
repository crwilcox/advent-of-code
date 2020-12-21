package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/2020/utils"
)

// Apply a bitmask to a provided int. X indicates to do nothing, 0 and 1
// each result in that position getting a 0 or 1 respectively
// 010 masked with x01 => 001
func applyBitmask(rhs int, mask string) int {
	binaryRepresentation := strconv.FormatInt(int64(rhs), 2)
	binaryRepresentation = fmt.Sprintf("%036v", binaryRepresentation)
	out := ""
	for k, v := range binaryRepresentation {
		if mask[k] == 'X' {
			out += string(v)
		} else {
			out += string(mask[k])
		}
	}

	i, _ := strconv.ParseInt(out, 2, 64)
	return int(i)
}

// Part 1 Mask Values, Sum Stored Values
func maskInputAndSumStoredValuesV1(lines []string) int {
	mem := make(map[int]int)
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	for _, line := range lines {
		split := strings.Split(line, " = ")
		lhs := split[0]
		rhs := split[1]

		if lhs == "mask" {
			// update mask
			mask = rhs
		} else {
			// store a value in memory
			index, _ := strconv.Atoi(lhs[4 : len(lhs)-1])
			rhs, _ := strconv.Atoi(rhs)
			calculatedValue := applyBitmask(rhs, mask)
			mem[index] = calculatedValue
		}
	}

	// all values processed, sum values
	sum := 0
	for _, v := range mem {
		sum += v
	}
	return sum
}

// Given a bitmask contianing 'X' return all possible bitmasks.
// an "X" represents both a 1 and a 0 state
func expandBitmasks(bitmask string) []string {
	bitmasks := make([]string, 0)
	// create an array containing offsets in address to floating bits.
	floating := make([]int, 0)
	for i, v := range bitmask {
		if v == 'X' {
			floating = append(floating, i)
		}
	}

	// loop over the different 0/1 combinations, assigning the X to the
	// corresponding values.
	loopMax := int(math.Pow(2, float64(len(floating))))
	for i := 0; i < loopMax; i++ {
		j := i
		maskCopy := []byte(bitmask)
		for _, v := range floating {
			if j&0x1 == 1 {
				maskCopy[v] = '1'
			} else {
				maskCopy[v] = '0'
			}
			j = j >> 1
		}
		bitmasks = append(bitmasks, string(maskCopy))
	}
	return bitmasks
}

// Given a value (memory address), Apply the rules for bitmasks for part 2,
// returns expanded bitmasks.
func findFloatingBitmaskedAddresses(val int, mask string) []string {
	binaryRepresentation := strconv.FormatInt(int64(val), 2)
	binaryRepresentation = fmt.Sprintf("%036v", binaryRepresentation)

	v2Mask := ""
	for k, v := range binaryRepresentation {
		if mask[k] == 'X' {
			v2Mask += "X" // leave an x here. this let's us expand.
		} else if mask[k] == '1' {
			v2Mask += string(mask[k])
		} else if mask[k] == '0' {
			v2Mask += string(v)
		}
	}

	// the "x" above can be expanded now to a list of precise addresses
	bitmasks := expandBitmasks(v2Mask)
	return bitmasks
}

// Part 2 Mask Memory Addresses, Sum Stored Values
func maskInputAndSumStoredValuesV2(lines []string) int {
	mem := make(map[int64]int)

	mask := ""
	for _, line := range lines {
		split := strings.Split(line, " = ")
		lhs := split[0]
		rhs := split[1]

		if lhs == "mask" {
			// update mask
			mask = rhs
		} else {
			// store a value in memory
			index, _ := strconv.Atoi(lhs[4 : len(lhs)-1])
			rhs, _ := strconv.Atoi(rhs)
			calculatedAddresses := findFloatingBitmaskedAddresses(index, mask)
			for _, address := range calculatedAddresses {
				intAddress, _ := strconv.ParseInt(address, 2, 64)
				mem[intAddress] = rhs
			}
		}
	}

	// all values processed, sum values
	sum := 0
	for _, v := range mem {
		sum += int(v)
	}
	return sum
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	lines, err := utils.ReadFileToLines(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println("🎄 Part 1 🎁: ") // Answer: 17934269678453
	fmt.Println(maskInputAndSumStoredValuesV1(lines))
	fmt.Println("🎄 Part 2 🎁: ") // Answer: 3440662844064
	fmt.Println(maskInputAndSumStoredValuesV2(lines))
}
