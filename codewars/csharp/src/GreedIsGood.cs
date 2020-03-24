using System.Linq;

namespace GreedIsGood
{
    public static class Solution
    {
        public static int Score(int[] dice)
        {
            int total = 0;
            int c;
            c = dice.Count(n => n == 1);
            if (c >= 3)
            {
                total += 1000 + (c - 3) * 100;
            }
            else if (c >= 1)
            {
                total += 100 * c;
            }
            c = dice.Count(n => n == 5);
            if (c >= 3)
            {
                total += 500 + (c - 3) * 50;
            }
            else if (c >= 1)
            {
                total += 50 * c;
            }
            for (int i = 2; i <= 6; i++)
            {
                if (i == 5)
                {
                    continue;
                }
                c = dice.Count(n => n == i);
                if (c >= 3)
                {
                    total += i * 100;
                }
            }
            return total;
        }
    }
}
