package main

func isBoomerang(points [][]int) bool {
	slope1 := (points[0][0] - points[1][0]) / (points[0][1] - points[1][1])
	slope2 := (points[1][0] - points[2][0]) / (points[1][1] - points[2][1])
	return slope1 != slope2
}

func main() {}
