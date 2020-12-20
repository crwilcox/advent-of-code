package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func readFileToLines(path string) ([]string, error) {
	rootDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	file, err := os.Open(filepath.Join(rootDir, path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

const tileEdgeLength = 10

func readFileToTiles(path string) (map[int]Tile, error) {
	lines, err := readFileToLines(path)
	if err != nil {
		return nil, err
	}

	tileMap := make(map[int]Tile)

	// Tile 3079:
	// #.#.#####.
	// .#..######
	// ..#.......
	// ######....
	// ####.#..#.
	// .#...#.##.
	// #.#####.##
	// ..#.###...
	// ..#.......
	// ..#.###...
	for t := 0; t < len(lines); t += 2 + tileEdgeLength {
		tileNumberString := lines[t][5 : len(lines[t])-1]
		// regexp.MatchString("Tile ([0-9]+):", lines[t])
		tileNumber, err := strconv.Atoi(tileNumberString)
		if err != nil {
			return nil, err
		}

		tile := Tile{}
		tile.number = tileNumber
		tile.grid = make([][]byte, 0)
		for r := 1; r <= tileEdgeLength; r++ {
			tile.grid = append(tile.grid, []byte(lines[t+r]))
		}
		tileMap[tileNumber] = tile
	}
	return tileMap, nil
}

type Tile struct {
	number int
	grid   [][]byte
	//edgesMatched int
}

func (tile Tile) getEdge(edgeDirection string) string {
	ret := []byte{}

	// read all edges clockwise to make them comparable
	switch edgeDirection {
	case "T":
		// left to right
		return string(tile.grid[0])
	case "B":
		// right to left
		for i := tileEdgeLength - 1; i >= 0; i-- {
			ret = append(ret, tile.grid[tileEdgeLength-1][i])
		}
	case "L":
		// bottom up
		for i := tileEdgeLength - 1; i >= 0; i-- {
			ret = append(ret, tile.grid[i][0])
		}
	case "R":
		// top down
		for i := 0; i < tileEdgeLength; i++ {
			ret = append(ret, tile.grid[i][tileEdgeLength-1])
		}
	default:
		panic("Unknown Edge Direction")
	}
	return string(ret)
}

func findCornerTilesProduct(tiles map[int]Tile) int {
	// we don't actually have to assemble the image. we just need to find the
	// 4 tiles that have 2 edges that won't pair with any other tile
	// for each tile in the map we need to compare each edge with every other edge,
	// until we find a match. if a match is found for an edge, we don't really need to compare it.

	cornerTiles := []int{}
	for k, tile := range tiles {
		if len(cornerTiles) == 4 {
			//break
		}
		edgesUsed := countMatchingEdgeTilesForTile(tiles, tile)
		if edgesUsed == 2 {
			fmt.Println("Found Corner Tile:", k)
			cornerTiles = append(cornerTiles, k)
		} else if edgesUsed == 3 {
			fmt.Println("Found Edge Tile:", k)
		} else if edgesUsed == 4 {
			fmt.Println("Found Center Tile:", k)
		} else {
			fmt.Println("SEEMS A BUG:", k)
		}
	}
	product := 1
	for _, v := range cornerTiles {
		product *= v
	}
	return product
}

func reverse(s string) string {
	result := ""
	for _, v := range s {
		result = string(v) + result
	}
	return result
}

func countMatchingEdgeTilesForTile(tiles map[int]Tile, tile Tile) int {
	// Top Edge
	topEdge := tile.getEdge("T")
	bottomEdge := tile.getEdge("B")
	leftEdge := tile.getEdge("L")
	rightEdge := tile.getEdge("R")

	tileMatches := 0
	for _, otherTile := range tiles {
		if otherTile.number != tile.number {
			//fmt.Println("tile:", tile.number, "other:", otherTile.number)
			edges := []string{
				otherTile.getEdge("T"), otherTile.getEdge("B"),
				otherTile.getEdge("L"), otherTile.getEdge("R")}
			for _, edge := range edges {
				// to compare, the other edge needs to be inverted.
				//edge = reverse(edge)
				if edge == topEdge ||
					edge == bottomEdge ||
					edge == leftEdge ||
					edge == rightEdge {
					tileMatches++
				}
				edge = reverse(edge)
				if edge == topEdge ||
					edge == bottomEdge ||
					edge == leftEdge ||
					edge == rightEdge {
					tileMatches++
				}
			}
		}

	}
	return tileMatches
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "path/to/input/file")
		return
	}
	filePath := os.Args[1]
	tiles, err := readFileToTiles(filePath)
	if err != nil {
		panic(err)
	}

	fmt.Println(tiles)

	product := findCornerTilesProduct(tiles)
	fmt.Println("🎄 Part 1 🎁:", product) // Answer: 4006801655873

	fmt.Println("🎄 Part 2 🎁: ") // Answer:
}
