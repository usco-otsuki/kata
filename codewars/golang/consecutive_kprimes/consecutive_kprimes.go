package consecutive_kprimes

func isKPrime(n, k int) bool {
	cnt := 0
	num := n
	if n == 1 {
		return false
	}
	for num%2 == 0 {
		num >>= 1
		cnt++
	}
	for d := 3; d*d <= n; d += 2 {
		for num%d == 0 {
			num /= d
			cnt++
		}
	}
	if num != 1 {
		cnt++
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
