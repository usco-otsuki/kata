package main

import (
	"strconv"
)

func Atoi(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func isValid(num string, first, second string) bool {
	fmt.Println("===")
	if (len(first) > 1 && first[0] == '0') || (len(second) > 1 && second[0] == '0') {
		return false
	}
	expected := strconv.Itoa(Atoi(first) + Atoi(second))
	prev := Atoi(second)
	for i := 0; i < len(num); {
		fmt.Println(expected)
		if !strings.HasPrefix(num[i:], expected) {
			return false
		}
		i += len(expected)
		prev, expected = Atoi(expected), strconv.Itoa(prev+Atoi(expected))
	}
	return true
}
func isAdditiveNumber(num string) bool {
	for i := 1; i < len(num); i++ {
		for j := i + 1; j < len(num); j++ {
			if isValid(num[j:], num[0:i], num[i:j]) {
				return true
			}
		}
	}
	return false
}

func main() {}
