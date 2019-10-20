package main

func isValid(s string) bool {
	stack := []rune{}
	matchTable := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, c := range s {
		switch c {
		case '(', '{', '[':
			stack = append(stack, c)
		case ')', '}', ']':
			if !(len(stack) > 0 && stack[len(stack)-1] == matchTable[c]) {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {}
