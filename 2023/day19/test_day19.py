import unittest
import day19 as d


class TestDay(unittest.TestCase):
    def test_part_2(self):
        workflows = {
            "in": ["A"],
        }
        self.assertEqual(d.part2(workflows), 4000**4)

        workflows = {
            "in": ["x<6:A", "next"],
            "next": ["x>5:A"],
        }
        self.assertEqual(d.part2(workflows), 4000**4)

        workflows = {
            "in": ["x<6:m", "x>3995:mm"],
            "m": ["m<6:a"],
            "a": ["a<6:s"],
            "s": ["s<6:A"],
            "mm": ["m>3995:aa"],
            "aa": ["a>3995:ss"],
            "ss": ["s>3995:A"],
        }
        self.assertEqual(d.part2(workflows), 5**4 + 5**4)

        workflows = {
            "in": ["s<1351:A", "qqz"],
            "qqz": ["m<1801:A", "R"],
        }
        # (1,4000), (1,1800), (1,4000), (1, 1350)    ==> 38,880,000,000,000
        # (1,4000), (1801,4000), (1,4000), (1, 1350) ==> 47,520,000,000,000
        # (1,4000), (1,1800), (1,4000), (1351, 4000) ==> 76,320,000,000,000
        #                                         Total:162,720,000,000,000
        self.assertEqual(d.part2(workflows), 162_720_000_000_000)

    def test_get_goal_states_from_workflow(self):
        # Simple Input
        workflows = {
            "in": ["s<1351:A", "qqz"],
            "qqz": ["m<1801:A", "R"],
        }
        expect = [
            ["s<1351"],
            ["m<1801", "s>1350"],
        ]
        self.assertEqual(d.get_goal_states_from_workflow(workflows), expect)

        # all less than 6
        workflows = {
            "in": ["x<6:m"],
            "m": ["m<6:a"],
            "a": ["a<6:s"],
            "s": ["s<6:A"],
        }
        expect = [
            ["s<6", "a<6", "m<6", "x<6"],
        ]
        self.assertEqual(d.get_goal_states_from_workflow(workflows), expect)

        # less than 6 or over 4995
        workflows = {
            "in": ["x<6:m", "x>3995:mm"],
            "m": ["m<6:a"],
            "a": ["a<6:s"],
            "s": ["s<6:A"],
            "mm": ["m>3995:aa"],
            "aa": ["a>3995:ss"],
            "ss": ["s>3995:A"],
        }
        expect = [
            ["s<6", "a<6", "m<6", "x<6"],
            ["s>3995", "a>3995", "m>3995", "x>3995", "x>5"],
        ]
        self.assertEqual(d.get_goal_states_from_workflow(workflows), expect)

        # Test Input
        workflows = {
            "px": ["a<2006:qkq", "m>2090:A", "rfg"],
            "pv": ["a>1716:R", "A"],
            "lnx": ["m>1548:A", "A"],
            "rfg": ["s<537:gd", "x>2440:R", "A"],
            "qs": ["s>3448:A", "lnx"],
            "qkq": ["x<1416:A", "crn"],
            "crn": ["x>2662:A", "R"],
            "in": ["s<1351:px", "qqz"],
            "qqz": ["s>2770:qs", "m<1801:hdj", "R"],
            "gd": ["a>3333:R", "R"],
            "hdj": ["m>838:A", "pv"],
        }
        expect = [
            ["m>2090", "a>2005", "s<1351"],
            ["s>536", "x<2441", "a>2005", "m<2091", "s<1351"],
            ["s>3448", "s>2770", "s>1350"],
            ["x<1416", "a<2006", "s<1351"],
            ["m>838", "m<1801", "s<2771", "s>1350"],
            ["a<1717", "m<839", "m<1801", "s<2771", "s>1350"],
            ["m>1548", "s<3449", "s>2770", "s>1350"],
            ["m<1549", "s<3449", "s>2770", "s>1350"],
            ["x>2662", "x>1415", "a<2006", "s<1351"],
        ]
        self.assertEqual(d.get_goal_states_from_workflow(workflows), expect)

    def test_reduce_condition_to_goal_state(self):
        goal_states = [["s<1351"], ["m<1801"]]
        expect = [
            {"x": (1, 4000), "m": (1, 4000), "a": (1, 4000), "s": (1, 1350)},
            {"x": (1, 4000), "m": (1, 1800), "a": (1, 4000), "s": (1, 4000)},
        ]
        self.assertEqual(d.reduce_condition_to_goal_state(goal_states), expect)

        goal_states = [["x<1416", "a<2006", "s>1351"]]
        expect = [{"x": (1, 1415), "m": (1, 4000), "a": (1, 2005), "s": (1352, 4000)}]
        self.assertEqual(d.reduce_condition_to_goal_state(goal_states), expect)

        goal_states = [["x>320", "x>19"]]
        expect = [{"x": (321, 4000), "m": (1, 4000), "a": (1, 4000), "s": (1, 4000)}]
        self.assertEqual(d.reduce_condition_to_goal_state(goal_states), expect)

        goal_states = [["a<2006", "a<2010"]]
        expect = [{"x": (1, 4000), "m": (1, 4000), "a": (1, 2005), "s": (1, 4000)}]
        self.assertEqual(d.reduce_condition_to_goal_state(goal_states), expect)

        # Value not stated, value more than one setting
        goal_states = [["x<1416", "x>320", "x>19", "a<2006", "a<2010", "s>1351"]]
        expect = [{"x": (321, 1415), "m": (1, 4000), "a": (1, 2005), "s": (1352, 4000)}]
        self.assertEqual(d.reduce_condition_to_goal_state(goal_states), expect)

        goal_states = [
            ["m>2090", "a>2006", "s<1351"],
            ["s>537", "x<2440", "a>2006", "m<2090", "s<1351"],
            ["s>3448", "s>2770", "s>1351"],
            ["x<1416", "a<2006", "s<1351"],
            ["m>838", "m<1801", "s<2770", "s>1351"],
            ["a<1716", "m<838", "m<1801", "s<2770", "s>1351"],
            ["m>1548", "s<3448", "s>2770", "s>1351"],
            ["m<1548", "s<3448", "s>2770", "s>1351"],
            ["x>2662", "x>1416", "a<2006", "s<1351"],
        ]
        expect = [
            {"x": (1, 4000), "m": (2091, 4000), "a": (2007, 4000), "s": (1, 1350)},
            {"x": (1, 2439), "m": (1, 2089), "a": (2007, 4000), "s": (538, 1350)},
            {"x": (1, 4000), "m": (1, 4000), "a": (1, 4000), "s": (3449, 4000)},
            {"x": (1, 1415), "m": (1, 4000), "a": (1, 2005), "s": (1, 1350)},
            {"x": (1, 4000), "m": (839, 1800), "a": (1, 4000), "s": (1352, 2769)},
            {"x": (1, 4000), "m": (1, 837), "a": (1, 1715), "s": (1352, 2769)},
            {"x": (1, 4000), "m": (1549, 4000), "a": (1, 4000), "s": (2771, 3447)},
            {"x": (1, 4000), "m": (1, 1547), "a": (1, 4000), "s": (2771, 3447)},
            {"x": (2663, 4000), "m": (1, 4000), "a": (1, 2005), "s": (1, 1350)},
        ]
        # print(d.reduce_condition_to_goal_state(goal_states))
        self.assertEqual(d.reduce_condition_to_goal_state(goal_states), expect)

    def test_get_split_points(self):
        reduced_goals = []
        expect_splits = {"a": [], "m": [], "s": [], "x": []}
        self.assertEqual(d.get_split_points(reduced_goals), expect_splits)

        reduced_goals = [
            {"x": (1, 4000), "m": (2091, 4000), "a": (1, 4000), "s": (1, 1350)},
        ]
        expect_splits = {
            "x": [1, 4000],
            "m": [2091, 4000],
            "a": [1, 4000],
            "s": [1, 1350],
        }
        self.assertEqual(d.get_split_points(reduced_goals), expect_splits)

        reduced_goals = [
            {"x": (1, 4000), "m": (1, 4000), "a": (1, 4000), "s": (1, 1350)},
            {"x": (1, 4000), "m": (1, 1800), "a": (1, 4000), "s": (1, 4000)},
        ]
        expect_splits = {
            "x": [1, 4000],
            "m": [1, 1800, 4000],
            "a": [1, 4000],
            "s": [1, 1350, 4000],
        }
        self.assertEqual(d.get_split_points(reduced_goals), expect_splits)

        reduced_goals = [
            {"x": (1, 4000), "m": (2091, 4000), "a": (1, 4000), "s": (1, 1350)},
            {"x": (1, 2000), "m": (1, 4000), "a": (1, 4000), "s": (1, 1350)},
            {"x": (1, 3000), "m": (1, 4000), "a": (1, 4000), "s": (2771, 4000)},
            {"x": (10, 1415), "m": (1, 4000), "a": (1, 2005), "s": (1, 1350)},
            {"x": (100, 4000), "m": (839, 1800), "a": (1, 4000), "s": (1, 4000)},
            {"x": (1, 4000), "m": (1, 1800), "a": (1, 4000), "s": (1, 4000)},
            {"x": (1, 4000), "m": (1549, 4000), "a": (1, 4000), "s": (2771, 4000)},
            {"x": (1, 4000), "m": (1, 4000), "a": (1, 4000), "s": (2771, 4000)},
            {"x": (2663, 4000), "m": (1, 4000), "a": (1, 2005), "s": (1, 1350)},
        ]
        expect_splits = {
            "x": [1, 10, 100, 1415, 2000, 2663, 3000, 4000],
            "m": [1, 839, 1549, 1800, 2091, 4000],
            "a": [1, 2005, 4000],
            "s": [1, 1350, 2771, 4000],
        }
        self.assertEqual(d.get_split_points(reduced_goals), expect_splits)

    def test_split_state(self):
        states = {"x": (1, 4000), "m": (), "a": (), "s": ()}
        splits = {"x": [1, 3, 5, 4000], "m": [], "a": [], "s": []}
        expect = {
            "x": {(1, 3), (4, 5), (6, 4000)},
            "m": set(),
            "a": set(),
            "s": set(),
        }
        self.assertEqual(d.split_state(states, splits), expect)

        states = {"x": (1, 4000), "m": (1, 4000), "a": (1, 4000), "s": (1, 1350)}
        splits = {
            "x": [1, 4000],
            "m": [1, 1800, 4000],
            "a": [1, 4000],
            "s": [1, 1350, 4000],
        }
        expect = {
            "x": {(1, 4000)},
            "m": {(1, 1800), (1801, 4000)},
            "a": {(1, 4000)},
            "s": {(1, 1350)},
        }
        self.assertEqual(d.split_state(states, splits), expect)

        states = {"x": (1, 4000), "m": (1800, 4000), "a": (1, 4000), "s": (1, 4000)}
        splits = {
            "x": [1, 4000],
            "m": [1, 1800, 4000],
            "a": [1, 4000],
            "s": [1, 1350, 4000],
        }
        expect = {
            "x": {(1, 4000)},
            "m": {(1800, 4000)},
            "a": {(1, 4000)},
            "s": {(1, 1350), (1351, 4000)},
        }
        self.assertEqual(d.split_state(states, splits), expect)

        states = {"x": (1, 4000), "m": (1549, 2091), "a": (1, 2005), "s": (2771, 4000)}
        splits = {
            "x": [1, 1415, 2663, 4000],
            "m": [1, 839, 1549, 1800, 2091, 4000],
            "a": [1, 2005, 4000],
            "s": [1, 1350, 2771, 4000],
        }
        expect = {
            "x": {(1, 1415), (1416, 2663), (2664, 4000)},
            "m": {(1549, 1800), (1801, 2091)},
            "a": {(1, 2005)},
            "s": {(2771, 4000)},
        }
        self.assertEqual(d.split_state(states, splits), expect)

    def test_get_sets(self):
        # Simple
        goals = [{"x": (321, 4000), "m": (1, 4000), "a": (1, 4000), "s": (1, 4000)}]
        splits = {"x": [321, 4000], "m": [1, 4000], "a": [1, 4000], "s": [1, 4000]}
        sets = [
            ((321, 4000), (1, 4000), (1, 4000), (1, 4000)),
        ]
        out = list(d.get_sets(goals, splits))
        out.sort()
        sets.sort()
        self.assertEqual(out, sets)

        goals = [
            {"x": (1, 4000), "m": (1, 4000), "a": (1, 4000), "s": (1, 1350)},
            {"x": (1, 4000), "m": (1, 1800), "a": (1, 4000), "s": (2200, 4000)},
        ]
        splits = {
            "x": [1, 4000],
            "m": [1, 1800, 4000],
            "a": [1, 4000],
            "s": [1, 1350, 2200, 4000],
        }
        sets = [
            ((1, 4000), (1, 1800), (1, 4000), (1, 1350)),
            ((1, 4000), (1, 1800), (1, 4000), (2200, 4000)),
            ((1, 4000), (1801, 4000), (1, 4000), (1, 1350)),
        ]
        out = list(d.get_sets(goals, splits))
        out.sort()
        sets.sort()
        self.assertEqual(out, sets)

        # Test Input
        goals = [
            {"x": (1, 4000), "m": (2091, 4000), "a": (2006, 4000), "s": (1, 1350)},
            {"x": (1, 2440), "m": (1, 2090), "a": (2006, 4000), "s": (537, 1350)},
            {"x": (1, 4000), "m": (1, 4000), "a": (1, 4000), "s": (3449, 4000)},
            {"x": (1, 1415), "m": (1, 4000), "a": (1, 2005), "s": (1, 1350)},
            {"x": (1, 4000), "m": (839, 1800), "a": (1, 4000), "s": (1351, 2770)},
            {"x": (1, 4000), "m": (1, 838), "a": (1, 1716), "s": (1351, 2770)},
            {"x": (1, 4000), "m": (1549, 4000), "a": (1, 4000), "s": (2771, 3448)},
            {"x": (1, 4000), "m": (1, 1548), "a": (1, 4000), "s": (2771, 3448)},
            {"x": (2663, 4000), "m": (1, 4000), "a": (1, 2005), "s": (1, 1350)},
        ]
        splits = {
            "x": [1, 1415, 2440, 2663, 4000],
            "m": [1, 838, 839, 1548, 1549, 1800, 2090, 2091, 4000],
            "a": [1, 1716, 2005, 2006, 4000],
            "s": [1, 537, 1350, 1351, 2770, 2771, 3448, 3449, 4000],
        }

        # TODO: why is the middle el duplicated?
        out = list(d.get_sets(goals, splits))

        self.assertEqual(len(out), 376)


if __name__ == "__main__":
    unittest.main()
