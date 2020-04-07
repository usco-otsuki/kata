from typing import List
from collections import Counter

class Solution:
    def countElements(self, arr: List[int]) -> int:
        counts = Counter(arr)
        return sum(count for value, count in counts.items() if value + 1 in counts)

if __name__ == "__main__":
    solver = Solution()
    tests = [
        [[1,2,3], 2],
        [[1,3,2,3,5,0], 3],
        [[1,1,2,2], 2]
    ]
    for test in tests:
        assert solver.countElements(test[0]) == test[1], f"actual = {solver.countElements(test[0])}, expected = {test[1]}"
