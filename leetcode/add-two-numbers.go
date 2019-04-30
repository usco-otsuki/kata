package main

import "math/big"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := big.NewInt(0)
	for digits, p := 1, l1; p != nil; p, digits = p.Next, digits*10 {
		sum += big.NewInt(p.Val).digits
	}
	for digits, p := 1, l2; p != nil; p, digits = p.Next, digits*10 {
		sum += p.Val * digits
	}
	var head *ListNode = nil
	p := &head
	for {
		val := sum % 10
		newNode := &ListNode{Val: val, Next: nil}
		*p = newNode
		p = &newNode.Next
		sum /= 10
		if sum == 0 {
			break
		}
	}
	return head
}

func main() {

}
