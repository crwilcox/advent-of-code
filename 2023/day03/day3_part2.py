total = 0
board = []
with open("input") as input:
    for line in input.readlines():
        board.append(line.strip())


def find_number(y, x):
    curr_x = x
    number = board[y][x]
    # start by going left
    while True:
        curr_x -= 1
        try:
            int(board[y][curr_x])
            number = board[y][curr_x] + number
        except:
            break
    # then go right of the base after too
    curr_x = x
    while True:
        curr_x += 1
        try:
            int(board[y][curr_x])
            number = number + board[y][curr_x]
        except:
            break
    return int(number)


# Given a gear location, find two values and multiple
def find_gear_ratio(y, x):
    numbers = []
    for y_offset in [-1, 0, 1]:
        for x_offset in [-1, 0, 1]:
            # check that these values are in range
            if (
                y + y_offset >= 0
                and x + x_offset >= 0
                and x + x_offset < len(board[0])
                and y + y_offset < len(board)
            ):
                try:
                    int(board[y + y_offset][x + x_offset])
                    number = find_number(y + y_offset, x + x_offset)
                    numbers.append(number)
                except:
                    # nothing to do, not a number
                    continue

    s = list(set(numbers))
    if len(set(numbers)) == 2:
        return s[0] * s[1]
    return 0


# find all '*' which represent gears. If there are
# two numbers adjacent, that is a gear ratio

for y, line in enumerate(board):
    for x, c in enumerate(line):
        if c == "*":
            # find two values, if there are two, multiple, return value.
            total += find_gear_ratio(y, x)


# Note: as authored, fairly sure if the gear had the same #
# on both sides, it wouldn't work.
print(f"🎄 Part 2 🎁: {total}")
