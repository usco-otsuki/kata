package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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
	m := make(map[point]int)
	for i := 1; i <= 300; i++ {
		for j := 1; j <= 300; j++ {
			power := ((i+10)*j + sn) * (i + 10)
			power = power / 100 % 10
			power -= 5
			m[point{i, j}] = power
		}
	}
	var maxPoint point
	maxPower := 0
	for i := 1; i <= 300-3; i++ {
		for j := 1; j <= 300-3; j++ {
			power := 0
			for dx := 0; dx < 3; dx++ {
				for dy := 0; dy < 3; dy++ {
					power += m[point{i + dx, j + dy}]
				}
			}
			if power > maxPower {
				maxPoint = point{i, j}
				maxPower = power
			}
		}
	}
	fmt.Print(maxPoint)
}
