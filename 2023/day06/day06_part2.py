file_lines = []
with open("input") as input:
    file_lines = input.readlines()

time = int(file_lines[0].split(":")[1].strip().replace(" ", ""))
distance = int(file_lines[1].split(":")[1].strip().replace(" ", ""))


def ways_to_win(time, distance):
    valid_way = 0
    for i in range(time):
        time_of_movement = time - i
        actual_distance = time_of_movement * i
        if actual_distance > distance:
            valid_way += 1
            # print(f"WIN: {i} hold for {actual_distance}")
    return valid_way


ways = ways_to_win(time, distance)

print(f"🎄 Part 2 🎁: {ways}")
