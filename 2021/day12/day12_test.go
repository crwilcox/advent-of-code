package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 10
	got := part1("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 3761
	got = part1("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func TestPart2(t *testing.T) {
	want := 36
	got := part2("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 99138
	got = part2("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
