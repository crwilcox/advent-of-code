def find_delta_patterns(deltas):
    # Iterate over sensor until all zeros
    while True:
        delta = []
        prev_val = deltas[-1][0]
        for val in deltas[-1][1:]:
            delta.append(int(val) - int(prev_val))
            prev_val = val
        deltas.append(delta)
        # print(deltas)

        def all_zeros(arr):
            for a in arr:
                if a != 0:
                    return False
            return True

        if all_zeros(delta):
            break
    return deltas


def part1(sensors):
    total = 0
    for sensor in sensors:
        deltas = [sensor]

        deltas = find_delta_patterns(deltas)

        def add_deltas(deltas):
            # from the back, find the last element to add to the last element of the one above.
            # do this until we are to the top.
            last_elem = 0
            for delta in deltas[::-1]:
                last_elem = delta[-1] + last_elem
            return last_elem

        extrapolated_elem = add_deltas(deltas)
        total += extrapolated_elem
    return total


def part2(sensors):
    total = 0
    for sensor in sensors:
        deltas = [sensor]

        deltas = find_delta_patterns(deltas)

        def sub_deltas(deltas):
            last_elem = 0
            for delta in deltas[::-1]:
                last_elem = delta[0] - last_elem

            return last_elem

        extrapolated_elem = sub_deltas(deltas)
        total += extrapolated_elem
    return total


def process_input(filename):
    file_lines = []
    with open(filename) as input:
        file_lines = input.readlines()
    return [[int(s) for s in sensor.strip().split(" ")] for sensor in file_lines]


sensors = process_input("input")
print(f"🎄 Part 1 🎁: {part1(sensors)}")
print(f"🎄 Part 2 🎁: {part2(sensors)}")
