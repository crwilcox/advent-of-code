from math import gcd


def process_input(filename):
    file_lines = []
    with open(filename) as input:
        file_lines = input.readlines()
    directions = file_lines[0].strip()
    return directions, file_lines[2:]


directions, nodes = process_input("input")


class Node:
    def __init__(self, current, left, right):
        self.current = current
        self.left = left
        self.right = right


node_map = {}
for line in nodes:
    # AAA = (BBB, CCC)
    s0 = line.replace(" ", "").replace("(", "").replace(")", "")
    s1 = s0.split("=")
    node = s1[0].strip()
    s2 = s1[1].strip().split(",")
    left = s2[0]
    right = s2[1]
    node_map[node] = Node(node, left, right)


def solve_lcm(steps):
    lcm = 1
    for s in steps:
        lcm = lcm * s // gcd(lcm, s)
    k = lcm
    # print("STEPS TO REACH LCM:", steps)
    return k


def get_steps_to_destination():
    steps_to_end = []
    curr_nodes = []
    # discover curr_nodes
    for k in node_map:
        if k.endswith("A"):
            curr_nodes.append(k)

    # Brute force is going to take too long.
    # However, we can use LCM when one is found, then use that to
    # determine the end number.
    steps = 0
    while True:
        for dir in directions:
            steps += 1
            # print("STEP:", steps, curr_nodes)
            next_nodes = []
            for curr_node in curr_nodes:
                if dir == "L":
                    curr_node = node_map[curr_node].left
                elif dir == "R":
                    curr_node = node_map[curr_node].right

                # If ended, don't add to next nodes, record steps to end.
                if not curr_node.endswith("Z"):
                    next_nodes.append(curr_node)
                else:
                    steps_to_end.append(steps)

            curr_nodes = next_nodes
            if len(curr_nodes) == 0:
                return steps_to_end


def part1():
    curr_node = "AAA"
    steps = 0
    while True:
        for dir in directions:
            steps += 1
            if dir == "L":
                curr_node = node_map[curr_node].left
            elif dir == "R":
                curr_node = node_map[curr_node].right
            else:
                print("WAT!")

            # print("NODE", curr_node)
            if curr_node == "ZZZ":
                return steps


def part2():
    steps = get_steps_to_destination()
    print(f"Able to hit an end in {steps[-1]} steps")
    lcm = solve_lcm(steps)
    return lcm


print(f"🎄 Part 1 🎁: {part1()}")
print(f"🎄 Part 2 🎁: {part2()}")
