using System.Numerics;

namespace BigNumbers
{
    public class Solver
    {
        public static string Add(string a, string b)
        {
            var aNum = BigInteger.Parse(a);
            var bNum = BigInteger.Parse(b);
            var calcResult = aNum + bNum;
            return calcResult.ToString();
        }
    }
}