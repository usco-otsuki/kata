from typing import List

class Solution:
    def traverse(self, x: int, y: int, grid: List[List[str]]) -> None:
        if x >= len(grid) or x < 0 or y >= len(grid[x]) or y < 0:
            return

        if grid[x][y] == "0":
            return

        grid[x][y] = "0"

        for dx, dy in [[-1, 0], [1, 0], [0, -1], [0, 1]]:
            self.traverse(x + dx, y + dy, grid)


    def numIslands(self, grid: List[List[str]]) -> int:
        cnt = 0
        for i in range(len(grid)):
            for j in range(len(grid[i])):
                if grid[i][j] == "1":
                    self.traverse(i, j, grid)
                    cnt += 1

        return cnt


if __name__ == "__main__":
    solver = Solution()
    assert solver.numIslands(
        [
            ["1","1","0","0","0"],
            ["1","1","0","0","0"],
            ["0","0","1","0","0"],
            ["0","0","0","1","1"]
        ]
    ) == 3
