package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 26
	got := part1("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 409
	got = part1("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func TestPart2(t *testing.T) {
	want := 61229
	got := part2("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 1024649
	got = part2("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
