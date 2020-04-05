from typing import List

class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        if len(prices) == 0:
            return 0

        return sum(prices[i] - prices[i-1] for i in range(1, len(prices)) if prices[i] - prices[i-1] > 0)


def main():
    solver = Solution()
    assert solver.maxProfit([7, 1, 5, 3, 6, 4]) == 7, solver.maxProfit([7, 1, 5, 3, 6, 4])
    assert solver.maxProfit([1, 2, 3, 4, 5]) == 4
    assert solver.maxProfit([7, 6, 4, 3, 1]) == 0

if __name__ == "__main__":
    main()
