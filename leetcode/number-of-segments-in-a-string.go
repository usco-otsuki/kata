package main

import (
	"regexp"
	"strings"
)

func countSegments(s string) int {
	trimmed := strings.TrimSpace(s)
	if len(trimmed) == 0 {
		return 0
	}
	segs := regexp.MustCompile("\\s+").Split(trimmed, -1)
	return len(segs)
}

func main() {}
