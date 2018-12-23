package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numbers = []int{}

func Calc(index int) (int, int) {
	numChildren := numbers[index]
	numMeta := numbers[index+1]
	sum := 0
	startIndex := index + 2
	values := []int{}
	for i := 0; i < numChildren; i++ {
		var value int
		value, startIndex = Calc(startIndex)
		values = append(values, value)
	}
	for i := 0; i < numMeta; i++ {
		if numChildren > 0 {
			if numChildren >= numbers[startIndex+i] {
				sum += values[numbers[startIndex+i]-1]
			}
		} else {
			sum += numbers[startIndex+i]
		}
	}
	return sum, startIndex + numMeta
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, num)
	}
	sum, _ := Calc(0)
	fmt.Println(sum)
}
