public static class TrailingZeros
{
    public static int Solve(int n)
    {
        var sum = 0;
        while (n / 5 > 0)
        {
            sum += n / 5;
            n /= 5;
        }
        return sum;
    }
}