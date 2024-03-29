package main

import (
	"fmt"
	"os"
	"sort"

	"github.com/crwilcox/advent-of-code/utils"
)

// Given a boarding pass string, decode that to the unique seat ID
// For example:FBFBBFFRLR
//  The first 7 characters define the row, the last 3 the seat in the row.
//  F indicates front half of the remaining plane, B the opposited.
//  L indicates the left half of the remaining row, R the opposite.
//  So, the seat we arrive at is row 44, column 5.
// A unique ID can be found by multiplying the row by 8 and adding the seat.
// The example seat has ID 44 * 8 + 5 = 357.
func decodeSeatFromBoardingPass(pass string) int {
	// Calculate the row position
	firstRow := 0
	lastRow := 127
	rowPortion := pass[0:7]

	for _, v := range rowPortion {
		currentSize := lastRow - firstRow
		newSize := currentSize / 2
		if v == rune("F"[0]) {
			lastRow -= newSize + 1
		} else {
			firstRow += newSize + 1
		}
	}
	seatRow := -1
	if rowPortion[6] == byte("F"[0]) {
		seatRow = firstRow
	} else {
		seatRow = lastRow
	}

	// Calculate the seat position
	firstSeat := 0
	lastSeat := 7
	seatPortion := pass[7:]
	for _, v := range seatPortion {
		currentSize := lastSeat - firstSeat
		newSize := currentSize / 2
		if v == rune("L"[0]) {
			lastSeat -= newSize + 1
		} else {
			firstSeat += newSize + 1
		}
	}
	seat := -1
	if seatPortion[2] == byte("L"[0]) {
		seat = firstSeat
	} else {
		seat = lastSeat
	}

	// Every seat also has a unique seat ID:
	// multiply the row by 8, then add the column. In this example,
	uniqueID := (seatRow * 8) + seat

	return uniqueID
}

// As a sanity check, look through your list of boarding passes. What is the
// highest seat ID on a boarding pass?
func part1(input []string) int {
	maxUID := 0
	for _, seatString := range input {
		uid := decodeSeatFromBoardingPass(seatString)
		if uid > maxUID {
			maxUID = uid
		}
	}
	return maxUID
}

// Given an array of seat IDs, identify which seat is missing, the seat missing
// will not be the very first or very last seat.
func part2(input []string) int {
	var allSeatIds []int
	for _, seatString := range input {
		uid := decodeSeatFromBoardingPass(seatString)
		allSeatIds = append(allSeatIds, uid)
	}

	sortedSeats := allSeatIds[:]
	sort.Ints(sortedSeats)
	currentSeat := sortedSeats[0]
	for _, v := range sortedSeats {
		if v != currentSeat {
			return currentSeat
		}
		currentSeat++
	}
	return -1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	seatStrings, err := utils.ReadFileToLines(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Print("Part 1: ")
	fmt.Println(part1(seatStrings)) // 858

	fmt.Print("Part 2: ")
	fmt.Println(part2(seatStrings)) // 557
}
