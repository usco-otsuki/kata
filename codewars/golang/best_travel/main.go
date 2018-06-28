package best_travel

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func chooseBestSum(t, k, s int, ls []int) int {
	if t < 0 {
		return -1
	}
	if k == 0 {
		return s
	}
	if len(ls) == 0 {
		return -1
	}

	m := max(chooseBestSum(t-ls[0], k-1, s+ls[0], ls[1:]), chooseBestSum(t, k, s, ls[1:]))
	return m
}

func ChooseBestSum(t, k int, ls []int) int {
	return chooseBestSum(t, k, 0, ls)
}
