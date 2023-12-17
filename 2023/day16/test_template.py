import unittest
import template as d

pt_1_1 = """.....
.S-7.
.|.|.
.L-J.
.....""".split(
    "\n"
)


class TestDay(unittest.TestCase):
    def test_part1(self):
        p1 = d.part1(pt_1_1)
        self.assertEqual(p1, 4)


if __name__ == "__main__":
    unittest.main()
