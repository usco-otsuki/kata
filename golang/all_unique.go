package main

import "fmt"

func HasUniqueChar(str string) bool {
	exists := map[byte]bool{}
	for i := 0; i < len(str); i++ {
		if exists[str[i]] {
			return false
		}
		exists[str[i]] = true
	}
	return true
}

func main() {
	fmt.Println(HasUniqueChar("abac"))
}
