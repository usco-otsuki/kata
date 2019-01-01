// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// )

// type point struct {
// 	x, y int
// }

// type cart struct {
// 	point      point
// 	direction  byte
// 	interCount int
// }

// type carts []cart

// func (c carts) Len() int {
// 	return len(c)
// }
// func (c carts) Swap(i, j int) {
// 	c[i], c[j] = c[j], c[i]
// }
// func (c carts) Less(i, j int) bool {
// 	return c[i].point.x < c[j].point.x || (c[i].point.x == c[j].point.x && c[i].point.y < c[j].point.y)
// }

// func PrintMap(track map[point]byte, carts carts, row, column int) {
// 	m := make(map[point]byte)
// 	for key, value := range track {
// 		m[key] = value
// 	}
// 	for _, cart := range carts {
// 		m[cart.point] = cart.direction
// 	}
// 	for i := 0; i < row; i++ {
// 		for j := 0; j < column; j++ {
// 			fmt.Printf("%c", m[point{i, j}])
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println()
// }

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	row := 0
// 	column := 0
// 	m := make(map[point]byte)
// 	carts := carts{}
// 	cartPos := make(map[point]bool)
// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		for column = 0; column < len(line); column++ {
// 			ch := line[column]
// 			if ch == '<' || ch == '>' || ch == '^' || ch == 'v' {
// 				carts = append(carts, cart{point{row, column}, ch, 0})
// 				if ch == '<' || ch == '>' {
// 					m[point{row, column}] = '-'
// 				} else {
// 					m[point{row, column}] = '|'
// 				}
// 			} else {
// 				m[point{row, column}] = ch
// 			}
// 		}
// 		row++
// 	}
// 	for {
// 		// PrintMap(m, carts, row, column)
// 		sort.Sort(carts)
// 		cartPos = make(map[point]bool)
// 		for i := 0; i < len(carts); i++ {
// 			cart := &carts[i]
// 			p := cart.point
// 			switch cart.direction {
// 			case '<':
// 				ch := m[point{p.x, p.y - 1}]
// 				cart.point = point{p.x, p.y - 1}
// 				switch ch {
// 				case '+':
// 					if cart.interCount == 0 {
// 						cart.direction = 'v'
// 					} else if cart.interCount == 2 {
// 						cart.direction = '^'
// 					}
// 					cart.interCount = (cart.interCount + 1) % 3
// 				case '/':
// 					cart.direction = 'v'
// 				case '\\':
// 					cart.direction = '^'
// 				}
// 			case '>':
// 				ch := m[point{p.x, p.y + 1}]
// 				cart.point = point{p.x, p.y + 1}
// 				switch ch {
// 				case '+':
// 					if cart.interCount == 0 {
// 						cart.direction = '^'
// 					} else if cart.interCount == 2 {
// 						cart.direction = 'v'
// 					}
// 					cart.interCount = (cart.interCount + 1) % 3
// 				case '/':
// 					cart.direction = '^'
// 				case '\\':
// 					cart.direction = 'v'
// 				}
// 			case '^':
// 				ch := m[point{p.x - 1, p.y}]
// 				cart.point = point{p.x - 1, p.y}
// 				switch ch {
// 				case '+':
// 					if cart.interCount == 0 {
// 						cart.direction = '<'
// 					} else if cart.interCount == 2 {
// 						cart.direction = '>'
// 					}
// 					cart.interCount = (cart.interCount + 1) % 3
// 				case '/':
// 					cart.direction = '>'
// 				case '\\':
// 					cart.direction = '<'
// 				}

// 			case 'v':
// 				ch := m[point{p.x + 1, p.y}]
// 				cart.point = point{p.x + 1, p.y}
// 				switch ch {
// 				case '+':
// 					if cart.interCount == 0 {
// 						cart.direction = '>'
// 					} else if cart.interCount == 2 {
// 						cart.direction = '<'
// 					}
// 					cart.interCount = (cart.interCount + 1) % 3
// 				case '/':
// 					cart.direction = '<'
// 				case '\\':
// 					cart.direction = '>'
// 				}

// 			}
// 			if _, ok := cartPos[cart.point]; ok {
// 				fmt.Println(cart.point)
// 				return
// 			} else {
// 				cartPos[cart.point] = true
// 			}
// 		}
// 	}
// }
