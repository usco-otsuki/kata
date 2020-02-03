using System;

public class PileOfCubes
{
    public static long findNb(long m)
    {
        long sum = 0;
        long i = 1;
        while (sum < m)
        {
            sum += i * i * i;
            i++;
        }
        return sum == m ? i - 1 : -1;
    }
}