def find_galaxies(space: list[str]) -> list[tuple]:
    ret = []
    for row, r in enumerate(space):
        for col, square in enumerate(r):
            if square == "#":
                ret.append((row, col))

    return ret


def adjust_coordinate_for_expanded_space(space, row, col, multiplier=1_000_000):
    # Expand Rows
    row_adjust = 0
    for r in space[:row]:
        if r.count(".") == len(r):
            row_adjust += 1
    row_adjust = (row_adjust * multiplier) - row_adjust

    # Expand Columns
    idx = 0
    col_adjust = 0
    while idx < col:
        all_dots = True
        for r in space:
            if r[idx] != ".":
                all_dots = False
                break
        if all_dots:
            col_adjust += 1
        idx += 1
    col_adjust = (col_adjust * multiplier) - col_adjust

    return row + row_adjust, col + col_adjust


def calculate_distance_between_galaxies(galaxy_1, galaxy_2):
    g1r, g1c = galaxy_1
    g2r, g2c = galaxy_2
    d = abs(g1r - g2r) + abs(g1c - g2c)
    return d


def get_distance_between_all_galaxies(galaxies):
    distances_between_galaxies = []
    for idx, g1 in enumerate(galaxies[:-1]):
        for _, g2 in enumerate(galaxies[idx + 1 :]):
            distances_between_galaxies.append(
                calculate_distance_between_galaxies(g1, g2)
            )
    return distances_between_galaxies


def expand_space(space: list[str]):
    """
    This is only used for Part 1, it is redundant with part 2 approach
    """
    ret_space = []
    for r in space:
        if r.count(".") == len(r):
            ret_space.append(r)
        ret_space.append(r)

    idx = 0
    while idx < len(ret_space[0]):
        all_dots = True
        for r in ret_space:
            if r[idx] != ".":
                all_dots = False
        if all_dots:
            for row, val in enumerate(ret_space):
                ret_space[row] = val[0:idx] + "." + val[idx:]
            idx += 1
        idx += 1
    return ret_space


def part1(input):
    space = expand_space(input)
    galaxies = find_galaxies(space)
    distances_between_galaxies = get_distance_between_all_galaxies(galaxies)
    return sum(distances_between_galaxies)


def part2(input, multiplier=1_000_000):
    galaxies = find_galaxies(input)
    galaxies = [
        adjust_coordinate_for_expanded_space(input, row, col, multiplier)
        for row, col in galaxies
    ]
    distances_between_galaxies = get_distance_between_all_galaxies(galaxies)

    return sum(distances_between_galaxies)


def process_input(filename):
    file_lines = []
    with open(filename) as input:
        file_lines = input.readlines()
    return [i.strip() for i in file_lines]


if __name__ == "__main__":
    file_lines = process_input("input")

    print(f"🎄 Part 1 🎁: {part1(file_lines)}")
    print(f"🎄 Part 1 🎁: {part2(file_lines, 2)} (Using Part 2 Implementation)")
    print(f"🎄 Part 2 🎁: {part2(file_lines)}")
