package recreation_two

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Prod2Sum(a, b, c, d int) [][]int {
	var e, f int
	e = Abs(a*c + b*d)
	f = Abs(a*d - b*c)
	t1 := []int{Min(e, f), Max(e, f)}
	e = Abs(a*d + b*c)
	f = Abs(a*c - b*d)
	t2 := []int{Min(e, f), Max(e, f)}
	if t1[0] == t2[0] && t1[1] == t2[1] {
		return [][]int{t1}
	}
	if t1[0] < t2[0] {
		return [][]int{t1, t2}
	}
	return [][]int{t2, t1}
}
