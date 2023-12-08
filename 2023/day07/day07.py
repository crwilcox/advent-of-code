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
            order = ["A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"]
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
        hand = {}
        has_five = False
        has_four = False
        has_three = False
        has_two = False
        has_two_pair = False
        for card in self.hand:
            if hand.get(card):
                hand[card] += 1
            else:
                hand[card] = 1

        for _, h in hand.items():
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

        if has_five:
            return "five"
        if has_four:
            return "four"
        if has_three and has_two:
            return "full"
        if has_three:
            return "three"
        if has_two_pair:
            return "two_pair"
        if has_two:
            return "one_pair"

        return "high"


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
print(f"🎄 Part 1 🎁: {scores_totaled}")
