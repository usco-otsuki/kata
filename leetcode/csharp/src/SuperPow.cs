using System.Numerics;
using System.Linq;

namespace SuperPow
{
    public class Solution
    {
        static readonly int MOD = 1337;
        public int SuperPow(int a, int[] b)
        {
            var exp = BigInteger.Parse(string.Join("", b.Select(digit => digit.ToString())));
            int baseval = a % MOD;
            int r = 1;
            while (exp > 0)
            {
                if (exp % 2 == 1)
                {
                    r = (r * baseval) % MOD;
                }
                exp = exp >> 1;
                baseval = (baseval * baseval) % MOD;
            }
            return r;
        }
    }
}
