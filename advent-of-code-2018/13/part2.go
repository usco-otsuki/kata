package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	x, y int
}

type Cart struct {
	point      Point
	direction  byte
	interCount int
}

type Carts []Cart

func (c Carts) Len() int {
	return len(c)
}
func (c Carts) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Carts) Less(i, j int) bool {
	return c[i].point.x < c[j].point.x || (c[i].point.x == c[j].point.x && c[i].point.y < c[j].point.y)
}

func PrintMap(track map[Point]byte, carts Carts, row, column int) {
	m := make(map[Point]byte)
	for key, value := range track {
		m[key] = value
	}
	for _, cart := range carts {
		m[cart.point] = cart.direction
	}
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			fmt.Printf("%c", m[Point{i, j}])
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	row := 0
	column := 0
	m := make(map[Point]byte)
	carts := Carts{}
	cartPos := make(map[Point][]*Cart)
	for scanner.Scan() {
		line := scanner.Text()
		for column = 0; column < len(line); column++ {
			ch := line[column]
			if ch == '<' || ch == '>' || ch == '^' || ch == 'v' {
				carts = append(carts, Cart{Point{row, column}, ch, 0})
				if ch == '<' || ch == '>' {
					m[Point{row, column}] = '-'
				} else {
					m[Point{row, column}] = '|'
				}
			} else {
				m[Point{row, column}] = ch
			}
		}
		row++
	}
	for t := 0; ; t++ {
		// if t%1000 == 0 {
		// 	PrintMap(m, carts, row, column)
		// }
		sort.Sort(carts)
		cartPos = make(map[Point][]*Cart)
		newCarts := Carts{}
		for i := 0; i < len(carts); i++ {
			cart := &carts[i]
			p := cart.point
			switch cart.direction {
			case '<':
				ch := m[Point{p.x, p.y - 1}]
				cart.point = Point{p.x, p.y - 1}
				switch ch {
				case '+':
					if cart.interCount == 0 {
						cart.direction = 'v'
					} else if cart.interCount == 2 {
						cart.direction = '^'
					}
					cart.interCount = (cart.interCount + 1) % 3
				case '/':
					cart.direction = 'v'
				case '\\':
					cart.direction = '^'
				}
			case '>':
				ch := m[Point{p.x, p.y + 1}]
				cart.point = Point{p.x, p.y + 1}
				switch ch {
				case '+':
					if cart.interCount == 0 {
						cart.direction = '^'
					} else if cart.interCount == 2 {
						cart.direction = 'v'
					}
					cart.interCount = (cart.interCount + 1) % 3
				case '/':
					cart.direction = '^'
				case '\\':
					cart.direction = 'v'
				}
			case '^':
				ch := m[Point{p.x - 1, p.y}]
				cart.point = Point{p.x - 1, p.y}
				switch ch {
				case '+':
					if cart.interCount == 0 {
						cart.direction = '<'
					} else if cart.interCount == 2 {
						cart.direction = '>'
					}
					cart.interCount = (cart.interCount + 1) % 3
				case '/':
					cart.direction = '>'
				case '\\':
					cart.direction = '<'
				}

			case 'v':
				ch := m[Point{p.x + 1, p.y}]
				cart.point = Point{p.x + 1, p.y}
				switch ch {
				case '+':
					if cart.interCount == 0 {
						cart.direction = '>'
					} else if cart.interCount == 2 {
						cart.direction = '<'
					}
					cart.interCount = (cart.interCount + 1) % 3
				case '/':
					cart.direction = '<'
				case '\\':
					cart.direction = '>'
				}

			}
			if _, ok := cartPos[cart.point]; !ok {
				cartPos[cart.point] = []*Cart{}
			}
			cartPos[cart.point] = append(cartPos[cart.point], cart)
		}
		for _, value := range cartPos {
			if len(value) == 1 {
				newCarts = append(newCarts, *(value[0]))
			}
		}
		carts = newCarts
		if len(carts) == 1 {
			PrintMap(m, carts, row, column)
			fmt.Println(carts[0].point)
			return
		}
	}
}
