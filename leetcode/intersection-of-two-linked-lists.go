package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	countA, countB := 0, 0
	for p := headA; p != nil; p = p.Next {
		countA++
	}
	for p := headB; p != nil; p = p.Next {
		countB++
	}
	var bigP, smallP *ListNode = nil, nil
	bigLen, smallLen := 0, 0
	if countA > countB {
		bigLen = countA
		smallLen = countB
		bigP = headA
		smallP = headB
	} else {
		bigLen = countB
		smallLen = countA
		bigP = headB
		smallP = headA
	}

	for bigLen > smallLen {
		bigP = bigP.Next
		bigLen--
	}

	for bigP != nil && smallP != nil && bigP != smallP {
		bigP = bigP.Next
		smallP = smallP.Next
	}
	if bigP == smallP {
		return bigP
	}
	return nil
}
