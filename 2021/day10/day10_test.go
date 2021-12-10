package main

import (
	"testing"
)

func TestPart1(t *testing.T) {

	want := 26397
	got := part1("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 243939
	got = part1("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func TestPart2(t *testing.T) {
	want := 288957
	got := part2("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 2421222841
	got = part2("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func Test_findMissingCharacter(t *testing.T) {
	tests := []struct {
		line string
		want rune
	}{
		{"{([(<{}[<>[]}>{[]{[(<()>", '}'},
		{"[[<[([]))<([[{}[[()]]]", ')'},
		{"[{[{({}]{}}([{[{{{}}([]", ']'},
		{"[<(<(<(<{}))><([]([]()", ')'},
		{"<{([([[(<>()){}]>(<<{{", '>'},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			if got, _ := findMissingCharacter(tt.line); got != tt.want {
				t.Errorf("findMissingCharacter(%v) = %v, want %v", tt.line, got, tt.want)
			}
		})
	}
}
