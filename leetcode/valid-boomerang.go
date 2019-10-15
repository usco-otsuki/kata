package main

func isBoomerang(points [][]int) bool {
	val := (points[0][0]-points[2][0])*(points[1][1]-points[2][1]) - (points[1][0]-points[2][0])*(points[0][1]-points[2][1])
	if val < 0 {
		val = -val
	}
	fmt.Println(val)
	area := float64(val) / 2.0
	return area > 0
}

func main() {}
