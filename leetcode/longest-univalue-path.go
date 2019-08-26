package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func longestUnivaluePath(root *TreeNode) int {
	maxValue := 0
	var getMaxValue func(node *TreeNode) int
	getMaxValue = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftMax := getMaxValue(node.Left)
		rightMax := getMaxValue(node.Right)

		if node.Left != nil && node.Left.Val == node.Val {
			leftMax++
		} else {
			leftMax = 0
		}
		if node.Right != nil && node.Right.Val == node.Val {
			rightMax++
		} else {
			rightMax = 0
		}
		if leftMax+rightMax > maxValue {
			maxValue = leftMax + rightMax
		}
		if leftMax > rightMax {
			return leftMax
		}
		return rightMax
	}
	getMaxValue(root)
	return maxValue
}

func main() {

}
