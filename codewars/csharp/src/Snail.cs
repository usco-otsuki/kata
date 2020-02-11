using System.Collections.Generic;
public class SnailSolution
{
    public static int[] Snail(int[][] array)
    {
        int n = array.Length;
        int start = 0;
        var list = new List<int>();
        if (n == 1 && array[0].Length == 0)
        {
            return new int[] { };
        }
        while (n > 0)
        {
            for (int i = 0; i < n; i++)
            {
                list.Add(array[start][i + start]);
            }
            for (int i = 1; i < n; i++)
            {
                list.Add(array[i + start][n - 1 + start]);
            }
            for (int i = n - 2; i >= 0; i--)
            {
                list.Add(array[n - 1 + start][i + start]);
            }
            for (int i = n - 2; i >= 1; i--)
            {
                list.Add(array[i + start][start]);
            }
            n -= 2;
            start++;
        }
        return list.ToArray();
    }
}