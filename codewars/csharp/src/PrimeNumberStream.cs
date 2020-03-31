using System.Collections.Generic;

namespace PrimeNumberStream
{
    public class Primes
    {
        public static IEnumerable<int> Stream()
        {
            yield return 2;
            int i = 3;
            while (true)
            {
                if (isPrime(i))
                {
                    yield return i;
                }
                i += 2;
            }
        }

        public static bool isPrime(int x)
        {
            if (x == 2)
            {
                return true;
            }
            for (int i = 3; i * i <= x; i += 2)
            {
                if (x % i == 0)
                {
                    return false;
                }
            }
            return true;
        }
    }
}