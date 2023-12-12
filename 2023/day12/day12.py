from functools import cache
from itertools import groupby


def is_pattern_valid(pattern_replaced, spring_groups):
    groups = groupby(pattern_replaced)
    result = [(sum(1 for _ in count)) for c, count in groups if c == "#"]
    return result == spring_groups


def count_valid_slow(pattern, spring_groups):
    """
    This is how I initially authored Part 1. I realized it was too
    slow and had to try a different approach.
    """

    def replace_unknown(pattern, bini):
        bini_idx = 0
        ret = ""
        for i in pattern:
            if i == "?":
                if bini[bini_idx] == "0":
                    ret += "."
                else:
                    ret += "#"
                bini_idx += 1
            else:
                ret += i
        return ret

    spring_count = pattern.count("?")
    valid_patterns = 0
    for i in range(pow(2, spring_count)):
        b = format(i, f"0{spring_count}b")
        pattern_replaced = replace_unknown(pattern, b)
        if is_pattern_valid(pattern_replaced, spring_groups):
            valid_patterns += 1
    return valid_patterns


@cache
def count_valid(pattern: str, spring_groups: tuple[int], capture_group_size=0):
    """
    The way of expanding all possible gets too big.
    Move to trying to do an iterative, recursive solution.
    """
    if not pattern:
        # Termination case. We've exhausted the pattern.
        if len(spring_groups) == 0 and capture_group_size == 0:
            # If we have no remaining groups to capture, valid solution
            return 1
        elif len(spring_groups) == 1 and spring_groups[0] == capture_group_size:
            # We had one group left and we can match it, solution.
            return 1
        else:
            # must not have a valid match despite exhausting the pattern
            return 0

    count = 0
    next = pattern[0]
    if next == "?":
        next = [".", "#"]
    for c in next:
        if c == "#":
            count += count_valid(pattern[1:], spring_groups, capture_group_size + 1)
        elif capture_group_size == 0:
            count += count_valid(pattern[1:], spring_groups)
        elif spring_groups and spring_groups[0] == capture_group_size:
            count += count_valid(pattern[1:], spring_groups[1:])

    return count


def parse_line(line):
    pattern, spring_arrangements = line.split(" ")
    return pattern, tuple([int(i) for i in spring_arrangements.split(",")])


def part1(input):
    total = 0
    for line in input:
        pattern, spring_groups = parse_line(line)
        count = count_valid(pattern, spring_groups)
        total += count
    return total


def unfold(pattern, spring_groups) -> (str, list[int]):
    folds = 5
    p = f"{pattern}?" * folds
    s = spring_groups * folds
    return p[:-1], s


def part2(input):
    total = 0
    for line in input:
        pattern, spring_groups = parse_line(line)
        p, s = unfold(pattern, spring_groups)
        count = count_valid(p, s)
        total += count
    return total


def process_input(filename):
    file_lines = []
    with open(filename) as input:
        file_lines = input.readlines()
    return file_lines


if __name__ == "__main__":
    file_lines = process_input("input")

    print(f"🎄 Part 1 🎁: {part1(file_lines)}")
    print(f"🎄 Part 2 🎁: {part2(file_lines)}")
