using System.Collections.Generic;
using System.Linq;
using System;

public class Kata
{
  public static string SpinWords(string sentence) {
    var spinned = new List<string>();
    string[] words = sentence.Split();
    for (int i = 0; i < words.Length; i++) {
      string word = words[i];
      if (word.Length >= 5) {
        char[] chars = word.ToCharArray();
        Array.Reverse<char>(chars);
        spinned.Add(new string(chars));
      } else {
        spinned.Add(word);
      }
    }
    return string.Join(' ', spinned);
  }	  
}
