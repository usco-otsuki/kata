using System.Collections.Generic;
using System;
using System.Linq;
public class PyramidSlideDown
{
    public static int LongestSlideDown(int[][] pyramid)
    {
        var dp = new List<List<int>>();
        for (int i = 0; i < pyramid.Length; i++ ) {
            dp.Add(new List<int>());
            for (int j = 0; j <= i ; j++) {
                dp[i].Add(0);
            }
        }
        dp[0][0] = pyramid[0][0];
        Console.WriteLine(dp.Count);
        for (int i = 1; i < pyramid.Length; i++) {
            for (int j = 1; j <= i; j++ ) {
                dp[i][j] = Math.Max(dp[i-1][j-1], dp[i-1][j]) + pyramid[i][j];
            }

            Console.WriteLine(i);
        }
        return dp[pyramid.Length-1].Max();
    }
}
