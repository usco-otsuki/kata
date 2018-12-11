package main

import (
	"bufio"
	"fmt"
	"os"
)

func isOpposite(a, b byte) bool {
	return (a-'A' == b-'a') || (a-'a' == b-'A')
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
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
	// fmt.Println(string(chars))
	fmt.Println(len(chars))
}
