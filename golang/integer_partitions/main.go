package integer_partitions

import (
	"sort"
)

type Partitions [][]int

func (p Partitions) Len() int {
	return len(p)
}

func (p Partitions) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Partitions) Less(i, j int) bool {
	a, b := p[i], p[j]
	n := len(b)
	if len(a) < len(b) {
		n = len(a)
	}
	for i := 0; i < n; i++ {
		if a[i] < b[i] {
			return false
		} else if a[i] > b[i] {
			return true
		}
	}
	return true // true of false does not matter in this specifi problem
}

func IntSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func EnumPart(n int) [][]int {
	parts := Partitions{{n}}
	for i := 1; i <= n>>1; i++ {
		first := EnumPart(i)
		second := EnumPart(n - i)
		for _, val1 := range second {
			for _, val2 := range first {
				concat := append(val1, val2...)
				sort.Sort(sort.Reverse(sort.IntSlice(concat)))
				parts = append(parts, concat)
			}
		}
	}
	sort.Sort(parts)
	result := Partitions{parts[0]}
	for i := 1; i < len(parts); i++ {
		val1 := result[len(result)-1]
		val2 := parts[i]
		if !IntSliceEqual(val1, val2) {
			result = append(result, val2)
		}
	}
	return result
}

func Part(n int) string {
	// your code
	return ""
}
