using System.Collections.Generic;
using System;
using System.Linq;
public class PyramidSlideDown
{
    public static int LongestSlideDown(int[][] pyramid)
    {
        var dp = new List<List<int>>();

        for (int i = 0; i < pyramid.Length; i++)
        {
            dp.Add(new List<int>());
            for (int j = 0; j <= i; j++)
            {
                dp[i].Add(0);
            }
            dp[i].Add(0); // extra value for preventing index error
        }
        dp[0][0] = pyramid[0][0];
        for (int i = 1; i < pyramid.Length; i++)
        {
            for (int j = 0; j <= i; j++)
            {
                dp[i][j] = Math.Max(dp[i - 1][Math.Max(0, j - 1)], dp[i - 1][j]) + pyramid[i][j];
            }
        }
        return dp[pyramid.Length - 1].Max();
    }
}
