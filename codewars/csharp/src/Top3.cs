using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

public class TopWords
{
    public static List<string> Top3(string s)
    {
        Regex rx = new Regex(@"('*[a-zA-Z]'*)+", RegexOptions.Compiled | RegexOptions.IgnoreCase);
        MatchCollection matches = rx.Matches(s);
        var list = new List<string>();
        foreach (Match match in matches)
        {
            list.Add(match.Value.ToLower());
        }
        return list.GroupBy(w => w).OrderByDescending(g => g.Count()).Take(3).Select(g => g.Key).ToList();
    }
}