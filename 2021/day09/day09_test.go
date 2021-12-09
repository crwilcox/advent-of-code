package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 15
	got := part1("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 498
	got = part1("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func TestPart2(t *testing.T) {
	want := 1134
	got := part2("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 1071000
	got = part2("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
