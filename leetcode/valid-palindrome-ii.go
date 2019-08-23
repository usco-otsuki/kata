package main

import "fmt"

func validPalindrome(s string) bool {
	deleted := false
	p1, p2 := 0, len(s)-1
	for p1 <= p2 {
		if s[p1] != s[p2] {
			if deleted {
				return false
			}
			if p1+1 <= p2 && s[p1+1] == s[p2] && (p1+2 > p2-1 || (p1+2 <= p2-1 && s[p1+2] == s[p2-1])) {
				deleted = true
				p1++
			} else if p1 <= p2-1 && s[p1] == s[p2-1] && (p1+1 > p2-2 || (p1+1 <= p2-2 && s[p1+1] == s[p2-2])) {
				deleted = true
				p2--
			} else {
				return false
			}
		}
		p1++
		p2--
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("dmaadedaeeddeeadedafad") == false)
	fmt.Println(validPalindrome("cuucu") == true)
	fmt.Println(validPalindrome("abbad") == true)
	fmt.Println(validPalindrome("abbadc") == false)
}
