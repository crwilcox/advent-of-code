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


locations = []
for seed in seeds:
    soil = find_lowest_map(seed_to_soil_map, int(seed))
    fertilizer = find_lowest_map(soil_to_fertilizer_map, soil)
    water = find_lowest_map(fertilizer_to_water_map, fertilizer)
    light = find_lowest_map(water_to_light_map, water)
    temperature = find_lowest_map(light_to_temperature_map, light)
    humidity = find_lowest_map(temperature_to_humidity_map, temperature)
    location = find_lowest_map(humidity_to_location_map, humidity)
    print(
        f"Seed {seed}, soil {soil}, fertilizer {fertilizer}, water {water}, light {light}, temperature {temperature}, humidity {humidity}, location {location}."
    )
    locations.append(location)

min_location = min(locations)

print(f"🎄 Part 1 🎁: {min_location}")
