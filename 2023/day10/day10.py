def find_start(grid) -> (int, int):
    for idx, line in enumerate(grid):
        if "S" in line:
            # print("START:", idx, line.index("S"))
            return idx, line.index("S")
    return -1, -1


def find_next(grid, row, col) -> (int, int):
    """
    | is a vertical pipe connecting north and south.
    - is a horizontal pipe connecting east and west.
    L is a 90-degree bend connecting north and east.
    J is a 90-degree bend connecting north and west.
    7 is a 90-degree bend connecting south and west.
    F is a 90-degree bend connecting south and east.
    . is ground; there is no pipe in this tile.
    S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
    """
    current_square = grid[row][col]

    end_north = ["|", "L", "J", "S"]
    end_east = ["-", "L", "F", "S"]
    end_south = ["|", "7", "F", "S"]
    end_west = ["-", "J", "7", "S"]

    # north
    if row > 0:
        # check if current square can go north.
        if (
            current_square in end_north
            and grid[row - 1][col] in end_south
            and (row - 1, col) not in visited
        ):
            next = (row - 1, col)
            visited[next] = 1
            return next
    # east
    if len(grid[0]) - 1 > col:
        if (
            current_square in end_east
            and grid[row][col + 1] in end_west
            and (row, col + 1) not in visited
        ):
            next = (row, col + 1)
            visited[next] = 1
            return next
    # south
    if len(grid) > row:
        if (
            current_square in end_south
            and grid[row + 1][col] in end_north
            and (row + 1, col) not in visited
        ):
            next = (row + 1, col)
            visited[next] = 1
            return next
    # west
    if col > 0:
        if (
            current_square in end_west
            and grid[row][col - 1] in end_east
            and (row, col - 1) not in visited
        ):
            next = (row, col - 1)
            visited[next] = 1
            return next

    # must have completed the loop
    return -1, -1


def get_unvisited_square(grid):
    for row, r in enumerate(grid):
        for col, _ in enumerate(r):
            if (row, col) not in visited:
                return row, col
    return -1, -1


visited = {}


def part1(grid):
    global visited
    visited = {}

    # Find loop start "S"
    start_row, start_col = find_start(grid)
    visited[start_row, start_col] = 1
    row, col = start_row, start_col
    count = 0
    while True:
        # print("row:", row, "col:", col)
        # walk each square, count our way around.
        count += 1
        row, col = find_next(grid, row, col)
        if row == -1 and col == -1:
            return int(count / 2)


def part2(grid):
    def print_grid(grid):
        for row in grid:
            for col in row:
                print(col, end="")
            print()

    def is_path(row, col):
        return visited.get((row, col), -1) == 1

    def char_grid(grid):
        """
        Takes a grid and makes it individiual chars. Also just convert all non maze chars to '.'
        """
        new_grid = []
        for row, rval in enumerate(grid):
            new_row = []
            for col, _ in enumerate(rval):
                if is_path(row, col):
                    new_row.append(grid[row][col])
                else:
                    new_row.append(".")
            new_grid.append(new_row)
        return new_grid

    enclosed_grid = char_grid(grid)

    # Walk the maze by row. Determine if we are in and out by watching for turns.
    count = 0
    for row, row_val in enumerate(enclosed_grid):
        out = True
        for col, val in enumerate(row_val):
            if is_path(row, col):
                if val in ["F", "S", "|", "7"]:
                    out = not out
            elif not out:
                enclosed_grid[row][col] = "1"
                count += 1
    # print_grid(enclosed_grid)

    return count


def process_input(filename):
    file_lines = []
    with open(filename) as input:
        file_lines = [i.strip() for i in input.readlines()]
    return file_lines


if __name__ == "__main__":
    file_lines = process_input("input")

    print(f"🎄 Part 1 🎁: {part1(file_lines)}")
    print(f"🎄 Part 2 🎁: {part2(file_lines)}")
