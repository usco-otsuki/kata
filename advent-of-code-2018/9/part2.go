package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value      int
	next, prev *Node
}

func (n *Node) InsertNext(value int) *Node {
	newNode := &Node{value: value, next: n.next, prev: n}
	n.next.prev = newNode
	n.next = newNode
	return newNode
}

func (n *Node) Delete() *Node {
	n.prev.next = n.next
	return n.prev.next
}

func (n *Node) Next(num int) *Node {
	current := n
	for i := 0; i < num; i++ {
		current = current.next
	}
	return current
}

func (n *Node) Prev(num int) *Node {
	current := n
	for i := 0; i < num; i++ {
		current = current.prev
	}
	return current
}

type Marbles struct {
	root            *Node
	scores          map[int]int
	nextMarbleValue int
	currentMarble   *Node
	numPlayers      int
}

func (m *Marbles) String() string {
	if m.currentMarble == nil {
		return ""
	}
	p := m.root
	values := []string{}
	for {
		values = append(values, strconv.Itoa(p.value))
		p = p.next
		if p == m.root {
			break
		}
	}
	return strings.Join(values, ", ") + ":" + strconv.Itoa(m.currentMarble.value)
}

func (m *Marbles) Add() {
	if m.currentMarble == nil {
		newNode := &Node{value: m.nextMarbleValue}
		newNode.prev = newNode
		newNode.next = newNode
		m.currentMarble = newNode
		m.root = newNode
	} else {
		if m.nextMarbleValue%23 == 0 {
			m.scores[m.nextMarbleValue%m.numPlayers] += m.nextMarbleValue + m.currentMarble.Prev(7).value
			m.currentMarble = m.currentMarble.Prev(7).Delete()
		} else {
			m.currentMarble = m.currentMarble.Next(1).InsertNext(m.nextMarbleValue)
		}
	}
	m.nextMarbleValue++
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var n, last int
	fmt.Sscanf(scanner.Text(), "%d players; last marble is worth %d points", &n, &last)
	marbles := &Marbles{scores: make(map[int]int), numPlayers: n}
	for marbles.nextMarbleValue < last*100 {
		marbles.Add()
	}
	maxScore := 0
	for _, score := range marbles.scores {
		if maxScore < score {
			maxScore = score
		}
	}
	fmt.Println(maxScore)
}
