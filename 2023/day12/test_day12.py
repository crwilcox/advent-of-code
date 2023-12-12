import unittest
import day12 as d


class TestDay(unittest.TestCase):
    def test_different_arrangements(self):
        tests = [
            ("???.### 1,1,3", 1),
            (".??..??...?##. 1,1,3", 4),
            ("?#?#?#?#?#?#?#? 1,3,1,6", 1),
            ("????.#...#... 4,1,1", 1),
            ("????.######..#####. 1,6,5", 4),
            ("?###???????? 3,2,1", 10),
        ]
        for test, want in tests:
            p, s = d.parse_line(test)
            got = d.count_valid(p, s)
            self.assertEqual(got, want, test)

    def test_pattern_valid(self):
        self.assertTrue(d.is_pattern_valid("#.#.#.#", [1, 1, 1, 1]))
        self.assertTrue(d.is_pattern_valid("######.#.#.#.", [6, 1, 1, 1]))
        self.assertTrue(d.is_pattern_valid("####.#...#...", [4, 1, 1]))
        # too much last group
        self.assertFalse(d.is_pattern_valid("######.#.#.########", [6, 1, 1, 1]))
        # extra group
        self.assertFalse(d.is_pattern_valid("######.#.#.#.######", [6, 1, 1, 1]))

    def test_unfold(self):
        scenarios = [(".#", [1], ".#?.#?.#?.#?.#", [1, 1, 1, 1, 1])]
        for p, s, ep, es in scenarios:
            p, s = d.unfold(p, s)
            self.assertEqual(p, ep)
            self.assertEqual(s, es)

    def test_part2_arrangments(self):
        scenarios = [(".#??", (1, 1), 16), ("...?", (1,), 16)]
        for p, s, expected in scenarios:
            p, s = d.unfold(p, s)
            count = d.count_valid(p, tuple(s))
            self.assertEqual(count, expected, p)

    def test_part2(self):
        lines = d.process_input("test_input")
        self.assertEqual(d.part2(lines), 525152)


if __name__ == "__main__":
    unittest.main()
