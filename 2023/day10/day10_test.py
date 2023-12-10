import unittest
import day10

pt1_1_map = """.....
.S-7.
.|.|.
.L-J.
.....""".split(
    "\n"
)

pt1_2_map = """..F7.
.FJ|.
SJ.L7
|F--J
LJ...""".split(
    "\n"
)


# Expect 4 Tiles Enclosed
pt2_3_map = """..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........""".split(
    "\n"
)

# Expect 8 Tiles Enclosed
pt2_4_map = """OF----7F7F7F7F-7OOOO
O|F--7||||||||FJOOOO
O||OFJ||||||||L7OOOO
FJL7L7LJLJ||LJIL-7OO
L--JOL7IIILJS7F-7L7O
OOOOF-JIIF7FJ|L7L7L7
OOOOL7IF7||L7|IL7L7|
OOOOO|FJLJ|FJ|F7|OLJ
OOOOFJL-7O||O||||OOO
OOOOL---JOLJOLJLJOOO""".split(
    "\n"
)

# Expect 10 tiles enclosed.
pt2_5_map = """FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJIF7FJ-
L---JF-JLJIIIIFJLJJ7
|F|F-JF---7IIIL7L|7|
|FFJF7L7F-JF7IIL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L""".split(
    "\n"
)


class TestDay(unittest.TestCase):
    def test_part1_1(self):
        p1 = day10.part1(pt1_1_map)
        self.assertEqual(p1, 4)

    def test_part1_2(self):
        p1 = day10.part1(pt1_2_map)
        self.assertEqual(p1, 8)

    def test_part2_3(self):
        day10.part1(pt2_3_map)
        p2 = day10.part2(pt2_3_map)
        self.assertEqual(p2, 4)

    def test_part2_4(self):
        day10.part1(pt2_4_map)
        p2 = day10.part2(pt2_4_map)
        self.assertEqual(p2, 8)

    def test_part2_5(self):
        day10.part1(pt2_5_map)
        p2 = day10.part2(pt2_5_map)
        self.assertEqual(p2, 10)


if __name__ == "__main__":
    unittest.main()
