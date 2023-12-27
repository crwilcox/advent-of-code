import math
import queue

from enum import Enum


class NodeType(Enum):
    BROADCASTER = 1
    FLIPFLOP = 2
    CONJUNCTION = 3


class Node:
    def __init__(self, name, children, type, state):
        self.name = name
        self.children = children
        self.type = type
        self.state = state
        self.input_states = {}


def press_button(nodes, button_press_count) -> (int, int, dict):
    """
    Return low count, high count, was feed to kz high (this is the in to rx, used for lcm test)
    """
    q = queue.Queue()

    low_send_count = 0
    high_send_count = 0
    high_to_kz = {}
    # False -> Low
    # True -> High
    q.put(("button", "broadcaster", False))

    while not q.empty():
        src, curr_node, pulse = q.get()
        # print(f"{src} -{pulse}-> {curr_node}")
        if pulse:
            high_send_count += 1
        else:
            low_send_count += 1

        if curr_node not in nodes:
            if curr_node == "rx" and not pulse:
                return low_send_count, high_send_count, True
            # not a real node, or rx getting high, just skip
            continue

        type = nodes[curr_node].type
        if type == NodeType.CONJUNCTION:
            # update pulse
            nodes[curr_node].input_states[src] = pulse
            # if it remembers high pulses for all inputs, it sends a low pulse;
            # otherwise, it sends a high pulse.
            send = False

            # for solving part 2. Return cycles to calc LCM
            if curr_node == "kz":
                for key, state in nodes[curr_node].input_states.items():
                    if state:
                        high_to_kz[key] = True

            for _, state in nodes[curr_node].input_states.items():
                if not state:
                    send = True
                    break
            for child in nodes[curr_node].children:
                q.put((curr_node, child, send))
        elif type == NodeType.FLIPFLOP:
            # Flip-flop modules (prefix %) are either on or off; they are initially off.
            # If a flip-flop module receives a high pulse, it is ignored and nothing happens.
            # However, if a flip-flop module receives a low pulse, it flips between on and off.
            # If it was off, it turns on and sends a high pulse. If it was on, it turns off and sends a low pulse.
            if pulse == False:
                nodes[curr_node].state = not nodes[curr_node].state
                pulse_to_send = False
                if nodes[curr_node].state:
                    pulse_to_send = True

                for child in nodes[curr_node].children:
                    q.put((curr_node, child, pulse_to_send))
        else:  # broadcaster
            for child in nodes[curr_node].children:
                q.put((curr_node, child, pulse))

    # print(f"High: {high_send_count}, Low: {low_send_count}")
    return low_send_count, high_send_count, high_to_kz


def part1(nodes: dict[Node]):
    low_send_count = 0
    high_send_count = 0

    for press in range(1000):
        low, high, _ = press_button(nodes, press)
        low_send_count += low
        high_send_count += high

    return low_send_count * high_send_count


def part2(nodes):
    """
    I tried letting this happen exhaustively but realized after 10 minutes, I can
    likely look at the input and be clever. This is a bit bespokely coded as is but
    could be made more general. I initially added a print statement
    to the solver, and aborted running once we saw the cycles. Then I calc'd LCM out of band
    sj: 3918, 7837 ==> 3919
    qq: 4002, 8005, 12008  ==> 4003
    ls: 3796, 7593 ==> 3797
    bg:  3738, 7477 ==> 3739
    # LCM : 222718819437131
    As authored, it now does the LCM calc, though it still depends on kz being the
    node that goes to rx, and that it is a conjuction.
    """
    low_send_count = 0
    high_send_count = 0

    # NOTE: this could be generalized to find the conjunction node that
    # points to rx, then find it's sources. For my input, it is "kz" and
    # kz has 4 inputs. Once we find those 4, we can calc LCM
    conjunction_el_count = 4
    cycles = {}
    for press in range(10_000):
        low, high, high_to_kz = press_button(nodes, press)
        low_send_count += low
        high_send_count += high
        if high_to_kz != {}:
            for element, _ in high_to_kz.items():
                c = cycles.get(element, [])
                c.append(press)
                cycles[element] = c

            LCMS = []
            for k, v in cycles.items():
                if len(v) >= 2:
                    LCMS.append(v[1] - v[0])
            if len(LCMS) == conjunction_el_count:
                return math.lcm(*LCMS)


def process_input(filename):
    file_lines = []
    with open(filename) as input:
        file_lines = input.readlines()
    lines = [i.strip() for i in file_lines]

    nodes = {}
    for l in lines:
        s = l.split(" -> ")
        name = s[0]
        children = s[1].split(", ")
        type = ""
        on = False
        if name[0] == "%":
            type = NodeType.FLIPFLOP
            name = name[1:]
        elif name[0] == "&":
            type = NodeType.CONJUNCTION
            name = name[1:]
        else:
            type = NodeType.BROADCASTER
        nodes[name] = Node(name, children, type, on)
    # walk over nodes, so we can put the srcs to each node in the list.
    for _, node in nodes.items():
        for child in node.children:
            if child in nodes:
                nodes[child].input_states[node.name] = False
    # print(nodes)
    return nodes


if __name__ == "__main__":
    input_name = "input"

    node_map = process_input(input_name)
    print(f"🎄 Part 1 🎁: {part1(node_map)}")

    node_map = process_input(input_name)
    print(f"🎄 Part 2 🎁: {part2(node_map)}")
