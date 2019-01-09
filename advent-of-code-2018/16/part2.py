import fileinput
from typing import Tuple, NewType, List
import re
from functools import reduce

Instruction = NewType("Instruction", Tuple[int, int, int, int])
Register = NewType("Instruction", Tuple[int, int, int, int])

class Operation:

    @staticmethod
    def get_ops_list():
        return {
            "addr": Operation.addr,
            "addi": Operation.addi,
            "mulr": Operation.mulr,
            "muli": Operation.muli,
            "banr":Operation.banr,
            "bani": Operation.bani,
            "borr": Operation.borr,
            "bori": Operation.bori,
            "setr": Operation.setr,
            "seti": Operation.seti,
            "gtir": Operation.gtir,
            "gtri": Operation.gtri,
            "gtrr": Operation.gtrr,
            "eqir": Operation.eqir,
            "eqri": Operation.eqri,
            "eqrr": Operation.eqrr
        }

    @staticmethod
    def addr(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = reg[in1] + reg[in2]
        return tuple(reg)

    @staticmethod
    def addi(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = reg[in1] + in2
        return tuple(reg)

    @staticmethod
    def mulr(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = reg[in1] * reg[in2]
        return tuple(reg)

    @staticmethod
    def muli(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = reg[in1] * in2
        return tuple(reg)

    @staticmethod
    def banr(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = reg[in1] & reg[in2]
        return tuple(reg)

    @staticmethod
    def bani(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = reg[in1] & in2
        return tuple(reg)

    @staticmethod
    def borr(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = reg[in1] | reg[in2]
        return tuple(reg)

    @staticmethod
    def bori(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = reg[in1] | in2
        return tuple(reg)

    @staticmethod
    def setr(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = reg[in1]
        return tuple(reg)

    @staticmethod
    def seti(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = in1
        return tuple(reg)

    @staticmethod
    def gtir(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = 1 if in1 > reg[in2] else 0
        return tuple(reg)

    @staticmethod
    def gtri(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = 1 if reg[in1] > in2 else 0
        return tuple(reg)

    @staticmethod
    def gtrr(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = 1 if reg[in1] > reg[in2] else 0
        return tuple(reg)

    @staticmethod
    def eqir(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = 1 if in1 == reg[in2] else 0
        return tuple(reg)

    @staticmethod
    def eqri(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = 1 if reg[in1] == in2 else 0
        return tuple(reg)

    @staticmethod
    def eqrr(reg, in1, in2, out):
        reg = list(reg)
        reg[out] = 1 if reg[in1] == reg[in2] else 0
        return tuple(reg)


def possible_ops(before: Register, after: Register, instruction: Instruction) -> List[str]:
    ops = []
    for name, method in Operation.get_ops_list().items():
        result = method(before, instruction[1], instruction[2], instruction[3])
        if result == after:
            ops.append(name)

    return ops

def identify_ops(candidates_lists):
    names_lists = [None] * len(candidates_lists)
    names = [None] * len(candidates_lists)
    ones = set()
    for i, candidates in enumerate(candidates_lists):
        intersection = reduce(lambda c1, c2: c1 & c2, candidates)
        if len(intersection) == 1:
            names[i] = intersection.pop()
            ones.add(names[i])
            names_lists[i] = []
        else:
            names_lists[i] = intersection

    while len(ones) < len(candidates_lists):
        for i, names_list in enumerate(names_lists):
            if len(names_list) <= 1:
                continue
            
            names_lists[i] = names_list - ones
            if len(names_lists[i]) == 1:
                names[i] = names_lists[i].pop() 
                ones.add(names[i])

    assert(len(set(names)) == len(names))
    print(names)
    return names

def main():
    status = 0
    before: Register = (0, 0, 0, 0)
    after: Register = (0, 0, 0, 0)
    instruction: Instruction = (0, 0, 0, 0)
    candidates_lists = [[] for i in range(16)]
    sample_tests = []
    for line in fileinput.input():
        line = line.strip()
        if status == 0 and line.startswith("Before"): # Expect Before
            before = tuple(map(int, re.search(r"(\d+), (\d+), (\d+), (\d+)", line).groups()))
        elif status == 1 and re.match(r"^\d", line): # Expect instruction
            instruction = tuple(map(int, line.split()))
        elif status == 2 and line.startswith("After"): # Expect After
            after = tuple(map(int, re.search(r"(\d+), (\d+), (\d+), (\d+)", line).groups()))
            candidates = possible_ops(before, after, instruction)
            candidates_lists[instruction[0]].append(set(candidates))

        elif status == 3 and len(line) == 0: # Expect empty line
            pass
        else:
            if line == "":
                continue
            
            sample_tests.append(list(map(int, line.split())))
            continue


        status = (status + 1) % 4

    op_names = identify_ops(candidates_lists)
    reg = (0, 0, 0, 0)
    for op_num, in1, in2, out in sample_tests:
        line = line.strip()
        if line == "":
            continue

        op = op_names[op_num]
        reg = Operation.get_ops_list()[op](reg, in1, in2, out)

    print(reg)

if __name__ == "__main__":
    main()