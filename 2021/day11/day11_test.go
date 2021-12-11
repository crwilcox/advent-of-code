package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 1656
	got := part1("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 1659
	got = part1("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func TestPart2(t *testing.T) {
	want := 195
	got := part2("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 227
	got = part2("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
