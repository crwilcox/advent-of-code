import unittest
import day17 as d


class TestDay(unittest.TestCase):
    def test_part1_1(self):
        p = d.part1("test_input")
        self.assertEqual(p, 102)
    def test_part1_2(self):
        p = d.part1("test_input_2")
        self.assertEqual(p, 59)


    def test_part2_1(self):
        p = d.part2("test_input")
        self.assertEqual(p, 94)
    def test_part2_2(self):
        p = d.part2("test_input_2")
        self.assertEqual(p, 71)
    def test_part2_3(self):
        p = d.part2("test_input_3")
        self.assertEqual(p, 18)
    def test_part2_4(self):
        p = d.part2("test_input_4")
        self.assertEqual(p, 36)


    def test_part1(self):
        p = d.part1("input")
        self.assertEqual(p, 1128)
    def test_part2(self):
        p = d.part2("input")
        self.assertEqual(p, 1268)

if __name__ == "__main__":
    unittest.main()
