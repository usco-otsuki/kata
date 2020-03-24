using System;
using System.Linq;

namespace Pangram
{
    public static class Solution
    {
        public static bool IsPangram(string str)
        {
            return str.Where(c => Char.ToLower(c) >= 'a' && Char.ToLower(c) <= 'z').GroupBy(c => Char.ToLower(c)).Count() == 26;
        }
    }
}
