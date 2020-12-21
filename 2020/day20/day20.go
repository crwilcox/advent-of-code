package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/crwilcox/advent-of-code/2020/utils"
)

const tileEdgeLength = 10

func readFileToTiles(path string) (map[int]Tile, error) {
	lines, err := utils.ReadFileToLines(path)
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

// Tile represents a square grid to be constructed into an image
type Tile struct {
	number int
	grid   [][]byte
	//edgesMatched int
}

func rotateClockwise(grid [][]byte) [][]byte {
	n := len(grid)
	newGrid := make([][]byte, n)
	for i := 0; i < n; i++ {
		newGrid[i] = make([]byte, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newGrid[i][j] = grid[n-j-1][i]
		}
	}
	return newGrid
}

func (tile *Tile) rotateClockwise() {
	tile.grid = rotateClockwise(tile.grid)
}

func (tile *Tile) getEdge(edgeDirection string) string {
	ret := []byte{}
	n := len(tile.grid)
	// read all edges clockwise to make them comparable
	switch edgeDirection {
	case "T":
		// left to right
		return string(tile.grid[0])
	case "B":
		// right to left
		for i := n - 1; i >= 0; i-- {
			ret = append(ret, tile.grid[n-1][i])
		}
	case "L":
		// bottom up
		for i := n - 1; i >= 0; i-- {
			ret = append(ret, tile.grid[i][0])
		}
	case "R":
		// top down
		for i := 0; i < n; i++ {
			ret = append(ret, tile.grid[i][n-1])
		}
	default:
		panic("Unknown Edge Direction")
	}
	return string(ret)
}

func flipVertical(grid [][]byte) [][]byte {
	n := len(grid)
	newGrid := make([][]byte, n)

	for k, v := range grid {
		newGrid[n-k-1] = v
	}
	return newGrid
}

func (tile *Tile) flipVertical() {
	tile.grid = flipVertical(tile.grid)
}

func flipHorizontal(grid [][]byte) [][]byte {
	n := len(grid)
	newGrid := make([][]byte, n)
	for i := 0; i < n; i++ {
		newGrid[i] = make([]byte, n)
	}

	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			newGrid[r][c] = grid[r][n-c-1]
		}
	}
	return newGrid
}

func (tile *Tile) flipHorizontal() {
	tile.grid = flipHorizontal(tile.grid)
}

func (tile *Tile) rotateUntilEdgeMatchInDirection(edge string, direction string) {
	// count is not strictly needed, but in case a bad edge was given it is
	// better to not infinitely loop
	for count := 0; reverse(edge) != tile.getEdge(direction) && count < 8; count++ {
		if edge == tile.getEdge(direction) {
			// if this happens, the tile needs to be flipped
			if direction == "T" || direction == "B" {
				// flip horizontal
				tile.flipHorizontal()
			} else {
				// flip vertical
				tile.flipVertical()
			}
		} else {
			tile.rotateClockwise()
		}
	}
}

func isEdgeInTileEdges(edges []TileEdge, edge string) bool {
	for _, v := range edges {
		if v.edge == edge || v.edge == reverse(edge) {
			return true
		}
	}
	return false
}

