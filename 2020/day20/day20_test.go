package main

import (
	"testing"
)

func compareByteArrays(expected [][]byte, actual [][]byte) bool {
	for x, v := range expected {
		for y, val := range v {
			if val != actual[x][y] {
				return false
			}
		}
	}
	return true
}
func TestFlipVertical(t *testing.T) {
	start := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}
	result := flipVertical(start)
	expected := [][]byte{
		{'g', 'h', 'i'},
		{'d', 'e', 'f'},
		{'a', 'b', 'c'},
	}
	if !compareByteArrays(expected, result) {
		t.Errorf("%s; want %s", result, expected)
	}
}

func TestFlipHorizontal(t *testing.T) {
	start := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}
	result := flipHorizontal(start)
	expected := [][]byte{
		{'c', 'b', 'a'},
		{'f', 'e', 'd'},
		{'i', 'h', 'g'},
	}
	if !compareByteArrays(expected, result) {
		t.Errorf("%s; want %s", result, expected)
	}
}

func TestRotateClockwise(t *testing.T) {
	start := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}
	result := rotateClockwise(start)
	expected := [][]byte{
		{'g', 'd', 'a'},
		{'h', 'e', 'b'},
		{'i', 'f', 'c'},
	}
	if !compareByteArrays(expected, result) {
		t.Errorf("%s; want %s", result, expected)
	}
}

func TestRemoveEdgesFromImage(t *testing.T) {
	start := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}
	result := removeEdgesFromImage(start)
	expected := [][]byte{
		{'e'},
	}
	if !compareByteArrays(expected, result) {
		t.Errorf("%s; want %s", result, expected)
	}
}

func TestEdgeRemoval(t *testing.T) {
	a := [][]byte{
		{'a', 'b', 'c'},
		{'d', '0', 'f'},
		{'g', 'h', 'i'},
	}
	b := [][]byte{
		{'a', 'b', 'c'},
		{'d', '1', 'f'},
		{'g', 'h', 'i'},
	}
	c := [][]byte{
		{'a', 'b', 'c'},
		{'d', '2', 'f'},
		{'g', 'h', 'i'},
	}
	d := [][]byte{
		{'a', 'b', 'c'},
		{'d', '3', 'f'},
		{'g', 'h', 'i'},
	}
	tiles := [][]Tile{
		[]Tile{
			Tile{0, a}, Tile{1, b}},
		[]Tile{
			Tile{2, c}, Tile{3, d}}}

	expected := [][]byte{
		{'0', '1'},
		{'2', '3'}}

	result := removeEdgesFromTilesAndMerge(tiles)
	if !compareByteArrays(expected, result) {
		t.Errorf("%s; want %s", result, expected)
	}
}

func TestTileRotation(t *testing.T) {
	tile := Tile{0, [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
	}}

	expected := [][]byte{
		{'i', 'f', 'c'}, // a
		{'h', 'e', 'b'}, // b
		{'g', 'd', 'a'}, // c
	}
	// should require a rotation, then a horizontal flip.
	tile.rotateUntilEdgeMatchInDirection("abc", "R")
	if !compareByteArrays(expected, tile.grid) {
		t.Errorf("%s; want %s", tile.grid, expected)
	}
}
