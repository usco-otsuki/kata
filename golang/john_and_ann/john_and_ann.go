package john_and_ann

// import "fmt"

func AnnJohn(n int) ([]int, []int) {
	a := []int{1}
	j := []int{0}
	for i := 1; i < n; i++ {
		john := i - a[j[i-1]]
		j = append(j, john)
		ann := i - j[a[i-1]]
		a = append(a, ann)
	}
	return a, j
}

func Ann(n int) []int {
	list, _ := AnnJohn(n)
	return list
}
func John(n int) []int {
	_, list := AnnJohn(n)
	return list
}
func SumJohn(n int) int {
	sum := 0
	for _, val := range John(n) {
		sum += val
	}
	return sum
}

func SumAnn(n int) int {
	sum := 0
	for _, val := range Ann(n) {
		sum += val
	}
	return sum
}
