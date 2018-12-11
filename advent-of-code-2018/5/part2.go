package main

import (
	"bufio"
	"fmt"
	"os"
)

func isOpposite(a, b byte) bool {
	return (a-'A' == b-'a') || (a-'a' == b-'A')
}

func reduce(line string) string {
	chars := make([]byte, 0, len(line))
	for {
		changed := false
		for i := 0; i < len(line); i++ {
			if i == len(line)-1 {
				chars = append(chars, line[i])
				continue
			}
			if isOpposite(line[i], line[i+1]) {
				i++
				changed = true
			} else {
				chars = append(chars, line[i])
			}
		}
		line = string(chars)

		if !changed {
			break
		}
		chars = []byte{}
	}
	return string(chars)
}

func removeType(line string, lower rune) string {
	result := []rune{}
	upper := lower - 'a' + 'A'
	for _, c := range line {
		if c != lower && c != upper {
			result = append(result, c)
		}
	}
	return string(result)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	min := len(line)
	for c := 'a'; c < 'z'; c++ {
		l := len(reduce(removeType(line, c)))
		if l < min {
			min = l
		}
	}
	fmt.Println(min)
}
