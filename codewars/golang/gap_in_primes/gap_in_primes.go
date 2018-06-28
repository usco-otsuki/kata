package gap_in_primes

func isPrime(n int) bool {
	if n == 1 || n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func EnumPrimes(m, n int) []int {
	primes := []int{}
	if m%2 == 0 {
		m++
	}
	for i := m; i <= n; i += 2 {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	return primes
}

func Gap(g, m, n int) []int {
	primes := EnumPrimes(m, n)
	for i := 0; i < len(primes)-1; i++ {
		if primes[i+1]-primes[i] == g {
			return []int{primes[i], primes[i+1]}
		}
	}
	return nil
}
