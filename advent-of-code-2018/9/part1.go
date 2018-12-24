package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var n, last int
	fmt.Sscanf(scanner.Text(), "%d players; last marble is worth %d points", &n, &last)
	scores := make(map[int]int)
	currentScore := -1
	currentMarble := 2
	currentIndex := 1
	marbles := []int{0, 1}
	for currentScore != last {
		fmt.Println(marbles)
		if currentMarble%23 == 0 {
			removeIndex := (currentIndex + len(marbles) - 7) % len(marbles)
			currentScore = currentMarble + marbles[removeIndex]
			scores[currentMarble%n] += currentScore
			marbles = append(marbles[:removeIndex], marbles[removeIndex+1:]...)
			currentMarble++
			currentIndex = removeIndex
			continue
		}
		newIndex := (currentIndex+1)%len(marbles) + 1
		if (currentIndex+2)%len(marbles) == 0 {
			marbles = append(marbles[:(currentIndex+1)%len(marbles)+1], currentMarble)
		} else {
			marbles = append(marbles[:(currentIndex+1)%len(marbles)+1], append([]int{currentMarble}, marbles[(currentIndex+2)%len(marbles):]...)...)
		}
		currentIndex = newIndex
		currentScore = 0
		currentMarble++
	}
	maxScore := 0
	for _, score := range scores {
		if maxScore < score {
			maxScore = score
		}
	}
	fmt.Println(maxScore)
}
