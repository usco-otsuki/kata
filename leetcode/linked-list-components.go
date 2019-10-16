package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func numComponents(head *ListNode, G []int) int {
	count := 0
	isInGroup := false
	exists := map[int]bool{}
	for _, val := range G {
		exists[val] = true
	}
	for p := head; p != nil; p = p.Next {
		if _, ok := exists[p.Val]; ok {
			isInGroup = true
		} else {
			if isInGroup {
				count++
				isInGroup = false
			}
		}
	}
	if isInGroup {
		count++
	}
	return count
}

func main() {}
