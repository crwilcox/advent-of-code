package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/crwilcox/advent-of-code/2020/utils"
)

// Move struct
type Move struct {
	direction string
	value     int
}

func abs(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

func readFileToMoves(path string) ([]Move, error) {
	lines, err := utils.ReadFileToLines(path)
	if err != nil {
		return nil, err
	}

	moves := make([]Move, 0)
	for _, line := range lines {

		m := Move{}
		m.direction = string(line[0])
		m.value, err = strconv.Atoi(line[1:])
		if err != nil {
			return nil, err
		}
		moves = append(moves, m)
	}

	return moves, nil
}

// Calculate the distance traveled from origin. Moves treated as follows:
// N - Move Ship N by value
// W - Move Ship W by value
// S - Move Ship S by value
// E - Move Ship E by value
// F - Move Ship forward, based on directionFace (starts by facing East)
// R - Turn ship to the right by value degrees (90,180,270,360,...)
// L - Turn ship to the left by value degrees (90,180,270,360,...)
func calculateManhattanDistance(moves []Move) int {
	eastWest := 0
	northSouth := 0

	directionFaced := "E"
	for _, move := range moves {
		if move.direction == "L" {
			for turn := 0; turn < move.value; turn += 90 {
				switch directionFaced {
				case "N":
					directionFaced = "W"
				case "W":
					directionFaced = "S"
				case "S":
					directionFaced = "E"
				case "E":
					directionFaced = "N"
				}
			}
		} else if move.direction == "R" {
			// rotate ship to right by value degrees
			for turn := 0; turn < move.value; turn += 90 {
				switch directionFaced {
				case "N":
					directionFaced = "E"
				case "W":
					directionFaced = "N"
				case "S":
					directionFaced = "W"
				case "E":
					directionFaced = "S"
				}
			}
		} else {
			if move.direction == "F" {
				move.direction = directionFaced
			}

			switch move.direction {
			case "N":
				northSouth -= move.value
			case "S":
				northSouth += move.value
			case "E":
				eastWest -= move.value
			case "W":
				eastWest += move.value
			}
		}

	}

	return abs(northSouth) + abs(eastWest)
}

// Calculate the distance traveled from origin. Moves treated as follows:
// N - Move the waypoint north by the given value.
// W - move the waypoint west by the given value.
// S - move the waypoint south by the given value.
// E - move the waypoint east by the given value.
// F - move forward to the waypoint a number of times equal to the given value.
// R - Move the waypoint around the ship to the right (clockwise) by the
//     given degrees.
// L - Move the waypoint around the ship to the left (counter-clockwise) by the
//     given degrees.
func calculateManhattanDistanceMovingWaypoint(moves []Move) int {
	waypointEastWest := -10  // represent east as negative, west as positive
	waypointNorthSouth := -1 // represent north as negative, south as positive
	eastWest := 0
	northSouth := 0

	for _, move := range moves {
		// fmt.Println("wpNS:", waypointNorthSouth, "wpEW:", waypointEastWest,
		// 	"move:", move, "locEW:", eastWest, "locNS:", northSouth)
		if move.direction == "L" {
			// move waypoint counter-clockwise
			for turn := 0; turn < move.value; turn += 90 {
				previousEW := waypointEastWest
				previousNS := waypointNorthSouth

				waypointNorthSouth = previousEW
				waypointEastWest = -previousNS
			}
		} else if move.direction == "R" {
			// move waypoint clockwise
			for turn := 0; turn < move.value; turn += 90 {
				previousEW := waypointEastWest
				previousNS := waypointNorthSouth

				waypointNorthSouth = -previousEW
				waypointEastWest = previousNS
			}
		} else {
			if move.direction == "F" {
				// move ship by a multiple of value and the waypoint
				northSouth += waypointNorthSouth * move.value
				eastWest += waypointEastWest * move.value
			}

			// move waypoint in direction specified by value
			switch move.direction {
			case "N":
				waypointNorthSouth -= move.value
			case "S":
				waypointNorthSouth += move.value
			case "E":
				waypointEastWest -= move.value
			case "W":
				waypointEastWest += move.value
			}
		}
	}

	return abs(northSouth) + abs(eastWest)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	moves, err := readFileToMoves(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Print("Part 1: ") // Answer: 441
	fmt.Println(calculateManhattanDistance(moves))

	fmt.Print("Part 2: ") // Answer: 40014
	fmt.Println(calculateManhattanDistanceMovingWaypoint(moves))
}
