package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const SIZE = 100

type point struct {
	x, y int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	sn, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	m := [3]map[point]int{make(map[point]int), make(map[point]int), make(map[point]int)}
	m1 := make(map[point]int)
	for i := 1; i <= SIZE; i++ {
		for j := 1; j <= SIZE; j++ {
			power := ((i+10)*j + sn) * (i + 10)
			power = power / 100 % 10
			power -= 5
			m[1][point{i, j}] = power
			m1[point{i, j}] = power
		}
	}
	var maxPoint point
	maxPower := 0
	maxSize := 0
	for size := 2; size <= SIZE; size++ {
		m[size%3] = make(map[point]int)
		pM := m[(size-1)%3]  // prev M
		ppM := m[(size-2)%3] // prev prev M
		for i := 1; i <= SIZE-size; i++ {
			for j := 1; j <= SIZE-size; j++ {
				power := pM[point{i, j}] + pM[point{i + 1, j}] + pM[point{i, j + 1}] + pM[point{i + 1, j + 1}] -
					(ppM[point{i + 1, j + 1}] + ppM[point{i + 1, j}] + ppM[point{i, j + 1}] + ppM[point{i + SIZE - size, j + 1}] + ppM[point{i + 1, j + SIZE - size}])
				m[size%3][point{i, j}] = power
				if power > maxPower {
					maxPoint = point{i, j}
					maxPower = power
					maxSize = size
				}
			}
		}
	}
	fmt.Print(maxPoint, maxSize, maxPower)
}

// *****
// *****
// *****
// *****
// *****