func arrangeTilesToGrid(tiles map[int]Tile) [][]Tile {
	edgeMap := createTileToConnectedTileMap(tiles)

	startingTileNo := -1
	// from the tile edge map, pick a number that is a corner
	for k, edges := range edgeMap {
		if len(edges) == 2 {
			// this is a corner tile. this is our starting tile.
			startingTileNo = k
			// rotate the starting tile so that the edges are down and right.
			startingTile := tiles[k]

			// TODO: was using this to verify behavior. can remove
			//startingTile.flipVertical()

			oneEdge := edges[0].edge
			twoEdge := edges[1].edge

			for r := 0; ; r++ {
				tileBottom := startingTile.getEdge("B")
				tileRight := startingTile.getEdge("R")
				// if isEdgeInTileEdges(edges, tileBottom) && isEdgeInTileEdges(edges, tileRight) {
				// 	break
				// }
				if (tileBottom == oneEdge || tileBottom == twoEdge) &&
					(tileRight == oneEdge || tileRight == twoEdge) {
					break
				}
				startingTile.rotateClockwise()
			}
			tiles[startingTileNo] = startingTile
			break
		}
	}
	fmt.Println("STARTING TILE:", startingTileNo)

	// now that we have the first element, we can begin to place the other elements,
	// row by row. find based on the tile to the left, what tile this could be.
	// Then place our tile, but rotate it to match the adjacent edge.  then
	// move on to the next one. once all tiles are exhausted, :+1:
	imageArrangement := make([][]Tile, 0)
	for row := 0; ; row++ {
		rowArr := []Tile{}
		for col := 0; ; col++ {
			if row == 0 && col == 0 {
				// this is the case we have the starting element picked, just insert it.
				rowArr = append(rowArr, tiles[startingTileNo])
			} else if col == 0 {
				// in the first position, we look up to determine this square.
				// look to the square above
				adjacentTile := imageArrangement[row-1][col]
				adjacentEdge := adjacentTile.getEdge("B")
				edges := edgeMap[adjacentTile.number]
				for _, edge := range edges {
					if edge.edge == adjacentEdge || reverse(edge.edge) == adjacentEdge {
						// this is the matching square.
						// rotate the tile to ensure it is facing the right way
						newTile := tiles[edge.toTile]
						newTile.rotateUntilEdgeMatchInDirection(adjacentEdge, "T")
						tiles[edge.toTile] = newTile
						rowArr = append(rowArr, newTile)

						// TODO: Test code, shouldn't ever print
						thisTileEdges := edgeMap[rowArr[col].number]
						edge := newTile.getEdge("L")
						if isEdgeInTileEdges(thisTileEdges, edge) {
							fmt.Println("there should not be a connection to the left of the left side")
						}
					}
				}

			} else {
				// otherwise we need to find the next square. look to the square to the left
				tileToLeft := rowArr[col-1]
				adjacentEdge := tileToLeft.getEdge("R")
				edges := edgeMap[tileToLeft.number]
				foundEdge := false
				for _, edge := range edges {
					if edge.edge == adjacentEdge || reverse(edge.edge) == adjacentEdge {
						// this is the matching square.
						// rotate the tile to ensure it is facing the right way
						newTile := tiles[edge.toTile]
						newTile.rotateUntilEdgeMatchInDirection(adjacentEdge, "L")
						tiles[edge.toTile] = newTile
						rowArr = append(rowArr, newTile)
						foundEdge = true

						// TEST CODE: THIS SHOULD NEVER PRINT ANYTHING
						// verify that the tile, provided it isn't in row 0
						// has a match above, not just to it's left
						if row > 0 {
							aboveTile := imageArrangement[row-1][col]
							aboveBottom := aboveTile.getEdge("B")
							currentTop := newTile.getEdge("T")
							if reverse(aboveBottom) != currentTop {
								fmt.Println("SEEMS TILES DON'T QUITE MATCH")
							}
						}
					}
				}
				if !foundEdge {
					fmt.Println("SOMEHOW DIDN'T FIND CONNECTED: edges:", edges, "adj:", adjacentEdge)
				} else {
					// check if this square has an edge that connects to it's right.
					// if it doesn't row is complete
					thisTileEdges := edgeMap[rowArr[col].number]
					thisTile := rowArr[col]
					edge := thisTile.getEdge("R")
					if !isEdgeInTileEdges(thisTileEdges, edge) {
						break
					}
				}
			}
		}
		imageArrangement = append(imageArrangement, rowArr)

		// check if this square has an edge that connects to it's bottom.
		// if it doesn't we are is complete
		tileInRowTileEdges := edgeMap[rowArr[0].number]
		tileInRow := rowArr[0]
		edge := tileInRow.getEdge("B")
		if !isEdgeInTileEdges(tileInRowTileEdges, edge) {
			break
		}
	}

	verifyTileGrid(imageArrangement)
	return imageArrangement

}

