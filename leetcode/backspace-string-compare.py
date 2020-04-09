from typing import List
from functools import reduce


class Solution:
    def backspaceCompare(self, S: str, T: str) -> bool:
        s_result = reduce(lambda a, b: a[:-1] if b == "#" else a + b, "*" + S)
        t_result = reduce(lambda a, b: a[:-1] if b == "#" else a + b, "*" + T)
        if s_result[0:1] == "*":
            s_result = s_result[1:]

        if t_result[0:1] == "*":
            t_result = t_result[1:]

        return s_result == t_result


if __name__ == "__main__":
    solver = Solution()
    tests = [
        ["ab#c", "ad#c", True],
        ["ab##", "c#d#", True],
        ["a##c", "#a#c", True],
        ["a#c", "b", False],
        ["y#fo##f", "y#f#o##f", True]
    ]
    for test in tests:
        actual = solver.backspaceCompare(test[0], test[1])
        assert  actual == test[2], f"actual = {actual}, expected = {test[2]}"
