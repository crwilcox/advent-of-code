package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/2020/utils"
)

// Rule contains a name and ranges, each which have a lower and upper bound
type Rule struct {
	name string
	// range for 1-3 or 5-7 is [1,3], [5,7]
	// 0th is the lower, 1th is the upper
	ranges [][]int
}

// TicketData contains a list of rules, a ticket, and nearby tickets that
// can be used to decipher which rule maps to which ticket column
type TicketData struct {
	rules         []Rule
	ticket        []int
	nearbyTickets [][]int
}

// NewTicketDataFromFile populates a TicketData from a filepath
func NewTicketDataFromFile(path string) TicketData {
	lines, _ := utils.ReadFileToLines(path)
	td := TicketData{}
	td.rules = make([]Rule, 0)
	td.nearbyTickets = make([][]int, 0)

	// first portion of file, until a blank line, are rules
	yourTicketStart := 0
	for i, line := range lines {
		if line == "" {
			yourTicketStart = i + 2
			break
		}

		r := Rule{}
		split := strings.Split(line, ": ")
		r.name = split[0]
		r.ranges = make([][]int, 0)
		ranges := strings.Split(split[1], " or ")
		for _, v := range ranges {
			split = strings.Split(v, "-")
			lower, _ := strconv.Atoi(split[0])
			upper, _ := strconv.Atoi(split[1])
			r.ranges = append(r.ranges, []int{lower, upper})
		}
		td.rules = append(td.rules, r)
	}

	// next, is your ticket:
	nearbyTicketsStart := 0
	for i, line := range lines[yourTicketStart:] {
		if line == "" {
			nearbyTicketsStart = yourTicketStart + i + 2
			break
		}

		ticket := make([]int, 0)
		for _, v := range strings.Split(line, ",") {
			vi, _ := strconv.Atoi(v)
			ticket = append(ticket, vi)
		}
		td.ticket = ticket
	}

	// then until the end are nearby tickets
	for _, line := range lines[nearbyTicketsStart:] {
		ticket := make([]int, 0)
		for _, v := range strings.Split(line, ",") {
			vi, _ := strconv.Atoi(v)
			ticket = append(ticket, vi)
		}
		td.nearbyTickets = append(td.nearbyTickets, ticket)
	}

	return td
}

func (td TicketData) findTicketInvalidInputs(ticket []int) int {
	for _, ticketNumber := range ticket {
		inRange := false
	RulesLoop:
		for _, rule := range td.rules {
			for _, r := range rule.ranges {
				if r[0] <= ticketNumber && ticketNumber <= r[1] {
					inRange = true
					break RulesLoop
				}
			}
		}
		if !inRange {
			return ticketNumber
		}
	}
	return -1
}

func (td TicketData) findTicketErrorRate() int {
	invalidTicketsSum := 0
	for _, nearbyTicket := range td.nearbyTickets {
		invalidInput := td.findTicketInvalidInputs(nearbyTicket)
		if invalidInput >= 0 {
			invalidTicketsSum += invalidInput
		}
	}
	return invalidTicketsSum
}

func (td TicketData) getValidNearbyTickets() [][]int {
	valid := make([][]int, 0)
	for _, nearby := range td.nearbyTickets {
		res := td.findTicketInvalidInputs(nearby)
		if res == -1 {
			valid = append(valid, nearby)
		}
	}
	return valid
}

func (td TicketData) getRuleNamesForNumber(number int) map[string]bool {
	res := make(map[string]bool)
	for _, rule := range td.rules {
		for _, r := range rule.ranges {
			if r[0] <= number && number <= r[1] {
				res[rule.name] = true
			}
		}
	}
	return res
}

func (td TicketData) decipherFieldMapping() []string {
	// Find all possible fields for each ticket.
	ticketsAsPossibleFieldsMap := make([][]map[string]bool, 0)
	for _, ticket := range td.nearbyTickets {
		mapping := make([]map[string]bool, 0)
		for _, n := range ticket {
			possibleNameMap := td.getRuleNamesForNumber(n)
			mapping = append(mapping, possibleNameMap)
		}

		ticketsAsPossibleFieldsMap = append(ticketsAsPossibleFieldsMap, mapping)
	}

	// populate a list of mappings that remain to be used. Once a column
	// is verified, remove a field from others.
	remainingMappings := make(map[string]bool)
	for _, v := range td.rules {
		remainingMappings[v.name] = true
	}

	// Initialize each column on each ticket to be any ticket value type
	ticketValueCount := len(td.ticket)
	possibleMappings := make([]map[string]bool, ticketValueCount)
	for i := 0; i < ticketValueCount; i++ {
		possibleMappings[i] = make(map[string]bool, len(td.rules))
		for j := 0; j < len(td.rules); j++ {
			possibleMappings[i][td.rules[j].name] = true
		}
	}

	// While it is the case we have remaining rules to map to, keep looping
	// over all tickets. This will reduce to 1 per column.
	for len(remainingMappings) > 0 {
		// for each ticket, compare the running possible mapping to this ticket's
		// possibilities. Take the intersection as we go.
		for _, ticket := range ticketsAsPossibleFieldsMap {
			for index, fieldPossibilities := range ticket {
				if len(possibleMappings[index]) > 1 {
					inter := intersection(fieldPossibilities, possibleMappings[index])
					inter = intersection(remainingMappings, inter)
					possibleMappings[index] = inter
					if len(inter) == 1 {
						// we have found a particular columns value, this field name
						// should be excluded from being available in the future
						for k := range inter {
							delete(remainingMappings, k)
						}
					}
				}
				//fmt.Print(index, possibleMappings[index])
			}
			//fmt.Println()
		}
	}
	// if all tickets match to a single field all is well.
	ret := make([]string, 0)
	for _, v := range possibleMappings {
		for k := range v {
			ret = append(ret, k)
		}
	}
	return ret
}

func (td TicketData) multiplyDepartureFields() int {
	// Once you work out which field is which, look for the six fields on your
	// ticket that start with the word departure. What do you get if you
	// multiply those six values together?
	mapping := td.decipherFieldMapping()

	product := 1
	for i, v := range mapping {
		if strings.HasPrefix(v, "departure") {
			product *= td.ticket[i]
		}
	}
	return product
}

func intersection(a, b map[string]bool) map[string]bool {
	intersect := make(map[string]bool)

	for entry := range b {
		if _, ok := a[entry]; ok {
			intersect[entry] = true
		}
	}

	return intersect
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]

	ticketData := NewTicketDataFromFile(filePath)

	fmt.Println("🎄 Part 1 🎁: ") // Answer: 20058
	fmt.Println(ticketData.findTicketErrorRate())

	fmt.Println("🎄 Part 2 🎁: ") // Answer:366871907221
	// Get the valid tickets and overwrite the invalid
	validTickets := ticketData.getValidNearbyTickets()
	ticketData.nearbyTickets = validTickets

	fmt.Println(ticketData.multiplyDepartureFields())
}
