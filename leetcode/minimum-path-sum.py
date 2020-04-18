from typing import List

class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        dp = [[0] * len(row) for row in grid]
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                min_comp_target = []
                if i - 1 >= 0:
                    min_comp_target.append(dp[i-1][j])

                if j - 1 >= 0:
                    min_comp_target.append(dp[i][j-1])

                dp[i][j] = grid[i][j] + min(min_comp_target, default=0)

        return dp[len(grid) - 1][len(grid[0]) - 1]


if __name__ == "__main__":
    solver = Solution()
    assert solver.minPathSum(
        [
           [1,3,1],
           [1,5,1],
           [4,2,1]
        ]
    ) == 7
