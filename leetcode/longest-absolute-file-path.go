package main

import (
	"path/filepath"
)

func lengthLongestPath(input string) int {
	maxLen := 0
	names := strings.Split(input, "\n")
	if len(names) == 0 {
		return 0
	}
	context := []string{}
	for _, name := range names {
		depth := strings.Count(name, "\t")
		if depth < len(context) {
			context = context[:depth]
		}
		context = append(context, strings.TrimLeft(name, "\t"))
		fmt.Println(strings.Join(context, "/"))
		if path := strings.Join(context, "/"); filepath.Ext(name) != "" && maxLen < len(path) {
			maxLen = len(path)
		}
	}
	return maxLen
}

func main() {}
