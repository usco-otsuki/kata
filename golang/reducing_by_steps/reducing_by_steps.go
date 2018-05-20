package reducing_by_steps

func Abs(x int) int {
	if x < 0 {
		x = -x
	}
	return x
}
func Gcdi(x, y int) int {
	x = Abs(x)
	y = Abs(y)
	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}
func Som(x, y int) int {
	return x + y
}
func Maxi(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func Mini(x, y int) int {
	if x > y {
		return y
	}
	return x
}
func Lcmu(x, y int) int {
	x = Abs(x)
	y = Abs(y)
	return x * y / Gcdi(x, y)
}

type FParam func(int, int) int

func OperArray(f FParam, arr []int, init int) []int {
	result := []int{}
	for _, val := range arr {
		init = f(init, val)
		result = append(result, init)
	}
	return result
}
