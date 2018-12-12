package main

import (
	"bufio"
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var x, y, maxX, maxY int
	coordinates := []point{}
	for scanner.Scan() {
		fmt.Sscanf(scanner.Text(), "%d, %d", &x, &y)
		coordinates = append(coordinates, point{x, y})
		if maxX < x {
			maxX = x
		}
		if maxY < y {
			maxY = y
		}
	}
	isInfinite := make(map[int]bool)
	counter := make(map[int]int)
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			closestID := -1
			closestDist := maxX + maxY
			for cID, c := range coordinates {
				dist := abs(c.x-i) + abs(c.y-j)
				if closestDist >= dist {
					if closestDist == dist {
						closestID = -1
					} else {
						closestID = cID
					}
					closestDist = dist
				}
			}
			if closestID == -1 {
				continue
			}
			if i == 0 || j == 0 || i == maxX || j == maxY {
				isInfinite[closestID] = true
			}
			counter[closestID]++
		}
	}
	maxCount := 0
	for i := 0; i < len(coordinates); i++ {
		if _, ok := isInfinite[i]; ok {
			continue
		}
		if maxCount < counter[i] {
			maxCount = counter[i]
		}
	}
	fmt.Println(maxCount)
}
