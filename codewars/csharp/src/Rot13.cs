using System;
using System.Collections.Generic;

public class Rot13
{
  public static string Solve(string input)
  {
      var chars = new List<char>();
      foreach( char c in input ) {
          if (!Char.IsLetter(c)) {
            chars.Add(c);
            continue;
          }
          char b = 'a';
          if ( 'A' <= c && c <= 'Z') {
              b = 'A';
          } 
          chars.Add((char)(((c - b)  + 13) % 26 + b));
      }
      return new string(chars.ToArray());
  }
}