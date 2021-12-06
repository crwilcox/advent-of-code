package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 5934
	got := part1("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 373378
	got = part1("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func TestPart2(t *testing.T) {
	want := 26984457539
	got := part2("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 1682576647495
	got = part2("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
