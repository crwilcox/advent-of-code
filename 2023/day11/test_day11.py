import unittest
import day11 as d

test_input = """...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....""".split(
    "\n"
)


class TestDay(unittest.TestCase):
    def test_expand_space(self):
        r = d.expand_space(test_input)
        self.assertEqual(len(r), len(test_input) + 2)

    def test_find_galaxies(self):
        g = d.find_galaxies(test_input)
        self.assertEqual(
            g, [(0, 3), (1, 7), (2, 0), (4, 6), (5, 1), (6, 9), (8, 7), (9, 0), (9, 4)]
        )

        s = d.expand_space(test_input)
        g = d.find_galaxies(s)
        self.assertEqual(
            g,
            [
                (0, 4),
                (1, 9),
                (2, 0),
                (5, 8),
                (6, 1),
                (7, 12),
                (10, 9),
                (11, 0),
                (11, 5),
            ],
        )

    def test_distance_between_galaxies(self):
        self.assertEqual(d.calculate_distance_between_galaxies((0, 4), (10, 9)), 15)
        self.assertEqual(d.calculate_distance_between_galaxies((2, 0), (7, 12)), 17)
        self.assertEqual(d.calculate_distance_between_galaxies((11, 0), (11, 5)), 5)
        self.assertEqual(d.calculate_distance_between_galaxies((6, 1), (11, 5)), 9)

    def test_part1(self):
        p1 = d.part1(test_input)
        self.assertEqual(p1, 374)

    def test_space_adjust(self):
        # find galaxies, verify we are starting where we expect
        space = test_input
        g = d.find_galaxies(space)
        self.assertEqual(
            g, [(0, 3), (1, 7), (2, 0), (4, 6), (5, 1), (6, 9), (8, 7), (9, 0), (9, 4)]
        )

        g2 = []
        for gal in g:
            row, col = gal
            a = d.adjust_coordinate_for_expanded_space(space, row, col, multiplier=2)
            g2.append(a)

        self.assertEqual(
            g2,
            [
                (0, 4),
                (1, 9),
                (2, 0),
                (5, 8),
                (6, 1),
                (7, 12),
                (10, 9),
                (11, 0),
                (11, 5),
            ],
        )

    def test_part2(self):
        p = d.part2(test_input, multiplier=10)
        self.assertEqual(p, 1030)

        p = d.part2(test_input, multiplier=100)
        self.assertEqual(p, 8410)

        # p = d.part2(input, multiplier=1_000_000)
        # self.assertEqual(p, 678728808158)


if __name__ == "__main__":
    unittest.main()
