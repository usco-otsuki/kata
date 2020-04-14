from typing import List
from functools import reduce

class Solution:
    def stringShift(self, s: str, shift: List[List[int]]) -> str:
        delta = reduce(lambda total, x: total + (x[1] if x[0] == 0 else -x[1]), shift, 0) % len(s)
        return s[delta:]+ s[:delta]

if __name__ == "__main__":
    solver = Solution()
    assert solver.stringShift("abcdefg", [[1,1],[1,1],[0,2],[1,3]]) == "efgabcd"
    assert solver.stringShift("xqgwkiqpif", [[1,4],[0,7],[0,8],[0,7],[0,6],[1,3],[0,1],[1,7],[0,5],[0,6]]) == "qpifxqgwki"
