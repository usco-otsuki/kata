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

func calcDistSum(p point, coordinates []point) int {
	sum := 0
	for _, c := range coordinates {
		sum += abs(c.x-p.x) + abs(c.y-p.y)
	}
	return sum
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
	count := 0
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			if calcDistSum(point{i, j}, coordinates) < 10000 {
				count++
			}
		}
	}
	fmt.Println(count)
}
