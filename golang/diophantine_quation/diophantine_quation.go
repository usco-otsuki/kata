package diophantine_quation

func Solequa(n int) [][]int {
	result := [][]int{}
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			a := i
			b := n / i
			if b > a {
				a, b = b, a
			}
			if (a+b)%2 != 0 || (a-b)%4 != 0 {
				continue
			}
			x := (a + b) / 2
			y := (a - b) / 4
			result = append(result, []int{x, y})
		}
	}
	return result
}

/*
(x - 2y)*(x+2y) = a * b = n
x - 2y = a
x + 2y = b

-4y = a - b => y = (b - a) / 4
x = b + (a-b)/2 = (a + b) / 2

x - 2y =
*/
