package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/utils"
)

func inputParse(path string) (paper [][]bool, folds []string) {
	lines, _ := utils.ReadFileToLines(path)
	maxX := 0
	maxY := 0
	for _, line := range lines {
		if line != "" && !strings.Contains(line, "fold") {
			coord := strings.Split(line, ",")
			x, _ := strconv.Atoi(coord[0])
			y, _ := strconv.Atoi(coord[1])
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	// make paper
	paper = make([][]bool, maxX+1)
	for x := 0; x <= maxX; x++ {
		paper[x] = make([]bool, maxY+1)
	}

	fmt.Println(len(paper), len(paper[0]))
	for _, line := range lines {
		if strings.Index(line, "fold") == 0 {
			fold := strings.Split(line, " ")[2]
			folds = append(folds, fold)
		} else if line != "" {
			coord := strings.Split(line, ",")
			x, _ := strconv.Atoi(coord[0])
			y, _ := strconv.Atoi(coord[1])
			paper[x][y] = true
		}
	}
	return
}

func horizontalFold(paper [][]bool, position int) [][]bool {
	for i := 1; i < len(paper)-position; i++ {
		for idx, v := range paper[position+i] {
			if v {
				paper[position-i][idx] = paper[position+i][idx]
			}
		}
	}
	paper = paper[:position]
	return paper
}

func verticalFold(paper [][]bool, position int) [][]bool {
	for ri, row := range paper {
		newRow := row[0:position]
		for i := 1; i < len(row)-position; i++ {
			if row[position+i] {
				newRow[position-i] = row[position+i]
			}
		}
		paper[ri] = newRow
	}
	return paper
}

func part1(path string) int {
	paper, folds := inputParse(path)

	firstFold := folds[0]
	directions := strings.Split(firstFold, "=")
	dir := directions[0]
	position, _ := strconv.Atoi(directions[1])
	if dir == "y" {
		paper = verticalFold(paper, position)
	} else if dir == "x" {
		paper = horizontalFold(paper, position)
	} else {
		fmt.Println("ERR")
	}

	count := 0
	for _, r := range paper {
		for _, c := range r {
			if c {
				count++
			}
		}
	}
	return count

}

func part2(path string) {
	paper, folds := inputParse(path)

	for _, fold := range folds {
		directions := strings.Split(fold, "=")
		dir := directions[0]
		position, _ := strconv.Atoi(directions[1])
		if dir == "y" {
			paper = verticalFold(paper, position)
		} else if dir == "x" {
			paper = horizontalFold(paper, position)
		} else {
			fmt.Println("ERR")
		}
	}
	printPaper(paper)
}

func printPaper(paper [][]bool) {
	// printed backwards so do reversed
	for i := len(paper) - 1; i >= 0; i-- {
		r := paper[i]
		for _, v := range r {

			if v {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
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
	part2(filePath)
}
