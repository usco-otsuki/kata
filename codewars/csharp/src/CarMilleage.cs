using System.Collections.Generic;
using System.Linq;
using System;
public static class CarMileage
{
    public static bool IsPalindrome(string str)
    {
        for (int i = 0; i < str.Length / 2; i++)
        {
            if (str[i] != str[str.Length - i - 1])
            {
                return false;
            }
        }
        return true;
    }
    public static bool IsSequential(int number, int direction)
    {
        int next = number.ToString()[0] - '0';
        string numStr = number.ToString();
        for (int i = 0; i < numStr.Length; i++)
        {
            var ch = numStr[i];
            if (next == 0 && i < numStr.Length - 1)
            {
                return false;
            }
            var n = ch - '0';
            if (n != next)
            {
                return false;
            }

            next = (next + direction + 10) % 10;
        }
        return true;
    }
    public static int IsInteresting(int number, List<int> awesomePhrases)
    {
        if (_IsInteresting(number, awesomePhrases))
        {
            return 2;
        }
        if (_IsInteresting(number + 1, awesomePhrases) || _IsInteresting(number + 2, awesomePhrases))
        {
            return 1;
        }
        return 0;
    }
    public static bool _IsInteresting(int number, List<int> awesomePhrases)
    {
        if (number < 100)
        {
            return false;
        }


        if (number.ToString().Substring(1).Count(c => c == '0') == number.ToString().Length - 1)
        {
            return true;
        }
        if (number.ToString().Count(c => c == number.ToString()[0]) == number.ToString().Length)
        {
            return true;
        }
        if (IsSequential(number, 1) || IsSequential(number, -1))
        {
            return true;
        }
        if (IsPalindrome(number.ToString()))
        {
            return true;
        }
        foreach (var phrase in awesomePhrases)
        {
            if (phrase == number)
            {
                return true;
            };
        }
        return false;
    }
}