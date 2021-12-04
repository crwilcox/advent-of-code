package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/utils"
)

type BingoSquare struct {
	number int
	marked bool
}

type BingoBoard [][]BingoSquare

func (b BingoBoard) Print() {
	for _, v := range b {
		for _, w := range v {
			fmt.Print(w.number, " ")
		}
		fmt.Println()
	}
}

// return true if winning now.
func (b BingoBoard) MarkBoard(number int) bool {
	for i, v := range b {
		for j, w := range v {
			if w.number == number {
				b[i][j].marked = true

				// check row
				foundUnmarked := false
				for _, v := range b[i] {
					if !v.marked {
						foundUnmarked = true
					}
				}
				if !foundUnmarked {
					return true
				}

				// check column
				foundUnmarked = false
				for i := 0; i < len(b); i++ {
					if !b[i][j].marked {
						foundUnmarked = true
					}
				}

				return !foundUnmarked
			}
		}
	}
	return false
}

func (b BingoBoard) ScoreBoard(lastCalledNumber int) int {
	// Start by finding the sum of all unmarked
	// numbers on that board; Then, multiply that sum by
	// the number that was just called when the board won.
	sum := 0
	for _, v := range b {
		for _, w := range v {
			if !w.marked {
				sum += w.number
			}
		}
	}
	return sum * lastCalledNumber
}

func buildBoards(lines []string) []BingoBoard {
	boards := make([]BingoBoard, 0)
	// get each board
	board := BingoBoard{}
	for _, line := range lines[2:] {
		line := strings.TrimSpace(line)
		if line == "" {
			boards = append(boards, board)
			board = BingoBoard{}
			continue
		}
		row := make([]BingoSquare, 0)
		for _, v := range strings.Split(line, " ") {
			v := strings.TrimSpace(v)
			if v == "" {
				continue
			}
			n, err := strconv.Atoi(v)
			if err != nil {
				fmt.Println("err: bingo square:", err)
			}
			row = append(row, BingoSquare{number: n, marked: false})
		}
		board = append(board, row)

	}
	boards = append(boards, board)

	// Possibly helpful, print boards
	// for _, board := range boards {
	// 	board.Print()
	// 	fmt.Println()
	// }

	return boards
}

func part1(path string) int {
	lines, _ := utils.ReadFileToLines(path)

	// first line is picks
	picks := lines[0]
	boards := buildBoards(lines[2:])

	for _, v := range strings.Split(picks, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			fmt.Println("err: pick:", err)
		}

		for _, board := range boards {
			winner := board.MarkBoard(n)
			if winner {
				score := board.ScoreBoard(n)
				return score
			}
		}
	}

	// We should find a winning board and never get here
	return -1
}

func part2(path string) int {
	lines, _ := utils.ReadFileToLines(path)

	// first line is picks
	picks := lines[0]
	boards := buildBoards(lines[2:])

	lastWinner := BingoBoard{}
	lastWinningPick := -1

	for _, v := range strings.Split(picks, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(v))
		if err != nil {
			fmt.Println("err: pick:", err)
		}

		nextRoundBoards := make([]BingoBoard, 0)
		for _, board := range boards {
			winner := board.MarkBoard(n)
			if winner {
				// we don't pick this yet, we want to keep going if it
				// isn't the last winning
				lastWinner = board
				lastWinningPick = n
			} else {
				nextRoundBoards = append(nextRoundBoards, board)
			}
		}
		boards = nextRoundBoards
	}

	return lastWinner.ScoreBoard(lastWinningPick)
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
