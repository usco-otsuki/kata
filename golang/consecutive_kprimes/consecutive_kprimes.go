package consecutive_kprimes

func isKPrime(n, k int) bool {
	cnt := 0
	for n%2 == 0 {
		n >>= 1
		cnt++
	}
	for d := 3; d*d <= n; d++ {
		for n%d == 0 {
			n /= d
			cnt++
		}
	}
	return cnt == k
}

func ConsecKprimes(k int, arr []int) int {
	isPrevKPrime := false
	cnt := 0
	for _, n := range arr {
		if isKPrime(n, k) {
			if isPrevKPrime {
				cnt++
			}
			isPrevKPrime = true
		} else {
			isPrevKPrime = false
		}
	}
	return cnt
}
