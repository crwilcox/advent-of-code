package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/utils"
)

// Stack wraps a slice to act as a stack
type Stack struct {
	stack []string
}

func (s *Stack) push(r string) {
	s.stack = append(s.stack, r)
}

func (s *Stack) pop() string {
	n := len(s.stack) - 1
	ret := s.stack[n]
	s.stack = s.stack[:n]
	return ret
}

func (s *Stack) isEmpty() bool {
	return len(s.stack) == 0
}

func leftRightMath(input string) int {
	input = input + "+" // add an extra token so we do a final loop
	total := 0
	lastNumber := ""
	lastOperator := '+'
	for _, v := range input {
		if v == '*' || v == '+' {
			n, _ := strconv.Atoi(lastNumber)
			if lastOperator == '+' {
				total += n
			} else {
				total *= n
			}
			lastOperator = v
			lastNumber = ""
		} else {
			lastNumber += string(v)
		}
	}

	return total
}

func precedenceMath(input string) int {
	// addition is evaluated before multiplication.
	product := 1
	multSplit := strings.Split(input, "*")
	for _, v := range multSplit {
		// only addition operators, or a number in this region anymore.
		addSum := 0
		valuesToSum := strings.Split(v, "+")
		for _, add := range valuesToSum {
			v, _ := strconv.Atoi(add)
			addSum += v
		}
		product *= addSum
	}

	return product
}

func calcLeftRight(input string) int {
	return calc(input, true)
}

func calcPrecedence(input string) int {
	return calc(input, false)
}

func calc(input string, leftRight bool) int {
	//fmt.Println("calc:", input)
	input = strings.Replace(input, " ", "", -1)

	stack := Stack{}
	stack.stack = make([]string, 0)

	// pop off the stack until a '(' is encountered. stop, do sub calc, put back
	for _, char := range input {
		switch char {
		case ')':
			// find the nearest '('
			// then do math for inner.
			// then push that all back :)
			innerString := ""
			curr := string(char)
			for {
				curr = stack.pop()
				if curr == "(" {
					break
				}
				innerString = string(curr) + innerString
			}
			m := 0
			if leftRight {
				m = leftRightMath(innerString)
			} else {
				m = precedenceMath(innerString)
			}
			stack.push(strconv.Itoa(m))
		default:
			// found a number, operator, etc. Push to stack.
			stack.push(string(char))
		}
	}

	finalString := ""
	for !stack.isEmpty() {
		finalString = stack.pop() + finalString
	}
	if leftRight {
		return leftRightMath(finalString)
	}

	return precedenceMath(finalString)
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

	fmt.Println("🎄 Part 1 🎁: ") // Answer: 21993583522852
	sum := 0
	for _, v := range lines {
		sum += calcLeftRight(v)
	}
	fmt.Println(sum)

	fmt.Println("🎄 Part 2 🎁: ") // Answer:122438593522757
	sum = 0
	for _, v := range lines {
		sum += calcPrecedence(v)
	}
	fmt.Println(sum)
}
