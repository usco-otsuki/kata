package integer_partitions

import (
	"fmt"
	"sort"
)

var cache = map[int][][]int{1: {{1}}}

func EnumPart(n int) [][]int {
	if val, ok := cache[n]; ok {
		return val
	}

	parts := [][]int{{n}}
	for i := 1; i < n; i++ {
		for _, val := range EnumPart(i) {
			if val[0] <= n-i {
				parts = append(parts, append([]int{n - i}, val...))
			}
		}
	}
	cache[n] = parts
	return parts
}

func Prod(parts [][]int) []int {
	exists := map[int]bool{}
	newParts := []int{}
	for _, part := range parts {
		val := 1
		for _, num := range part {
			val *= num
		}
		if _, ok := exists[val]; !ok {
			exists[val] = true
			newParts = append(newParts, val)
		}
	}
	sort.Ints(newParts)
	return newParts
}

func Average(a []int) float64 {
	sum := 0.0
	for _, val := range a {
		sum += float64(val)
	}
	return sum / float64(len(a))
}

func Median(a []int) float64 {
	n := len(a)
	if n%2 == 0 {
		return (float64(a[n>>1]) + float64(a[n>>1-1])) / 2
	} else {
		return float64(a[n>>1])
	}
}

func Part(n int) string {
	values := Prod(EnumPart(n))
	return fmt.Sprintf("Range: %d Average: %.2f Median: %.2f", values[len(values)-1]-values[0], Average(values), Median(values))
}
