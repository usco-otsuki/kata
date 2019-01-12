import fileinput
from typing import List, NewType, Tuple
from collections import Counter

Field = List[List[str]]

directions: List[Tuple[int, int]] = [(-1, 0), (1, 0), (0, -1), (0, 1), (-1, -1), (-1, 1), (1, -1), (1, 1)]

s = set()

def print_field(field: Field):
    for row in field[1:len(field)-1]:
        print("".join(row[1:len(row)-1]))

def solve(field: Field, minutes: int) -> int:
    for i in range(minutes):
        # print(f"Processing {i+1}-th minute")
        process_minute(field)
        count = Counter([cell for row in field for cell in row])
        total = count["|"] * count["#"]
        s.add(total)
        print(f'{i+1}, {total}, {len(s)}')
        # print_field(field)

    count = Counter([cell for row in field for cell in row])
    print(count)
    return count["|"] * count["#"]


def process_minute(field: Field):
    changes: List[Tuple[int, int, int]] = [] # (row, column, new value)
    for row in range(1, len(field) - 1):
        assert len(field[row]) == len(field[0])
        for column in range(1, len(field[row]) - 1):
            count = Counter(field[row + direction[0]][column + direction[1]] for direction in directions)
            assert sum(count.values()) == 8, count
            value = field[row][column]
            newvalue = value
            assert len(value) >= 1, value
            if value == ".":
                if count["|"] >= 3:
                    newvalue = "|"
            elif value == "|":
                if count["#"] >= 3:
                    newvalue = "#"
            else:
                assert value == "#", value
                if count["#"] >= 1 and count["|"] >= 1:
                    newvalue = "#"
                else:
                    newvalue = "."

            changes.append((row, column, newvalue))

    for row, column, newval in changes:
        field[row][column] = newval

def main():
    field: Field = []
    for line in fileinput.input():
        line = line.strip()
        row = ["w"] + list(line) + ["w"]
        field.append(row)

    field.insert(0, ["w"] * len(field[0]))
    field.append(["w"] * len(field[0]))

    print(solve(field, 1000))
    # You can see that after 598 minutes the sequence of length 28 repeats

if __name__ == "__main__":
    main()
