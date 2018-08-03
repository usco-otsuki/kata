package sum_of_positives

func PositiveSum(numbers []int) int {
	sum := 0
	for _, val := range numbers {
		if val > 0 {
			sum += val
		}
	}
	return sum
}
