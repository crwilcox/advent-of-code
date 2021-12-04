package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 4512
	got := part1("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 11774
	got = part1("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func TestPart2(t *testing.T) {
	want := 1924
	got := part2("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 4495
	got = part2("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
