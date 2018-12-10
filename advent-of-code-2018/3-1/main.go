package main

import (
	"bufio"
	"fmt"
	"os"
)

type coordinate struct {
	x, y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	m := make(map[coordinate]int)
	var id, x, y, w, h int
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		var _ = id
		for i := 0; i < w; i++ {
			for j := 0; j < h; j++ {
				m[coordinate{x + i, y + j}]++
			}
		}
	}
	sum := 0
	for _, count := range m {
		if count > 1 {
			sum++
		}
	}
	fmt.Println(sum)
}
