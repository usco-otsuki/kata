package fibo_akin

func getList(n int) []int {
	list := []int{0, 1, 1}
	if n == 1 {
		return list[1:2]
	}
	if n == 2 {
		return list[1:3]
	}
	for i := 3; i <= n; i++ {
		val := list[i-list[i-1]] + list[i-list[i-2]]
		list = append(list, val)
	}
	return list[1:]
}

func LengthSupUk(n, k int) int {
	u := getList(n)
	count := 0
	for _, val := range u {
		if val >= k {
			count++
		}
	}
	return count
}
func Comp(n int) int {
	u := getList(n)
	count := 0
	for i := 0; i < len(u)-1; i++ {
		if u[i+1] < u[i] {
			count++
		}
	}
	return count
}
