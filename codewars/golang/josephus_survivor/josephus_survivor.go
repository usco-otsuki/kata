package josephus_survivor

func JosephusSurvivor(n, k int) int {
	survivor := 0
	for i := 1; i < n; i++ {
		survivor = (survivor + k) % (i + 1)
	}
	return survivor + 1
}
