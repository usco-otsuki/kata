package kata

func Multiple3And5(number int) int {
	sum := 0
	for i := 3; i < number; i += 3 {
		sum += i
	}

	for i := 5; i < number; i += 5 {
		if i%3 > 0 {
			sum += i
		}
	}

	return sum
}
