package main

import (
	"fmt"
	"io"
)

func TwoAndThree(str string) (bool, bool) {
	twoFlag := false
	threeFlag := false
	count := map[rune]int{}
	for _, char := range str {
		count[char]++
	}
	for _, freq := range count {
		if freq == 2 {
			twoFlag = true
		} else if freq == 3 {
			threeFlag = true
		}
	}
	return twoFlag, threeFlag
}

func main() {
	var val string
	twoCount := 0
	threeCount := 0
	for {
		_, err := fmt.Scan(&val)
		if err == io.EOF {
			break
		}
		two, three := TwoAndThree(val)
		if two {
			twoCount++
		}
		if three {
			threeCount++
		}
	}
	fmt.Println(twoCount * threeCount)
}
