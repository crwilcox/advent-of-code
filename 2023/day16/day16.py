import sys


energized_locations = {}


def get_next_loc(location, direction):
    row, col = location
    next_row, next_col = row, col
    if direction == "l":
        next_col = col - 1
    elif direction == "r":
        next_col = col + 1
    elif direction == "u":
        next_row = row - 1
    elif direction == "d":
        next_row = row + 1

    # print(f"Next:{next_row}, {next_col}")
    return (next_row, next_col)


# direction is l,r,u,d
def walk_grid(grid, location, direction):
    row, col = location
    if row >= len(grid) or col >= len(grid[0]) or row < 0 or col < 0:
        # left the bounds of the grid.
        return

    # check if we've energized this square in the same direction
    energized = energized_locations.get(location, [])
    if direction in energized:
        # we've already done this square heading this direction, short circuit
        return
    energized.append(direction)
    energized_locations[location] = energized

    square = grid[row][col]
    if square == ".":
        walk_grid(grid, get_next_loc(location, direction), direction)
    elif square == "/":
        if direction == "r":
            direction = "u"
        elif direction == "d":
            direction = "l"
        elif direction == "l":
            direction = "d"
        elif direction == "u":
            direction = "r"
        walk_grid(grid, get_next_loc(location, direction), direction)
    elif square == "\\":
        if direction == "r":
            direction = "d"
        elif direction == "d":
            direction = "r"
        elif direction == "l":
            direction = "u"
        elif direction == "u":
            direction = "l"
        walk_grid(grid, get_next_loc(location, direction), direction)
    elif square == "|":
        if direction == "u" or direction == "d":
            walk_grid(grid, get_next_loc(location, direction), direction)
        else:
            # split
            walk_grid(grid, get_next_loc(location, "u"), "u")
            walk_grid(grid, get_next_loc(location, "d"), "d")
    elif square == "-":
        if direction == "r" or direction == "l":
            walk_grid(grid, get_next_loc(location, direction), direction)
        else:
            # split
            walk_grid(grid, get_next_loc(location, "l"), "l")
            walk_grid(grid, get_next_loc(location, "r"), "r")
    else:
        raise Exception("UNKNOWN ELEMENT:" + square)


def print_energized(grid, energized_locations):
    for r_idx, r in enumerate(grid):
        for c_idx, c in enumerate(r):
            if energized_locations.get((r_idx, c_idx), False):
                print("#", end="")
            else:
                print(".", end="")
        print()


def part1(grid):
    beam_row, beam_col = 0, 0
    beam_direction = "r"
    walk_grid(grid, (beam_row, beam_col), beam_direction)
    # print("ENERGIZED COUNT:", len(energized_locations))
    # print(energized_locations)
    # print_energized(grid, energized_locations)
    return len(energized_locations)


def run_grid(grid, start_row, start_col, dir) -> int:
    global energized_locations
    energized_locations = {}
    walk_grid(grid, (start_row, start_col), dir)
    return len(energized_locations)


def part2(grid):
    # find best starting point.
    max_energy = 0
    for idx, _ in enumerate(grid[0]):
        # top row, dir down
        max_energy = max(max_energy, run_grid(grid, 0, idx, "d"))

        # bottom row, dir up
        max_energy = max(max_energy, run_grid(grid, len(grid) - 1, idx, "u"))

    rightmost = len(grid[0]) - 1
    for idx, _ in enumerate(grid):
        # left, dir right
        max_energy = max(max_energy, run_grid(grid, idx, 0, "r"))

        # right, dir left
        max_energy = max(max_energy, run_grid(grid, idx, rightmost, "l"))

    return max_energy


def process_input(filename):
    file_lines = []
    with open(filename) as input:
        file_lines = input.readlines()
    return [i.strip() for i in file_lines]


if __name__ == "__main__":
    file_lines = process_input("input")
    sys.setrecursionlimit(10000)
    print(f"🎄 Part 1 🎁: {part1(file_lines)}")
    print(f"🎄 Part 2 🎁: {part2(file_lines)}")
