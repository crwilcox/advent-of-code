package main

import (
	"fmt"
	"strconv"
	"strings"
)

func playCupGameUnoptimized(cups []int, rounds int) string {
	currentCup := 0
	currentCupLabel := cups[currentCup]
	for round := 1; round <= rounds; round++ {
		// fmt.Println("-- move", round, "--")
		// fmt.Println("cups:", cups)
		// fmt.Println("current:", currentCupLabel, "idx:", currentCup)

		// 	Before the crab starts, it will designate the first cup in your list as the current cup.
		// The crab is then going to do 100 moves.
		// Each move, the crab does the following actions:

		// The crab picks up the three cups that are immediately clockwise of the current cup. They
		// are removed from the circle; cup spacing is adjusted as necessary to maintain the circle.
		// identify cups to remove.
		removedCups := []int{}
		for i := 1; i <= 3; i++ {
			cupToRemove := (currentCup + i) % len(cups)
			removedCups = append(removedCups, cups[cupToRemove])
		}

		// fmt.Println("pick up:", removedCups)

		newCups := []int{}

		// remove cups
		for i := 0; i < len(cups); i++ {
			if i == (currentCup+1)%len(cups) || i == (currentCup+2)%len(cups) || i == (currentCup+3)%len(cups) {
				// DON'T ADD THEY ARE REMOVED
			} else {
				newCups = append(newCups, cups[i])
			}
		}
		// fmt.Println("newcups:", newCups)
		destinationCupLabel := currentCupLabel - 1

		// The crab selects a destination cup: the cup with a label equal to the current cup's label
		// minus one. If this would select one of the cups that was just picked up, the crab will
		// keep subtracting one until it finds a cup that wasn't just picked up. If at any point
		// in this process the value goes below the lowest value on any cup's label, it wraps
		// around to the highest value on any cup's label instead.
		// The crab places the cups it just picked up so that they are immediately clockwise of
		// the destination cup. They keep the same order as when they were picked up.
		minCup := newCups[0]
		maxCup := newCups[0]
		for _, v := range newCups {
			if v < minCup {
				minCup = v
			}
			if v > maxCup {
				maxCup = v
			}
		}
	foundInsertLocation:
		for {
			for i := 0; i < len(newCups); i++ {
				if newCups[i] == destinationCupLabel {
					// found the destination
					//fmt.Println("destination:", destinationCupLabel)
					insertedCups := make([]int, i+1)
					copy(insertedCups, newCups[0:i+1])

					insertedCups = append(insertedCups, removedCups...)
					insertedCups = append(insertedCups, newCups[i+1:]...)
					cups = insertedCups
					break foundInsertLocation
				}
			}
			destinationCupLabel--
			if destinationCupLabel < minCup {
				destinationCupLabel = maxCup
			}
		}

		// The crab selects a new current cup: the cup which is immediately clockwise of the current cup.
		for i := 0; i < len(cups); i++ {
			if cups[i] == currentCupLabel {
				currentCup = (i + 1) % len(cups)
				currentCupLabel = cups[currentCup]
				break
			}
		}
	}
	// fmt.Println("-- final --")
	// fmt.Println("cups:", cups)
	// fmt.Println("current:", currentCupLabel, "idx:", currentCup)

	// calculate clocwise from one
	cwFromOne := []int{}
	for k, v := range cups {
		if v == 1 {
			cwFromOne = append(cwFromOne, cups[k+1:]...)
			cwFromOne = append(cwFromOne, cups[0:k]...)
		}
	}

	out := fmt.Sprintf("%v", cwFromOne)
	out = strings.TrimPrefix(out, "[")
	out = strings.TrimSuffix(out, "]")
	out = strings.Join(strings.Split(out, " "), "")
	return out
}

// Cup struct
type Cup struct {
	next  *Cup
	value int
}

func (c *Cup) printCups() {
	startingCup := c
	iter := c
	for {
		fmt.Print(iter.value)
		iter = iter.next
		if iter == startingCup {
			break
		}
	}
	fmt.Println()
}

