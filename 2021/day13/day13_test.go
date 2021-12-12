package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	want := 17
	got := part1("test_input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}

	want = 693
	got = part1("input")
	if got != want {
		t.Error("got:", got, "want:", want)
	}
}
