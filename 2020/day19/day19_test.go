package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	// expect 2
	/// should result in  aaab, aaba, bbab, bbba, abaa, abbb, baaa, or babb.
	// Rule 0, therefore, matches a (rule 4), then any of the eight options from rule 1, then b (rule 5):
	//aaaabb, aaabab, abbabb, abbbab, aabaab, aabbbb, abaaab, or ababbb.
	rules := readFileToRulesInput("test-input-1")
	rules.simplifyRules()
	count := rules.countMatchingInputs(0)
	if count != 2 {
		t.Errorf("count = %d; want %d", count, 2)
	}
}
