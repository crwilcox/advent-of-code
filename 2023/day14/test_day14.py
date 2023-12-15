import unittest
import day14 as d


def split_lines(lines):
    return [i.strip() for i in lines.split("\n")]


input = split_lines(
    """O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#...."""
)

output = split_lines(
    """OOOO.#.O..
OO..#....#
OO..O##..O
O..#.OO...
........#.
..#....#.#
..O..#.O.O
..O.......
#....###..
#....#...."""
)


def str_to_list(s):
    return [i for i in s]


class TestDay(unittest.TestCase):
    def test_column_slide(self):
        res = d.tilt_col_north("OO.O.O..##")
        self.assertEqual(res, str_to_list("OOOO....##"))

        res = d.tilt_col_north(".O...#O..O")
        self.assertEqual(res, str_to_list("O....#OO.."))

    def test_tilt_north(self):
        # p1 = d.part1(input)
        tilt = d.tilt_north(input)
        self.assertEqual(tilt, output)

    def test_tilt_west(self):
        tilt = d.tilt_west(["..O#.O#..O"])
        self.assertEqual(tilt, [str_to_list("O..#O.#O..")])

    def test_tilt_east(self):
        tilt = d.tilt_east(["O..#O.#O.."])
        self.assertEqual(tilt, [str_to_list("..O#.O#..O")])

    def test_calc_load(self):
        self.assertEqual(d.calc_load(output), 136)

    def test_part1(self):
        self.assertEqual(d.part1(input), 136)


if __name__ == "__main__":
    unittest.main()