func verifyTileGrid(imageArrangement [][]Tile) {
	// verify tile grid. each inner tile's edges should match the edge to it's
	// neighbors. NOt necessary, but was debugigng and added it.
	for row := 1; row < len(imageArrangement)-1; row++ {
		for col := 1; col < len(imageArrangement)-1; col++ {
			topNeighbor := imageArrangement[row-1][col].getEdge("B")
			rightNeighbor := imageArrangement[row][col+1].getEdge("L")
			bottomNeighbor := imageArrangement[row+1][col].getEdge("T")
			leftNeighbor := imageArrangement[row][col-1].getEdge("R")

			current := imageArrangement[row][col]
			top := current.getEdge("T")
			bottom := current.getEdge("B")
			left := current.getEdge("L")
			right := current.getEdge("R")
			if topNeighbor != reverse(top) {
				fmt.Println("TOP MISMATCH")
			}

			if bottomNeighbor != reverse(bottom) {
				fmt.Println("BOTTOM MISMATCH")
			}

			if leftNeighbor != reverse(left) {
				fmt.Println("LEFT MISMATCH")
			}

			if rightNeighbor != reverse(right) {
				fmt.Println("RIGHT MISMATCH")
			}
		}
	}
}

// returns a map of tile to other tiles numbers
func createTileToConnectedTileMap(tiles map[int]Tile) map[int][]TileEdge {
	tileMap := make(map[int][]TileEdge)
	for _, tile := range tiles {
		edgesUsed := findMatchingEdgeTilesForTile(tiles, tile)
		tileMap[tile.number] = edgesUsed
	}
	// 	1951    2311    3079
	// 2729    1427    2473
	// 2971    1489    1171
	// - any direction works -
	// 	3079 2473 1171
	// 2311 1427 1489
	// 1951 2729 2971
	return tileMap
}

