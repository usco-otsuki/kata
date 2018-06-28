package perimeter_of_squares_in_a_rectangle

func Perimeter(n int) int {
	val := [2]int{1, 1}
	sum := 1
	if n+1 >= 2 {
		sum += 1
	}

	for i := 3; i <= n+1; i++ {
		val[i%2] = val[0] + val[1]
		sum += val[i%2]
	}

	return sum * 4
}
