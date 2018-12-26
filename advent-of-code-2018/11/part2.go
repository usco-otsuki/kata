package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const SIZE = 300

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
	var maxPoint point
	maxPower := 0
	maxSize := 1
	m := [2]map[point]int{make(map[point]int), make(map[point]int)}
	m1 := make(map[point]int)
	for i := 1; i <= SIZE; i++ {
		for j := 1; j <= SIZE; j++ {
			power := ((i+10)*j + sn) * (i + 10)
			power = power / 100 % 10
			power -= 5
			m[1][point{i, j}] = power
			m1[point{i, j}] = power
			if maxPower < power {
				maxPower = power
				maxPoint = point{i, j}
			}
			// fmt.Printf("%d ", power)
		}
		// fmt.Println()
	}
	for size := 2; size <= SIZE; size++ {
		m[size%2] = make(map[point]int)
		curM := m[size%2]
		pM := m[(size-1)%2] // prev M
		for i := 1; i <= SIZE-size+1; i++ {
			for j := 1; j <= SIZE; j++ {
				curM[point{i, j}] = pM[point{i, j}] + m1[point{i, j + size - 1}]
			}
		}
		for i := 1; i <= SIZE-size; i++ {
			power := 0
			for j := 1; j <= size; j++ {
				power += curM[point{i, j}]
			}
			if maxPower < power {
				maxPower = power
				maxPoint = point{i, 1}
				maxSize = size
			}
			for j := 2; j <= SIZE-size; j++ {
				power = power + curM[point{i, j + size - 1}] - curM[point{i, j - 1}]
				if maxPower < power {
					maxPoint = point{i, j}
					maxPower = power
					maxSize = size
				}
			}
		}
	}
	fmt.Println(maxStrip)
	fmt.Print(maxPoint, maxSize, maxPower)
}
