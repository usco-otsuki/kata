using System;
using System.Collections.Generic;

public class Kata {
    public static bool IsIsogram(string str) {
        str = str.ToLower();
        Dictionary<char, bool> exists = new Dictionary<char, bool>();
	for (int i = 0; i < str.Length; i++) {
	    if (exists.ContainsKey(str[i])) {
	        return false;
	    }
	    exists[str[i]] = true;
	}
	return true;
    }
}
