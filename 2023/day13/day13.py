def find_reflection(map, allow_smudge=False) -> (int, int):
    # for line in map:
    #     print(line)

    def delta(l, r):
        delta = 0
        for idx, _ in enumerate(l):
            if l[idx] != r[idx]:
                delta += 1
        return delta

    def check_horizontal(map) -> int:
        horizontal_reflection_pt = 0
        smudge_used = False
        for idx, _ in enumerate(map):
            split_point = idx + 1  # offset because we start in row 1.

            offset = 0
            mirroring = False
            while split_point - offset - 1 >= 0 and split_point + offset < len(map):
                l, r = map[split_point - offset - 1], map[split_point + offset]
                diff = delta(l, r)

                if allow_smudge and not smudge_used and diff <= 1:
                    mirroring = True
                    if diff == 1:
                        smudge_used = True
                elif diff == 0:
                    mirroring = True
                else:
                    mirroring = False
                    smudge_used = False
                    break
                offset += 1
            if mirroring:
                if allow_smudge and not smudge_used:
                    # not really mirroring.
                    pass
                else:
                    above = split_point

                    horizontal_reflection_pt = above
                    if allow_smudge:
                        if smudge_used:
                            # if smudge is used, we should immediately return this as the result.
                            return horizontal_reflection_pt

                    break

        return horizontal_reflection_pt

    def check_vertical(map) -> int:
        smudge_used = False

        vertical_reflection_pt = 0

        def get_column(map, col):
            res = []
            for i in map:
                res.append(i[col])
            return res

        for idx, _ in enumerate(map[0]):
            split_point = idx + 1  # offset because we start in row 1.

            offset = 0
            mirroring = False
            while split_point - offset - 1 >= 0 and split_point + offset < len(map[0]):
                l = get_column(map, split_point - offset - 1)
                r = get_column(map, split_point + offset)

                diff = delta(l, r)
                if allow_smudge and not smudge_used and diff <= 1:
                    mirroring = True
                    if diff == 1:
                        smudge_used = True

                elif diff == 0:
                    mirroring = True
                else:
                    mirroring = False
                    smudge_used = False
                    break
                offset += 1
            if mirroring:
                if allow_smudge and not smudge_used:
                    # not really mirroring.
                    pass
                else:
                    vertical_reflection_pt = split_point
                    break

        return vertical_reflection_pt

    horizontal_reflection_pt = check_horizontal(map)
    vertical_reflection_pt = check_vertical(map)

    return (
        horizontal_reflection_pt,
        vertical_reflection_pt,
    )


def part1(maps):
    horizontal = 0
    vertical = 0
    for map in maps:
        h, v = find_reflection(map)
        horizontal += h
        vertical += v

    return horizontal * 100 + vertical


def part2(input):
    horizontal = 0
    vertical = 0
    for map in maps:
        h, v = find_reflection(map, allow_smudge=True)
        if h and v:
            print("WARNING: THIS SHOULD NOT HAPPEN:", h, v)
        horizontal += h
        vertical += v

    return horizontal * 100 + vertical


def process_input(filename):
    file_lines = []
    with open(filename) as input:
        file_lines = input.readlines()

    maps = []
    map = []
    for l in file_lines:
        line = l.strip()
        if line != "":
            map.append(line)
        else:
            maps.append(map)
            map = []
    maps.append(map)
    return maps


if __name__ == "__main__":
    maps = process_input("input")
    # 35538
    print(f"🎄 Part 1 🎁: {part1(maps)}")
    # 30442
    print(f"🎄 Part 2 🎁: {part2(maps)}")
