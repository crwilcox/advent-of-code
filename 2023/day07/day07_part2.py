file_lines = []
with open("input") as input:
    file_lines = input.readlines()


class Player:
    def __init__(self, hand, bid):
        self.hand = hand
        self.bid = int(bid)

    def __repr__(self) -> str:
        return f"{self.hand} {self.bid} {self._hand_type()}"

    def __lt__(self, other):
        # five of kind, four of kind, full house, three of kind, two pair, one pair, high card
        self_hand = self._hand_type()
        other_hand = other._hand_type()

        if self_hand == other_hand:
            # Note, order is different, j is now low and joker.
            order = ["A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"]
            for idx, self_card in enumerate(self.hand):
                other_card = other.hand[idx]
                if self_card == other_card:
                    continue
                else:
                    # comparator is backwards because order is backwards also
                    return order.index(self_card) > order.index(other_card)
            print("IDENTICAL HANDS. WHAT?")

        ranking = ["five", "four", "full", "three", "two_pair", "one_pair", "high"]
        self_rank = ranking.index(self_hand)
        other_rank = ranking.index(other_hand)
        return self_rank > other_rank

    def _hand_type(self):
        card_counts = {}

        has_five = False
        has_four = False
        has_three = False
        has_two = False
        has_two_pair = False
        has_one = False

        for card in self.hand:
            if card_counts.get(card):
                card_counts[card] += 1
            else:
                card_counts[card] = 1

        if "J" in card_counts:
            del card_counts["J"]
        # jokers shouldn't be used as a standalone card count, we'll use to boost score.

        # print(f"{self.hand}: {card_counts}")
        for k, h in card_counts.items():
            if h == 5:
                has_five = True
            elif h == 4:
                has_four = True
            elif h == 3:
                has_three = True
            elif h == 2:
                if has_two:
                    has_two_pair = True
                has_two = True
            elif h == 1:
                has_one = True

        # Unmodified
        ret = "jokers"
        if has_five:
            ret = "five"
        elif has_four:
            ret = "four"
        elif has_three and has_two:
            ret = "full"
        elif has_three:
            ret = "three"
        elif has_two_pair:
            ret = "two_pair"
        elif has_two:
            ret = "one_pair"
        elif has_one:
            ret = "high"

        # Jokers can modify the result
        joker_count = self.hand.count("J")
        if joker_count == 0:
            return ret

        if ret == "jokers":
            # this is unusual case of all "J"
            return "five"
        elif ret == "two_pair":
            if joker_count == 1:
                # if two pair and one J, it's a full house now.
                return "full"
            else:
                return "two_pair"

        promote = ["jokers", "high", "one_pair", "three", "four", "five", "five"]
        # get the index of a hand, then promote by joker count.

        curr = promote.index(ret)
        if curr < 0:
            # this is a full house, which will be promoted to four.
            print("CURR IS UNSET:", ret)
            curr = promote.index("three")
        if curr + joker_count >= len(promote):
            return "five"
        ret = promote[curr + joker_count]
        # print("RESULT:", ret)
        return ret


players = []
for line in file_lines:
    el = line.split(" ")
    players.append(Player(el[0], int(el[1])))

players.sort()
scores_totaled = 0

for idx, p in enumerate(players):
    rank = idx + 1
    # print(f"rank:{rank} bid:{p.bid}")
    scores_totaled += p.bid * rank


# print(sorted(players))
print(f"🎄 Part 2 🎁: {scores_totaled}")
