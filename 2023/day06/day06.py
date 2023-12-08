file_lines = []
with open("input") as input:
    file_lines = input.readlines()

times = [int(t) for t in file_lines[0].split(":")[1].strip().split(" ") if t != ""]
distances = [int(d) for d in file_lines[1].split(":")[1].strip().split(" ") if d != ""]


def ways_to_win(time, distance):
    valid_way = 0
    for i in range(time):
        time_of_movement = time - i
        actual_distance = time_of_movement * i
        if actual_distance > distance:
            valid_way += 1
            # print(f"WIN: {i} hold for {actual_distance}")
    return valid_way


result = 1
for idx, time in enumerate(times):
    print("Race:", idx)
    distance = distances[idx]
    ways = ways_to_win(time, distance)
    result *= ways

print(f"🎄 Part 1 🎁: {result}")
