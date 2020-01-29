using System.Collections.Generic;
using System.Linq;
using System;

public class FindParityOutlier
{
    public static int Find(int[] integers)
    {
        var evens = integers.Where((integer) => integer % 2 == 0);
        var odds = integers.Where((integer) => integer % 2 != 0);
        if (evens.Count() > 1) {
            return odds.First();
        }
        return evens.First();
    }
}