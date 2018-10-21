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

func main() {
	scanner := NewScanner()
	n := scanner.Read()
	m := scanner.Read()
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			d[i][j] = 1000
		}
		d[i][i] = 0
	}

	edges := []struct{ from, to, len int }{}

	for i := 0; i < m; i++ {
		a := scanner.Read()
		b := scanner.Read()
		c := scanner.Read()
		a--
		b--
		d[a][b] = c
		d[b][a] = c
		edges = append(edges, struct{ from, to, len int }{a, b, c})
	}

	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				if d[i][j] > d[i][k]+d[j][k] {
					d[i][j] = d[i][k] + d[j][k]
					d[j][i] = d[i][k] + d[j][k]
				}
			}
		}
	}

	count := 0
	for _, edge := range edges {
		i, j, val := edge.from, edge.to, edge.len
		isIncluded := false
		for s := 0; s < n; s++ {
			if d[s][i]+val == d[s][j] || d[s][j]+val == d[s][i] {
				isIncluded = true
				break
			}
		}
		if !isIncluded {
			count++
		}
	}

	fmt.Println(count)
}
