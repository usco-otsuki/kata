package main

import "fmt"

func main() {
	var val int
	var sum int = 0
	for {
		_, err := fmt.Scan(&val)
		if err != nil {
			break
		}
		sum += val
	}
	fmt.Println(sum)
}