func findCornerTilesProduct(tiles map[int]Tile) int {
	// we don't actually have to assemble the image. we just need to find the
	// 4 tiles that have 2 edges that won't pair with any other tile
	// for each tile in the map we need to compare each edge with every other edge,
	// until we find a match. if a match is found for an edge, we don't really need to compare it.

	cornerTiles := []int{}
	for k, tile := range tiles {
		if len(cornerTiles) == 4 {
			break
		}
		edgesUsed := findMatchingEdgeTilesForTile(tiles, tile)
		if len(edgesUsed) == 2 {
			// fmt.Println("Found Corner Tile:", k)
			cornerTiles = append(cornerTiles, k)
		} else if len(edgesUsed) == 3 {
			// fmt.Println("Found Edge Tile:", k)
		} else if len(edgesUsed) == 4 {
			// fmt.Println("Found Center Tile:", k)
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

// TileEdge represents a single edge on a Tile and which tile it links to.
type TileEdge struct {
	edge string
	//fromTile int
	toTile int
}

func findMatchingEdgeTilesForTile(tiles map[int]Tile, tile Tile) []TileEdge {
	// Top Edge
	topEdge := tile.getEdge("T")
	bottomEdge := tile.getEdge("B")
	leftEdge := tile.getEdge("L")
	rightEdge := tile.getEdge("R")

	tileMatches := []TileEdge{}
	for _, otherTile := range tiles {
		if otherTile.number != tile.number {
			//fmt.Println("tile:", tile.number, "other:", otherTile.number)
			edges := []string{
				otherTile.getEdge("T"), otherTile.getEdge("B"),
				otherTile.getEdge("L"), otherTile.getEdge("R")}
			for _, edge := range edges {
				// to compare, the other edge needs to be inverted.
				if edge == topEdge ||
					edge == bottomEdge ||
					edge == leftEdge ||
					edge == rightEdge {

					tileEdge := TileEdge{}
					tileEdge.edge = edge
					tileEdge.toTile = otherTile.number
					tileMatches = append(tileMatches, tileEdge)
				}
				edge = reverse(edge)
				if edge == topEdge ||
					edge == bottomEdge ||
					edge == leftEdge ||
					edge == rightEdge {
					tileEdge := TileEdge{}
					tileEdge.edge = edge
					tileEdge.toTile = otherTile.number
					tileMatches = append(tileMatches, tileEdge)
				}
			}
		}

	}
	return tileMatches
}

func removeEdgesFromImage(image [][]byte) [][]byte {
	n := len(image) - 2
	newImage := make([][]byte, n)
	for i := 0; i < n; i++ {
		newImage[i] = make([]byte, n)
	}

	for row := 1; row < len(image)-1; row++ {
		for col := 1; col < len(image)-1; col++ {
			newImage[row-1][col-1] = image[row][col]
		}
	}

	return newImage
}

// removes the outer most row or column from a [][]byte object as the edges
// aren't part of the image
func removeEdgesFromTilesAndMerge(tileGrid [][]Tile) [][]byte {
	tileSize := len(tileGrid[0][0].grid) - 2

	// initialize an array for the image.
	rowsLength := len(tileGrid) * tileSize
	newImage := make([][]byte, rowsLength)
	for i := 0; i < rowsLength; i++ {
		newImage[i] = make([]byte, rowsLength)
	}

	for rowTileIndex, rows := range tileGrid {
		for colTileIndex, tile := range rows {
			tileGrid := removeEdgesFromImage(tile.grid)
			for rowPixelIndex, v := range tileGrid {
				for colPixelIndex, pixel := range v {
					newImageRow := (rowTileIndex * tileSize) + rowPixelIndex
					newImageCol := (colTileIndex * tileSize) + colPixelIndex
					newImage[newImageRow][newImageCol] = pixel
				}
			}
		}
	}

	return newImage
}

// Coord is a row/col pair offset used by findMonsters
type Coord struct {
	rowOffset, colOffset int
}

// inner method
func findMonsters(image [][]byte) int {
	// ..................O..                                       (-1, 18)
	// O....OO....OO....OOO. (0,0) (0,5)(0,6) (0,11)(0,12) (0,17)(0,18)(0,19)
	// .O..O..O..O..O..O....   (1,1)(1,4)(1,7)(1,10) (1,13) (1,16)
	nessieOffsets := []Coord{
		Coord{-1, 18}, Coord{0, 0}, Coord{0, 5}, Coord{0, 6}, Coord{0, 11},
		Coord{0, 12}, Coord{0, 17}, Coord{0, 18}, Coord{0, 19}, Coord{1, 1},
		Coord{1, 4}, Coord{1, 7}, Coord{1, 10}, Coord{1, 13}, Coord{1, 16},
	}

	// start one row in, and stop one before end, since nessie is 3 rows long,
	// and we index off of the middle row.
	monsterCount := 0
	for row := 1; row < len(image)-1; row++ {
		lengthOfNessie := 20
		for col := 0; col < len(image[row])-lengthOfNessie; col++ {
			if image[row][col] == '#' {
				// this might be nessie. Check other tiles relative to this one
				isNessie := true
				for _, rowCol := range nessieOffsets {
					if image[row+rowCol.rowOffset][col+rowCol.colOffset] != '#' {
						// seems not nessie. :(
						isNessie = false
						break
					}
				}

				if isNessie {
					// We found nessie, add to monsterCount
					monsterCount++
				}
			}
		}
	}

	// for _, v := range image {
	// 	fmt.Println(string(v))
	// }
	// when we find a # we assume that is a sea monster tail. Check other
	// relative positions. if it is a sea monster, convert the # to be O.
	if monsterCount > 0 {
		sum := 0
		// we found a monster, so now we need to count the squares
		for _, row := range image {
			for _, col := range row {
				if col == '#' {
					sum++
				}
			}
		}
		return sum - (monsterCount * 15)
	}
	// If we have gotten here, we didn't find a sea monster, return -1.
	return -1
}

// Given an image byte array consisting of '#' and '.', find sea monsters.
// Return the count of '#' not in monster. Monsters are the following pattern
// within a 2d byte array
func findMonstersAndCountSquares(image [][]byte) int {
	// This can be both flipped or rotated. Iterate over that option while looking.
	// for vertFlips := 0; vertFlips < 2; vertFlips++ {
	for flips := 0; flips < 2; flips++ {
		// check all rotations
		for rotations := 0; rotations < 4; rotations++ {
			notMonsterSquares := findMonsters(image)
			if notMonsterSquares >= 0 {
				return notMonsterSquares
			}
			image = rotateClockwise(image)
		}
		// if we didn't figure it out after rotating, maybe we need a flip?
		image = flipVertical(image)
	}

	return -1
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

	product := findCornerTilesProduct(tiles)
	fmt.Println("🎄 Part 1 🎁:", product) // Answer: 4006801655873

	// Part 2
	tileGrid := arrangeTilesToGrid(tiles)
	// once constructed each tile's edges need to be carved off the tile
	combinedImage := removeEdgesFromTilesAndMerge(tileGrid)
	// Now, search for monsters and determine water roughness
	roughWater := findMonsters(combinedImage)
	fmt.Println("🎄 Part 2 🎁:", roughWater) // Answer:
}
