package main

import (
	"fmt"
	"io"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`#\d+ @ (\d+),(\d+): (\d+)x(\d+)`)
	var line string
	for {
		_, err := fmt.Scanln(&line)
		if err == io.EOF {
			matches := r.FindStringSubmatch(line)
			break
		}
		x, y, w, h := matches[1:]
	}
}
