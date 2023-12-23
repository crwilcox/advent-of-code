import queue


def find_accepted_parts(workflows, parts):
    accepted = []
    for part in parts:
        debugprint("part:", part)
        curr_workflow = "in"

        while curr_workflow:
            debugprint("workflow:", curr_workflow)
            if curr_workflow == "R":
                break
            elif curr_workflow == "A":
                accepted.append(part)
                break

            for cond in workflows[curr_workflow]:
                if cond == "R" or cond == "A" or ":" not in cond:
                    curr_workflow = cond
                    break
                else:
                    condition = cond.split(":")[0]
                    next_workflow = cond.split(":")[1]

                    s = condition.replace("<", ">").split(">")
                    attribute = s[0]
                    val = int(s[1])
                    gt_or_lt = condition[len(attribute) : len(attribute) + 1]

                    debugprint("attr:", attribute, "gt_or_lt:", gt_or_lt, "val:", val)

                    if gt_or_lt == ">" and part[attribute] > val:
                        curr_workflow = next_workflow
                        break
                    elif gt_or_lt == "<" and part[attribute] < val:
                        curr_workflow = next_workflow
                        break

    debugprint("ACCEPTED, Len:", len(accepted), accepted)
    return accepted


def part1(workflows, parts):
    accepted_parts = find_accepted_parts(workflows, parts)

    sum = 0
    for p in accepted_parts:
        for _, val in p.items():
            sum += val
    return sum


DEBUG = False


def debugprint(*args):
    if DEBUG:
        print(*args)


# need to track end state searching for, conditions already known to get there
# which workflow_name, which state saerching for, known conditions already.
def get_goal_states_from_workflow(workflows):
    """
    Produces list of goal states, still expressed as strs
    [['m>2090', 's<1351'], ['s<1351'], ['s>3448', 's>2770'], ['x<1416', 'a<2006', 's<1351'], ['m>838', 'm<1801'], ['m<1801'], ['m>1548', 's>2770'], ['s>2770'], ['x>2662', 'a<2006', 's<1351']]
    """
    q = queue.Queue()

    q.put(("A", "", []))
    goal_states = []
    while not q.empty():
        target, rule_to_search, conditions_to_meet = q.get()
        debugprint(
            f"q: target:{target}, rule:{rule_to_search}, cond:{conditions_to_meet}"
        )
        """
        # Simple Input
        workflows = {
            "in": ["s<1351:A", "qqz"],
            "qqz": ["m<1801:A", "R"],
        }
        expect = [
            ["s<1351"],
            ["m<1801"],
        ]
        """
        # if goal_state == "in":
        if target == "in":
            debugprint(f"-> Adding to goal_states: {conditions_to_meet}")
            if conditions_to_meet in goal_states:
                print("WHY DUPE?")
                pass
            else:
                goal_states.append(conditions_to_meet)
            continue

        workflows_to_search = {}
        if rule_to_search:
            workflows_to_search[target] = workflows[target]
        else:
            workflows_to_search = workflows.items()

        for workflow_name, workflow_rules in workflows.items():
            for idx, rule in enumerate(workflow_rules):
                # find rule that matches our goal state, append it's condition if it has one.
                if rule == target or rule.endswith(":" + target):
                    c = conditions_to_meet.copy()
                    if ":" in rule:
                        c.append(rule.split(":")[0])
                    # we also need to take the inverse rule of all of the rules before this one.
                    other_rules = workflow_rules[:idx]
                    for other_rule in other_rules:
                        other_rule = other_rule.split(":")[0]
                        if ">" in other_rule:
                            before = other_rule[: other_rule.index(">")]
                            number = int(other_rule[other_rule.index(">") + 1 :])
                            debugprint(
                                f"other rule inverse: before:{other_rule} after:{before}<{number+1}"
                            )
                            other_rule = f"{before}<{number+1}"

                        elif "<" in other_rule:
                            before = other_rule[: other_rule.index("<")]
                            number = int(other_rule[other_rule.index("<") + 1 :])
                            debugprint(
                                f"other rule inverse: before:{other_rule} after:{before}>{number-1}"
                            )
                            other_rule = f"{before}>{number-1}"

                        c.append(other_rule)
                    debugprint("-> qput", workflow_name, target, c)
                    q.put((workflow_name, target, c))

    debugprint("#######################################")
    debugprint("goal states:", goal_states)
    debugprint("#######################################")
    return goal_states


def reduce_condition_to_goal_state(goal_states):
    """
    BEFORE: ['x<1416', 'a<2006', 's<1351']
    AFTER: {'x': (1, 1415), 'm': (1, 4000), 'a': (1, 2005), 's': (1, 1350)}
    """
    reduced_goal_states = []

    # goals states need to be reduced. Each char has a < and > than possibility.
    for goal_state in goal_states:
        reduced_goal_state = {
            "x": (1, 4000),
            "m": (1, 4000),
            "a": (1, 4000),
            "s": (1, 4000),
        }

        for condition in goal_state:
            split = condition.replace("<", ">").split(">")
            ltr = split[0]
            val = int(split[1])
            gt_or_lt = condition[len(ltr) : len(ltr) + 1]

            ltr_min, ltr_max = reduced_goal_state[ltr]
            if gt_or_lt == ">":
                ltr_min = max(ltr_min, val + 1)
            elif gt_or_lt == "<":
                ltr_max = min(ltr_max, val - 1)

            reduced_goal_state[ltr] = (ltr_min, ltr_max)

        reduced_goal_states.append(reduced_goal_state)
    debugprint(reduced_goal_states)
    return reduced_goal_states


