
# Target Game
game = {"red": 12, "green":13, "blue":14}

sum = 0
# with open("input") as input:
with open("input") as input:
    for line in input.readlines():
        # Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
        s = line.split(":")
        game_id = int(s[0][5:])
        draws = s[1].split(";")

        possible = True
        game = {"red": 0, "green":0, "blue":0}
        for draw in draws:
            dice = draw.split(",")
            for die in dice:
                d = die.strip().split(" ")
                count = int(d[0])
                color = d[1]

                game[color] = max(game[color], count)
        powers = game["red"] * game["green"] * game["blue"]
       
        sum += powers

print(f"🎄 Part 2 🎁: {sum}")
