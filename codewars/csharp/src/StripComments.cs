using System;
using System.Linq;
public class StripCommentsSolution
{
    public static string StripComments(string text, string[] commentSymbols)
    {
        string[] texts = text.Split('\n');
        return string.Join("\n", texts.Select(t => t.Split(commentSymbols, StringSplitOptions.None).First().TrimEnd()));
    }
}