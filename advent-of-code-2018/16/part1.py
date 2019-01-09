import fileinput
from typing import Tuple, NewType
import re

Instruction = NewType("Instruction", Tuple[int, int, int, int])
Register = NewType("Instruction", Tuple[int, int, int, int])

class Operation:

    @staticmethod
    def get_ops_list():
        return [
            Operation.addr,
            Operation.addi,
            Operation.mulr,
            Operation.muli,
            Operation.banr,
            Operation.bani,
            Operation.borr,
            Operation.bori,
            Operation.setr,
            Operation.seti,
            Operation.gtir,
            Operation.gtri,
            Operation.gtrr,
            Operation.eqir,
            Operation.eqri,
            Operation.eqrr
        ]

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


def count_possible_ops(before: Register, after: Register, instruction: Instruction) -> int:
    count = 0
    for method in Operation.get_ops_list():
        result = method(before, instruction[1], instruction[2], instruction[3])
        if result == after:
            count += 1

    return count

def main():
    status = 0
    count = 0
    before: Register = (0, 0, 0, 0)
    after: Register = (0, 0, 0, 0)
    instruction: Instruction = (0, 0, 0, 0)
    for line in fileinput.input():
        line = line.strip()
        if status == 0 and line.startswith("Before"): # Expect Before
            before = tuple(map(int, re.search(r"(\d+), (\d+), (\d+), (\d+)", line).groups()))
        elif status == 1 and re.match(r"^\d", line): # Expect instruction
            instruction = tuple(map(int, line.split()))
        elif status == 2 and line.startswith("After"): # Expect After
            after = tuple(map(int, re.search(r"(\d+), (\d+), (\d+), (\d+)", line).groups()))
            if count_possible_ops(before, after, instruction) >= 3:
                count += 1

        elif status == 3 and len(line) == 0: # Expect empty line
            pass
        else:
            break

        status = (status + 1) % 4

    print(count)

if __name__ == "__main__":
    main()





