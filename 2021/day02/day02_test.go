package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 150
	got := part1("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}

func TestPart2(t *testing.T) {
	want := 900
	got := part2("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
