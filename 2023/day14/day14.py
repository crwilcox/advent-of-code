def get_column(grid, col_idx):
    res = []
    for row in grid:
        res.append(row[col_idx])
    return res


def put_column(grid, col_idx, col_vals) -> tuple[str]:
    if len(col_vals) != len(grid):
        raise Exception("Not enough colvals")
    out_grid = []
    for idx, val in enumerate(col_vals):
        if type(grid[idx]) is not str:
            val = [val]
        replaced = grid[idx][:col_idx] + val + grid[idx][col_idx + 1 :]
        out_grid.append(str(replaced))

    return tuple(out_grid)


ROUNDROCK = "O"
CUBEDROCK = "#"


def tilt_col_north(col) -> list[str]:
    new_col = []
    for idx, val in enumerate(col):
        if val == ROUNDROCK:
            new_col.append(ROUNDROCK)
        if val == CUBEDROCK:
            # we can't move rocks more from here.
            # we need to pad out and place the cube rock
            empty_squares = idx - len(new_col)
            for _ in range(empty_squares):
                new_col.append(".")
            new_col.append(CUBEDROCK)

    # check padding. make sure they are the same length
    empty_squares = idx - len(new_col)
    while len(new_col) != len(col):
        new_col.append(".")
    return new_col


def tilt_north(grid) -> tuple[str]:
    # get columns, work to collapse them north.
    for idx in range(len(grid[0])):
        col = get_column(grid, idx)
        new_col = tilt_col_north(col)
        grid = put_column(grid, idx, new_col)

    return tuple(grid)


def tilt_south(grid) -> tuple[str]:
    for idx in range(len(grid[0])):
        col = get_column(grid, idx)
        new_col = tilt_col_north(col[::-1])
        grid = put_column(grid, idx, new_col[::-1])

    return tuple(grid)


def tilt_row_west(row) -> tuple[str]:
    new_row = ""
    for row_idx, val in enumerate(row):
        if val == ROUNDROCK:
            new_row += ROUNDROCK
        if val == CUBEDROCK:
            empty_squares = row_idx - len(new_row)
            for _ in range(empty_squares):
                new_row += "."
            new_row += CUBEDROCK
    while len(new_row) != len(row):
        new_row += "."
    return new_row


def tilt_west(grid) -> tuple[str]:
    new_grid = []
    for _, row in enumerate(grid):
        new_row = tilt_row_west(row)
        new_grid.append(new_row)
    return tuple(new_grid)


def tilt_east(grid) -> tuple[str]:
    new_grid = []
    for _, row in enumerate(grid):
        new_row = tilt_row_west(row[::-1])
        new_grid.append(new_row[::-1])
    return tuple(new_grid)


def calc_load(grid) -> int:
    load_total = 0
    row_weight = len(grid)
    for row in grid:
        load_total += row.count("O") * row_weight
        row_weight -= 1

    return load_total


def print_grid(grid):
    for r in grid:
        print(r)


def part1(grid):
    tilted_grid = tilt_north(grid)
    print_grid(tilted_grid)
    return calc_load(tilted_grid)


def execute_cycle(grid) -> tuple[str]:
    grid = tilt_north(grid)
    grid = tilt_west(grid)
    grid = tilt_south(grid)
    grid = tilt_east(grid)
    return grid


def part2(grid):
    seen = {}
    grid = tuple(grid)
    i = 0
    max_val = 1_000_000_000
    while i < max_val:
        if i % 10_000_000 == 0:
            print("i:", i)
        hashval = str.join("", [str(i) for i in grid])
        looped = seen.get(hashval, False)
        if looped and i + looped < max_val:
            print("CYCLE DETECTED:", i, "JUMP:", looped)
            loop_size = i - looped
            loop_multiples = int((max_val - i - 1) / loop_size)
            i += loop_size * (loop_multiples)
            seen[hashval] = i

        else:
            grid = execute_cycle(grid)
            # print("full process:", i, "calc load:", calc_load(grid))
            seen[hashval] = i
            i += 1

    return calc_load(grid)


def process_input(filename):
    file_lines = []
    with open(filename) as input:
        file_lines = input.readlines()
    return [i.strip() for i in file_lines]


if __name__ == "__main__":
    file_lines = process_input("input")

    print(f"🎄 Part 1 🎁: {part1(file_lines)}")
    print(f"🎄 Part 2 🎁: {part2(file_lines)}")
