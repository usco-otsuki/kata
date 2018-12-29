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

const NUM_GEN = 20

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
	for i := 0; i < NUM_GEN; i++ {
		state = []byte{'.', '.'}
		prevState = append([]byte("...."), append(prevState, []byte("....")...)...)
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
					state[j] = n.after
					break
				}
			}
		}
		state = append(state, []byte{'.', '.'}...)
		prevState = state
	}
	sum := 0
	for i, c := range state {
		if c == '#' {
			sum += -NUM_GEN*4 + i
		}
	}
	fmt.Println(sum)
}
