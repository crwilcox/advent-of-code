def hash_func(input) -> int:
    """
    Determine the ASCII code for the current character of the string.
    Increase the current value by the ASCII code you just determined.
    Set the current value to itself multiplied by 17.
    Set the current value to the remainder of dividing itself by 256.
    """
    current_value = 0
    for char in input:
        current_value += ord(char)
        current_value *= 17
        current_value %= 256

    return current_value


def part1(input):
    sum = 0
    for line in input:
        for segment in line.split(","):
            seg_hash = hash_func(segment)
            # print("Seg:", segment, "hash:", seg_hash)
            sum += seg_hash
    return sum


def part2(input):
    boxes = {}
    for line in input:
        for segment in line.split(","):
            segment_label = segment.replace("-", "=").split("=")[0]
            seg_hash = hash_func(segment_label)
            symbol = segment[len(segment_label)]

            # print(
            #     "Seg:",
            #     segment,
            #     "label:",
            #     segment_label,
            #     "symbol:",
            #     symbol,
            #     "hash:",
            #     seg_hash,
            # )

            box = boxes.get(seg_hash, {})
            if symbol == "=":
                segment_value = segment[len(segment_label) + 1 :]
                # place in box, if no other one exists
                # if not box.get(segment_label, False):
                box[segment_label] = int(segment_value)
            elif symbol == "-":
                # remove from box if exists.
                if box.get(segment_label, False):
                    del box[segment_label]
            else:
                raise Exception("input invalid?:" + symbol)
            boxes[seg_hash] = box
            # print(boxes)

        sum = 0
        for boxno, box in boxes.items():
            for slot, focal_length in enumerate(box.values()):
                s = (boxno + 1) * (slot + 1) * focal_length
                # print(
                #     f"score:{s}, box:{boxno+1}, slot:{slot+1}, focal_length:{focal_length}"
                # )
                sum += s
    return sum


def process_input(filename):
    file_lines = []
    with open(filename) as input:
        file_lines = input.readlines()
    return file_lines


if __name__ == "__main__":
    file_lines = process_input("input")

    print(f"🎄 Part 1 🎁: {part1(file_lines)}")
    print(f"🎄 Part 2 🎁: {part2(file_lines)}")