// playCupGame: Play with a set of cups, set of rounds. expandCupsToCount
// expands the given cup array to the given size. For instance, given [2,1,3]
// and expandCupsToCount of 5 => [2,1,3,4,5]
func playCupGame(cups []int, rounds int, expandCupsToCount int) *Cup {
	firstCup := &Cup{value: cups[0]}
	previousCup := firstCup

	// store cup values to a map of cups to speed up lookups
	cupValueMap := make(map[int]*Cup)
	cupValueMap[cups[0]] = firstCup

	maxCup := cups[0]
	for _, v := range cups {
		if v > maxCup {
			maxCup = v
		}
	}

	// create cup objects from provided list
	for i := 1; i < len(cups); i++ {
		cup := &Cup{
			value: cups[i],
		}
		previousCup.next = cup
		cupValueMap[cups[i]] = cup
		previousCup = cup
	}

	// Extend this to the requested cup count.
	for i := len(cups); i < expandCupsToCount; i++ {
		cup := &Cup{
			value: i + 1,
		}
		previousCup.next = cup
		cupValueMap[i+1] = cup
		previousCup = cup
	}
	// make list circular
	previousCup.next = firstCup

	currentCup := firstCup
	currentCupLabel := firstCup.value
	for round := 1; round <= rounds; round++ {
		// fmt.Println("-- move", round, "--")
		// fmt.Println("cups:", cups)
		// fmt.Println("current:", currentCupLabel, "idx:", currentCup).

		// The crab picks up the three cups that are immediately clockwise of the current cup. They
		// are removed from the circle; cup spacing is adjusted as necessary to maintain the circle.
		removedCups := []*Cup{
			currentCup.next, currentCup.next.next, currentCup.next.next.next}
		currentCup.next = currentCup.next.next.next.next

		// fmt.Println("pick up:", removedCups)
		destinationCupLabel := currentCupLabel - 1

		// The crab selects a destination cup: the cup with a label equal to the current cup's label
		// minus one. If this would select one of the cups that was just picked up, the crab will
		// keep subtracting one until it finds a cup that wasn't just picked up. If at any point
		// in this process the value goes below the lowest value on any cup's label, it wraps
		// around to the highest value on any cup's label instead.
		// The crab places the cups it just picked up so that they are immediately clockwise of
		// the destination cup. They keep the same order as when they were picked up.
		minCup := 1
		maxCup := expandCupsToCount

	foundInsertLocation:
		for {
			if destinationCupLabel < minCup {
				destinationCupLabel = maxCup
			}
			if destinationCupLabel == removedCups[0].value ||
				destinationCupLabel == removedCups[1].value ||
				destinationCupLabel == removedCups[2].value {
				// do nothing
			} else if destCup, ok := cupValueMap[destinationCupLabel]; ok {
				oldNext := destCup.next
				destCup.next = removedCups[0]
				removedCups[2].next = oldNext
				break foundInsertLocation
			}
			destinationCupLabel--
		}
		//fmt.Println("dest:", destinationCupLabel)

		// The crab selects a new current cup: the cup immediately clockwise of the current cup.
		currentCup = currentCup.next
		currentCupLabel = currentCup.value

	}
	// fmt.Println("-- final --")
	// fmt.Println("current:", currentCupLabel, "idx:", currentCup)

	return cupValueMap[1]
}
func main() {
	input := "215694783" // Input
	//input := "389125467" // Test Input

	// Take input string to int[]
	cups := []int{}
	for _, v := range input {
		cup, err := strconv.Atoi(string(v))
		if err != nil {
			panic("Failed Converstion of input")
		}
		cups = append(cups, cup)
	}

	fmt.Println("🎄 Part 1 (Unoptimized)🎁:") // Answer: 46978532
	fmt.Println(playCupGameUnoptimized(cups, 100))

	fmt.Println("🎄 Part 1 🎁, LL implementation")
	cup := playCupGame(cups, 100, 9)
	for iterCup := cup.next; iterCup != cup; iterCup = iterCup.next {
		fmt.Print(iterCup.value)
	}
	fmt.Println()

	// Expeced for Test: 149245887792
	fmt.Println("🎄 Part 2 🎁:") // Answer: 163035127721
	oneCup := playCupGame(cups, 10_000_000, 1_000_000)

	// The Product of the two values clockwise from 1.
	iterCup := oneCup
	product := iterCup.next.value * iterCup.next.next.value
	fmt.Println(product)
}
