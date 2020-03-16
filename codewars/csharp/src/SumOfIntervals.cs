using System;

public class Intervals
{
    public static int SumIntervals((int, int)[] intervals)
    {
        var sum = 0;
        if (intervals.Length == 0)
        {
            return 0;
        }
        Array.Sort(intervals);
        var currentInterval = intervals[0];
        foreach (var interval in intervals)
        {
            if (currentInterval.Item2 < interval.Item1)
            {
                sum += currentInterval.Item2 - currentInterval.Item1;
                currentInterval = interval;
            }
            else
            {
                currentInterval = (currentInterval.Item1, Math.Max(currentInterval.Item2, interval.Item2));
            }
        }
        sum += currentInterval.Item2 - currentInterval.Item1;
        return sum;
    }
}