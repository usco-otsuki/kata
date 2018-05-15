package product_of_consecutive_fib_numbers

func ProductFib(prod uint64) [3]uint64 {
	if prod < 1 {
		panic("prod should be >= 1")
	}
	// 'a' contains the head element of fib sequence computed so far
	// 'b' contains element right before 'a'
	var a, b, tmp uint64 = 1, 1, 0
	for a*b < prod {
		tmp = a + b
		b = a
		a = tmp
	}
	if a*b == prod {
		return [3]uint64{b, a, 1}
	}
	return [3]uint64{b, a, 0}
}
