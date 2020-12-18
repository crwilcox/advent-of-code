package main

import (
	"testing"
)

func verifyLeftRight(t *testing.T, input string, expected int) {
	val := calcLeftRight(input)
	if val != expected {
		t.Error("unexpected result:", input, "exp:", expected, "act:", val)
	}
}

func verifyPrecedence(t *testing.T, input string, expected int) {
	val := calcPrecedence(input)
	if val != expected {
		t.Error("unexpected result:", input, "exp:", expected, "act:", val)
	}
}

func TestPart1(t *testing.T) {
	verifyLeftRight(t, "1 + (2 * 3) + (4 * (5 + 6))", 51)
	verifyLeftRight(t, "2 * 3 + (4 * 5)", 26)
	verifyLeftRight(t, "5 + (8 * 3 + 9 + 3 * 4 * 3)", 437)
	verifyLeftRight(t, "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 12240)
	verifyLeftRight(t, "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 ", 13632)
}

func TestPart2(t *testing.T) {
	verifyPrecedence(t, "1 + (2 * 3) + (4 * (5 + 6))", 51)
	verifyPrecedence(t, "2 * 3 + (4 * 5)", 46)
	verifyPrecedence(t, "5 + (8 * 3 + 9 + 3 * 4 * 3)", 1445)
	verifyPrecedence(t, "5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", 669060)
	verifyPrecedence(t, "((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2 ", 23340)
}
