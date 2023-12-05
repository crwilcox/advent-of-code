cards = []
with open("input") as input:
    for line in input.readlines():
        # Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
        card_and_numbers = line.split(":")[1].strip().split("|")
        card = card_and_numbers[0].strip().split(" ")
        drawn = card_and_numbers[1].strip().split(" ")
        card = [c for c in card if c != ""]
        drawn = [c for c in drawn if c != ""]
        cards.append((card, drawn, 1))


def wins_from_card(card, drawn):
    matches = 0
    for n in card:
        if n in drawn:
            print("MATCH:", n)
            matches += 1
    return matches


for idx, _ in enumerate(cards):
    card, drawn, copies = cards[idx]
    matches = wins_from_card(card, drawn)

    for i in range(matches):
        if idx + i + 1 < len(cards):
            a, b, c = cards[idx + i + 1]
            c += copies  # increase count of cards.
            cards[idx + i + 1] = (a, b, c)

# Card Count
sum = 0
for _, _, count in cards:
    sum += count

print(f"🎄 Part 2 🎁: {sum}")
