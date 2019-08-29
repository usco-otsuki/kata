package main

import "fmt"

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	sum := 0
	step := 0
	for ; sum < target; step++ {
		sum += step
	}
	step--
	if sum == target {
		return step
	}
	if (sum-target)%2 == 0 {
		return step
	}
	if step%2 == 0 {
		return step + 1
	}
	return step + 2
}

func main() {
	fmt.Println(reachNumber(3))
	fmt.Println(reachNumber(2))
	fmt.Println(reachNumber(4))
	fmt.Println(reachNumber(9))
}
