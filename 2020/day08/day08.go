package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/utils"
)

type instruction struct {
	name    string
	offset  int
	visited bool
}

func readFileToInstructions(path string) ([]instruction, error) {
	lines, err := utils.ReadFileToLines(path)
	if err != nil {
		return nil, err
	}

	instructions := []instruction{}
	for _, line := range lines {
		split := strings.Split(line, " ")
		i := instruction{}
		i.name = split[0]
		i.offset, err = strconv.Atoi(split[1])
		if err != nil {
			return nil, err
		}
		instructions = append(instructions, i)
	}

	return instructions, nil
}

// Runs the provided instructions and returns the accumulators. If the provided
// code contains an infinite loop, before any instruction is
// executed a second time, returns the state of the accumulator,
func executeInstructions(instructions []instruction) (int, bool) {
	currentInstruction := 0
	accumulator := 0
	for currentInstruction < len(instructions) {
		inst := instructions[currentInstruction]
		if inst.visited {
			return accumulator, false
		}
		instructions[currentInstruction].visited = true
		switch inst.name {
		case "acc":
			accumulator += inst.offset
			currentInstruction++
		case "jmp":
			currentInstruction += inst.offset
		case "nop":
			currentInstruction++
		}
	}
	return accumulator, true
}

// Given instructions that have an infinite loop, it attempts to find a
// mutation that will allow running to completion. Mutates jmp to nop calls and
// nop to jmp calls. Once a mutation is found that allows to run to completion,
// return the accumulator value of that mutation.
func findDefectAndRunInstructions(instructions []instruction) int {
	for idx, inst := range instructions {
		instr := make([]instruction, len(instructions))
		copy(instr, instructions)
		instr[0].visited = false

		switch inst.name {
		case "jmp":
			instr[idx].name = "nop"
		case "nop":
			instr[idx].name = "jmp"
		default:
			continue
		}

		acc, terminated := executeInstructions(instr)

		if terminated {
			return acc
		}
	}
	// Never found a terminating case
	return -1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	instructions, err := readFileToInstructions(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Print("Part 1: ")
	part1Instructions := make([]instruction, len(instructions))
	copy(part1Instructions, instructions)
	accumulator, _ := executeInstructions(part1Instructions)
	fmt.Println(accumulator) // 2051

	fmt.Print("Part 2: ")
	fmt.Println(findDefectAndRunInstructions(instructions)) // 2304
}
