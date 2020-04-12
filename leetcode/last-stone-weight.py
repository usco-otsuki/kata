from typing import List


class Solution:
    def lastStoneWeight(self, stones: List[int]) -> int:
        while len(stones) > 1:
            stones.sort()
            new_stone = stones[-1] - stones[-2]
            stones.pop()
            stones.pop()
            stones.append(new_stone)

        return stones[0]


if __name__ == "__main__":
    solver = Solution()
    assert solver.lastStoneWeight([2,7,4,1,8,1]) == 1
