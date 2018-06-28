package diophantine_equation

import "sort"

type Pairs [][]int

func (p Pairs) Len() int {
	return len(p)
}

func (p Pairs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Pairs) Less(i, j int) bool {
	return p[i][0] > p[j][0]
}

func Solequa(n int) [][]int {
	result := Pairs{}
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
	sort.Sort(result)
	return result
}
