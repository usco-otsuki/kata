import fileinput
from typing import List, Tuple, NewType, Dict, Callable

Instruction = Tuple[int, int, int, int]
Register = Tuple[int, int, int, int, int, int]
OpsFunction = Callable[[Register, int, int, int], Register]

class Operation:

    @staticmethod
    def get_ops_list() -> Dict[str, OpsFunction]:
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
    def addr(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = reg_list[in1] + reg_list[in2]
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def addi(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = reg_list[in1] + in2
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def mulr(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = reg_list[in1] * reg_list[in2]
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def muli(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = reg_list[in1] * in2
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def banr(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = reg_list[in1] & reg_list[in2]
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def bani(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = reg_list[in1] & in2
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def borr(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = reg_list[in1] | reg_list[in2]
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def bori(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = reg_list[in1] | in2
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def setr(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = reg_list[in1]
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def seti(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = in1
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def gtir(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = 1 if in1 > reg_list[in2] else 0
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def gtri(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = 1 if reg_list[in1] > in2 else 0
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def gtrr(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = 1 if reg_list[in1] > reg_list[in2] else 0
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def eqir(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = 1 if in1 == reg_list[in2] else 0
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def eqri(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = 1 if reg_list[in1] == in2 else 0
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

    @staticmethod
    def eqrr(reg: Register, in1: int, in2: int, out: int) -> Register:
        reg_list = list(reg)
        reg_list[out] = 1 if reg_list[in1] == reg_list[in2] else 0
        return (reg_list[0], reg_list[1], reg_list[2], reg_list[3], reg_list[4], reg_list[5])

def main() -> None:
    lines = [line.strip() for line in fileinput.input()]
    ip_reg = int(lines[0].split()[1])
    ip = 0
    lines = lines[1:]
    instructions: List[Tuple[str, int, int, int]] = []
    for line in lines:
        ops, in1_str, in2_str, out_str = line.split()
        instructions.append((ops, int(in1_str), int(in2_str), int(out_str)))

    reg: Register = (0, 0, 0, 0, 0, 0)
    while ip < len(instructions):
        ops, in1, in2, out = instructions[ip]
        reg = reg[:ip_reg] + (ip,) + reg[ip_reg+1:]
        reg = Operation.get_ops_list()[ops](reg, in1, in2, out)
        ip = reg[ip_reg] + 1

    print(reg[0])

if __name__ == "__main__":
    main()
