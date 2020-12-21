package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// BusSchedule struct
type BusSchedule struct {
	earliestDeparture int
	buses             []int
}

func readFileToBusSchedule(path string) (*BusSchedule, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := strings.TrimSpace(scanner.Text())
	earliestDeparture, err := strconv.Atoi(line)
	if err != nil {
		return nil, err
	}

	schedule := BusSchedule{}
	schedule.earliestDeparture = earliestDeparture
	schedule.buses = make([]int, 0)

	scanner.Scan()
	line = strings.TrimSpace(scanner.Text())
	for _, v := range strings.Split(line, ",") {
		if v != "x" {
			bus, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}

			schedule.buses = append(schedule.buses, bus)
		} else {
			schedule.buses = append(schedule.buses, 0)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &schedule, nil
}

// Finds the earliest bus that can be departed on after the given
// Earliest departure time. Returns the result of: time waited * bus ID
func findEarliestBusDepartureTime(earliestDeparture int, buses []int) int {
	minimumWaitTime := -1
	busID := -1

	for _, bus := range buses {
		if bus > 0 {
			loops := (earliestDeparture / bus) + 1
			busArrival := loops * bus
			waitTime := busArrival - earliestDeparture

			//fmt.Println("Bus:", bus, "wait:", waitTime)
			if minimumWaitTime == -1 || minimumWaitTime > waitTime {
				busID = bus
				minimumWaitTime = waitTime
			}
		}
	}

	return minimumWaitTime * busID
}

// Find the earliest timestamp such that all of the listed bus IDs depart
// at offsets matching their positions in the list
func findTimeWhereBusesArriveEachMinuteAfter(buses []int) int {
	time := 0
	step := buses[0]

	for k, bus := range buses[1:] {
		offset := k + 1
		if bus > 0 {
			// fmt.Println("Bus:", bus, "Step:", step, "t:", t)
			for ; ; time += step {
				if (time+offset)%bus == 0 {
					break
				}
			}
			step *= bus
		}
	}
	return time
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	schedule, err := readFileToBusSchedule(filePath)
	if err != nil {
		panic(err)
	}

	//fmt.Println(schedule)

	fmt.Println("🎄 Part 1 🎁:") //Answer: 4782
	fmt.Println(findEarliestBusDepartureTime(schedule.earliestDeparture, schedule.buses))

	fmt.Println("🎄 Part 2 🎁:") // Answer: 1118684865113056
	fmt.Println(findTimeWhereBusesArriveEachMinuteAfter(schedule.buses))
}
