using System.Collections.Generic;
using System.Linq;

public class Hamming
{
    public static long hamming(int n)
    {
        var numbers = new List<long> { 1 };
        int cnt = 1;
        int i = 0, j = 0, k = 0;
        while (cnt < n)
        {
            while ( numbers[i] * 2 <= numbers.Last()) {
                i++;
            }
            while ( numbers[j] * 3 <= numbers.Last()) {
                j++;
            }
            while ( numbers[k] * 5 <= numbers.Last()) {
                k++;
            }
            numbers.Add(new List<long>{numbers[i] * 2, numbers[j] * 3, numbers[k] * 5}.Min());
            int min = 
            cnt++;
        }
        return numbers.Last();
    }
}