using System;
using System.Collections.Generic;

public class WeightSort
{
    private static long sum(long x)
    {
        long sum = 0;
        while (x > 0)
        {
            sum += x % 10;
            x /= 10;
        }
        return sum;
    }
    public static string orderWeight(string strng)
    {
        var nums = strng.Split(new char[] { ' ' });
        var list = new List<string>(nums);
        list.Sort(delegate (string x, string y)
        {
            var xs = sum(long.Parse(x));
            var ys = sum(long.Parse(y));
            if (xs == ys)
            {
                return x.CompareTo(y);
            }
            if (xs < ys)
            {
                return -1;
            }
            return 1;
        });
        return string.Join(" ", list);
    }
}
