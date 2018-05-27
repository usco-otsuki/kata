package friend_cheating

func RemovNb(n uint64) [][2]uint64 {
	result := [][2]uint64{}
	var sum, i uint64
	sum = n * (n + 1) / 2
	for i = 1; i <= n; i++ {
		j := (sum - i) / (i + 1)
		if i == j || j > n {
			continue
		}
		target := sum - i - j
		if i*j == target {
			result = append(result, [2]uint64{i, j})
		}
	}
	if len(result) == 0 {
		return nil
	}
	return result
}
