package main

import (
	"bufio"
	"fmt"
	"os"
)

type fabric struct {
	id, x, y, w, h int
}

func overlap(f1, f2 fabric) bool {
	hOverlap := (f1.x+f1.w-1 >= f2.x) && (f2.x+f2.w-1 >= f1.x)
	vOverlap := (f1.y+f1.h-1 >= f2.y) && (f2.y+f2.h-1 >= f1.y)
	return hOverlap && vOverlap
}

func main() {
	fabrics := []fabric{}
	scanner := bufio.NewScanner(os.Stdin)
	var id, x, y, w, h int
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &id, &x, &y, &w, &h)
		fabrics = append(fabrics, fabric{id, x, y, w, h})
	}
	noOverlapIndex := -1
	for i := 0; i < len(fabrics); i++ {
		hasOverlap := false
		for j := 0; j < len(fabrics); j++ {
			if i == j {
				continue
			}
			if hasOverlap = overlap(fabrics[i], fabrics[j]); hasOverlap {
				// fmt.Printf("%d and %d has overlap\n", i+1, j+1)
				break
			}
		}
		if !hasOverlap {
			noOverlapIndex = i
			break
		}
	}
	fmt.Println(noOverlapIndex + 1)
}
