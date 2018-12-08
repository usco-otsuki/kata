package main

import "fmt"

func main() {
	var val int
	m := map[int]bool{0: true}
	a := []int{}
	sum := 0
	for {
		_, err := fmt.Scan(&val)
		if err != nil {
			break
		}
		a = append(a, val)
	}

	for {
		for _, el := range a {
			sum += el
			if _, ok := m[sum]; ok {
				fmt.Println(sum)
				return
			}
			m[sum] = true
		}
	}
}
