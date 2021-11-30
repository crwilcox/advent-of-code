package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/crwilcox/advent-of-code/utils"
)

// Rule struct
type Rule struct {
	branches []string
	concrete bool
}

// Rules struct
type Rules struct {
	rules map[int]Rule
	input []string
}

func getRulesInString(str string) []int {
	found := make([]int, 0)
	for _, v := range strings.Split(str, " | ") {
		for _, w := range strings.Split(v, " ") {
			i, err := strconv.Atoi(w)
			if err != nil {
				panic(err)
			}
			found = append(found, i)
		}
	}
	return found
}

// Transforms the rules on a rules object to the concrete form
func (rules *Rules) simplifyRules() {
	// find rules that are complete first.
	concreteRules := make(map[int]Rule, 0)

	for len(concreteRules) < len(rules.rules) {
		for k, rule := range rules.rules {
			if rule.concrete {
				concreteRules[k] = rule
			}
		}

		for ruleIndex, rule := range rules.rules {
			if !rule.concrete {
				// attempt to solve a rule using the concrete rules we have
				ruleConcrete := true
				for _, branch := range rule.branches {
					rulesInBranch := getRulesInString(branch)
					//fmt.Println(rulesInBranch)
					// if all rules are concrete, we should update this rule.
					for _, concreteRule := range rulesInBranch {
						if _, ok := concreteRules[concreteRule]; !ok {
							ruleConcrete = false
							break
						}
					}
				}

				// if this is a concretable rule, take the opportunity to make it into one :)
				if ruleConcrete {
					// for each branch in our rule, work to expand the contents.
					for branchIndex, branch := range rule.branches {
						rulesInBranch := getRulesInString(branch)
						newRule := ""
						// for each rule in this branch, fetch the actual content, place it in place.
						newRule = ""
						for _, refRule := range rulesInBranch {
							joinedBranches := strings.Join(rules.rules[refRule].branches, "|")
							newRule += "(" + joinedBranches + ")"
							// if ruleIndex == 11 { // TODO: maybe PART2
							// 	newRule += "+"
							// }
						}
						// if ruleIndex == 8 { // TODO: maybe PART2
						// 	// 8 is recursive. just need to add a +
						// 	newRule += "+"
						// }
						rule.branches[branchIndex] = newRule
					}

					rule.concrete = true
					rules.rules[ruleIndex] = rule
				}
			}
		}
	}
}

func (rules Rules) countMatchingInputs(ruleIndex int) int {
	passedRuleCount := 0
	ruleToCompare := rules.rules[ruleIndex]

	joinedBranches := strings.Join(ruleToCompare.branches, "|")
	compiled, err := regexp.Compile("^(" + joinedBranches + ")$")
	if err != nil {
		panic(err)
	}

	for _, line := range rules.input {
		matched := compiled.Match([]byte(line))
		if matched {
			passedRuleCount++
		}
	}
	return passedRuleCount
}

func readFileToRulesInput(path string) Rules {
	// 0: 4 1 5
	// 1: 2 3 | 3 2
	// 2: 4 4 | 5 5
	// 3: 4 5 | 5 4
	// 4: "a"
	// 5: "b"

	// ababbb
	// bababa
	// abbbab
	// aaabbb
	// aaaabbb
	lines, err := utils.ReadFileToLines(path)
	if err != nil {
		panic(err)
	}
	rules := Rules{}
	rules.rules = make(map[int]Rule)
	rules.input = make([]string, 0)

	parsingRules := true
	for _, line := range lines {
		if line == "" {
			parsingRules = false
		} else if parsingRules {
			// parse rules
			// 0: 4 1 5
			split := strings.Split(line, ": ")
			index, err := strconv.Atoi(split[0])
			if err != nil {
				panic(err)
			}

			r := Rule{}
			if strings.HasPrefix(split[1], "\"") {
				// this rule is already complete
				r.concrete = true
				r.branches = []string{split[1][1 : len(split[1])-1]}
			} else {
				r.concrete = false

				orRules := make([]string, 0)
				for _, v := range strings.Split(split[1], " | ") {
					orRules = append(orRules, v)
				}
				r.branches = orRules
			}
			rules.rules[index] = r
		} else {
			// parse input
			// ababbb
			rules.input = append(rules.input, line)
		}
	}

	return rules
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	rules := readFileToRulesInput(filePath)

	// Part 1
	rules.simplifyRules()
	countPt1 := rules.countMatchingInputs(0)
	fmt.Println("🎄 Part 1 🎁:", countPt1) // Answer: 156

	// Part 2
	// Turns out we didn't get two rules right.

	rules42 := "(" + strings.Join(rules.rules[42].branches, "|") + ")"

	// Rule 8 Can be replaced by a reentrant rule
	// 8: 42 | 42 8
	newRule8 := "(" + rules42 + ")+"
	rule8 := rules.rules[8]
	rule8.branches = []string{newRule8}
	rules.rules[8] = rule8

	// Rule 11 is harder. Go doesn't support the regex style I would like,
	// since this needs to enforce an equal nuber of 42 and 31.
	// Just expand them by hand though and let the simplification do its thing.
	// 11: 42 31 | 42 11 31
	rule11 := rules.rules[11]
	rule11.branches = []string{
		"42 31", "42 42 31 31", "42 42 42 31 31 31", "42 42 42 42 31 31 31 31",
		"42 42 42 42 42 31 31 31 31 31"}

	rule11.concrete = false
	rules.rules[11] = rule11

	// Then, we can zero out 0, (replace it with "8 11" so we can recalculate)
	rule0 := rules.rules[0]
	rule0.concrete = false
	rule0.branches = []string{"8 11"}
	rules.rules[0] = rule0

	rules.simplifyRules()
	countPt2 := rules.countMatchingInputs(0)

	fmt.Println("🎄 Part 2 🎁:", countPt2) // Answer: 363
}
