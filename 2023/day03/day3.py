total = 0
board = []
with open("test_input") as input:
    for line in input.readlines():
        board.append(line.strip())


def is_adjacent(y, x):
    for y_offset in [-1, 0, 1]:
        for x_offset in [-1, 0, 1]:
            # check that these values are in range
            if (
                y + y_offset >= 0
                and x + x_offset >= 0
                and x + x_offset < len(board[0])
                and y + y_offset < len(board)
            ):
                c = board[y + y_offset][x + x_offset]
                try:
                    int(c)
                except:
                    if c != ".":
                        # we found a gear
                        return True
    return False


complete_number = ""
adjacent = False
for y, line in enumerate(board):
    for x, c in enumerate(line):
        try:
            int(c)
            complete_number += c
            # this is a number
            if not adjacent:
                adjacent = is_adjacent(y, x)
        except:
            # found the end of a number.
            if adjacent:
                total += int(complete_number)
            complete_number = ""
            adjacent = False


# Note: as authored, fairly sure if the gear had the same #
# on both sides, it wouldn't work.
print(f"🎄 Part 1 🎁: {total}")
