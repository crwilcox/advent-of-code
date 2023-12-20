from functools import cache
import queue
import sys

sys.setrecursionlimit(1_000_000)

MAXINT = 999_999_999_999


def part1():
    global cost_to_square
    cost_to_square = {}
    # cost = walk_to_exit(0, 0, 0, None, 0)
    cost = walk_to_exit_manager(0, 0, 0, None, 0)

    def p(r, c, d, note=""):
        print(f"{r},{c},{d} = {cost_to_square[(r,c,d)]}  {note}")

    """"
      0123456789012
    0 2>>34^>>>1323
    1 32v>>>35v5623
    2 32552456v>>54
    3 3446585845v52
    4 4546657867v>6
    5 14385987984v4
    6 44578769877v6
    7 36378779796v>
    8 465496798688v
    9 456467998645v
    0 12246868655<v
    1 25465488877v5
    2 43226746555v>
    """
    """
      0123456789012

    0 2413432311323
    1 3215453535623
    2 3255245654254
    3 3446585845452
    4 4546657867536
    5 1438598798454
    6 4457876987766
    7 3637877979653
    8 4654967986887
    9 4564679986453
    0 1224686865563
    1 2546548887735
    2 4322674655533
    """
    if DEBUG:  # test input
        print(
            "EXPECTED PATH:",
            "[(0, 0), (0, 1), (0, 2), (1, 2), (1, 3), (1, 4), (1,5),",
            "(0,5), (0,6), (0,7), (0,8), (1,8), (2,8), (2,9), (2,10),",
            "(3,10), (4,10), (4,11), (5,11)])",
        )
        p(0, 1, "r")
        p(0, 2, "r")
        p(1, 2, "d")
        p(1, 3, "r")
        p(1, 4, "r")
        p(1, 5, "r")
        p(0, 5, "u")
        p(0, 6, "r", "expect:25")
        p(0, 7, "r", "expect:28")
        p(0, 8, "r")
        p(1, 8, "d", "expect:32")
        p(2, 8, "d", "expect:37")
        p(2, 9, "r")
        p(2, 10, "r")

        print()
        print()
        p(len(grid) - 1, len(grid[0]) - 1, "r")

        print()
        print()
        p(7, 12, "r")  # 74
        p(8, 12, "d")  # 80
        p(9, 12, "d")  # 96
        p(10, 12, "d")  # 97
        p(10, 11, "l")
        p(11, 11, "d")
        p(12, 11, "d")
        p(12, 12, "r")

    return cost


def part2():
    global cost_to_square
    cost_to_square = {}
    pass


def walk_to_exit_manager(
    curr_r, curr_c, curr_cost, prev_direction, repeat_direction_cnt
):
    global q
    q.put((curr_r, curr_c, curr_cost, prev_direction, repeat_direction_cnt, None))

    processed = 0
    while not q.empty():
        processed += 1
        if processed % 1_000_000 == 0:
            lowest_row, lowest_column = len(grid) - 1, len(grid[0]) - 1
            lowest = MAXINT
            for d in ["l", "r", "d", "u"]:
                lowest = min(
                    cost_to_square.get((lowest_row, lowest_column, d), (MAXINT,))[0],
                    lowest,
                )

            print(
                f"processed: {processed:_} queued:{q._qsize():_} lowest end:{lowest:_}"
            )

        args = q.get()
        curr_r, curr_c, curr_cost, prev_direction, repeat_direction_cnt, path = args
        res = walk_to_exit(
            curr_r, curr_c, curr_cost, prev_direction, repeat_direction_cnt, path
        )
        if res == MAXINT:
            pass
        elif res < 0 or res is None:
            pass
        else:
            print(f"res:{res}")
            # pass

    lowest_row, lowest_column = len(grid) - 1, len(grid[0]) - 1
    lowest = MAXINT
    for d in ["l", "r", "d", "u"]:
        lowest = min(
            cost_to_square.get((lowest_row, lowest_column, d), (MAXINT,))[0], lowest
        )
    return lowest


@cache
def walk_to_exit(
    curr_r, curr_c, curr_cost, prev_direction, repeat_direction_cnt, path=()
):
    global q

    if curr_r < 0 or curr_c < 0 or curr_r >= len(grid) or curr_c >= len(grid[0]):
        # this path has left the grid. consider it high enough to be rejected
        return MAXINT

    # check if we've been to this square before, if we have and cost was lower, no point trying.
    rep = cost_to_square.get(
        (curr_r, curr_c, prev_direction), (MAXINT, repeat_direction_cnt, [])
    )
    if rep[0] < curr_cost and rep[1] <= repeat_direction_cnt:
        # terminate, we've been here before at a lower cost. This path isn't going to need
        # to be walked
        return MAXINT

    lowest_row, lowest_column = len(grid) - 1, len(grid[0]) - 1
    end_square = cost_to_square.get(
        (lowest_row, lowest_column, prev_direction), (MAXINT,)
    )
    if end_square[0] < curr_cost:
        # terminate, we've already found a path to the end that is shorter than this.
        return MAXINT

    cost_to_square[(curr_r, curr_c, prev_direction)] = (
        curr_cost,
        repeat_direction_cnt,
        path,
    )
    if curr_r == len(grid) - 1 and curr_c == len(grid[0]) - 1:
        # reached end
        # print("REACHED END at cost:", curr_cost)
        return curr_cost

    # can go all directions but the one we came from
    directions = ["l", "r", "d", "u"]
    if prev_direction == "l":
        # no right
        directions = ["d", "u"]
    elif prev_direction == "r":
        # no left
        directions = ["d", "u"]
    elif prev_direction == "u":
        # no down
        directions = ["l", "r"]
    elif prev_direction == "d":
        # no up
        directions = ["l", "r"]
    # if we've went the same direction 3 times, we can't do it again.
    if repeat_direction_cnt < 3:  # todo: ensure this is the right offset?
        directions.append(prev_direction)

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
            q.put((n_r, n_c, curr_cost + grid[n_r][n_c], dir, next_count, path))

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


grid = process_input("test_input")
cost_to_square = {}
q = queue.Queue()
if __name__ == "__main__":
    print_grid(grid)
    from datetime import datetime

    print(datetime.now())
    print("Test:")
    DEBUG = True
    grid = process_input("test_input")

    print(f"🎄 Part 1 🎁: {part1()}")
    print(datetime.now())
    print(f"🎄 Part 2 🎁: {part2()}")
    print(datetime.now())

    walk_to_exit.cache_clear()
    print("Live:")
    cost_to_square = {}
    grid = process_input("input")
    DEBUG = False

    print(datetime.now())
    print(f"🎄 Part 1 🎁: {part1()}")
    print(datetime.now())
    print(f"🎄 Part 2 🎁: {part2()}")
    print(datetime.now())
