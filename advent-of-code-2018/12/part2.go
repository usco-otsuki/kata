package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type note struct {
	before string
	after  byte
}

// const NUM_GEN = 50000000000

const NUM_GEN = 200

func AddDots(state []byte) ([]byte, int) {
	newState := []byte{}
	var index int
	for index = 0; index < len(state); index++ {
		if state[index] == '#' {
			break
		}
	}
	dots := []byte{}
	for i := 0; i < 4-index; i++ {
		dots = append(dots, '.')
	}
	n := 4 - index
	if n < 0 {
		n = 0
	}
	newState = append(dots, state...)
	for index = len(state) - 1; index >= 0; index-- {
		if state[index] == '#' {
			break
		}
	}
	dots = []byte{}
	for i := 0; i < 4-(len(state)-1-index); i++ {
		dots = append(dots, '.')
	}
	newState = append(newState, dots...)

	return newState, n
}

func main() {
	notes := []note{}
	var prevStateString string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "initial state: %s", &prevStateString)
	scanner.Scan() // Read empty line
	prevState := []byte(prevStateString)
	for scanner.Scan() {
		var before string
		var after byte
		n, err := fmt.Sscanf(scanner.Text(), "%s => %c", &before, &after)
		if n != 2 {
			log.Fatalf("Scanned %d, want 2", n)
		}
		if err != nil {
			panic(err)
		}
		notes = append(notes, note{before, after})
	}
	var state []byte
	shift := 0
	for i := 0; i < NUM_GEN; i++ {
		// fmt.Println(i + 1)
		state = []byte{}
		var prependNum int
		prevState, prependNum = AddDots(prevState)
		// if prependNum-2 > 0 {
		shift -= prependNum - 2
		// }
		// fmt.Println(string(prevState))
		for j := 2; j < len(prevState)-2; j++ {
			state = append(state, '.')
			for _, n := range notes {
				hasMatch := true
				for k := 0; k < len(n.before); k++ {
					if n.before[k] != prevState[j-2+k] {
						hasMatch = false
						break
					}
				}
				if hasMatch {
					state[j-2] = n.after
					break
				}
			}
		}
		prevState = state
		sum := 0
		for l, c := range state {
			if c == '#' {
				sum += shift + l
			}
		}
		fmt.Println(i+1, sum, shift, string(state))
		// fmt.Println(i+1, sum)

	}
}
