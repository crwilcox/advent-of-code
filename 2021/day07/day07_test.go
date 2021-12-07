package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 37
	got := part1("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 341558
	got = part1("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func TestPart2(t *testing.T) {
	want := 168
	got := part2("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 93214037
	got = part2("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
