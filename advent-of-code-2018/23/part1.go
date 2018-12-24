package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type coordinate struct {
	x, y, z int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var x, y, z, r int
	maxR := -1
	maxIndex := -1
	index := 0
	coordinates := []coordinate{}
	for scanner.Scan() {
		n, err := fmt.Sscanf(scanner.Text(), "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
		if n != 4 {
			log.Fatalf("Scanned %d, want 4", n)
		}
		if err != nil {
			panic(err)
		}
		coordinates = append(coordinates, coordinate{x, y, z})
		if r > maxR {
			maxR = r
			maxIndex = index
		}
		index++
	}
	count := 0
	maxRobot := coordinates[maxIndex]
	for _, robot := range coordinates {
		dist := abs(robot.x-maxRobot.x) + abs(robot.y-maxRobot.y) + abs(robot.z-maxRobot.z)
		if dist <= maxR {
			count++
		}
	}
	fmt.Println(count)
}