def get_split_points(reduced_goal_states):
    """
    Walks reduced goal states and produces split points to divide goal states on
    {
        'x': [1, 1416, 2662, 4000], 'm': [1, 838, 1548, 1801, 2090, 4000],
        'a': [1, 2006, 4000], 's': [1, 1351, 2770, 4000]
    }
    """
    # print("REDUCED GOALS:", reduced_goal_states)
    splits = {"x": set(), "m": set(), "a": set(), "s": set()}
    for gs in reduced_goal_states:
        for ltr, val in gs.items():
            ltr_min, ltr_max = val
            splits[ltr].add(ltr_min)
            splits[ltr].add(ltr_max)

    for ltr, s in splits.items():
        s = list(s)
        s.sort()
        splits[ltr] = s
        # print("l:", ltr, "set:", s)

    # print("Splits:", splits)
    return splits


def split_state(state: dict[str, tuple], splits: dict[str, list]) -> set():
    """
    Takes a state and divides it on the splits, returning sets that represent the complete range

    ex:
        state:{
            'x': (1, 4000), 'm': (1549, 4000), 'a': (1, 4000), 's': (2771, 4000)}
        splits:{
            'x': [1, 1415, 2663, 4000],
            'm': [1, 839, 1549, 1800, 2091, 4000],
            'a': [1, 2005, 4000],
            's': [1, 1350, 2771, 4000]}
        returns: {
            'x': {(2664, 4000), (1, 1415), (1416, 2663)},
            'm': {(1549, 1800), (2092, 4000), (1801, 2091)},
            'a': {(1, 2005), (2006, 4000)},
            's': {(2771, 4000)}}
    """
    # print(f"splitstates: st:{state}, sp:{splits}")
    divided = {}
    for ltr in ["x", "m", "a", "s"]:
        new_states = set()

        prev_split = None
        for split in splits[ltr]:
            # find starting split
            x_min, x_max = state[ltr]
            if split < x_min:
                continue
            if prev_split is None:
                prev_split = split
                continue

            new_states.add((prev_split, split))
            prev_split = split + 1

            if split >= x_max:
                break
        divided[ltr] = new_states
        # print(f"OUT:\n  state:{state}\n  splits:{splits}\n  new_states:{new_states}")
        # print("divided:", divided)

    return divided


def get_sets(reduced_goal_states, splits):
    gs_count = len(reduced_goal_states)
    for idx, s in enumerate(reduced_goal_states):
        print(f"Expanding goal states to individual sets: {idx}/{gs_count}", end="\r")
        states = split_state(s, splits)
        debugprint("states:", states)

        # If sets are made first they don't need to be made in the loop.
        x_set = set(states["x"])
        m_set = set(states["m"])
        a_set = set(states["a"])
        s_set = set(states["s"])

        for x in x_set:
            for m in m_set:
                for a in a_set:
                    for s in s_set:
                        yield ((x, m, a, s))
    print()
    return


def part2(workflows: dict):
    """
    Brute forcing this is going to take too long. If we find all "A" and
    backtrack the combinations to get to exits can be known, those ranges
    can then be taken to figure out the total possible.
    """
    ##########################################################################
    # Get Goal State from conditions, walk back from "A" results.
    goal_states = get_goal_states_from_workflow(workflows)
    print("Found Goal States...", len(goal_states))
    debugprint("GOAL STATES:", goal_states)

    # Take the goal state produced, reduce it to tuple ranges,
    # ['m>2090', 's<1351'] => "m": (2091,4000), "s": (1,1350)
    reduced_goal_states = reduce_condition_to_goal_state(goal_states)
    print("Reduced Goal States...", len(reduced_goal_states))
    debugprint("REDUCED GOAL STATES:", reduced_goal_states)

    # the above splits are large. Find all split points of all goal
    # states, use that to break these up into sets that can be
    # matched. This avoids double counting.
    splits = get_split_points(reduced_goal_states)
    print("Found Splits...", len(splits))
    debugprint("SPLITS:", splits)

    sets = get_sets(reduced_goal_states, splits)
    print(
        "Retrieved Sets...",
    )

    # Test Calculating Here:
    total_poss = 0
    for v in sets:
        x, m, a, s = v
        x_min, x_max = x
        m_min, m_max = m
        a_min, a_max = a
        s_min, s_max = s

        # Because min and max are inclusive, a +1 is needed.
        inner_res = (
            (x_max - x_min + 1)
            * (m_max - m_min + 1)
            * (a_max - a_min + 1)
            * (s_max - s_min + 1)
        )
        debugprint(v, ":", inner_res)
        total_poss += inner_res
    return total_poss


def process_rule(line):
    s = line.split("{")
    workflow_name = s[0]
    condition_str = s[1][:-1]
    conditions = [cond for cond in condition_str.split(",")]

    return workflow_name, conditions


def process_input(filename):
    """
    Returns input in the form:
    workflows: dict[str] => list(rules)
    parts: dict[str]=>int
    """
    file_lines = []
    with open(filename) as input:
        file_lines = input.readlines()
    file_lines = [i.strip() for i in file_lines]

    workflows = {}
    for i in file_lines:
        if i == "":
            break
        workflow_name, rules = process_rule(i)
        workflows[workflow_name] = rules

    parts = []
    for part in file_lines[len(workflows) + 1 :]:
        part = part[1:-1]
        dict_part = {}
        for p in part.split(","):
            pp = p.split("=")
            dict_part[pp[0]] = int(pp[1])
        parts.append(dict_part)

    return workflows, parts


if __name__ == "__main__":
    workflows, parts = process_input("test_input")

    print(f"🎄 Test Part 1 🎁: {part1(workflows, parts)}")
    print(f"🎄 Test Part 2 🎁: {part2(workflows):_}")

    workflows, parts = process_input("input")
    print(f"🎄 Part 1 🎁: {part1(workflows, parts)}")
    print(f"🎄 Part 2 🎁: {part2(workflows):_}")
