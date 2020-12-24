package main

import (
	"fmt"
	"os"

	"github.com/crwilcox/advent-of-code/2020/utils"
)

const white = 0
const black = 1

// Floor structs
type Floor struct {
	tiles map[int]map[int]int
	days  int
}

func (f *Floor) walkPath(path string) {
	// Tile Directions (From Tile)
	// NW - -1, +1
	// NE - +1. +1
	// E  - +2,  0
	// SE - +1, -1
	// SW - -1, -1
	// W  - -2

	splitPath := parsePath(path)
	xPos := 0
	yPos := 0

	for _, v := range splitPath {
		switch v {
		case "nw":
			xPos--
			yPos++
		case "ne":
			xPos++
			yPos++
		case "e":
			xPos += 2
		case "se":
			xPos++
			yPos--
		case "sw":
			xPos--
			yPos--
		case "w":
			xPos -= 2
		}
	}

	currSide := f.getFloorTile(xPos, yPos)
	if currSide == black {
		f.tiles[xPos][yPos] = white
	} else {
		f.tiles[xPos][yPos] = black
	}
}

func (f *Floor) countBlackTiles() int {
	blackTileCount := 0
	for _, x := range f.tiles {
		for _, v := range x {
			if v == black {
				blackTileCount++
			}
		}
	}
	return blackTileCount
}

func (f *Floor) fillGapsInGrid() {
	xMin := 0
	xMax := 0
	yMin := 0
	yMax := 0

	for x, xMap := range f.tiles {
		for y, val := range xMap {
			if val == black {
				if xMin > x {
					xMin = x
				}
				if xMax < x {
					xMax = x
				}
				if yMin > y {
					yMin = y
				}
				if yMax < y {
					yMax = y
				}
			}
		}
	}

	for x := xMin - 2; x <= xMax+2; x++ {
		for y := yMin - 1; y <= yMax+1; y++ {
			f.getFloorTile(x, y)
		}
	}
}
func (f *Floor) getFloorTile(x int, y int) int {
	if _, ok := f.tiles[x]; !ok {
		f.tiles[x] = make(map[int]int)
	}

	if v, ok := f.tiles[x][y]; ok {
		return v
	} else {
		// will cause grid to enlarge as we run
		f.tiles[x][y] = white
	}
	return white
}

// The tile floor in the lobby is meant to be a living art exhibit. Every day,
// the tiles are all flipped according to the following rules:
// Any black tile with zero or more than 2 black tiles immediately adjacent
// to it is flipped to white.
// Any white tile with exactly 2 black tiles immediately adjacent to it
// is flipped to black.
func (f *Floor) updateFloorToNextDay() {
	// for part 1 the grid isn't actually complete. all black tiles need their
	// neighbors populated for this to work.
	f.fillGapsInGrid()
	newFloorMap := make(map[int]map[int]int)
	for x, xMap := range f.tiles {
		newFloorMap[x] = make(map[int]int)
		for y, value := range xMap {
			// look at surrounding tiles
			// Tile Directions (From Tile)
			// NW - -1, +1
			// NE - +1. +1
			// E  - +2,  0
			// SE - +1, -1
			// SW - -1, -1
			// W  - -2
			// black is 1, white is 0, so the return of getFloorTile can just
			// be added
			blackTilesSurrounding := 0
			blackTilesSurrounding += f.getFloorTile(x-1, y+1) // NW
			blackTilesSurrounding += f.getFloorTile(x+1, y+1) // NE
			blackTilesSurrounding += f.getFloorTile(x+2, y)   // E
			blackTilesSurrounding += f.getFloorTile(x+1, y-1) // SE
			blackTilesSurrounding += f.getFloorTile(x-1, y-1) // SW
			blackTilesSurrounding += f.getFloorTile(x-2, y)   // W

			newFloorMap[x][y] = f.tiles[x][y]
			//fmt.Println("current:", value, "surrounding:", blackTilesSurrounding)
			if value == black {
				// Any black tile with zero or more than 2 black tiles immediately adjacent
				// to it is flipped to white.
				if blackTilesSurrounding == 0 || blackTilesSurrounding > 2 {
					newFloorMap[x][y] = white
				}
			} else {
				// Any white tile with exactly 2 black tiles immediately adjacent to it
				// is flipped to black.
				if blackTilesSurrounding == 2 {
					newFloorMap[x][y] = black
				}
			}
		}
	}
	f.tiles = newFloorMap
}

func parsePath(path string) []string {
	out := make([]string, 0)
	for i := 0; i < len(path); i++ {
		switch path[i] {
		case 'w':
			out = append(out, "w")
		case 'e':
			out = append(out, "e")
		case 'n':
			i++
			if path[i] == 'w' {
				out = append(out, "nw")
			} else {
				out = append(out, "ne")
			}
		case 's':
			i++
			if path[i] == 'w' {
				out = append(out, "sw")
			} else {
				out = append(out, "se")
			}
		}
	}
	return out
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	lines, err := utils.ReadFileToLines(filePath)
	if err != nil {
		panic(err)
	}

	// Part 1
	floor := Floor{tiles: make(map[int]map[int]int)}

	for _, path := range lines {
		floor.walkPath(path)
	}
	blackTiles := floor.countBlackTiles()
	fmt.Println("🎄 Part 1 🎁:", blackTiles) // Answer: 495

	fmt.Println("🎄 Part 2 🎁: ") // Answer: 4012
	for day := 1; day <= 100; day++ {
		floor.updateFloorToNextDay()
		if day%10 == 0 {
			fmt.Println("Day:", day, "Count:", floor.countBlackTiles())
		}
	}
}
