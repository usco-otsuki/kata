package main

import (
	"fmt"
	"io"
)

func FindCommon(s1, s2 string) (string, error) {
	diffIndex := -1
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			if diffIndex >= 0 {
				return "", fmt.Errorf("They differ by more than one character")
			}
			diffIndex = i
		}
	}
	if diffIndex == -1 {
		return "", fmt.Errorf("They are exactyle the same")
	}
	return s1[:diffIndex] + s1[diffIndex+1:], nil
}

func main() {
	var s string
	strings := []string{}
	for {
		_, err := fmt.Scan(&s)
		if err == io.EOF {
			break
		}
		strings = append(strings, s)
	}
	for i := 0; i < len(strings); i++ {
		for j := i + 1; j < len(strings); j++ {
			common, err := FindCommon(strings[i], strings[j])
			if err == nil {
				fmt.Println(common)
				return
			}
		}
	}
}
