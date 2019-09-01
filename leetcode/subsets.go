package main

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{[]int{}}
	}
	powerset := [][]int{}
	for _, set := range subsets(nums[1:]) {
		powerset = append(powerset, set, append([]int{nums[0]}, set...))
	}
	return powerset
}

func main()
