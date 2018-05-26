package twice_linear

func DblLinear(n int) int {
	px, py := 0, 0
	numbers := []int{1}
	for len(numbers) < n+1 {
		x := 2*numbers[px] + 1
		y := 3*numbers[py] + 1
		if x < y {
			px++
			if x <= numbers[len(numbers)-1] {
				continue
			}
			numbers = append(numbers, x)
		} else {
			py++
			if y <= numbers[len(numbers)-1] {
				continue
			}
			numbers = append(numbers, y)
		}
	}
	return numbers[n]
}
