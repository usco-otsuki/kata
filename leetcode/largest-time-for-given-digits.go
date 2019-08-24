package main

import (
	"fmt"
)

func getPermutations(list []int) [][]int {
	if len(list) == 0 {
		return [][]int{}
	}
	if len(list) == 1 {
		return [][]int{list}
	}
	permutations := [][]int{}
	subPerms := getPermutations(list[1:])

	for _, subPerm := range subPerms {
		for i := 0; i <= len(subPerm); i++ {
			latter := append([]int{list[0]}, subPerm[i:]...)
			newPerm := append(append([]int{}, subPerm[:i]...), latter...)
			permutations = append(permutations, newPerm)
		}
	}
	return permutations
}

func largestTimeFromDigits(A []int) string {
	maxStr := ""
	permutations := getPermutations(A)
	for _, list := range permutations {
		if list[0] >= 3 {
			continue
		}
		if list[0] == 2 && list[1] >= 4 {
			continue
		}
		if list[2] >= 6 {
			continue
		}
		newStr := fmt.Sprintf("%d%d:%d%d", list[0], list[1], list[2], list[3])
		if newStr > maxStr {
			maxStr = newStr
		}
	}
	return maxStr
}

func main() {
	fmt.Println(largestTimeFromDigits([]int{1, 2, 3, 4}))
	fmt.Println(largestTimeFromDigits([]int{5, 5, 5, 5}))
	fmt.Println(largestTimeFromDigits([]int{0, 4, 0, 0}))
	fmt.Println(largestTimeFromDigits([]int{2, 0, 6, 6}))
}
