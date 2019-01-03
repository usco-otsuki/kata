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
	activated  bool
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

func findCollision(carts *Carts) (*Cart, *Cart) {
	cartPos := make(map[Point]*Cart)
	for i := 0; i < len(*carts); i++ {
		cart := &((*carts)[i])
		if !cart.activated {
			continue
		}
		if c, ok := cartPos[cart.point]; ok {
			return cart, c
		}
		cartPos[cart.point] = cart
	}
	return nil, nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	row := 0
	column := 0
	m := make(map[Point]byte)
	carts := Carts{}
	for scanner.Scan() {
		line := scanner.Text()
		for column = 0; column < len(line); column++ {
			ch := line[column]
			if ch == '<' || ch == '>' || ch == '^' || ch == 'v' {
				carts = append(carts, Cart{Point{row, column}, true, ch, 0})
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
	for {
		sort.Sort(carts)
		newCarts := Carts{}
		for i := 0; i < len(carts); i++ {
			cart := &carts[i]
			if !cart.activated {
				continue
			}
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
			if c1, c2 := findCollision(&carts); c1 != nil && c2 != nil {
				c1.activated = false
				c2.activated = false
			}
		}
		for _, cart := range carts {
			if cart.activated {
				newCarts = append(newCarts, cart)
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
