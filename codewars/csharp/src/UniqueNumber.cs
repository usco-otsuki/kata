using System.Collections.Generic;
using System.Linq;

public class UniqueNumber
{
    public static int GetUnique(IEnumerable<int> numbers)
    {
        if (numbers.Count(number => number == numbers.First()) == 1)
        {
            return numbers.First();
        }
        return numbers.Where(number => number != numbers.First()).First();
    }
}