package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 5
	got := part1("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 5197
	got = part1("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func TestPart2(t *testing.T) {
	want := 12
	got := part2("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 18605
	got = part2("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
