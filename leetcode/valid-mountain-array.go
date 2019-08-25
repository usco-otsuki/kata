package main

import (
	"fmt"
)

func validMountainArray(A []int) bool {
	if len(A) < 3 {
		return false
	}
	i := 0
	for ; i < len(A)-1 && A[i] < A[i+1]; i++ {
	}
	if i == 0 || i == len(A)-1 {
		return false
	}
	for ; i < len(A)-1; i++ {
		if A[i] <= A[i+1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(validMountainArray([]int{0, 1, 2, 3, 4}) == false)
	fmt.Println(validMountainArray([]int{3, 4, 5, 4, 3}) == true)
	fmt.Println(validMountainArray([]int{3, 4, 5, 4, 5}) == false)
}
