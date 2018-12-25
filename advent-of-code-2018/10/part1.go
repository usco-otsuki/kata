package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type point struct {
	x, y, vx, vy int
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func MovePoints(points *[]point, sign int) {
	for i := 0; i < len(*points); i++ {
		(*points)[i].x += sign * (*points)[i].vx
		(*points)[i].y += sign * (*points)[i].vy
	}
}

func FindRange(points []point) (int, int, int, int) {
	if len(points) == 0 {
		log.Fatalf("len(points) == 0")
	}
	minX, maxX := points[0].x, points[0].x
	minY, maxY := points[0].y, points[0].y
	for _, p := range points[1:] {
		if minX > p.x {
			minX = p.x
		}
		if maxX < p.x {
			maxX = p.x
		}
		if minY > p.y {
			minY = p.y
		}
		if maxY < p.y {
			maxY = p.y
		}
	}
	return minX, maxX, minY, maxY
}

func CalcCohesion(points []point) int {
	minX, maxX, minY, maxY := FindRange(points)
	return Abs(maxX-minX) + Abs(maxY-minY)
}

type coordinate struct {
	x, y int
}

func Print(points []point) {
	minX, maxX, minY, maxY := FindRange(points)
	fmt.Println(minX, maxX, minY, maxY)
	m := make(map[coordinate]bool)
	// shift each point so by minX so that each point is in the first quadrant
	for i := 0; i < len(points); i++ {
		points[i].x -= minX
		points[i].y -= minY
		m[coordinate{points[i].x, points[i].y}] = true
	}
	maxX, maxY = maxX-minX, maxY-minY
	fmt.Println(maxX, maxY)
	for i := 0; i <= maxX; i++ {
		for j := 0; j <= maxY; j++ {
			if _, ok := m[coordinate{i, j}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	points := []point{}
	var x0, y0, vx, vy int
	for scanner.Scan() {
		n, err := fmt.Sscanf(scanner.Text(), "position=<%d,%d> velocity=<%d,%d>", &x0, &y0, &vx, &vy)
		if n != 4 {
			log.Fatalf("Scanned %d, want 4", n)
		}
		if err != nil {
			panic(err)
		}
		points = append(points, point{x0, y0, vx, vy})
	}
	prev := 10000000000
	for i := 0; i < 1000000; i++ {
		cohesion := CalcCohesion(points)
		if prev < cohesion {
			MovePoints(&points, -1)
			break
		}
		prev = cohesion
		MovePoints(&points, 1)
	}
	Print(points)
}
