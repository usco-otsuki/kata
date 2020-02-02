using System;
using System.Linq;
using System.Collections.Generic;

public class ValidBraces {
    public static bool validBraces(String braces) {
      var stack = new List<char>();
      for (int i = 0; i < braces.Count(); i++) {
        switch (braces[i]) {
            case '{':
            case '[':
            case '(':
                stack.Insert(0, braces[i]);
                break;
            default:
                if (stack.Count() == 0 ) {
                    return false;
                }
                var ch = stack[0];
                stack.RemoveAt(0);
                if (braces[i] == '}' && ch != '{') {
                    return false;
                }
                if (braces[i] == ']' && ch != '[') {
                    return false;
                }
                if (braces[i] == ')' && ch != '(') {
                    return false;
                }
                break;
        }
      }
      if (stack.Count() > 0) {
          return false;
      }
      return true;
    }
}