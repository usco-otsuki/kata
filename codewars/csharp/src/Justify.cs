using System.Linq;
using System.Collections.Generic;
using System;
public class Justifier
{
    public static string JustifyLine(List<string> words, int len)
    {
        if (words.Count == 1)
        {
            return words[0];
        }
        string currentStr = words[0];
        int lenForWhiteSpace = len - words.Sum(word => word.Length);
        int minSpaceLen = lenForWhiteSpace / (words.Count - 1);
        int extraSpaceCnt = lenForWhiteSpace % (words.Count - 1);
        for (int i = 1; i < words.Count; i++)
        {
            for (int j = 0; j < minSpaceLen; j++)
            {
                currentStr += " ";
            }
            if (extraSpaceCnt > 0)
            {
                currentStr += " ";
                extraSpaceCnt--;
            }
            currentStr += words[i];
        }
        return currentStr;
    }
    public static string Justify(string str, int len)
    {
        if (str == null) {
            return "";
        }
        var words = str.Split(new char[] { ' ', '\n' });
        int i = 0;
        var lineWords = new List<string> { };
        string finalResult = "";
        while (i < words.Count())
        {
            lineWords.Add(words[i]);
            if (string.Join(" ", lineWords).Length > len)
            {
                lineWords.RemoveAt(lineWords.Count() - 1);
                string justifiedLine = JustifyLine(lineWords, len);
                finalResult += justifiedLine + "\n";
                lineWords = new List<string> { };
            } else {
                i++;
            }
        }
        if (lineWords.Count > 0) {
            finalResult += string.Join(" ", lineWords);
        }
        return finalResult;
    }
}