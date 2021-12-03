package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := uint64(198)
	got := part1("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 2250414
	got = part1("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func TestPart2(t *testing.T) {
	want := uint64(230)
	got := part2("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 6085575
	got = part2("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
