namespace AsFarFromLandAsPossible
{
    public class Solution
    {
        public int MaxDistance(int[][] grid)
        {
            int [,] dp = new int[grid.GetLength(0), grid.GetLength(1)];
            for(int i = 0; i < grid.GetLength(0); i++) {
                for (int j = 0; j < grid.GetLength(1); j++) {
                    if (grid[i][j] == 1) {
                        dp[i,j] = 0;
                        continue;
                    }
                }
            }
        }
    }
}