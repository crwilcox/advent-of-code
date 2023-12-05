total = 0
seeds = []
seed_to_soil_map = []
soil_to_fertilizer_map = []
fertilizer_to_water_map = []
water_to_light_map = []
light_to_temperature_map = []
temperature_to_humidity_map = []
humidity_to_location_map = []

file_lines = []
with open("input") as input:
    file_lines = input.readlines()


def get_map(file_lines, idx):
    ret_map = []
    # ignore first line, it's the tithle
    idx += 1

    while idx < len(file_lines) and len(file_lines[idx].strip()) > 0:
        ret_map.append(file_lines[idx].strip().split(" "))
        idx += 1

    return ret_map


idx = 0
while idx < len(file_lines):
    line = file_lines[idx]
    if line.startswith("seeds:"):
        seeds = line.split(":")[1].strip().split(" ")
        print("Seeds:", seeds)
    elif line.startswith("seed-to-soil"):
        seed_to_soil_map = get_map(file_lines, idx)
        idx += len(seed_to_soil_map)
    elif line.startswith("soil-to-fertilizer"):
        soil_to_fertilizer_map = get_map(file_lines, idx)
        idx += len(soil_to_fertilizer_map)
    elif line.startswith("fertilizer-to-water"):
        fertilizer_to_water_map = get_map(file_lines, idx)
        idx += len(fertilizer_to_water_map)
    elif line.startswith("water-to-light"):
        water_to_light_map = get_map(file_lines, idx)
        idx += len(water_to_light_map)
    elif line.startswith("light-to-temperature"):
        light_to_temperature_map = get_map(file_lines, idx)
        idx += len(light_to_temperature_map)
    elif line.startswith("temperature-to-humidity"):
        temperature_to_humidity_map = get_map(file_lines, idx)
        idx += len(temperature_to_humidity_map)
    elif line.startswith("humidity-to-location"):
        humidity_to_location_map = get_map(file_lines, idx)
        idx += len(humidity_to_location_map)
    idx += 1


def find_lowest_map(map, value):
    lowest = value
    for entry in map:
        dest_range_start = int(entry[0])
        source_range_start = int(entry[1])
        range = int(entry[2])

        if source_range_start <= value and source_range_start + range > value:
            lowest = value + dest_range_start - source_range_start

    return lowest


def find_lowest_map_reverse(map, value):
    lowest = value
    for entry in map:
        dest_range_start = int(entry[0])
        source_range_start = int(entry[1])
        range = int(entry[2])

        if dest_range_start <= value and dest_range_start + range > value:
            lowest = value + source_range_start - dest_range_start

    return lowest


def seed_exists(seed):
    i = 0
    while i < len(seeds):
        seed_val = int(seeds[i])
        count = int(seeds[i + 1])
        if seed >= seed_val and seed <= seed_val + count:
            return True
        i += 2
    return False


print("This may take a moment...")
location = 0
while True:
    # instead of starting from seeds, work backwards be starting at a min
    # location, seeing if a seed could exist to get there.

    humidity = find_lowest_map_reverse(humidity_to_location_map, location)
    temperature = find_lowest_map_reverse(temperature_to_humidity_map, humidity)
    light = find_lowest_map_reverse(light_to_temperature_map, temperature)
    water = find_lowest_map_reverse(water_to_light_map, light)
    fertilizer = find_lowest_map_reverse(fertilizer_to_water_map, water)
    soil = find_lowest_map_reverse(soil_to_fertilizer_map, fertilizer)
    seed = find_lowest_map_reverse(seed_to_soil_map, soil)

    if seed_exists(seed):
        print("LOCATION:", location)
        break
    location += 1
    if location % 10000 == 0:
        print(f"\r{location}", end="")

# 2520479.
print(f"🎄 Part 2 🎁: {location}")
