package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

type Scanner struct {
	scanner *bufio.Scanner
}

func NewScanner() *Scanner {
	scanner := &Scanner{scanner: bufio.NewScanner(os.Stdin)}
	scanner.scanner.Split(bufio.ScanWords)
	return scanner
}

func (s *Scanner) Read() int {
	s.scanner.Scan()
	val, _ := strconv.Atoi(s.scanner.Text())
	return val
}

type Pair struct {
	first, second int
}

func findIntersection(as, bs [][]int) [][]int {
	intersection := [][]int{}
	for _, a := range as {
		for _, b := range bs {
			if a[0] == b[0] && a[1] == b[1] {
				intersection = append(intersection, a)
			}
		}
	}
	return intersection
}

func main() {
	aTable := make(map[int][][]int)
	bTable := make(map[int][][]int)
	for t1 := 0; t1 <= 3; t1++ {
		for t2 := 0; t2 <= 3; t2++ {
			a := t1 | t2
			b := t1 & t2
			aTable[a] = append(aTable[a], []int{t1, t2})
			bTable[b] = append(bTable[b], []int{t1, t2})
		}
	}
	sc := NewScanner()
	n := sc.Read()
	prevA, prevB := sc.Read(), sc.Read()
	prevCandidates := findIntersection(aTable[prevA], bTable[prevB])
	for i := 0; i < n; i++ {
		newCandidates := [][]int{}
		a, b := sc.Read(), sc.Read()
		candidates := findIntersection(aTable[a], bTable[b])
		for _, prevCandidate := range prevCandidates {
			for _, candidate := range candidates {
				if prevCandidate[len(prevCandidates)-1] != candidate[0] {
					continue
				}
				newCandidates = append(newCandidates, prevCandidate, prevCandidate[1:2])
			}
		}

		prevCandidates = newCandidates
	}
	if len(prevCandidates) == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println(prevCandidates[0])
	return
}
