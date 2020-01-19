using System;
using System.Collections.Generic;

namespace myApp
{
    public static class Kata {
        public static string AlphabetPosition(string text)
        {
            List<string> nums = new List<string>();
            for (int i = 0; i < text.Length; i++) {
                char c = text[i].ToString().ToLower()[0];
                int diff = c - 'a' + 1;
                if (diff >= 1 && diff <= 26) {
                    nums.Add(diff.ToString());
                }
            }
            return string.Join(" ", nums);
        }
    }
    class ReplaceWithAlphabetPosition 
    {
        static void Main(string[] args)
        {
            Console.WriteLine(Kata.AlphabetPosition("abc"));
        }
    }
}
