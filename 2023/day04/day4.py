total = 0
with open("input") as input:
    for line in input.readlines():
        # Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
        card_and_numbers = line.split(":")[1].strip().split("|")
        card = card_and_numbers[0].strip().split(" ")
        drawn = card_and_numbers[1].strip().split(" ")
        card = [c for c in card if c != ""]
        drawn = [c for c in drawn if c != ""]
        print("CARD:", card)
        print("DRAWN:", drawn)
        matches = 0
        for n in card:
            if n in drawn:
                print("MATCH:", n)
                matches += 1

        print(matches)
        if matches > 0:
            total += pow(2, matches - 1)


print(f"🎄 Part 1 🎁: {total}")
