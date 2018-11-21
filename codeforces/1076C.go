package main

import (
	"bufio"
	"fmt"
	"math"
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

// a + b = d
// a  = d / b
/*
 d / b + b = d
 d + b**2 = d * b
 b ** 2 - d * b + d = 0
 b =  (d + sqrt(d**2 - 4 * d) ) / 2
 d > 4
*/

func solve(d float64) (bool, float64, float64) {
	if 0 < d && d < 4 {
		return false, 0, 0
	}
	b := (d + math.Sqrt(d*d-4*d)) / 2
	a := d - b
	if a >= 0 {
		return true, a, b
	}
	return false, 0, 0
}

func main() {
	scanner := NewScanner()
	t := scanner.Read()
	for i := 0; i < t; i++ {
		d := scanner.Read()
		ok, a, b := solve(float64(d))
		if a < b {
			a, b = b, a
		}
		if ok {
			fmt.Printf("Y %.10f %.10f\n", a, b)
		} else {
			fmt.Printf("N\n")
		}
	}
}
