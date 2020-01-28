using System;
using System.Collections.Generic;

public class DnaStrand {
  public static readonly Dictionary<char, char> dict = new Dictionary<char, char>() {
    {'A', 'T'},
    {'T', 'A'},
    {'G', 'C'},
    {'C', 'G'},
  };
  public static string MakeComplement(string dna){
    var list = new List<char>{};
    for (int i = 0; i < dna.Length; i++) {
      list.Add(dict[dna[i]]);
    }
    return string.Join("", list);
  }
}
