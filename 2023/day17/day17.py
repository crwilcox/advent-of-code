from functools import cache
import queue
import sys
from datetime import datetime
import heapq

sys.setrecursionlimit(1_000_000)

MAXINT = 999_999_999_999


def part1(filename):
    cost = walk_to_exit_manager(filename, 1, 3)
    return cost


def part2(filename):
    cost = walk_to_exit_manager(filename, 4, 10)
    return cost


def walk_to_exit_manager(
    filename,
    min_count,
    max_count,
):
    global heap
    global q
    global cost_to_square
    global grid
    global walk_to_exit
    walk_to_exit.cache_clear()

    cost_to_square = {}
    grid = process_input(filename)

    # cost, row, col, prev_direction, direction_count, path
    heapq.heappush(heap, (0, 0, 0, None, 0))
    lowest = MAXINT

    # stored_path = None
    while heap:
        # cost, row, col, prev_direction, direction_count
        curr_cost, curr_r, curr_c, prev_direction, repeat_direction_cnt = heapq.heappop(
            heap
        )
        res = walk_to_exit(
            curr_r,
            curr_c,
            curr_cost,
            prev_direction,
            repeat_direction_cnt,
            min_count,
            max_count,
        )
        if res == MAXINT:
            pass
        elif res < 0 or res is None:
            pass
        else:
            # print(f"res:{res}")
            # if res <= lowest: stored_path = path
            lowest = min(lowest, res)

    # print(stored_path)
    return lowest


@cache
def walk_to_exit(
    curr_r,
    curr_c,
    curr_cost,
    prev_direction,
    repeat_direction_cnt,
    min_count=3,
    max_count=MAXINT,
):
    global q
    if curr_r < 0 or curr_c < 0 or curr_r >= len(grid) or curr_c >= len(grid[0]):
        # this path has left the grid. consider it high enough to be rejected
        return MAXINT

    # check if we've been to this square before, if we have and cost was lower, no point trying.
    rep = cost_to_square.get(
        (curr_r, curr_c, prev_direction, repeat_direction_cnt),
        MAXINT
    )
    if rep < curr_cost:
        # terminate, we've been here before at a lower cost. This path isn't going to need
        # to be walked
        return MAXINT

    cost_to_square[(curr_r, curr_c, prev_direction, repeat_direction_cnt)] = curr_cost

    if curr_r == len(grid) - 1 and curr_c == len(grid[0]) - 1:
        # reached end
        # print("REACHED END at cost:", curr_cost)
        if repeat_direction_cnt < min_count:
            # ultracrucible needs to go in one direction at least 4 squares to stop.
            return MAXINT
        else:
            return curr_cost

    # can go all directions but the one we came from
    # if min hasn't been met, continue forward only.
    directions = []
    if repeat_direction_cnt < min_count:
        directions = [prev_direction]
    else:
        if prev_direction in ["l", "r"]:
            directions = ["d", "u"]
        elif prev_direction in ["u", "d"]:
            directions = ["l", "r"]

    # as long as max isn't reached, continue forward
    if repeat_direction_cnt < max_count:
        directions.append(prev_direction)

    # for base case, start with two
    if prev_direction == None:
        # basecase.
        directions = ["r", "d"]

    for dir in directions:
        n_r, n_c = curr_r, curr_c
        if dir == "l":
            n_c -= 1
        elif dir == "r":
            n_c += 1
        elif dir == "u":
            n_r -= 1
        elif dir == "d":
            n_r += 1

        next_count = repeat_direction_cnt
        if dir == prev_direction:
            next_count += 1
        else:
            next_count = 1

        if n_r < 0 or n_c < 0 or n_r >= len(grid) or n_c >= len(grid[0]):
            # this path has left the grid. consider it high enough to be rejected
            # do nothing.
            pass
        else:
            heapq.heappush(
                heap, (curr_cost + grid[n_r][n_c], n_r, n_c, dir, next_count)
            )

    return MAXINT


def process_input(filename):
    file_lines = []
    with open(filename) as input:
        file_lines = input.readlines()
    lines = []
    for i in file_lines:
        lines.append([int(n) for n in i.strip()])
    return lines


def print_grid(grid):
    for r in grid:
        for c in r:
            print(c, end="")
        print()


cost_to_square = {}
heap = []
DEBUG = False

if __name__ == "__main__":
    print(f"🎄 Part 1 🎁: {part1('input')}")
    print(f"🎄 Part 2 🎁: {part2('input')}")
